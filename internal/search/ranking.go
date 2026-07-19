package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"sort"
)

func FindTopDocuments(
	documents []*domain.Document,
	engine *Engine,
	questionVector map[string]float64,
	analysis *nlp.QueryAnalysis,
	limit int,
) []Result {
	if len(documents) == 0 || len(questionVector) == 0 || limit <= 0 {
		return []Result{}
	}

	literalResults := findLiteralDocumentMatches(documents, analysis)
	if len(literalResults) > 0 {
		if limit > len(literalResults) {
			limit = len(literalResults)
		}

		return literalResults[:limit]
	}

	results := make([]Result, 0, len(documents))

	for _, document := range documents {
		documentVectors, exists := engine.DocumentVectors[document.ID]

		if !exists {
			continue
		}

		similarity := CosineSimilarity(questionVector, documentVectors)
		boost := CalculateIntentBoost(analysis, document)

		results = append(results, Result{
			Document:   document,
			Similarity: similarity + boost,
		})
	}

	sort.Slice(results, func(firstIndex int, secondIndex int) bool {
		if results[firstIndex].Similarity == results[secondIndex].Similarity {
			return results[firstIndex].Document.ID < results[secondIndex].Document.ID
		}

		return results[firstIndex].Similarity > results[secondIndex].Similarity
	})

	results = FilterRelevantResults(results, analysis, engine.MinimumSimilarity)

	if limit > len(results) {
		limit = len(results)
	}

	return results[:limit]
}

func findLiteralDocumentMatches(documents []*domain.Document, analysis *nlp.QueryAnalysis) []Result {
	results := make([]Result, 0)

	for _, document := range documents {
		if !questionContainsDocumentContent(document, analysis) {
			continue
		}

		results = append(results, Result{
			Document:   document,
			Similarity: 1,
		})
	}

	sort.Slice(results, func(firstIndex int, secondIndex int) bool {
		return len(results[firstIndex].Document.Content) > len(results[secondIndex].Document.Content)
	})

	return results
}
