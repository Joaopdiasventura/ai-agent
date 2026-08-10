package language

func EditSimilarity(left, right string) float64 {
	leftRunes := []rune(Normalize(left))
	rightRunes := []rune(Normalize(right))

	if len(leftRunes) == 0 && len(rightRunes) == 0 {
		return 1
	}

	if len(leftRunes) == 0 || len(rightRunes) == 0 {
		return 0
	}

	distance := levenshteinDistance(
		leftRunes,
		rightRunes,
	)

	maxLength := len(leftRunes)

	if len(rightRunes) > maxLength {
		maxLength = len(rightRunes)
	}

	return 1 - float64(distance)/float64(maxLength)

}

func FuzzySimilarity(left string, right string) float64 {
	ngram := CharacterNGramSimilarity(
		left,
		right,
		3,
		4,
	)

	edit := EditSimilarity(
		left,
		right,
	)

	if edit > ngram {
		return edit
	}

	return ngram
}

func levenshteinDistance(
	left []rune,
	right []rune,
) int {
	if len(left) == 0 {
		return len(right)
	}

	if len(right) == 0 {
		return len(left)
	}

	previous := make(
		[]int,
		len(right)+1,
	)

	current := make(
		[]int,
		len(right)+1,
	)

	for index := range previous {
		previous[index] = index
	}

	for leftIndex := 1; leftIndex <= len(left); leftIndex++ {
		current[0] = leftIndex

		for rightIndex := 1; rightIndex <= len(right); rightIndex++ {
			cost := 0

			if left[leftIndex-1] != right[rightIndex-1] {
				cost = 1
			}

			deletion :=
				previous[rightIndex] + 1

			insertion :=
				current[rightIndex-1] + 1

			substitution :=
				previous[rightIndex-1] + cost

			current[rightIndex] = min(
				deletion,
				insertion,
				substitution,
			)
		}

		previous, current =
			current, previous
	}

	return previous[len(right)]
}
