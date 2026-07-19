package evaluation

import (
	"ai-agent/internal/agent"
	"ai-agent/internal/app"
	"ai-agent/internal/domain"
	"ai-agent/internal/embedding"
	"ai-agent/internal/generation"
	"ai-agent/internal/ranking"
	"ai-agent/internal/retrieval"
	"ai-agent/internal/vectorindex"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Case struct {
	Name                   string   `json:"name"`
	Question               string   `json:"question"`
	Language               string   `json:"language"`
	ExpectedDocumentIDs    []string `json:"expected_document_ids"`
	ExpectedCategory       string   `json:"expected_category"`
	ExpectedTemporalStatus string   `json:"expected_temporal_status"`
	RequiredTerms          []string `json:"required_terms"`
	ForbiddenTerms         []string `json:"forbidden_terms"`
}

type Metrics struct {
	Total                  int
	RecallAt1              float64
	RecallAt3              float64
	RecallAt5              float64
	MRR                    float64
	LanguageAccuracy       float64
	CategoryAccuracy       float64
	TemporalAccuracy       float64
	CorrectResponseRate    float64
	FalsePositiveRate      float64
	AverageRetrievalMillis float64
	Failures               []Failure
}

type Failure struct {
	CaseName string
	Reason   string
}

type rankedDocument struct {
	ID             string
	Category       string
	Language       string
	TemporalStatus string
}

const evaluationSearchLimit = 5

func LoadCases(path string) ([]Case, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cases []Case
	if err := json.Unmarshal(content, &cases); err != nil {
		return nil, err
	}

	return cases, nil
}

func Run(cases []Case) Metrics {
	service, err := evaluationService()
	if err != nil {
		return Metrics{
			Total:    len(cases),
			Failures: []Failure{{CaseName: "setup", Reason: err.Error()}},
		}
	}

	var metrics Metrics
	metrics.Total = len(cases)

	if len(cases) == 0 {
		return metrics
	}

	var recallAt1, recallAt3, recallAt5, reciprocalRank float64
	var languageHits, categoryHits, temporalHits, correctResponses, falsePositives int
	var totalRetrievalDuration time.Duration
	var noAnswerCases int

	for _, testCase := range cases {
		start := time.Now()
		retrievalResult, err := service.Retrieve(context.Background(), testCase.Question, evaluationSearchLimit)
		totalRetrievalDuration += time.Since(start)
		if err != nil {
			metrics.Failures = append(metrics.Failures, Failure{CaseName: testCase.Name, Reason: err.Error()})
			continue
		}

		documents := rankedEvidence(retrievalResult.Evidence)
		rank := firstExpectedRank(documents, testCase.ExpectedDocumentIDs)

		if rank > 0 {
			reciprocalRank += 1 / float64(rank)
		}
		if rank > 0 && rank <= 1 {
			recallAt1++
		}
		if rank > 0 && rank <= 3 {
			recallAt3++
		}
		if rank > 0 && rank <= 5 {
			recallAt5++
		}

		if retrievalResult.Query.Language == testCase.Language {
			languageHits++
		} else {
			metrics.Failures = append(metrics.Failures, Failure{CaseName: testCase.Name, Reason: "language mismatch"})
		}

		if len(testCase.ExpectedDocumentIDs) == 0 || categoryMatches(documents, testCase.ExpectedCategory) {
			categoryHits++
		} else if testCase.ExpectedCategory != "" {
			metrics.Failures = append(metrics.Failures, Failure{CaseName: testCase.Name, Reason: "category mismatch"})
		}

		if len(testCase.ExpectedDocumentIDs) == 0 || temporalMatches(documents, testCase.ExpectedTemporalStatus) {
			temporalHits++
		} else if testCase.ExpectedTemporalStatus != "" {
			metrics.Failures = append(metrics.Failures, Failure{CaseName: testCase.Name, Reason: "temporal mismatch"})
		}

		response, hasResponse, language := app.AgentResponse(testCase.Question)
		responseCorrect := responseMatches(testCase, response, hasResponse, language)

		if len(testCase.ExpectedDocumentIDs) == 0 {
			noAnswerCases++
			if hasResponse {
				falsePositives++
				metrics.Failures = append(metrics.Failures, Failure{CaseName: testCase.Name, Reason: "false positive response"})
			}
		}

		if responseCorrect {
			correctResponses++
		} else {
			metrics.Failures = append(metrics.Failures, Failure{CaseName: testCase.Name, Reason: "response mismatch"})
		}
	}

	total := float64(metrics.Total)
	metrics.RecallAt1 = recallAt1 / total
	metrics.RecallAt3 = recallAt3 / total
	metrics.RecallAt5 = recallAt5 / total
	metrics.MRR = reciprocalRank / total
	metrics.LanguageAccuracy = float64(languageHits) / total
	metrics.CategoryAccuracy = float64(categoryHits) / total
	metrics.TemporalAccuracy = float64(temporalHits) / total
	metrics.CorrectResponseRate = float64(correctResponses) / total
	metrics.AverageRetrievalMillis = float64(totalRetrievalDuration.Microseconds()) / 1000 / total

	if noAnswerCases > 0 {
		metrics.FalsePositiveRate = float64(falsePositives) / float64(noAnswerCases)
	}

	sort.Slice(metrics.Failures, func(firstIndex int, secondIndex int) bool {
		if metrics.Failures[firstIndex].CaseName == metrics.Failures[secondIndex].CaseName {
			return metrics.Failures[firstIndex].Reason < metrics.Failures[secondIndex].Reason
		}

		return metrics.Failures[firstIndex].CaseName < metrics.Failures[secondIndex].CaseName
	})

	return metrics
}

func Print(metrics Metrics) {
	fmt.Printf("cases: %d\n", metrics.Total)
	fmt.Printf("recall@1: %.4f\n", metrics.RecallAt1)
	fmt.Printf("recall@3: %.4f\n", metrics.RecallAt3)
	fmt.Printf("recall@5: %.4f\n", metrics.RecallAt5)
	fmt.Printf("mrr: %.4f\n", metrics.MRR)
	fmt.Printf("language_accuracy: %.4f\n", metrics.LanguageAccuracy)
	fmt.Printf("category_accuracy: %.4f\n", metrics.CategoryAccuracy)
	fmt.Printf("temporal_accuracy: %.4f\n", metrics.TemporalAccuracy)
	fmt.Printf("correct_response_rate: %.4f\n", metrics.CorrectResponseRate)
	fmt.Printf("false_positive_rate: %.4f\n", metrics.FalsePositiveRate)
	fmt.Printf("average_retrieval_ms: %.4f\n", metrics.AverageRetrievalMillis)

	if len(metrics.Failures) == 0 {
		return
	}

	fmt.Println("failures:")
	for _, failure := range metrics.Failures {
		fmt.Printf("- %s: %s\n", failure.CaseName, failure.Reason)
	}
}

func rankedDocuments(results []domain.SearchResult) []rankedDocument {
	documents := make([]rankedDocument, 0, len(results))
	for _, result := range results {
		if result.Document == nil {
			continue
		}

		documents = append(documents, rankedDocument{
			ID:             result.Document.ID,
			Category:       result.Document.Category,
			Language:       result.Document.Language,
			TemporalStatus: result.Document.TemporalStatus,
		})
	}

	return documents
}

func rankedEvidence(evidence []domain.Evidence) []rankedDocument {
	documents := make([]rankedDocument, 0, len(evidence))
	for _, item := range evidence {
		documents = append(documents, rankedDocument{
			ID:             item.DocumentID,
			Category:       item.Category,
			Language:       item.Language,
			TemporalStatus: item.TemporalStatus,
		})
	}

	return documents
}

func evaluationService() (*agent.Service, error) {
	index, err := vectorindex.LoadEmbedded()
	if err != nil {
		return nil, err
	}

	embedder, err := embedding.NewDeterministicEmbedder(index.Dimension)
	if err != nil {
		return nil, err
	}

	return agent.NewService(
		embedder,
		retrieval.NewHybridRetriever(index),
		ranking.DefaultMetadataReranker(),
		generation.NewGroundedGenerator(),
	)
}

func firstExpectedRank(documents []rankedDocument, expectedIDs []string) int {
	if len(expectedIDs) == 0 {
		return 0
	}

	expected := make(map[string]struct{}, len(expectedIDs))
	for _, id := range expectedIDs {
		expected[id] = struct{}{}
	}

	for index, document := range documents {
		if _, ok := expected[document.ID]; ok {
			return index + 1
		}
	}

	return 0
}

func categoryMatches(documents []rankedDocument, expectedCategory string) bool {
	if expectedCategory == "" {
		return len(documents) == 0
	}

	if len(documents) == 0 {
		return false
	}

	for _, document := range documents {
		if document.Category == expectedCategory || evaluationCompatibleCategory(expectedCategory, document.Category) {
			return true
		}
	}

	return false
}

func temporalMatches(documents []rankedDocument, expectedTemporalStatus string) bool {
	if expectedTemporalStatus == "" {
		return len(documents) == 0
	}

	if expectedTemporalStatus == "timeless" && len(documents) == 0 {
		return true
	}

	if len(documents) == 0 {
		return false
	}

	return documents[0].TemporalStatus == expectedTemporalStatus
}

func evaluationCompatibleCategory(expectedCategory string, documentCategory string) bool {
	if expectedCategory == "project" {
		return documentCategory == "impact" || documentCategory == "technology" || documentCategory == "comparison"
	}

	return false
}

func responseMatches(testCase Case, response string, hasResponse bool, language string) bool {
	if len(testCase.ExpectedDocumentIDs) == 0 {
		return !hasResponse
	}

	if !hasResponse || language != testCase.Language {
		return false
	}

	normalizedResponse := strings.ToLower(response)

	for _, term := range testCase.RequiredTerms {
		if !strings.Contains(normalizedResponse, strings.ToLower(term)) {
			return false
		}
	}

	for _, term := range testCase.ForbiddenTerms {
		if strings.Contains(normalizedResponse, strings.ToLower(term)) {
			return false
		}
	}

	return true
}
