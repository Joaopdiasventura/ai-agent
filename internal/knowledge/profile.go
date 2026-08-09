package knowledge

import "ai-agent/internal/domain"

func profileFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"profile-location",
			EntityJoao,
			domain.RelationLocatedIn,
			EntitySaoPaulo,
			domain.FactCategoryProfile,
			concepts("location"),
			nil,
			localized(
				"João Paulo Dias Ventura está localizado em São Paulo - SP.",
				"João Paulo Dias Ventura is located in São Paulo, Brazil.",
			),
			0.7,
			nil,
		),
		textFact(
			"profile-summary",
			EntityJoao,
			domain.RelationHasExperience,
			localized(
				"sistemas financeiros, processamento assíncrono, pipelines de dados e soluções de Inteligência Artificial",
				"financial systems, asynchronous processing, data pipelines and Artificial Intelligence solutions",
			),
			domain.FactCategoryProfile,
			concepts(
				"financial-systems",
				"async-processing",
				"data-pipelines",
				"artificial-intelligence",
			),
			nil,
			localized(
				"João possui experiência em sistemas financeiros, processamento assíncrono, pipelines de dados e soluções de Inteligência Artificial.",
				"João has experience with financial systems, asynchronous processing, data pipelines and Artificial Intelligence solutions.",
			),
			0.95,
			nil,
		),
		textFact(
			"profile-engineering-focus",
			EntityJoao,
			domain.RelationHasExperience,
			localized(
				"consistência transacional, confiabilidade, performance e eficiência operacional",
				"transactional consistency, reliability, performance and operational efficiency",
			),
			domain.FactCategoryProfile,
			concepts(
				"transactional-consistency",
				"reliability",
				"performance",
				"operational-efficiency",
			),
			nil,
			localized(
				"Seu trabalho tem foco em consistência transacional, confiabilidade, performance e eficiência operacional.",
				"His work focuses on transactional consistency, reliability, performance and operational efficiency.",
			),
			0.9,
			nil,
		),
		textFact(
			"profile-email",
			EntityJoao,
			domain.RelationIs,
			localized(
				"joaopdias.dev@gmail.com",
				"joaopdias.dev@gmail.com",
			),
			domain.FactCategoryContact,
			concepts("contact", "email"),
			nil,
			localized(
				"O e-mail de João é joaopdias.dev@gmail.com.",
				"João's email address is joaopdias.dev@gmail.com.",
			),
			1,
			nil,
		),
		textFact(
			"profile-phone",
			EntityJoao,
			domain.RelationIs,
			localized(
				"+55 (11) 93445-3236",
				"+55 (11) 93445-3236",
			),
			domain.FactCategoryContact,
			concepts("contact", "phone"),
			nil,
			localized(
				"O telefone informado no currículo de João é +55 (11) 93445-3236.",
				"The phone number listed on João's résumé is +55 (11) 93445-3236.",
			),
			1,
			nil,
		),
	}
}
