package memory

import "strings"

func referencesLastEntity(question string) bool {
	normalizedQuestion := strings.ToLower(question)

	replacer := strings.NewReplacer(
		"?", " ",
		"!", " ",
		".", " ",
		",", " ",
		";", " ",
		":", " ",
	)

	normalizedQuestion = replacer.Replace(normalizedQuestion)
	normalizedQuestion = " " + strings.Join(strings.Fields(normalizedQuestion), " ") + " "

	references := []string{
		" ele ",
		" ela ",
		" dele ",
		" dela ",
		" isso ",
		" esse ",
		" essa ",
		" este ",
		" esta ",
	}

	for _, reference := range references {
		if strings.Contains(normalizedQuestion, reference) {
			return true
		}
	}

	return false
}
