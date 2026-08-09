package ontology

import (
	"fmt"
	"sort"

	"ai-agent/internal/domain"
	"ai-agent/internal/language"
)

type WeightedConcept struct {
	ID     domain.ConceptID
	Weight float64
}

type Alias struct {
	ConceptID  domain.ConceptID
	Language   domain.Language
	Normalized string
	Weight     float64
}

type relationLink struct {
	ConceptID domain.ConceptID
	Weight    float64
}

type Ontology struct {
	concepts  map[domain.ConceptID]domain.Concept
	aliases   map[domain.Language]map[string][]domain.ConceptID
	relations map[domain.ConceptID][]relationLink
	bindings  map[domain.EntityID][]WeightedConcept
}

func New() (*Ontology, error) {
	result := &Ontology{
		concepts:  make(map[domain.ConceptID]domain.Concept),
		aliases:   make(map[domain.Language]map[string][]domain.ConceptID),
		relations: make(map[domain.ConceptID][]relationLink),
		bindings:  make(map[domain.EntityID][]WeightedConcept),
	}

	result.aliases[domain.LanguagePortuguese] =
		make(map[string][]domain.ConceptID)

	result.aliases[domain.LanguageEnglish] =
		make(map[string][]domain.ConceptID)

	result.aliases[domain.LanguageUnknown] =
		make(map[string][]domain.ConceptID)

	for _, current := range allConcepts() {
		if current.ID == "" {
			return nil, fmt.Errorf(
				"concept id is required",
			)
		}

		if _, exists := result.concepts[current.ID]; exists {
			return nil, fmt.Errorf(
				"duplicated concept id: %s",
				current.ID,
			)
		}

		if current.Name.PT == "" || current.Name.EN == "" {
			return nil, fmt.Errorf(
				"concept %s requires names in both languages",
				current.ID,
			)
		}

		result.concepts[current.ID] = current
	}

	for _, current := range result.concepts {
		for _, parentID := range current.ParentIDs {
			if _, exists := result.concepts[parentID]; !exists {
				return nil, fmt.Errorf(
					"concept %s references unknown parent %s",
					current.ID,
					parentID,
				)
			}
		}

		result.indexAlias(
			current.ID,
			domain.LanguagePortuguese,
			current.Name.PT,
		)

		result.indexAlias(
			current.ID,
			domain.LanguageEnglish,
			current.Name.EN,
		)

		for aliasLanguage, aliases := range current.Aliases {
			for _, alias := range aliases {
				result.indexAlias(
					current.ID,
					aliasLanguage,
					alias,
				)
			}
		}
	}

	for _, relation := range allRelations() {
		if relation.Weight <= 0 || relation.Weight > 1 {
			return nil, fmt.Errorf(
				"invalid relation weight from %s to %s",
				relation.From,
				relation.To,
			)
		}

		if _, exists := result.concepts[relation.From]; !exists {
			return nil, fmt.Errorf(
				"relation references unknown concept %s",
				relation.From,
			)
		}

		if _, exists := result.concepts[relation.To]; !exists {
			return nil, fmt.Errorf(
				"relation references unknown concept %s",
				relation.To,
			)
		}

		result.addRelation(
			relation.From,
			relation.To,
			relation.Weight,
		)

		if relation.Bidirectional {
			result.addRelation(
				relation.To,
				relation.From,
				relation.Weight,
			)
		}
	}

	for _, binding := range allEntityBindings() {
		if binding.EntityID == "" {
			return nil, fmt.Errorf(
				"entity binding requires entity id",
			)
		}

		if _, exists := result.bindings[binding.EntityID]; exists {
			return nil, fmt.Errorf(
				"duplicated entity binding: %s",
				binding.EntityID,
			)
		}

		values := make(
			[]WeightedConcept,
			0,
			len(binding.Concepts),
		)

		seen := make(
			map[domain.ConceptID]struct{},
		)

		for _, current := range binding.Concepts {
			if _, exists := result.concepts[current.ID]; !exists {
				return nil, fmt.Errorf(
					"entity %s references unknown concept %s",
					binding.EntityID,
					current.ID,
				)
			}

			if current.Weight <= 0 || current.Weight > 1 {
				return nil, fmt.Errorf(
					"invalid binding weight for entity %s and concept %s",
					binding.EntityID,
					current.ID,
				)
			}

			if _, exists := seen[current.ID]; exists {
				return nil, fmt.Errorf(
					"entity %s has duplicated concept %s",
					binding.EntityID,
					current.ID,
				)
			}

			seen[current.ID] = struct{}{}
			values = append(values, current)
		}

		result.bindings[binding.EntityID] = values
	}

	return result, nil
}

