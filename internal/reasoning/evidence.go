package reasoning

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ranking"
)

type EvidenceBuilder struct {
	base *knowledge.Knowledge
}

func NewEvidenceBuilder(
	base *knowledge.Knowledge,
) *EvidenceBuilder {
	return &EvidenceBuilder{
		base: base,
	}
}

func (b *EvidenceBuilder) Build(
	currentQuery domain.Query,
	ranked ranking.Result,
	limit int,
) []Evidence {
	if len(ranked.Candidates) == 0 {
		return nil
	}

	values := make(
		[]Evidence,
		0,
		len(ranked.Candidates),
	)

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, candidate := range ranked.Candidates {
		if _, exists :=
			seen[candidate.FactID]; exists {
			continue
		}

		fact, found :=
			b.base.Fact(
				candidate.FactID,
			)

		if !found {
			continue
		}

		seen[candidate.FactID] =
			struct{}{}

		directness :=
			evidenceDirectness(
				currentQuery,
				fact,
			)

		values = append(
			values,
			Evidence{
				FactID:     fact.ID,
				Score:      candidate.Score,
				Importance: fact.Importance,
				Directness: directness,
				MatchedEntities: copyEntities(
					candidate.MatchedEntities,
				),
				MatchedConcepts: copyConcepts(
					candidate.MatchedConcepts,
				),
			},
		)
	}

	sortEvidence(values)

	if limit > 0 &&
		len(values) > limit {
		values =
			values[:limit]

		for index := range values {
			values[index].Rank =
				index + 1
		}
	}

	return values
}

func evidenceDirectness(
	currentQuery domain.Query,
	fact domain.Fact,
) float64 {
	entityScore :=
		directEntityMatch(
			currentQuery,
			fact,
		)

	conceptScore :=
		directConceptMatch(
			currentQuery,
			fact,
		)

	switch {
	case entityScore > 0 &&
		conceptScore > 0:
		return clamp(
			0.55*entityScore +
				0.45*conceptScore,
		)

	case entityScore > 0:
		return clamp(
			0.85 * entityScore,
		)

	case conceptScore > 0:
		return clamp(
			0.9 * conceptScore,
		)

	default:
		return 0.25
	}
}

func directEntityMatch(
	currentQuery domain.Query,
	fact domain.Fact,
) float64 {
	maximum := 0.0

	for _, entity := range currentQuery.Entities {
		if !entity.Explicit {
			continue
		}

		if !factReferencesEntity(
			fact,
			entity.EntityID,
		) {
			continue
		}

		if entity.Score > maximum {
			maximum =
				entity.Score
		}
	}

	return clamp(maximum)
}

func directConceptMatch(
	currentQuery domain.Query,
	fact domain.Fact,
) float64 {
	maximum := 0.0

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		if !factHasConcept(
			fact,
			concept.ConceptID,
		) {
			continue
		}

		if concept.Score > maximum {
			maximum =
				concept.Score
		}
	}

	return clamp(maximum)
}

func factReferencesEntity(
	fact domain.Fact,
	entityID domain.EntityID,
) bool {
	if fact.Subject == entityID {
		return true
	}

	if fact.Object.Kind ==
		domain.FactObjectEntity &&
		fact.Object.EntityID ==
			entityID {
		return true
	}

	for _, contextID := range fact.Context {
		if contextID == entityID {
			return true
		}
	}

	return false
}

func factHasConcept(
	fact domain.Fact,
	conceptID domain.ConceptID,
) bool {
	for _, current := range fact.Concepts {
		if current == conceptID {
			return true
		}
	}

	return false
}

func copyEntities(
	values []domain.EntityID,
) []domain.EntityID {
	result := make(
		[]domain.EntityID,
		len(values),
	)

	copy(result, values)

	return result
}

func copyConcepts(
	values []domain.ConceptID,
) []domain.ConceptID {
	result := make(
		[]domain.ConceptID,
		len(values),
	)

	copy(result, values)

	return result
}
