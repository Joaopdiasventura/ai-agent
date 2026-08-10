package index

import (
	"sort"

	"ai-agent/internal/domain"
)

func (i *Index) indexFactStructure(
	fact domain.Fact,
) {
	i.addSubjectFact(
		fact.Subject,
		fact.ID,
	)

	i.addEntityFact(
		fact.Subject,
		fact.ID,
	)

	if fact.Object.Kind ==
		domain.FactObjectEntity {
		i.addEntityFact(
			fact.Object.EntityID,
			fact.ID,
		)
	}

	for _, contextID := range fact.Context {
		i.addContextFact(
			contextID,
			fact.ID,
		)

		i.addEntityFact(
			contextID,
			fact.ID,
		)
	}

	for _, conceptID := range fact.Concepts {
		i.addConceptFact(
			conceptID,
			fact.ID,
		)
	}

	i.addCategoryFact(
		fact.Category,
		fact.ID,
	)
}

func (i *Index) addSubjectFact(
	entityID domain.EntityID,
	factID domain.FactID,
) {
	i.subjectFacts[entityID] =
		appendUniqueFactID(
			i.subjectFacts[entityID],
			factID,
		)
}

func (i *Index) addEntityFact(
	entityID domain.EntityID,
	factID domain.FactID,
) {
	i.entityFacts[entityID] =
		appendUniqueFactID(
			i.entityFacts[entityID],
			factID,
		)
}

func (i *Index) addConceptFact(
	conceptID domain.ConceptID,
	factID domain.FactID,
) {
	i.conceptFacts[conceptID] =
		appendUniqueFactID(
			i.conceptFacts[conceptID],
			factID,
		)
}

func (i *Index) addCategoryFact(
	category domain.FactCategory,
	factID domain.FactID,
) {
	i.categoryFacts[category] =
		appendUniqueFactID(
			i.categoryFacts[category],
			factID,
		)
}

func (i *Index) addContextFact(
	entityID domain.EntityID,
	factID domain.FactID,
) {
	i.contextFacts[entityID] =
		appendUniqueFactID(
			i.contextFacts[entityID],
			factID,
		)
}

func appendUniqueFactID(
	values []domain.FactID,
	value domain.FactID,
) []domain.FactID {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(values, value)
}

func sortStructuralIndexes(
	currentIndex *Index,
) {
	for _, values := range currentIndex.subjectFacts {
		sortFactIDs(values)
	}

	for _, values := range currentIndex.entityFacts {
		sortFactIDs(values)
	}

	for _, values := range currentIndex.conceptFacts {
		sortFactIDs(values)
	}

	for _, values := range currentIndex.categoryFacts {
		sortFactIDs(values)
	}

	for _, values := range currentIndex.contextFacts {
		sortFactIDs(values)
	}
}

func sortFactIDs(
	values []domain.FactID,
) {
	sort.Slice(
		values,
		func(left int, right int) bool {
			return values[left] <
				values[right]
		},
	)
}
