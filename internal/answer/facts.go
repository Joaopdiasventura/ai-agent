package answer

func SelectFactsByDetail(facts []string, detailLevel DetailLevel) []string {
	if len(facts) == 0 {
		return []string{}
	}

	limit := len(facts)

	switch detailLevel {
	case DetailShort:
		limit = 1
	case DetailMedium:
		limit = 2
	case DetailDetailed:
		limit = len(facts)
	default:
		limit = 1
	}

	if limit > len(facts) {
		limit = len(facts)
	}

	selectedFacts := make([]string, limit)
	copy(selectedFacts, facts[:limit])

	return selectedFacts

}
