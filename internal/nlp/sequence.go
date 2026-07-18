package nlp

func containsTokenSequence(tokens []string, sequence []string) bool {
	if len(sequence) == 0 || len(sequence) > len(tokens) {
		return false
	}

	maximumStartIndex := len(tokens) - len(sequence)

	for startIndex := 0; startIndex <= maximumStartIndex; startIndex++ {
		matches := true

		for sequenceIndex, expectedToken := range sequence {
			if tokens[startIndex+sequenceIndex] != expectedToken {
				matches = false
				break
			}
		}

		if matches {
			return true
		}
	}

	return false
}
