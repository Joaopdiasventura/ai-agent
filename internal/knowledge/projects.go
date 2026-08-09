package knowledge

import "ai-agent/internal/domain"

func projectFacts() []domain.Fact {
	facts := make([]domain.Fact, 0)

	facts = append(facts, auronixFacts()...)
	facts = append(facts, xTubeFacts()...)
	facts = append(facts, ggCompressFacts()...)
	facts = append(facts, voxFacts()...)

	return facts
}

func auronixFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"project-auronix-joao",
			EntityJoao,
			domain.RelationWorkedOn,
			EntityAuronix,
			domain.FactCategoryProject,
			concepts("financial-systems"),
			[]domain.EntityID{EntityAuronix},
			localized(
				"João apresenta o Auronix como um de seus casos de estudo.",
				"João presents Auronix as one of his case studies.",
			),
			0.9,
			nil,
		),
		textFact(
			"project-auronix-description",
			EntityAuronix,
			domain.RelationIs,
			localized(
				"banco digital",
				"digital bank",
			),
			domain.FactCategoryProject,
			concepts("financial-systems"),
			nil,
			localized(
				"Auronix é uma plataforma financeira full stack apresentada como banco digital.",
				"Auronix is a full-stack financial platform presented as a digital bank.",
			),
			1,
			nil,
		),
		projectUses(
			"project-auronix-spring",
			EntityAuronix,
			EntitySpringBoot,
			"backend",
		),
		projectUses(
			"project-auronix-angular",
			EntityAuronix,
			EntityAngular,
			"frontend",
		),
		projectUses(
			"project-auronix-postgresql",
			EntityAuronix,
			EntityPostgreSQL,
			"database",
		),
		projectUses(
			"project-auronix-rabbitmq",
			EntityAuronix,
			EntityRabbitMQ,
			"messaging",
		),
		textFact(
			"project-auronix-transfers",
			EntityAuronix,
			domain.RelationDemonstrates,
			localized(
				"transferências com consistência transacional, versionamento otimista e ledger auditável",
				"transfers with transactional consistency, optimistic versioning and an auditable ledger",
			),
			domain.FactCategoryProject,
			concepts(
				"transactional-consistency",
				"optimistic-locking",
				"auditability",
				"financial-systems",
			),
			nil,
			localized(
				"O Auronix implementa transferências com consistência transacional, versionamento otimista e ledger auditável.",
				"Auronix implements transfers with transactional consistency, optimistic versioning and an auditable ledger.",
			),
			1,
			nil,
		),
		textFact(
			"project-auronix-distributed",
			EntityAuronix,
			domain.RelationDesigned,
			localized(
				"backend projetado para execução distribuída",
				"backend designed for distributed execution",
			),
			domain.FactCategoryProject,
			concepts(
				"distributed-systems",
				"backend",
				"scalability",
			),
			nil,
			localized(
				"O backend do Auronix foi projetado para execução distribuída.",
				"Auronix's backend was designed for distributed execution.",
			),
			1,
			nil,
		),
		projectUses(
			"project-auronix-redis",
			EntityAuronix,
			EntityRedis,
			"distributed-systems",
		),
		textFact(
			"project-auronix-async-notifications",
			EntityAuronix,
			domain.RelationUses,
			localized(
				"Redis para eventos e coordenação de notificações assíncronas",
				"Redis for events and asynchronous notification coordination",
			),
			domain.FactCategoryProject,
			concepts(
				"async-processing",
				"event-driven",
				"distributed-systems",
			),
			nil,
			localized(
				"O Auronix utiliza Redis para eventos e coordenação de notificações assíncronas.",
				"Auronix uses Redis for events and asynchronous notification coordination.",
			),
			0.95,
			nil,
		),
		projectUses(
			"project-auronix-docker",
			EntityAuronix,
			EntityDocker,
			"containerization",
		),
		projectUses(
			"project-auronix-kubernetes",
			EntityAuronix,
			EntityKubernetes,
			"orchestration",
		),
		entityFact(
			"project-auronix-eks",
			EntityAuronix,
			domain.RelationDeploysOn,
			EntityEKS,
			domain.FactCategoryProject,
			concepts(
				"aws",
				"cloud",
				"orchestration",
			),
			nil,
			localized(
				"A infraestrutura do Auronix utiliza AWS EKS.",
				"Auronix's infrastructure uses AWS EKS.",
			),
			0.9,
			nil,
		),
		projectUses(
			"project-auronix-terraform",
			EntityAuronix,
			EntityTerraform,
			"infrastructure-as-code",
		),
	}
}

func xTubeFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"project-xtube-joao",
			EntityJoao,
			domain.RelationWorkedOn,
			EntityXTube,
			domain.FactCategoryProject,
			concepts("streaming", "media-processing"),
			[]domain.EntityID{EntityXTube},
			localized(
				"João apresenta o X Tube como um de seus casos de estudo.",
				"João presents X Tube as one of his case studies.",
			),
			0.9,
			nil,
		),
		entityFact(
			"project-xtube-leadership",
			EntityJoao,
			domain.RelationLed,
			EntityXTube,
			domain.FactCategoryProject,
			concepts(
				"leadership",
				"architecture",
				"teamwork",
			),
			[]domain.EntityID{EntityXTube},
			localized(
				"João exerceu liderança técnica em uma equipe de três desenvolvedores no X Tube.",
				"João provided technical leadership for a three-developer team on X Tube.",
			),
			1,
			nil,
		),
		numberFact(
			"project-xtube-team-size",
			EntityXTube,
			domain.RelationIs,
			3,
			"developers",
			domain.NumberOperatorExact,
			domain.FactCategoryProject,
			concepts("leadership", "teamwork"),
			nil,
			localized(
				"A equipe do X Tube descrita no currículo era composta por três desenvolvedores.",
				"The X Tube team described in the résumé consisted of three developers.",
			),
			0.85,
			nil,
		),
		textFact(
			"project-xtube-flow-design",
			EntityJoao,
			domain.RelationDesigned,
			localized(
				"fluxo de upload, processamento e entrega",
				"upload, processing and delivery flow",
			),
			domain.FactCategoryProject,
			concepts(
				"architecture",
				"streaming",
				"media-processing",
			),
			[]domain.EntityID{EntityXTube},
			localized(
				"No X Tube, João participou do desenho do fluxo de upload, processamento e entrega.",
				"On X Tube, João worked on the design of the upload, processing and delivery flow.",
			),
			0.95,
			nil,
		),
		entityFact(
			"project-xtube-go",
			EntityXTube,
			domain.RelationUses,
			EntityGo,
			domain.FactCategoryProject,
			concepts(
				"backend",
				"media-processing",
				"performance",
			),
			nil,
			localized(
				"O serviço de processamento do X Tube foi desenvolvido em Go.",
				"X Tube's processing service was developed in Go.",
			),
			1,
			nil,
		),
		projectUses(
			"project-xtube-sqs",
			EntityXTube,
			EntitySQS,
			"async-processing",
		),
		projectUses(
			"project-xtube-ffmpeg",
			EntityXTube,
			EntityFFmpeg,
			"media-processing",
		),
		projectUses(
			"project-xtube-s3",
			EntityXTube,
			EntityS3,
			"storage",
		),
		projectUses(
			"project-xtube-kafka",
			EntityXTube,
			EntityKafka,
			"event-driven",
		),
		textFact(
			"project-xtube-progress-events",
			EntityXTube,
			domain.RelationImplemented,
			localized(
				"publicação de dezenas de eventos de progresso por vídeo via Kafka",
				"publication of dozens of progress events per video through Kafka",
			),
			domain.FactCategoryProject,
			concepts(
				"event-driven",
				"messaging",
				"observability",
			),
			nil,
			localized(
				"O X Tube publica dezenas de eventos de progresso por vídeo via Kafka durante o processamento.",
				"X Tube publishes dozens of progress events per video through Kafka during processing.",
			),
			0.95,
			nil,
		),
		projectUses(
			"project-xtube-prometheus",
			EntityXTube,
			EntityPrometheus,
			"observability",
		),
	}
}

func ggCompressFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"project-ggcompress-joao",
			EntityJoao,
			domain.RelationWorkedOn,
			EntityGGCompress,
			domain.FactCategoryProject,
			concepts(
				"compression",
				"concurrency",
				"performance",
			),
			[]domain.EntityID{EntityGGCompress},
			localized(
				"João apresenta o GGCompress como um de seus casos de estudo.",
				"João presents GGCompress as one of his case studies.",
			),
			0.9,
			nil,
		),
		entityFact(
			"project-ggcompress-go",
			EntityGGCompress,
			domain.RelationUses,
			EntityGo,
			domain.FactCategoryProject,
			concepts(
				"backend",
				"concurrency",
				"performance",
			),
			nil,
			localized(
				"O GGCompress é uma engine desenvolvida em Go.",
				"GGCompress is an engine developed in Go.",
			),
			1,
			nil,
		),
		textFact(
			"project-ggcompress-purpose",
			EntityGGCompress,
			domain.RelationIs,
			localized(
				"engine de compressão e arquivamento",
				"compression and archiving engine",
			),
			domain.FactCategoryProject,
			concepts(
				"compression",
				"file-processing",
			),
			nil,
			localized(
				"GGCompress é uma engine de compressão e arquivamento.",
				"GGCompress is a compression and archiving engine.",
			),
			1,
			nil,
		),
		textFact(
			"project-ggcompress-concurrency",
			EntityGGCompress,
			domain.RelationDemonstrates,
			localized(
				"processamento concorrente",
				"concurrent processing",
			),
			domain.FactCategoryProject,
			concepts(
				"concurrency",
				"performance",
			),
			nil,
			localized(
				"O GGCompress é orientado a processamento concorrente.",
				"GGCompress is designed around concurrent processing.",
			),
			1,
			nil,
		),
		textFact(
			"project-ggcompress-integrity",
			EntityGGCompress,
			domain.RelationDemonstrates,
			localized(
				"integridade determinística",
				"deterministic integrity",
			),
			domain.FactCategoryProject,
			concepts(
				"data-integrity",
				"reliability",
			),
			nil,
			localized(
				"O GGCompress foi desenvolvido com foco em integridade determinística.",
				"GGCompress was developed with a focus on deterministic integrity.",
			),
			0.95,
			nil,
		),
		textFact(
			"project-ggcompress-format",
			EntityGGCompress,
			domain.RelationImplemented,
			localized(
				"formato versionado .ggc com manifesto e indexação por chunks",
				"versioned .ggc format with manifest and chunk indexing",
			),
			domain.FactCategoryProject,
			concepts(
				"file-format",
				"chunking",
				"data-integrity",
			),
			nil,
			localized(
				"O GGCompress utiliza um formato versionado .ggc com manifesto e indexação por chunks.",
				"GGCompress uses a versioned .ggc format with a manifest and chunk indexing.",
			),
			0.95,
			nil,
		),
		textFact(
			"project-ggcompress-goroutines",
			EntityGGCompress,
			domain.RelationUses,
			localized(
				"goroutines com escrita ordenada",
				"goroutines with ordered writes",
			),
			domain.FactCategoryProject,
			concepts(
				"concurrency",
				"goroutines",
				"ordered-processing",
			),
			nil,
			localized(
				"O GGCompress utiliza goroutines e preserva escrita ordenada.",
				"GGCompress uses goroutines while preserving ordered writes.",
			),
			1,
			nil,
		),
		numberFact(
			"project-ggcompress-throughput",
			EntityGGCompress,
			domain.RelationAchieved,
			1.23,
			"GB/s",
			domain.NumberOperatorUpTo,
			domain.FactCategoryAchievement,
			concepts(
				"performance",
				"benchmark",
				"throughput",
				"concurrency",
			),
			nil,
			localized(
				"O GGCompress atingiu throughput de até 1,23 GB/s em benchmarks.",
				"GGCompress reached throughput of up to 1.23 GB/s in benchmarks.",
			),
			1,
			nil,
		),
		numberFact(
			"project-ggcompress-benchmark-file-size",
			EntityGGCompress,
			domain.RelationProcesses,
			9.77,
			"GB",
			domain.NumberOperatorExact,
			domain.FactCategoryAchievement,
			concepts(
				"performance",
				"benchmark",
			),
			nil,
			localized(
				"O benchmark de 1,23 GB/s foi realizado com arquivos de 9,77 GB.",
				"The 1.23 GB/s benchmark used 9.77 GB files.",
			),
			0.95,
			nil,
		),
		textFact(
			"project-ggcompress-checksum",
			EntityGGCompress,
			domain.RelationImplemented,
			localized(
				"checksum por chunk",
				"per-chunk checksum",
			),
			domain.FactCategoryProject,
			concepts(
				"data-integrity",
				"checksum",
				"reliability",
			),
			nil,
			localized(
				"O GGCompress implementa checksum por chunk.",
				"GGCompress implements a checksum for each chunk.",
			),
			0.95,
			nil,
		),
		entityFact(
			"project-ggcompress-sha256",
			EntityGGCompress,
			domain.RelationUses,
			EntitySHA256,
			domain.FactCategoryProject,
			concepts(
				"data-integrity",
				"checksum",
				"reliability",
			),
			nil,
			localized(
				"O GGCompress realiza verificação global utilizando SHA-256.",
				"GGCompress performs global verification using SHA-256.",
			),
			0.95,
			nil,
		),
		textFact(
			"project-ggcompress-safe-extraction",
			EntityGGCompress,
			domain.RelationImplemented,
			localized(
				"extração segura com operações atômicas e isolamento temporário",
				"safe extraction using atomic operations and temporary isolation",
			),
			domain.FactCategoryProject,
			concepts(
				"reliability",
				"data-integrity",
				"atomicity",
				"safe-extraction",
			),
			nil,
			localized(
				"O GGCompress implementa extração segura com operações atômicas e isolamento temporário.",
				"GGCompress implements safe extraction with atomic operations and temporary isolation.",
			),
			0.95,
			nil,
		),
	}
}

