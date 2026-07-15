package search

func FilterRelevantResults(results []Result, minimumSimilarity float64) []Result {
	relevantResults := make([]Result, 0, len(results))

	for _, result := range results {
		if result.Similarity < minimumSimilarity {
			continue
		}

		relevantResults = append(relevantResults, result)
	}

	return relevantResults
}
