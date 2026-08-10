package generation

import (
	"ai-agent/internal/domain"
)

type Answer struct {
	Text     string
	Language domain.Language
	FactIDs  []domain.FactID
}

func (a Answer) Empty() bool {
	return a.Text == ""
}

func appendFactID(
	values []domain.FactID,
	value domain.FactID,
) []domain.FactID {
	if value == "" {
		return values
	}

	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func appendFactIDs(
	values []domain.FactID,
	additions []domain.FactID,
) []domain.FactID {
	for _, value := range additions {
		values = appendFactID(
			values,
			value,
		)
	}

	return values
}
