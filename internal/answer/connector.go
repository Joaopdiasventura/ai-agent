package answer

import "ai-agent/internal/nlp"

var connectorsByIntent = map[nlp.Intent]weightedOpenings{
	nlp.IntentCurrentJob: {
		Options: []string{
			"na",
			"dentro da",
			"pela",
		},
		Weights: []float64{
			0.60,
			0.20,
			0.20,
		},
	},
	nlp.IntentFirstJob: {
		Options: []string{
			"na",
			"dentro da",
			"pela",
		},
		Weights: []float64{
			0.60,
			0.20,
			0.20,
		},
	},
	nlp.IntentEducation: {
		Options: []string{
			"na",
			"pela",
			"junto à",
		},
		Weights: []float64{
			0.60,
			0.20,
			0.20,
		},
	},
	nlp.IntentProject: {
		Options: []string{
			"com foco em",
			"voltado para",
			"construído para",
		},
		Weights: []float64{
			0.40,
			0.30,
			0.30,
		},
	},
	nlp.IntentTechnologies: {
		Options: []string{
			"incluindo",
			"como",
			"entre elas",
		},
		Weights: []float64{
			0.40,
			0.35,
			0.25,
		},
	},
}

func SelectConnector(intent nlp.Intent) string {
	connectors, exists := connectorsByIntent[intent]

	if !exists {
		return ""
	}

	return SelectWeightedOption(
		connectors.Options,
		connectors.Weights,
	)
}
