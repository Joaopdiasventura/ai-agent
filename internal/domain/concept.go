package domain

type ConceptID string

type Concept struct {
	ID        ConceptID
	Name      LocalizedText
	Aliases   map[Language][]string
	ParentIDs []ConceptID
}

func (c Concept) Label(language Language) string {
	return c.Name.For(language)
}

func (c Concept) AllAliases(language Language) []string {
	aliases := make([]string, 0)

	if label := c.Label(language); label != "" {
		aliases = append(aliases, label)
	}

	if languageAliases, ok := c.Aliases[language]; ok {
		aliases = append(aliases, languageAliases...)
	}

	if commonAliases, ok := c.Aliases[LanguageUnknown]; ok {
		aliases = append(aliases, commonAliases...)
	}

	return aliases
}

func (c Concept) IsZero() bool {
	return c.ID == ""
}
