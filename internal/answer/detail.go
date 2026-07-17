package answer

type DetailLevel string

const (
	DetailShort    DetailLevel = "short"
	DetailMedium   DetailLevel = "medium"
	DetailDetailed DetailLevel = "detailed"
)

func SelectDetailLevel() DetailLevel {
	options := []string{
		string(DetailShort),
		string(DetailMedium),
		string(DetailDetailed),
	}

	weights := []float64{
		0.25,
		0.50,
		0.25,
	}

	selected := SelectWeightedOption(options, weights)

	return DetailLevel(selected)
}
