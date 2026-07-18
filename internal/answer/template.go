package answer

import "math/rand/v2"

func SelectTemplate(templates []string) string {
	if len(templates) == 0 {
		return ""
	}

	if len(templates) == 1 {
		return templates[0]
	}

	randomIndex := rand.IntN(len(templates))

	return templates[randomIndex]
}
