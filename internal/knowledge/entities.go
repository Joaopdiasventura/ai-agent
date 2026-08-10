package knowledge

import "ai-agent/internal/domain"

func allEntities() []domain.Entity {
	return []domain.Entity{
		{
			ID:   EntityJoao,
			Type: domain.EntityTypePerson,
			Name: localized(
				"João Paulo Dias Ventura",
				"João Paulo Dias Ventura",
			),
			Aliases: aliasMap(
				[]string{
					"joao",
					"joao paulo",
					"joão",
					"joão paulo",
				},
				[]string{
					"ele",
				},
				[]string{
					"he",
					"him",
					"his",
				},
			),
		},
		{
			ID:   EntitySaoPaulo,
			Type: domain.EntityTypeLocation,
			Name: localized(
				"São Paulo - SP",
				"São Paulo, Brazil",
			),
		},
		{
			ID:   EntityUFind,
			Type: domain.EntityTypeCompany,
			Name: localized(
				"uFind Tecnologia",
				"uFind Tecnologia",
			),
			Aliases: aliasMap(
				[]string{"ufind"},
				nil,
				nil,
			),
		},
		{
			ID:   EntityRepresentaOnline,
			Type: domain.EntityTypeCompany,
			Name: localized(
				"Representa Online",
				"Representa Online",
			),
			Aliases: aliasMap(
				[]string{"representa"},
				nil,
				nil,
			),
		},
		{
			ID:   EntityRoleMid,
			Type: domain.EntityTypeRole,
			Name: localized(
				"Desenvolvedor Pleno",
				"Mid-level Software Developer",
			),
		},
		{
			ID:   EntityRoleJunior,
			Type: domain.EntityTypeRole,
			Name: localized(
				"Desenvolvedor Júnior",
				"Junior Software Developer",
			),
		},
		{
			ID:   EntityRoleIntern,
			Type: domain.EntityTypeRole,
			Name: localized(
				"Desenvolvedor de Sistemas - Estágio",
				"Systems Developer Intern",
			),
		},
		projectEntity(
			EntityAuronix,
			"Auronix",
			[]string{"banco digital", "digital bank"},
		),
		projectEntity(
			EntityXTube,
			"X Tube",
			[]string{"x tube", "xtube"},
		),
		projectEntity(
			EntityGGCompress,
			"GGCompress",
			[]string{"gg compress", "ggcompress"},
		),
		projectEntity(
			EntityVox,
			"Vox",
			[]string{"plataforma eleitoral", "electoral platform"},
		),
		technologyEntity(EntityAngular, "Angular", nil),
		technologyEntity(EntityReact, "React", nil),
		technologyEntity(EntityNextJS, "Next.js", []string{"nextjs", "next js"}),
		technologyEntity(EntityJava, "Java", nil),
		technologyEntity(EntitySpringBoot, "Spring Boot", []string{"spring"}),
		technologyEntity(EntityGo, "Go", []string{"golang"}),
		technologyEntity(EntityNodeJS, "Node.js", []string{"node", "nodejs"}),
		technologyEntity(EntityJavaScript, "JavaScript", []string{"javascript", "java script"}),
		technologyEntity(EntityTypeScript, "TypeScript", []string{"typescript", "type script"}),
		technologyEntity(EntityNestJS, "NestJS", []string{"nest", "nest.js"}),
		technologyEntity(EntityPostgreSQL, "PostgreSQL", []string{"postgres"}),
		technologyEntity(EntityMongoDB, "MongoDB", []string{"mongo"}),
		technologyEntity(EntityRedis, "Redis", nil),
		technologyEntity(EntityRabbitMQ, "RabbitMQ", []string{"rabbit"}),
		technologyEntity(EntityKafka, "Kafka", []string{"apache kafka"}),
		technologyEntity(EntitySQS, "Amazon SQS", []string{"sqs"}),
		technologyEntity(EntityDocker, "Docker", nil),
		technologyEntity(EntityTerraform, "Terraform", nil),
		technologyEntity(EntityKubernetes, "Kubernetes", []string{"k8s"}),
		technologyEntity(EntityAWS, "AWS", []string{"amazon web services"}),
		technologyEntity(EntityECS, "Amazon ECS", []string{"ecs"}),
		technologyEntity(EntityEKS, "Amazon EKS", []string{"eks"}),
		technologyEntity(EntityS3, "Amazon S3", []string{"s3"}),
		technologyEntity(EntityIAM, "AWS IAM", []string{"iam"}),
		technologyEntity(EntityCloudflare, "Cloudflare", nil),
		technologyEntity(EntityFFmpeg, "FFmpeg", nil),
		technologyEntity(EntityPrometheus, "Prometheus", nil),
		technologyEntity(EntityTauri, "Tauri", nil),
		technologyEntity(EntityJWT, "JWT", nil),
		technologyEntity(EntityOAuth2, "OAuth2", []string{"oauth"}),
		technologyEntity(
			EntityWorkerThread,
			"Worker Threads",
			[]string{"worker thread", "worker threads"},
		),
		technologyEntity(
			EntityWebWorker,
			"Web Workers",
			[]string{"web worker", "web workers"},
		),
		technologyEntity(EntitySHA256, "SHA-256", []string{"sha256"}),
		technologyEntity(EntityBase44, "Base44", nil),
		technologyEntity(
			EntityVectorStores,
			"OpenAI Vector Stores",
			[]string{"vector stores"},
		),
		technologyEntity(
			EntityFileSearch,
			"OpenAI File Search",
			[]string{"file search"},
		),
		{
			ID:   EntityFIAP,
			Type: domain.EntityTypeInstitution,
			Name: localized("FIAP", "FIAP"),
		},
		{
			ID:   EntityEtec,
			Type: domain.EntityTypeInstitution,
			Name: localized(
				"Etec de Guarulhos",
				"Etec de Guarulhos",
			),
		},
		{
			ID:   EntityCertAWS,
			Type: domain.EntityTypeCertification,
			Name: localized(
				"AWS - Implantação de microsserviços no Amazon EKS e Modelagem de arquiteturas orientadas a eventos",
				"AWS - Microservices deployment on Amazon EKS and event-driven architecture modeling",
			),
		},
		{
			ID:   EntityCertMongoDB,
			Type: domain.EntityTypeCertification,
			Name: localized(
				"MongoDB - Modelagem de Dados e Conhecimento do setor de serviços financeiros",
				"MongoDB - Data Modeling and Financial Services Industry Knowledge",
			),
		},
		{
			ID:   EntityCertEDB,
			Type: domain.EntityTypeCertification,
			Name: localized(
				"EDB - Postgres Distribuído",
				"EDB - Distributed Postgres",
			),
		},
		{
			ID:   EntityPortuguese,
			Type: domain.EntityTypeLanguage,
			Name: localized(
				"Português",
				"Portuguese",
			),
		},
		{
			ID:   EntityEnglish,
			Type: domain.EntityTypeLanguage,
			Name: localized(
				"Inglês",
				"English",
			),
		},
	}
}

func projectEntity(
	id domain.EntityID,
	name string,
	aliases []string,
) domain.Entity {
	return domain.Entity{
		ID:   id,
		Type: domain.EntityTypeProject,
		Name: localized(name, name),
		Aliases: aliasMap(
			aliases,
			nil,
			nil,
		),
	}
}

func technologyEntity(
	id domain.EntityID,
	name string,
	aliases []string,
) domain.Entity {
	return domain.Entity{
		ID:   id,
		Type: domain.EntityTypeTechnology,
		Name: localized(name, name),
		Aliases: aliasMap(
			aliases,
			nil,
			nil,
		),
	}
}
