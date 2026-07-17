package answer

import "ai-agent/internal/nlp"

type weightedOpenings struct {
	Options []string
	Weights []float64
}

var openingsByIntent = map[nlp.Intent]weightedOpenings{
	nlp.IntentCurrentJob: {
		Options: []string{
			"Atualmente,",
			"Hoje,",
			"No momento,",
		},
		Weights: []float64{
			0.50,
			0.30,
			0.20,
		},
	},
	nlp.IntentFirstJob: {
		Options: []string{
			"Inicialmente,",
			"No começo da carreira,",
			"Em sua primeira experiência profissional,",
		},
		Weights: []float64{
			0.30,
			0.35,
			0.35,
		},
	},
	nlp.IntentEducation: {
		Options: []string{
			"Atualmente,",
			"Na área acadêmica,",
			"Em relação à formação,",
		},
		Weights: []float64{
			0.40,
			0.30,
			0.30,
		},
	},
	nlp.IntentProject: {
		Options: []string{
			"Em resumo,",
			"De forma direta,",
			"Sobre o projeto,",
			"",
		},
		Weights: []float64{
			0.25,
			0.20,
			0.20,
			0.35,
		},
	},
	nlp.IntentTechnologies: {
		Options: []string{
			"Em relação à stack,",
			"Entre as tecnologias utilizadas,",
			"Na parte técnica,",
			"",
		},
		Weights: []float64{
			0.25,
			0.25,
			0.20,
			0.30,
		},
	},
}

func SelectOpening(intent nlp.Intent) string {
	openings, exists := openingsByIntent[intent]

	if !exists {
		return ""
	}

	return SelectWeightedOption(
		openings.Options,
		openings.Weights,
	)
}
