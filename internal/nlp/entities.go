package nlp

type EntityType string

const (
	EntityProject     EntityType = "project"
	EntityCompany     EntityType = "company"
	EntityInstitution EntityType = "institution"
	EntityTechnology  EntityType = "technology"
)

type Entity struct {
	Type  EntityType
	Value string
}

func DetectEntity(tokens []string) (Entity, bool) {
	if len(tokens) == 0 {
		return Entity{}, false
	}

	for _, alias := range entityAliases {
		if containsTokenSequence(tokens, alias.Tokens) {
			return alias.Entity, true
		}
	}

	return Entity{}, false
}
