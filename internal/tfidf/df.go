package tfidf

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/tokenizer"
)

func CalculateDF(documents []domain.Document) map[string]int {
	documentFrequency := make(map[string]int)

	for _, document := range documents {
		tokens := tokenizer.Tokenize(document.Content)
		uniqueTerms := make(map[string]bool)

		for _, token := range tokens {
			uniqueTerms[token] = true
		}

		for term := range uniqueTerms {
			documentFrequency[term]++
		}
	}

	return documentFrequency
}
