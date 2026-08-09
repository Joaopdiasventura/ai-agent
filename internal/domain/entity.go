package domain

type EntityID string

type EntityType string

const (
	EntityTypeUnknown       EntityType = "unknown"
	EntityTypePerson        EntityType = "person"
	EntityTypeCompany       EntityType = "company"
	EntityTypeProject       EntityType = "project"
	EntityTypeTechnology    EntityType = "technology"
	EntityTypeInstitution   EntityType = "institution"
	EntityTypeCertification EntityType = "certification"
	EntityTypeRole          EntityType = "role"
	EntityTypeLocation      EntityType = "location"
)

type Entity struct {
	ID      EntityID
	Type    EntityType
	Name    LocalizedText
	Aliases map[Language][]string
}

func (e Entity) Label(language Language) string {
	return e.Name.For(language)
}

func (e Entity) AllAliases(language Language) []string {
	aliases := make([]string, 0)

	if label := e.Label(language); label != "" {
		aliases = append(aliases, label)
	}

	if languageAliases, ok := e.Aliases[language]; ok {
		aliases = append(aliases, languageAliases...)
	}

	if commonAliases, ok := e.Aliases[LanguageUnknown]; ok {
		aliases = append(aliases, commonAliases...)
	}

	return aliases
}

func (e Entity) IsZero() bool {
	return e.ID == ""
}
