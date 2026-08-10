package planning

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/reasoning"
)

type EvidenceSelector struct {
	base *knowledge.Knowledge
}

func NewEvidenceSelector(
	base *knowledge.Knowledge,
) *EvidenceSelector {
	return &EvidenceSelector{
		base: base,
	}
}

func (s *EvidenceSelector) Select(
	evidence []reasoning.Evidence,
	limit int,
) []reasoning.Evidence {
	if limit <= 0 ||
		len(evidence) == 0 {
		return nil
	}

	if limit > len(evidence) {
		limit = len(evidence)
	}

	selected := make(
		[]reasoning.Evidence,
		0,
		limit,
	)

	selectedIDs := make(
		map[domain.FactID]struct{},
	)

	categories := make(
		map[domain.FactCategory]int,
	)

	predicates := make(
		map[domain.Relation]int,
	)

	for len(selected) < limit {
		bestIndex := -1
		bestScore := -1.0

		for index, current := range evidence {
			if _, exists :=
				selectedIDs[current.FactID]; exists {
				continue
			}

			score :=
				s.selectionScore(
					current,
					categories,
					predicates,
				)

			if score > bestScore {
				bestIndex = index
				bestScore = score
				continue
			}

			if score == bestScore &&
				bestIndex >= 0 &&
				current.FactID <
					evidence[bestIndex].FactID {
				bestIndex = index
			}
		}

		if bestIndex < 0 {
			break
		}

		current :=
			evidence[bestIndex]

		selected = append(
			selected,
			current,
		)

		selectedIDs[current.FactID] =
			struct{}{}

		fact, found :=
			s.base.Fact(
				current.FactID,
			)

		if found {
			categories[fact.Category]++

			predicates[fact.Predicate]++
		}
	}

	return selected
}

func (s *EvidenceSelector) selectionScore(
	current reasoning.Evidence,
	categories map[domain.FactCategory]int,
	predicates map[domain.Relation]int,
) float64 {
	score :=
		0.62*clamp(current.Score) +
			0.23*clamp(current.Directness) +
			0.15*clamp(current.Importance)

	fact, found :=
		s.base.Fact(
			current.FactID,
		)

	if !found {
		return score
	}

	if categories[fact.Category] == 0 {
		score += 0.08
	}

	if predicates[fact.Predicate] == 0 {
		score += 0.08
	}

	if categories[fact.Category] >= 2 {
		score -= 0.04
	}

	if predicates[fact.Predicate] >= 2 {
		score -= 0.06
	}

	return score
}

func evidenceIDs(
	evidence []reasoning.Evidence,
) []domain.FactID {
	result := make(
		[]domain.FactID,
		0,
		len(evidence),
	)

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, current := range evidence {
		if current.FactID == "" {
			continue
		}

		if _, exists :=
			seen[current.FactID]; exists {
			continue
		}

		seen[current.FactID] =
			struct{}{}

		result = append(
			result,
			current.FactID,
		)
	}

	return result
}
