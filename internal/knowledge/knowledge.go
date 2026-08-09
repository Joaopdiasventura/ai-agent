package knowledge

import (
	"fmt"
	"sort"

	"ai-agent/internal/domain"
)

type Knowledge struct {
	entities map[domain.EntityID]domain.Entity
	facts    map[domain.FactID]domain.Fact
}

func New() (*Knowledge, error) {
	entityList := allEntities()
	factList := allFacts()

	entities := make(map[domain.EntityID]domain.Entity, len(entityList))

	for _, entity := range entityList {
		if entity.ID == "" {
			return nil, fmt.Errorf("entity id is required")
		}

		if _, exists := entities[entity.ID]; exists {
			return nil, fmt.Errorf(
				"duplicated entity id: %s",
				entity.ID,
			)
		}

		entities[entity.ID] = entity
	}

	facts := make(map[domain.FactID]domain.Fact, len(factList))

	for _, fact := range factList {
		if err := fact.Validate(); err != nil {
			return nil, fmt.Errorf(
				"invalid fact %s: %w",
				fact.ID,
				err,
			)
		}

		if fact.Source == "" {
			return nil, fmt.Errorf(
				"fact %s has no source",
				fact.ID,
			)
		}

		if _, exists := facts[fact.ID]; exists {
			return nil, fmt.Errorf(
				"duplicated fact id: %s",
				fact.ID,
			)
		}

		if _, exists := entities[fact.Subject]; !exists {
			return nil, fmt.Errorf(
				"fact %s references unknown subject %s",
				fact.ID,
				fact.Subject,
			)
		}

		if fact.Object.Kind == domain.FactObjectEntity {
			if _, exists := entities[fact.Object.EntityID]; !exists {
				return nil, fmt.Errorf(
					"fact %s references unknown object %s",
					fact.ID,
					fact.Object.EntityID,
				)
			}
		}

		for _, contextID := range fact.Context {
			if _, exists := entities[contextID]; !exists {
				return nil, fmt.Errorf(
					"fact %s references unknown context %s",
					fact.ID,
					contextID,
				)
			}
		}

		for _, concept := range fact.Concepts {
			if concept == "" {
				return nil, fmt.Errorf(
					"fact %s contains empty concept",
					fact.ID,
				)
			}
		}

		if fact.Statement.PT == "" {
			return nil, fmt.Errorf(
				"fact %s has no portuguese statement",
				fact.ID,
			)
		}

		if fact.Statement.EN == "" {
			return nil, fmt.Errorf(
				"fact %s has no english statement",
				fact.ID,
			)
		}

		facts[fact.ID] = fact
	}

	return &Knowledge{
		entities: entities,
		facts:    facts,
	}, nil
}

func (k *Knowledge) Entity(id domain.EntityID) (domain.Entity, bool) {
	entity, found := k.entities[id]
	return entity, found
}

func (k *Knowledge) Fact(id domain.FactID) (domain.Fact, bool) {
	fact, found := k.facts[id]
	return fact, found
}

func (k *Knowledge) Entities() []domain.Entity {
	result := make([]domain.Entity, 0, len(k.entities))

	for _, entity := range k.entities {
		result = append(result, entity)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func (k *Knowledge) Facts() []domain.Fact {
	result := make([]domain.Fact, 0, len(k.facts))

	for _, fact := range k.facts {
		result = append(result, fact)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func (k *Knowledge) EntitiesByType(
	entityType domain.EntityType,
) []domain.Entity {
	result := make([]domain.Entity, 0)

	for _, entity := range k.entities {
		if entity.Type == entityType {
			result = append(result, entity)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func (k *Knowledge) FactsBySubject(
	subject domain.EntityID,
) []domain.Fact {
	result := make([]domain.Fact, 0)

	for _, fact := range k.facts {
		if fact.Subject == subject {
			result = append(result, fact)
		}
	}

	sortFacts(result)

	return result
}

func (k *Knowledge) FactsByCategory(
	category domain.FactCategory,
) []domain.Fact {
	result := make([]domain.Fact, 0)

	for _, fact := range k.facts {
		if fact.Category == category {
			result = append(result, fact)
		}
	}

	sortFacts(result)

	return result
}

func (k *Knowledge) FactsByConcept(
	concept domain.ConceptID,
) []domain.Fact {
	result := make([]domain.Fact, 0)

	for _, fact := range k.facts {
		for _, factConcept := range fact.Concepts {
			if factConcept == concept {
				result = append(result, fact)
				break
			}
		}
	}

	sortFacts(result)

	return result
}

func (k *Knowledge) FactsByContext(
	context domain.EntityID,
) []domain.Fact {
	result := make([]domain.Fact, 0)

	for _, fact := range k.facts {
		for _, contextID := range fact.Context {
			if contextID == context {
				result = append(result, fact)
				break
			}
		}
	}

	sortFacts(result)

	return result
}

func sortFacts(facts []domain.Fact) {
	sort.Slice(facts, func(i int, j int) bool {
		if facts[i].Importance != facts[j].Importance {
			return facts[i].Importance > facts[j].Importance
		}

		return facts[i].ID < facts[j].ID
	})
}
