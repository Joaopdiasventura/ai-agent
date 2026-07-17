package answer

import "ai-agent/internal/nlp"

var verbsByIntent = map[nlp.Intent]weightedOpenings{
	nlp.IntentCurrentJob: {
		Options: []string{
			"trabalha atualmente como",
			"atua como",
			"ocupa o cargo de",
		},
		Weights: []float64{
			0.45,
			0.35,
			0.20,
		},
	},
	nlp.IntentFirstJob: {
		Options: []string{
			"começou sua carreira como",
			"iniciou sua trajetória profissional como",
			"teve sua primeira experiência como",
		},
		Weights: []float64{
			0.40,
			0.30,
			0.30,
		},
	},
	nlp.IntentEducation: {
		Options: []string{
			"cursa",
			"estuda",
			"está se formando em",
		},
		Weights: []float64{
			0.45,
			0.35,
			0.20,
		},
	},
	nlp.IntentProject: {
		Options: []string{
			"é",
			"consiste em",
			"funciona como",
		},
		Weights: []float64{
			0.50,
			0.25,
			0.25,
		},
	},
	nlp.IntentTechnologies: {
		Options: []string{
			"utiliza principalmente",
			"foi construído com",
			"tem uma stack composta por",
		},
		Weights: []float64{
			0.40,
			0.35,
			0.25,
		},
	},
}

func SelectVerb(intent nlp.Intent) string {
	verbs, exists := verbsByIntent[intent]

	if !exists {
		return ""
	}

	return SelectWeightedOption(
		verbs.Options,
		verbs.Weights,
	)
}
