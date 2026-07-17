package tfidf

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/tokenizer"
	"math"
)

func CalculateIDF(documents []domain.Document) map[string]float64 {
	idf := make(map[string]float64)

	if len(documents) == 0 {
		return idf
	}

	documentFrequency := calculateDF(documents)
	totalDocuments := float64(len(documents))

	for term, frequency := range documentFrequency {
		idf[term] = math.Log(totalDocuments/float64(frequency)) + 1
	}

	return idf
}

func CalculateDocumentVectors(documents []domain.Document, idf map[string]float64) map[string]map[string]float64 {
	vectors := make(map[string]map[string]float64)

	for _, document := range documents {
		tokens := tokenizer.Tokenize(document.Content)
		vectors[document.ID] = CalculateTFIDF(tokens, idf)
	}

	return vectors
}

func CalculateTFIDF(tokens []string, idf map[string]float64) map[string]float64 {
	vector := make(map[string]float64)
	termFrequency := calculateTF(tokens)

	for term, frequency := range termFrequency {
		idfValue, exists := idf[term]

		if !exists {
			continue
		}

		vector[term] = frequency * idfValue
	}

	return vector
}

func calculateDF(documents []domain.Document) map[string]int {
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

func calculateTF(tokens []string) map[string]float64 {
	termFrequency := make(map[string]float64)

	if len(tokens) == 0 {
		return termFrequency
	}

	termCounts := countTerms(tokens)
	totalTokens := float64(len(tokens))

	for term, count := range termCounts {
		termFrequency[term] = float64(count) / totalTokens
	}

	return termFrequency
}

func countTerms(tokens []string) map[string]int {
	termsCounts := make(map[string]int)

	for _, token := range tokens {
		termsCounts[token]++
	}

	return termsCounts
}
