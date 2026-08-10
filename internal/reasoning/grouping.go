package reasoning

import (
	"math"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

type Grouper struct {
	base *knowledge.Knowledge
}

func NewGrouper(
	base *knowledge.Knowledge,
) *Grouper {
	return &Grouper{
		base: base,
	}
}

func (g *Grouper) Group(
	currentQuery domain.Query,
	evidence []Evidence,
	entityType domain.EntityType,
) []EntityGroup {
	type groupState struct {
		evidence []Evidence
	}

	states := make(
		map[domain.EntityID]*groupState,
	)

	for _, currentEvidence :=
		range evidence {
		fact, found :=
			g.base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			continue
		}

		entityIDs :=
			g.factEntitiesOfType(
				fact,
				entityType,
			)

		for _, entityID :=
			range entityIDs {
			state, exists :=
				states[entityID]

			if !exists {
				state =
					&groupState{}

				states[entityID] =
					state
			}

			state.evidence =
				appendEvidenceUnique(
					state.evidence,
					currentEvidence,
				)
		}
	}

	groups := make(
		[]EntityGroup,
		0,
		len(states),
	)

	for entityID, state :=
		range states {
		sortEvidence(
			state.evidence,
		)

		strength :=
			evidenceStrength(
				state.evidence,
			)

		coverage :=
			groupConceptCoverage(
				currentQuery,
				state.evidence,
				g.base,
			)

		diversity :=
			groupDiversity(
				state.evidence,
				g.base,
			)

		quantity :=
			groupQuantity(
				len(state.evidence),
			)

		score :=
			clamp(
				0.6*strength +
					0.2*coverage +
					0.12*diversity +
					0.08*quantity,
			)

		groups = append(
			groups,
			EntityGroup{
				EntityID:
					entityID,
				Score:
					score,
				Evidence:
					state.evidence,
				EvidenceStrength:
					strength,
				ConceptCoverage:
					coverage,
				Diversity:
					diversity,
				Quantity:
					quantity,
			},
		)
	}

	sortGroups(groups)

	return groups
}

func (g *Grouper) factEntitiesOfType(
	fact domain.Fact,
	entityType domain.EntityType,
) []domain.EntityID {
	result := make(
		[]domain.EntityID,
		0,
	)

	add := func(
		entityID domain.EntityID,
	) {
		if entityID == "" {
			return
		}

		entity, found :=
			g.base.Entity(entityID)

		if !found ||
			entity.Type != entityType {
			return
		}

		for _, existing :=
			range result {
			if existing == entityID {
				return
			}
		}

		result = append(
			result,
			entityID,
		)
	}

	add(fact.Subject)

	if fact.Object.Kind ==
		domain.FactObjectEntity {
		add(
			fact.Object.EntityID,
		)
	}

	for _, contextID :=
		range fact.Context {
		add(contextID)
	}

	return result
}

func appendEvidenceUnique(
	values []Evidence,
	value Evidence,
) []Evidence {
	for _, existing :=
		range values {
		if existing.FactID ==
			value.FactID {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func evidenceStrength(
	evidence []Evidence,
) float64 {
	if len(evidence) == 0 {
		return 0
	}

	weights := []float64{
		1,
		0.68,
		0.48,
		0.34,
		0.24,
		0.17,
		0.12,
	}

	total := 0.0
	totalWeight := 0.0

	for index, current :=
		range evidence {
		var weight float64

		if index < len(weights) {
			weight = weights[index]
		} else {
			weight =
				0.08 /
					float64(
						index-len(weights)+2,
					)
		}

		effectiveScore :=
			clamp(
				0.7*current.Score +
					0.2*current.Directness +
					0.1*current.Importance,
			)

		total +=
			effectiveScore * weight

		totalWeight +=
			weight
	}

	if totalWeight == 0 {
		return 0
	}

	return clamp(
		total / totalWeight,
	)
}

func groupConceptCoverage(
	currentQuery domain.Query,
	evidence []Evidence,
	base *knowledge.Knowledge,
) float64 {
	queryConcepts :=
		relevantQueryConcepts(
			currentQuery,
		)

	if len(queryConcepts) == 0 {
		return 0.5
	}

	totalWeight := 0.0
	matchedWeight := 0.0

	for _, queryConcept :=
		range queryConcepts {
		weight :=
			clamp(
				queryConcept.Score,
			)

		totalWeight += weight

		found := false

		for _, currentEvidence :=
			range evidence {
			fact, exists :=
				base.Fact(
					currentEvidence.FactID,
				)

			if !exists {
				continue
			}

			if factHasConcept(
				fact,
				queryConcept.ConceptID,
			) {
				found = true
				break
			}
		}

		if found {
			matchedWeight +=
				weight
		}
	}

	if totalWeight == 0 {
		return 0.5
	}

	return clamp(
		matchedWeight /
			totalWeight,
	)
}

func relevantQueryConcepts(
	currentQuery domain.Query,
) []domain.ConceptMatch {
	direct := make(
		[]domain.ConceptMatch,
		0,
	)

	for _, concept :=
		range currentQuery.Concepts {
		if concept.MatchedText != "" {
			direct = append(
				direct,
				concept,
			)
		}
	}

	if len(direct) > 0 {
		return direct
	}

	result := make(
		[]domain.ConceptMatch,
		0,
	)

	for _, concept :=
		range currentQuery.Concepts {
		if concept.Score < 0.5 {
			continue
		}

		result = append(
			result,
			concept,
		)
	}

	return result
}

func groupDiversity(
	evidence []Evidence,
	base *knowledge.Knowledge,
) float64 {
	if len(evidence) == 0 {
		return 0
	}

	categories := make(
		map[domain.FactCategory]struct{},
	)

	predicates := make(
		map[domain.Relation]struct{},
	)

	for _, currentEvidence :=
		range evidence {
		fact, found :=
			base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			continue
		}

		categories[
			fact.Category,
		] = struct{}{}

		predicates[
			fact.Predicate,
		] = struct{}{}
	}

	categoryScore :=
		math.Min(
			float64(len(categories))/3,
			1,
		)

	predicateScore :=
		math.Min(
			float64(len(predicates))/4,
			1,
		)

	return clamp(
		0.45*categoryScore +
			0.55*predicateScore,
	)
}

func groupQuantity(
	count int,
) float64 {
	if count <= 0 {
		return 0
	}

	return clamp(
		1 -
			math.Exp(
				-float64(count)/3,
			),
	)
}