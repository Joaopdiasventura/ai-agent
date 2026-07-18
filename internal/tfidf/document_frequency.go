package tfidf

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/tokenizer"
)

func calculateDF(documents []*domain.Document) map[string]int {
	documentFrequency := make(map[string]int)

	for _, document := range documents {
		tokens := tokenizer.Tokenize(document.Content)
		uniqueTerms := make(map[string]struct{})

		for _, token := range tokens {
			uniqueTerms[token] = struct{}{}
		}

		for term := range uniqueTerms {
			documentFrequency[term]++
		}
	}

	return documentFrequency
}
