package answer

import "strings"

func normalizeFactForSubject(subject string, fact string) string {
	subject = strings.TrimSpace(subject)
	fact = strings.TrimSpace(fact)

	if subject == "" || fact == "" {
		return fact
	}

	prefixes := []string{
		subject + " é ",
		subject + " foi ",
		subject + " utiliza ",
		subject + " usa ",
		subject + " trabalha ",
		subject + " atua ",
	}

	lowerFact := strings.ToLower(fact)

	for _, prefix := range prefixes {
		lowerPrefix := strings.ToLower(prefix)

		if !strings.HasPrefix(lowerFact, lowerPrefix) {
			continue
		}

		remainingFact := strings.TrimSpace(fact[len(prefix):])

		if remainingFact == "" {
			return fact
		}

		return strings.ToLower(remainingFact[:1] + remainingFact[1:])
	}

	return fact
}

func RenderTemplate(template string, plan Plan) string {
	fact := ""

	if len(plan.Facts) > 0 {
		fact = plan.Facts[0]
	}

	technologies := formatList(plan.Technologies)

	replacer := strings.NewReplacer(
		"{subject}", plan.Subject,
		"{fact}", fact,
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
