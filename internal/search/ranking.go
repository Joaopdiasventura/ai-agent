package search

import (
	"ai-agent/internal/domain"
	"sort"
)

func FindTopDocuments(
	documents []domain.Document,
	documentVectors map[string]map[string]float64,
	questionVector map[string]float64,
	limit int,
) []Result {
	if len(documents) == 0 || len(questionVector) == 0 || limit <= 0 {
		return []Result{}
	}

	results := make([]Result, 0, len(documents))

	for _, document := range documents {
		documentVectors, exists := documentVectors[document.ID]

		if !exists {
			continue
		}

		similarity := CosineSimilarity(questionVector, documentVectors)

		results = append(results, Result{
			Document:   document,
			Similarity: similarity,
		})
	}

	sort.Slice(results, func(firstIndex int, secondIndex int) bool {
		return results[firstIndex].Similarity > results[secondIndex].Similarity
	})

	if limit > len(results) {
		limit = len(results)
	}

	return results[:limit]
}
