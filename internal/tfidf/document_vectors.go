package tfidf

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/tokenizer"
)

func CalculateDocumentVectors(documents []domain.Document, idf map[string]float64) map[string]map[string]float64 {
	vectors := make(map[string]map[string]float64)

	for _, document := range documents {
		tokens := tokenizer.Tokenize(document.Content)
		vectors[document.ID] = CalculateTFIDF(tokens, idf)
	}

	return vectors
}
