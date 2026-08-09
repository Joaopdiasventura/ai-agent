package knowledge

import "ai-agent/internal/domain"

func certificationFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"certification-aws",
			EntityJoao,
			domain.RelationCertifiedIn,
			EntityCertAWS,
			domain.FactCategoryCertification,
			concepts(
				"aws",
				"kubernetes",
				"event-driven",
			),
			nil,
			localized(
				"O currículo de João lista certificação AWS sobre implantação de microsserviços no Amazon EKS e modelagem de arquiteturas orientadas a eventos.",
				"João's résumé lists AWS certification covering microservices deployment on Amazon EKS and event-driven architecture modeling.",
			),
			0.9,
			nil,
		),
		entityFact(
			"certification-mongodb",
			EntityJoao,
			domain.RelationCertifiedIn,
			EntityCertMongoDB,
			domain.FactCategoryCertification,
			concepts(
				"mongodb",
				"data-modeling",
				"financial-systems",
			),
			nil,
			localized(
				"O currículo de João lista certificação MongoDB em modelagem de dados e conhecimento do setor de serviços financeiros.",
				"João's résumé lists MongoDB certification in data modeling and financial services industry knowledge.",
			),
			0.85,
			nil,
		),
		entityFact(
			"certification-edb",
			EntityJoao,
			domain.RelationCertifiedIn,
			EntityCertEDB,
			domain.FactCategoryCertification,
			concepts(
				"postgresql",
				"distributed-systems",
				"database",
			),
			nil,
			localized(
				"O currículo de João lista certificação EDB em Postgres Distribuído.",
				"João's résumé lists an EDB certification in Distributed Postgres.",
			),
			0.85,
			nil,
		),
	}
}
