package knowledge

import "ai-agent/internal/domain"

func languageFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"language-portuguese",
			EntityJoao,
			domain.RelationSpeaks,
			EntityPortuguese,
			domain.FactCategoryLanguage,
			concepts("language"),
			nil,
			localized(
				"João possui português fluente.",
				"João is fluent in Portuguese.",
			),
			1,
			nil,
		),
		textFact(
			"language-portuguese-level",
			EntityJoao,
			domain.RelationIs,
			localized(
				"Fluente",
				"Fluent",
			),
			domain.FactCategoryLanguage,
			concepts("language", "fluency"),
			[]domain.EntityID{EntityPortuguese},
			localized(
				"O nível de português informado no currículo é fluente.",
				"The Portuguese level listed on the résumé is fluent.",
			),
			1,
			nil,
		),
		entityFact(
			"language-english",
			EntityJoao,
			domain.RelationSpeaks,
			EntityEnglish,
			domain.FactCategoryLanguage,
			concepts("language"),
			nil,
			localized(
				"João possui inglês avançado.",
				"João has advanced English proficiency.",
			),
			1,
			nil,
		),
		textFact(
			"language-english-level",
			EntityJoao,
			domain.RelationIs,
			localized(
				"Avançado",
				"Advanced",
			),
			domain.FactCategoryLanguage,
			concepts("language"),
			[]domain.EntityID{EntityEnglish},
			localized(
				"O nível de inglês informado no currículo é avançado.",
				"The English level listed on the résumé is advanced.",
			),
			1,
			nil,
		),
	}
}
