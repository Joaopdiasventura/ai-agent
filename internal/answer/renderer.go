package answer

import (
	"ai-agent/internal/nlp"
	"strings"
)

func RenderTemplate(template string, plan Plan) string {
	fact := ""

	if len(plan.Facts) > 0 {
		fact = plan.Facts[0]
	}

	technologies := formatList(plan.Technologies, plan.Language)

	formattedFacts := strings.TrimSpace(plan.FormattedFacts)

	if formattedFacts == "" {
		formattedFacts = fact
	}

	replacer := strings.NewReplacer(
		"{subject}", plan.Subject,
		"{fact}", fact,
		"{facts}", formattedFacts,
		"{technologies}", technologies,
	)

	return replacer.Replace(template)
}

func formatList(items []string, language nlp.Language) string {
	switch len(items) {
	case 0:
		if language == nlp.LanguageEnglish {
			return "technologies not specified"
		}

		return "tecnologias não especificadas"
	case 1:
		return items[0]
	}

	lastIndex := len(items) - 1
	firstItems := strings.Join(items[:lastIndex], ", ")

	if language == nlp.LanguageEnglish {
		return firstItems + " and " + items[lastIndex]
	}

	return firstItems + " e " + items[lastIndex]
}
