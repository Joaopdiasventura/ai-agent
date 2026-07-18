package answer

import (
	"ai-agent/internal/nlp"
	"strings"
)

func FormatFacts(facts []string, language nlp.Language) string {
	if len(facts) == 0 {
		return ""
	}

	formattedFacts := make([]string, 0, len(facts))

	for _, fact := range facts {
		fact = strings.TrimSpace(fact)
		fact = strings.TrimRight(fact, ".!?;")

		if fact != "" {
			formattedFacts = append(formattedFacts, fact)
		}
	}

	if len(formattedFacts) == 0 {
		return ""
	}

	if len(formattedFacts) == 1 {
		return formattedFacts[0] + "."
	}

	connectors := connectorsByLanguage(language)

	weights := []float64{
		0.40,
		0.25,
		0.20,
		0.15,
	}

	builder := strings.Builder{}
	builder.WriteString(formattedFacts[0])

	for _, fact := range formattedFacts[1:] {
		connector := SelectWeightedOption(connectors, weights)

		builder.WriteString(".")
		builder.WriteString(connector)
		builder.WriteString(strings.ToLower(fact[:1]))
		builder.WriteString(fact[1:])
	}

	builder.WriteString(".")

	return builder.String()
}

func connectorsByLanguage(language nlp.Language) []string {
	if language == nlp.LanguageEnglish {
		return []string{
			" Also, ",
			" Another relevant point is that ",
			" In addition, ",
			" On top of that, ",
		}
	}

	return []string{
		" Além disso, ",
		" Também ",
		" Outro ponto relevante é que ",
		" Somado a isso, ",
	}
}
