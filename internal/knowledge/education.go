package knowledge

import "ai-agent/internal/domain"

func educationFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"education-fiap-institution",
			EntityJoao,
			domain.RelationStudiedAt,
			EntityFIAP,
			domain.FactCategoryEducation,
			concepts(
				"education",
				"artificial-intelligence",
			),
			[]domain.EntityID{EntityFIAP},
			localized(
				"João possui formação em Inteligência Artificial na FIAP listada entre agosto de 2026 e junho de 2028.",
				"João has an Artificial Intelligence education entry at FIAP listed from August 2026 to June 2028.",
			),
			0.95,
			closedPeriod(2026, 8, 2028, 6),
		),
		textFact(
			"education-fiap-course",
			EntityJoao,
			domain.RelationHasExperience,
			localized(
				"Inteligência Artificial",
				"Artificial Intelligence",
			),
			domain.FactCategoryEducation,
			concepts(
				"education",
				"artificial-intelligence",
			),
			[]domain.EntityID{EntityFIAP},
			localized(
				"O curso listado para João na FIAP é Inteligência Artificial.",
				"The course listed for João at FIAP is Artificial Intelligence.",
			),
			0.95,
			closedPeriod(2026, 8, 2028, 6),
		),
		entityFact(
			"education-etec-institution",
			EntityJoao,
			domain.RelationStudiedAt,
			EntityEtec,
			domain.FactCategoryEducation,
			concepts(
				"education",
				"software-development",
			),
			[]domain.EntityID{EntityEtec},
			localized(
				"João estudou Desenvolvimento de Sistemas na Etec de Guarulhos entre fevereiro de 2023 e dezembro de 2025.",
				"João studied Systems Development at Etec de Guarulhos from February 2023 to December 2025.",
			),
			0.95,
			closedPeriod(2023, 2, 2025, 12),
		),
		textFact(
			"education-etec-course",
			EntityJoao,
			domain.RelationHasExperience,
			localized(
				"Desenvolvimento de Sistemas",
				"Systems Development",
			),
			domain.FactCategoryEducation,
			concepts(
				"education",
				"software-development",
			),
			[]domain.EntityID{EntityEtec},
			localized(
				"Na Etec de Guarulhos, João cursou Desenvolvimento de Sistemas.",
				"At Etec de Guarulhos, João studied Systems Development.",
			),
			0.9,
			closedPeriod(2023, 2, 2025, 12),
		),
	}
}
