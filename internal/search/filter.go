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

		if shouldFilterByEntity(analysis) && !strings.Contains(
			strings.ToLower(result.Document.Content),
			strings.ToLower(analysis.Entity.Value)) {
			continue
		}

		relevantResults = append(relevantResults, result)
	}

	return relevantResults
}

func shouldFilterByEntity(analysis *nlp.QueryAnalysis) bool {
	if !analysis.HasEntity {
		return false
	}

	return analysis.Entity.Type != nlp.EntityPerson
}
