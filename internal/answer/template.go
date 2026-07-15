package answer

func SelectTemplate(templates []string, lastTemplateIndex int) (string, int) {
	if len(templates) == 0 {
		return "", -1
	}

	if len(templates) == 1 {
		return templates[0], 0
	}

	nextTemplateIndex := lastTemplateIndex + 1

	if nextTemplateIndex >= len(templates) {
		nextTemplateIndex = 0
	}

	return templates[nextTemplateIndex], nextTemplateIndex
}
