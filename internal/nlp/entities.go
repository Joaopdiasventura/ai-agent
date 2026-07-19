package nlp

type EntityType string

const (
	EntityProject     EntityType = "project"
	EntityCompany     EntityType = "company"
	EntityInstitution EntityType = "institution"
	EntityTechnology  EntityType = "technology"
	EntityPerson      EntityType = "person"
)

type Entity struct {
	Type  EntityType
	Value string
}

func DetectEntity(tokens []string) (Entity, bool) {
	if len(tokens) == 0 {
		return Entity{}, false
	}

	bestAliasLength := 0
	bestPriority := -1
	bestEntity := Entity{}

	for _, alias := range entityAliases {
		if !containsTokenSequence(tokens, alias.Tokens) {
			continue
		}

		priority := entityTypePriority(alias.Entity.Type)
		if priority > bestPriority ||
			(priority == bestPriority && len(alias.Tokens) > bestAliasLength) {
			bestPriority = priority
			bestAliasLength = len(alias.Tokens)
			bestEntity = alias.Entity
		}
	}

	if bestPriority < 0 {
		return Entity{}, false
	}

	return bestEntity, true
}

func entityTypePriority(entityType EntityType) int {
	switch entityType {
	case EntityProject:
		return 5
	case EntityCompany:
		return 4
	case EntityInstitution:
		return 3
	case EntityTechnology:
		return 2
	case EntityPerson:
		return 1
	default:
		return 0
	}
}
