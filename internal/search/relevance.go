package search

func IsRelevant(similarity float64, minimumSimilarity float64) bool {
	return similarity >= minimumSimilarity
}