package tfidf

import (
	"ai-agent/internal/domain"
	"math"
)

func CalculateIDF(documents []domain.Document) map[string]float64 {
	inverseDocumentFrequency := make(map[string]float64)

	if len(documents) == 0 {
		return inverseDocumentFrequency
	}

	documentFrequency := CalculateDF(documents)
	totalDocuments := float64(len(documents))

	for term, frequency := range documentFrequency {
		inverseDocumentFrequency[term] = math.Log(
			totalDocuments/float64(frequency),
		) + 1
	}

	return inverseDocumentFrequency
}
