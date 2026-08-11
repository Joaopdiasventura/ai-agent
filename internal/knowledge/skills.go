package knowledge

import "ai-agent/internal/domain"

func skillFacts() []domain.Fact {
	return []domain.Fact{
		skillFact(EntityAngular, "framework", "frontend"),
		skillFact(EntityReact, "framework", "frontend"),
		skillFact(EntityNextJS, "framework", "frontend"),
		skillFact(EntityJava, "programming-language", "backend"),
		skillFact(EntityJavaScript, "programming-language", "frontend"),
		skillFact(EntityTypeScript, "programming-language", "frontend"),
		skillFact(EntitySpringBoot, "framework", "backend"),
		skillFact(EntityGo, "programming-language", "backend"),
		skillFact(EntityNodeJS, "runtime", "backend"),
		skillFact(EntityNestJS, "framework", "backend"),
		skillFact(EntityPostgreSQL, "database"),
		skillFact(EntityMongoDB, "database"),
		skillFact(EntityRedis, "database"),
		skillFact(EntityRabbitMQ, "messaging"),
		skillFact(EntityKafka, "messaging"),
		skillFact(EntitySQS, "messaging"),
		skillFact(EntityDocker, "devops"),
		skillFact(EntityTerraform, "infrastructure-as-code"),
		skillFact(EntityKubernetes, "orchestration"),
		entityFact(
			"skill-aws",
			EntityJoao,
			domain.RelationHasSkill,
			EntityAWS,
			domain.FactCategorySkill,
			concepts("cloud", "aws"),
			nil,
			localized(
				"João possui experiência com AWS.",
				"João has experience with AWS.",
			),
			0.9,
			nil,
		),
	}
}

func skillFact(
	technology domain.EntityID,
	conceptsValue ...string,
) domain.Fact {
	return entityFact(
		"skill-"+string(technology),
		EntityJoao,
		domain.RelationHasSkill,
		technology,
		domain.FactCategorySkill,
		concepts(conceptsValue...),
		nil,
		localized(
			"João possui experiência com "+entityDisplayPT(technology)+".",
			"João has experience with "+entityDisplayEN(technology)+".",
		),
		0.85,
		nil,
	)
}

func entityDisplayPT(
	id domain.EntityID,
) string {
	names := map[domain.EntityID]string{
		EntityAngular:    "Angular",
		EntityReact:      "React",
		EntityNextJS:     "Next.js",
		EntityJava:       "Java",
		EntityJavaScript: "JavaScript",
		EntityTypeScript: "TypeScript",
		EntitySpringBoot: "Spring Boot",
		EntityGo:         "Go",
		EntityNodeJS:     "Node.js",
		EntityNestJS:     "NestJS",
		EntityPostgreSQL: "PostgreSQL",
		EntityMongoDB:    "MongoDB",
		EntityRedis:      "Redis",
		EntityRabbitMQ:   "RabbitMQ",
		EntityKafka:      "Kafka",
		EntitySQS:        "SQS",
		EntityDocker:     "Docker",
		EntityTerraform:  "Terraform",
		EntityKubernetes: "Kubernetes",
	}

	return names[id]
}

func entityDisplayEN(
	id domain.EntityID,
) string {
	return entityDisplayPT(id)
}
