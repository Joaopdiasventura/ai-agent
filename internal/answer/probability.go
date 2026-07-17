package answer

import "math/rand/v2"

func SelectWeightedOption(options []string, weights []float64) string {
	if len(options) == 0 || len(options) != len(weights) {
		return ""
	}

	totalWeight := 0.0

	for _, weight := range weights {
		if weight > 0 {
			totalWeight += weight
		}
	}

	if totalWeight == 0 {
		return options[0]
	}

	target := rand.Float64() * totalWeight
	accumelatedWeight := 0.0

	for index, weight := range weights {
		if weight <= 0 {
			continue
		}

		accumelatedWeight += weight

		if target < accumelatedWeight {
			return options[index]
		}
	}

	return options[len(options) - 1]
}
