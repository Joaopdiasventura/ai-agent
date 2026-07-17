package answer

import "strings"

func RenderTemplate(template string, plan Plan) string {
	fact := ""

	if len(plan.Facts) > 0 {
		fact = plan.Facts[0]
	}

	technologies := formatList(plan.Technologies)

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

func formatList(items []string) string {
	switch len(items) {
	case 0:
		return "tecnologias não especificadas"
	case 1:
		return items[0]
	}

	lastIndex := len(items) - 1
	firstItems := strings.Join(items[:lastIndex], ", ")

	return firstItems + " e " + items[lastIndex]
}
