package query

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/ontology"
	"sort"
)

type ConceptExpander struct {
	currentOntology *ontology.Ontology
}

func NewconceptExpander(
	currentOntology *ontology.Ontology,
) *ConceptExpander {
	return &ConceptExpander{
		currentOntology: currentOntology,
	}
}

func (e *ConceptExpander) Expand(
	direct []domain.ConceptMatch,
	entities []domain.EntityMatch,
) []domain.ConceptMatch {
	result := make(map[domain.ConceptID]domain.ConceptMatch)

	directIDs := make(map[domain.ConceptID]struct{})

	for _, current := range direct {
		directIDs[current.ConceptID] = struct{}{}

		updateConceptMatch(result, current)
	}

	for _, current := range direct {
		expanded := e.currentOntology.Expand(current.ConceptID, 1)

		for _, candidate := range expanded {
			if candidate.ID == current.ConceptID {
				continue
			}

			score := current.Score * candidate.Weight * 0.72

			updateConceptMatch(
				result,
				domain.ConceptMatch{
					ConceptID: candidate.ID,
					Score: score,
				},
			)
		}
	}

	for _, entity := range entities {
		bindings := e.currentOntology.EntityConcepts(entity.EntityID)

		for _, binding := range bindings {
			score := entity.Score * binding.Weight * 0.65

			if _, direct := directIDs[binding.ID]; direct {
				score *= 0.9
			}

			updateConceptMatch(
				result,
				domain.ConceptMatch{
					ConceptID: binding.ID,
					Score: score,
				},
			)
		}
	}

	values := make([]domain.ConceptMatch, 0, len(result))

	for _, match := range result {
		values = append(values, match)
	}

	sort.Slice(
		values,
		func(i, j int) bool {
			if values[i].Score != values[j].Score {
				return values[i].Score > values[j].Score
			}

			return values[i].ConceptID < values[j].ConceptID
		},
	)

	return values
}