func (o *Ontology) Concept(
	id domain.ConceptID,
) (domain.Concept, bool) {
	current, found := o.concepts[id]
	return current, found
}

func (o *Ontology) Concepts() []domain.Concept {
	result := make(
		[]domain.Concept,
		0,
		len(o.concepts),
	)

	for _, current := range o.concepts {
		result = append(result, current)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func (o *Ontology) ResolveAlias(
	value string,
	targetLanguage domain.Language,
) []domain.ConceptID {
	normalized := language.Normalize(value)

	if normalized == "" {
		return nil
	}

	seen := make(
		map[domain.ConceptID]struct{},
	)

	result := make(
		[]domain.ConceptID,
		0,
	)

	appendMatches := func(
		indexLanguage domain.Language,
	) {
		index, exists := o.aliases[indexLanguage]

		if !exists {
			return
		}

		for _, conceptID := range index[normalized] {
			if _, exists := seen[conceptID]; exists {
				continue
			}

			seen[conceptID] = struct{}{}
			result = append(result, conceptID)
		}
	}

	switch targetLanguage {
	case domain.LanguagePortuguese:
		appendMatches(
			domain.LanguagePortuguese,
		)

		appendMatches(
			domain.LanguageUnknown,
		)

		if len(result) == 0 {
			appendMatches(
				domain.LanguageEnglish,
			)
		}

	case domain.LanguageEnglish:
		appendMatches(
			domain.LanguageEnglish,
		)

		appendMatches(
			domain.LanguageUnknown,
		)

		if len(result) == 0 {
			appendMatches(
				domain.LanguagePortuguese,
			)
		}

	default:
		appendMatches(
			domain.LanguagePortuguese,
		)

		appendMatches(
			domain.LanguageEnglish,
		)

		appendMatches(
			domain.LanguageUnknown,
		)
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func (o *Ontology) Aliases(
	targetLanguage domain.Language,
) []Alias {
	type languageWeight struct {
		language domain.Language
		weight   float64
	}

	languages := []languageWeight{
		{
			language: domain.LanguagePortuguese,
			weight:   1,
		},
		{
			language: domain.LanguageEnglish,
			weight:   1,
		},
		{
			language: domain.LanguageUnknown,
			weight:   1,
		},
	}

	switch targetLanguage {
	case domain.LanguagePortuguese:
		languages = []languageWeight{
			{
				language: domain.LanguagePortuguese,
				weight:   1,
			},
			{
				language: domain.LanguageUnknown,
				weight:   0.95,
			},
			{
				language: domain.LanguageEnglish,
				weight:   0.85,
			},
		}

	case domain.LanguageEnglish:
		languages = []languageWeight{
			{
				language: domain.LanguageEnglish,
				weight:   1,
			},
			{
				language: domain.LanguageUnknown,
				weight:   0.95,
			},
			{
				language: domain.LanguagePortuguese,
				weight:   0.85,
			},
		}
	}

	selected := make(
		map[string]Alias,
	)

	for _, item := range languages {
		index := o.aliases[item.language]

		for normalized, conceptIDs := range index {
			for _, conceptID := range conceptIDs {
				key :=
					normalized +
						"\x00" +
						string(conceptID)

				current, exists := selected[key]

				if exists && current.Weight >= item.weight {
					continue
				}

				selected[key] = Alias{
					ConceptID:  conceptID,
					Language:   item.language,
					Normalized: normalized,
					Weight:     item.weight,
				}
			}
		}
	}

	result := make(
		[]Alias,
		0,
		len(selected),
	)

	for _, alias := range selected {
		result = append(result, alias)
	}

	sort.Slice(result, func(i int, j int) bool {
		leftLength :=
			len([]rune(result[i].Normalized))

		rightLength :=
			len([]rune(result[j].Normalized))

		if leftLength != rightLength {
			return leftLength > rightLength
		}

		if result[i].Weight != result[j].Weight {
			return result[i].Weight > result[j].Weight
		}

		if result[i].Normalized != result[j].Normalized {
			return result[i].Normalized <
				result[j].Normalized
		}

		return result[i].ConceptID <
			result[j].ConceptID
	})

	return result
}

func (o *Ontology) EntityConcepts(
	entityID domain.EntityID,
) []WeightedConcept {
	values := o.bindings[entityID]

	result := make(
		[]WeightedConcept,
		len(values),
	)

	copy(result, values)

	sort.Slice(result, func(i int, j int) bool {
		if result[i].Weight != result[j].Weight {
			return result[i].Weight >
				result[j].Weight
		}

		return result[i].ID < result[j].ID
	})

	return result
}

func (o *Ontology) Expand(
	id domain.ConceptID,
	maxDepth int,
) []WeightedConcept {
	if maxDepth < 0 {
		return nil
	}

	if _, exists := o.concepts[id]; !exists {
		return nil
	}

	type state struct {
		id     domain.ConceptID
		weight float64
		depth  int
	}

	best := map[domain.ConceptID]float64{
		id: 1,
	}

	queue := []state{
		{
			id:     id,
			weight: 1,
			depth:  0,
		},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.depth >= maxDepth {
			continue
		}

		currentConcept := o.concepts[current.id]

		for _, parentID := range currentConcept.ParentIDs {
			nextWeight :=
				current.weight * 0.85

			if nextWeight <= best[parentID] {
				continue
			}

			best[parentID] = nextWeight

			queue = append(
				queue,
				state{
					id:     parentID,
					weight: nextWeight,
					depth:  current.depth + 1,
				},
			)
		}

		for _, link := range o.relations[current.id] {
			nextWeight :=
				current.weight * link.Weight

			if nextWeight <= best[link.ConceptID] {
				continue
			}

			best[link.ConceptID] = nextWeight

			queue = append(
				queue,
				state{
					id:     link.ConceptID,
					weight: nextWeight,
					depth:  current.depth + 1,
				},
			)
		}
	}

	result := make(
		[]WeightedConcept,
		0,
		len(best),
	)

	for conceptID, weight := range best {
		result = append(
			result,
			WeightedConcept{
				ID:     conceptID,
				Weight: weight,
			},
		)
	}

	sort.Slice(result, func(i int, j int) bool {
		if result[i].Weight != result[j].Weight {
			return result[i].Weight >
				result[j].Weight
		}

		return result[i].ID < result[j].ID
	})

	return result
}

func (o *Ontology) ValidateFacts(
	facts []domain.Fact,
) error {
	for _, fact := range facts {
		for _, conceptID := range fact.Concepts {
			if _, exists := o.concepts[conceptID]; !exists {
				return fmt.Errorf(
					"fact %s references unknown concept %s",
					fact.ID,
					conceptID,
				)
			}
		}
	}

	return nil
}

func (o *Ontology) indexAlias(
	id domain.ConceptID,
	targetLanguage domain.Language,
	value string,
) {
	normalized := language.Normalize(value)

	if normalized == "" {
		return
	}

	if _, exists := o.aliases[targetLanguage]; !exists {
		o.aliases[targetLanguage] =
			make(map[string][]domain.ConceptID)
	}

	values :=
		o.aliases[targetLanguage][normalized]

	for _, existing := range values {
		if existing == id {
			return
		}
	}

	o.aliases[targetLanguage][normalized] =
		append(values, id)
}

func (o *Ontology) addRelation(
	from domain.ConceptID,
	to domain.ConceptID,
	weight float64,
) {
	o.relations[from] =
		append(
			o.relations[from],
			relationLink{
				ConceptID: to,
				Weight:    weight,
			},
		)
}