func voxFacts() []domain.Fact {
	return []domain.Fact{
		entityFact(
			"project-vox-joao",
			EntityJoao,
			domain.RelationWorkedOn,
			EntityVox,
			domain.FactCategoryProject,
			concepts("electoral-system"),
			[]domain.EntityID{EntityVox},
			localized(
				"João apresenta o Vox como um de seus casos de estudo.",
				"João presents Vox as one of his case studies.",
			),
			0.9,
			nil,
		),
		textFact(
			"project-vox-description",
			EntityVox,
			domain.RelationIs,
			localized(
				"plataforma eleitoral multiplataforma",
				"cross-platform electoral platform",
			),
			domain.FactCategoryProject,
			concepts(
				"electoral-system",
				"cross-platform",
			),
			nil,
			localized(
				"Vox é uma plataforma eleitoral para desktop e dispositivos móveis.",
				"Vox is an electoral platform for desktop and mobile devices.",
			),
			1,
			nil,
		),
		projectUses(
			"project-vox-angular",
			EntityVox,
			EntityAngular,
			"frontend",
		),
		projectUses(
			"project-vox-tauri",
			EntityVox,
			EntityTauri,
			"cross-platform",
		),
		textFact(
			"project-vox-performance",
			EntityVox,
			domain.RelationDesigned,
			localized(
				"baixo consumo de recursos, inicialização rápida e desempenho",
				"low resource consumption, fast startup and performance",
			),
			domain.FactCategoryProject,
			concepts(
				"performance",
				"resource-efficiency",
			),
			nil,
			localized(
				"O Vox prioriza baixo consumo de recursos, inicialização rápida e desempenho.",
				"Vox prioritizes low resource consumption, fast startup and performance.",
			),
			0.9,
			nil,
		),
		projectUses(
			"project-vox-java",
			EntityVox,
			EntityJava,
			"backend",
		),
		projectUses(
			"project-vox-spring",
			EntityVox,
			EntitySpringBoot,
			"backend",
		),
		projectUses(
			"project-vox-postgresql",
			EntityVox,
			EntityPostgreSQL,
			"database",
		),
		textFact(
			"project-vox-vote-integrity",
			EntityVox,
			domain.RelationDemonstrates,
			localized(
				"integridade, consistência e persistência segura dos votos",
				"integrity, consistency and secure vote persistence",
			),
			domain.FactCategoryProject,
			concepts(
				"data-integrity",
				"consistency",
				"secure-persistence",
			),
			nil,
			localized(
				"O backend do Vox utiliza PostgreSQL para garantir integridade, consistência e persistência segura dos votos.",
				"Vox's backend uses PostgreSQL to ensure integrity, consistency and secure persistence of votes.",
			),
			0.95,
			nil,
		),
		numberFact(
			"project-vox-voters",
			EntityVox,
			domain.RelationAchieved,
			500,
			"voters",
			domain.NumberOperatorGreaterThan,
			domain.FactCategoryAchievement,
			concepts(
				"usage",
				"real-world-use",
			),
			[]domain.EntityID{EntityEtec},
			localized(
				"O Vox foi utilizado por mais de 500 eleitores em uma eleição realizada na Etec de Guarulhos.",
				"Vox was used by more than 500 voters in an election held at Etec de Guarulhos.",
			),
			1,
			nil,
		),
	}
}

func projectUses(
	id string,
	project domain.EntityID,
	technology domain.EntityID,
	concept string,
) domain.Fact {
	projectName := projectDisplayName(project)
	technologyName := entityDisplayProjectTechnology(technology)

	return entityFact(
		id,
		project,
		domain.RelationUses,
		technology,
		domain.FactCategoryProject,
		concepts(concept),
		nil,
		localized(
			projectName+" utiliza "+technologyName+".",
			projectName+" uses "+technologyName+".",
		),
		0.85,
		nil,
	)
}

func projectDisplayName(id domain.EntityID) string {
	names := map[domain.EntityID]string{
		EntityAuronix:    "Auronix",
		EntityXTube:      "X Tube",
		EntityGGCompress: "GGCompress",
		EntityVox:        "Vox",
	}

	return names[id]
}

func entityDisplayProjectTechnology(id domain.EntityID) string {
	names := map[domain.EntityID]string{
		EntitySpringBoot: "Spring Boot",
		EntityAngular:    "Angular",
		EntityPostgreSQL: "PostgreSQL",
		EntityRabbitMQ:   "RabbitMQ",
		EntityRedis:      "Redis",
		EntityDocker:     "Docker",
		EntityKubernetes: "Kubernetes",
		EntityTerraform:  "Terraform",
		EntitySQS:        "SQS",
		EntityFFmpeg:     "FFmpeg",
		EntityS3:         "S3",
		EntityKafka:      "Kafka",
		EntityPrometheus: "Prometheus",
		EntityTauri:      "Tauri",
		EntityJava:       "Java",
		EntityGo:         "Go",
		EntityEKS:        "Amazon EKS",
	}

	return names[id]
}
