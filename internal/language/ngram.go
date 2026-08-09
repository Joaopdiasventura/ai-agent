package language

func CharacterNGrams(
	value string,
	minSize int,
	maxSize int,
) []string {
	if minSize <= 0 || maxSize < minSize {
		return nil
	}

	tokens := Tokenize(value)

	if len(tokens) == 0 {
		return nil
	}

	result := make([]string, 0)
	seen := make(map[string]struct{})

	for _, token := range tokens {
		tokenRunes := []rune(token)

		if len(tokenRunes) == 0 {
			continue
		}

		if len(tokenRunes) < minSize {
			addUniqueNGram(&result, seen, string(tokenRunes))
			continue
		}

		limit := min(maxSize, len(tokenRunes))

		for size := minSize; size <= limit; size++ {
			for start := 0; start+size <= len(tokenRunes); start++ {
				gram := string(
					tokenRunes[start : start+size],
				)

				addUniqueNGram(&result, seen, gram)
			}
		}
	}

	return result
}

func CharacterNGramSimilarity(
	left string,
	right string,
	minSize int,
	maxSize int,
) float64 {
	leftGrams := CharacterNGrams(
		left,
		minSize,
		maxSize,
	)

	rightGrams := CharacterNGrams(
		right,
		minSize,
		maxSize,
	)

	if len(leftGrams) == 0 || len(rightGrams) == 0 {
		return 0
	}

	leftSet := make(
		map[string]struct{},
		len(leftGrams),
	)

	for _, gram := range leftGrams {
		leftSet[gram] = struct{}{}
	}

	rightSet := make(
		map[string]struct{},
		len(rightGrams),
	)

	for _, gram := range rightGrams {
		rightSet[gram] = struct{}{}
	}

	intersection := 0

	for gram := range leftSet {
		if _, exists := rightSet[gram]; exists {
			intersection++
		}
	}

	union := len(leftSet) + len(rightSet) - intersection

	if union == 0 {
		return 0
	}

	return float64(intersection) / float64(union)
}

func addUniqueNGram(
	result *[]string,
	seen map[string]struct{},
	gram string,
) {
	if gram == "" {
		return
	}

	if _, exists := seen[gram]; exists {
		return
	}

	seen[gram] = struct{}{}
	*result = append(*result, gram)
}
