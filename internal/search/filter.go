package search

import (
	"ai-agent/internal/nlp"
	"strings"
)

func FilterRelevantResults(results []Result, analysis *nlp.QueryAnalysis, minimumSimilarity float64) []Result {
	relevantResults := make([]Result, 0, len(results))

	for _, result := range results {
		if result.Similarity < minimumSimilarity {
			continue
		}

		if analysis.HasEntity &&
			strings.Contains(
				result.Document.Content,
				analysis.Entity.Value) {
			relevantResults = append(relevantResults, result)
		}
	}

	return relevantResults
}
