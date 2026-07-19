package retrieval

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/vectorindex"
	"errors"
	"sort"
)

var ErrDimensionMismatch = errors.New("query vector dimension does not match index")

func VectorSearch(index vectorindex.Index, queryVector []float32, limit int) ([]domain.SearchResult, error) {
	if limit <= 0 || len(index.Entries) == 0 {
		return []domain.SearchResult{}, nil
	}

	if len(queryVector) != index.Dimension {
		return nil, ErrDimensionMismatch
	}

	results := make([]domain.SearchResult, 0, len(index.Entries))

	for _, entry := range index.Entries {
		if len(entry.Embedding) != index.Dimension {
			return nil, ErrDimensionMismatch
		}

		document := entry.Document
		score := dotProduct(queryVector, entry.Embedding)
		results = append(results, domain.SearchResult{
			Document:   &document,
			Score:      score,
			VectorRank: 0,
			Sources:    []string{"vector"},
		})
	}

	sort.Slice(results, func(firstIndex int, secondIndex int) bool {
		if results[firstIndex].Score == results[secondIndex].Score {
			return results[firstIndex].Document.ID < results[secondIndex].Document.ID
		}

		return results[firstIndex].Score > results[secondIndex].Score
	})

	if limit > len(results) {
		limit = len(results)
	}

	results = results[:limit]
	for index := range results {
		results[index].VectorRank = index + 1
	}

	return results, nil
}

func dotProduct(first []float32, second []float32) float64 {
	var score float64

	for index := range first {
		score += float64(first[index] * second[index])
	}

	return score
}
