package tfidf

func CountTerms(tokens []string) map[string]int {
	terms := make(map[string]int)

	for _, token := range tokens {
		terms[token]++
	}

	return terms
}

func CalculateTF(tokens []string) map[string]float64 {
	termFrequency := make(map[string]float64)

	if len(tokens) == 0 {
		return termFrequency
	}

	termCounts := CountTerms(tokens)
	totalTokens := float64(len(tokens))

	for term, count := range termCounts {
		termFrequency[term] = float64(count) / totalTokens
	}

	return termFrequency
}