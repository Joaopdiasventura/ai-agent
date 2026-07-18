package tfidf

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
