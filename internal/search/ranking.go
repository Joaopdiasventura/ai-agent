package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"sort"
)

func FindTopDocuments(
	documents []domain.Document,
	engine *Engine,
	questionVector map[string]float64,
	analysis *nlp.QueryAnalysis,
	limit int,
) []Result {
	if len(documents) == 0 || len(questionVector) == 0 || limit <= 0 {
		return []Result{}
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
		return results[firstIndex].Similarity > results[secondIndex].Similarity
	})

	if limit > len(results) {
		limit = len(results)
	}

	return FilterRelevantResults(results[:limit], analysis, engine.MinimumSimilarity)
}
