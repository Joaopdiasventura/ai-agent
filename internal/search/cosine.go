package search

import "math"

func CosineSimilarity(firstVector map[string]float64, secondVector map[string]float64) float64 {
	if len(firstVector) == 0 || len(secondVector) == 0 {
		return 0
	}

	dotProduct := 0.0
	firstMagnitude := 0.0
	secondMagnitude := 0.0

	for term, firstWeight := range firstVector {
		secondWeight := secondVector[term]
		dotProduct += firstWeight * secondWeight
		firstMagnitude += firstWeight * firstWeight
	}

	for _, secondWeight := range secondVector {
		secondMagnitude += secondWeight * secondWeight
	}

	if firstMagnitude == 0 || secondMagnitude == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(firstMagnitude) * math.Sqrt(secondMagnitude))
}
