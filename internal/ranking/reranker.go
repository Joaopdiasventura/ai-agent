package ranking

import (
	"ai-agent/internal/domain"
	"sort"
	"strings"
)

type MetadataWeights struct {
	BaseScoreMultiplier         float64
	LanguageMatchBoost          float64
	LanguageMismatchPenalty     float64
	CategoryMatchBoost          float64
	CategoryMismatchPenalty     float64
	TemporalMatchBoost          float64
	TemporalMismatchPenalty     float64
	ProjectMatchBoost           float64
	ProjectMismatchPenalty      float64
	SpecificDocumentBoost       float64
	QueryTokenMatchBoost        float64
	GenericComparisonPenalty    float64
	BothRetrievalSourcesBoost   float64
	LexicalOnlyExactSignalBoost float64
	VectorOnlyWeakSignalPenalty float64
}

type MetadataReranker struct {
	Weights MetadataWeights
}

func DefaultMetadataReranker() MetadataReranker {
	return MetadataReranker{
		Weights: MetadataWeights{
			BaseScoreMultiplier:         100,
			LanguageMatchBoost:          40,
			LanguageMismatchPenalty:     500,
			CategoryMatchBoost:          80,
			CategoryMismatchPenalty:     250,
			TemporalMatchBoost:          80,
			TemporalMismatchPenalty:     300,
			ProjectMatchBoost:           24,
			ProjectMismatchPenalty:      140,
			SpecificDocumentBoost:       8,
			QueryTokenMatchBoost:        5,
			GenericComparisonPenalty:    25,
			BothRetrievalSourcesBoost:   12,
			LexicalOnlyExactSignalBoost: 6,
			VectorOnlyWeakSignalPenalty: 35,
		},
	}
}

func (reranker MetadataReranker) Rerank(query domain.Query, results []domain.SearchResult) []domain.SearchResult {
	reranked := make([]domain.SearchResult, 0, len(results))

	for _, result := range results {
		if result.Document == nil {
			continue
		}

		score := result.Score * reranker.Weights.BaseScoreMultiplier
		metadataScore := 0.0
		penalties := make([]string, 0)

		if query.Language != "" {
			if result.Document.Language == query.Language {
				metadataScore += reranker.Weights.LanguageMatchBoost
			} else {
				metadataScore -= reranker.Weights.LanguageMismatchPenalty
				penalties = append(penalties, "language_mismatch")
			}
		}

		if query.Category != "" {
			if result.Document.Category == query.Category {
				metadataScore += reranker.Weights.CategoryMatchBoost
			} else if !compatibleCategory(query.Category, result.Document.Category) {
				metadataScore -= reranker.Weights.CategoryMismatchPenalty
				penalties = append(penalties, "category_mismatch")
			}
		}

		if query.TemporalStatus != "" && query.TemporalStatus != domain.TemporalTimeless {
			if result.Document.TemporalStatus == query.TemporalStatus {
				metadataScore += reranker.Weights.TemporalMatchBoost
			} else if result.Document.TemporalStatus != domain.TemporalTimeless || query.Category == "education" || query.Category == "career" {
				metadataScore -= reranker.Weights.TemporalMismatchPenalty
				penalties = append(penalties, "temporal_mismatch")
			}
		}

		if query.Project != "" {
			if result.Document.Project == query.Project {
				metadataScore += reranker.Weights.ProjectMatchBoost
			} else if result.Document.Project != "" {
				metadataScore -= reranker.Weights.ProjectMismatchPenalty
				penalties = append(penalties, "project_mismatch")
			}
		}

		if result.Document.Subject != "" && result.Document.Subject != result.Document.Category {
			metadataScore += reranker.Weights.SpecificDocumentBoost
		}

		metadataScore += float64(queryTokenMatches(query, *result.Document)) * reranker.Weights.QueryTokenMatchBoost

		if result.Document.Category == "comparison" && query.Category == "project" {
			metadataScore -= reranker.Weights.GenericComparisonPenalty
			penalties = append(penalties, "generic_comparison")
		}

		if hasSource(result.Sources, "vector") && hasSource(result.Sources, "lexical") {
			metadataScore += reranker.Weights.BothRetrievalSourcesBoost
		}

		if result.VectorRank == 0 && result.LexicalRank > 0 && len(query.ExactTerms) > 0 {
			metadataScore += reranker.Weights.LexicalOnlyExactSignalBoost
		}

		if result.VectorRank > 0 && result.LexicalRank == 0 && query.Category == "" && len(query.ExactTerms) == 0 {
			metadataScore -= reranker.Weights.VectorOnlyWeakSignalPenalty
			penalties = append(penalties, "weak_vector_only")
		}

		result.MetadataScore = metadataScore
		result.FinalScore = score + metadataScore
		result.PenaltyReasons = append(result.PenaltyReasons, penalties...)
		reranked = append(reranked, result)
	}

	sort.Slice(reranked, func(firstIndex int, secondIndex int) bool {
		if reranked[firstIndex].FinalScore == reranked[secondIndex].FinalScore {
			return reranked[firstIndex].Document.ID < reranked[secondIndex].Document.ID
		}

		return reranked[firstIndex].FinalScore > reranked[secondIndex].FinalScore
	})

	return reranked
}

func hasSource(sources []string, source string) bool {
	for _, existing := range sources {
		if existing == source {
			return true
		}
	}

	return false
}

func compatibleCategory(queryCategory string, documentCategory string) bool {
	if queryCategory == "project" {
		return documentCategory == "impact" || documentCategory == "comparison" || documentCategory == "technology"
	}

	if queryCategory == "technology" {
		return documentCategory == "project"
	}

	return false
}

func queryTokenMatches(query domain.Query, document domain.Document) int {
	if len(query.Tokens) == 0 {
		return 0
	}

	documentTerms := make(map[string]struct{})
	for _, value := range append([]string{document.ID, document.Category, document.Subject, document.Project, document.Content}, document.Keywords...) {
		for _, token := range strings.Fields(normalize(value)) {
			documentTerms[strings.Trim(token, ".,;:!?()[]{}\"'")] = struct{}{}
		}
	}

	matches := 0
	for _, token := range query.Tokens {
		if len([]rune(token)) < 3 {
			continue
		}

		if _, exists := documentTerms[token]; exists {
			matches++
		}
	}

	return matches
}

func normalize(text string) string {
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "â", "a", "ã", "a", "ä", "a",
		"é", "e", "ê", "e", "è", "e", "ë", "e",
		"í", "i", "î", "i", "ì", "i", "ï", "i",
		"ó", "o", "ô", "o", "õ", "o", "ò", "o", "ö", "o",
		"ú", "u", "û", "u", "ù", "u", "ü", "u",
		"ç", "c",
	)

	return replacer.Replace(strings.ToLower(text))
}
