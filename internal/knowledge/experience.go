package knowledge

import "ai-agent/internal/domain"

func experienceFacts() []domain.Fact {
	facts := make([]domain.Fact, 0)

	facts = append(facts, currentExperienceFacts()...)
	facts = append(facts, juniorExperienceFacts()...)
	facts = append(facts, internshipFacts()...)

	return facts
}

func currentExperienceFacts() []domain.Fact {
	period := currentPeriod(2025, 6)
	context := []domain.EntityID{EntityUFind}

	return []domain.Fact{
		entityFact(
			"experience-current-company",
			EntityJoao,
			domain.RelationWorksAt,
			EntityUFind,
			domain.FactCategoryExperience,
			concepts("professional-experience"),
			context,
			localized(
				"João trabalha na uFind Tecnologia desde junho de 2025.",
				"João has worked at uFind Tecnologia since June 2025.",
			),
			1,
			period,
		),
		entityFact(
			"experience-current-role",
			EntityJoao,
			domain.RelationHasRole,
			EntityRoleMid,
			domain.FactCategoryExperience,
			concepts("professional-experience"),
			context,
			localized(
				"João atua como Desenvolvedor Pleno na uFind Tecnologia.",
				"João works as a Mid-level Software Developer at uFind Tecnologia.",
			),
			1,
			period,
		),
		textFact(
			"experience-current-financial-systems",
			EntityJoao,
			domain.RelationHasExperience,
			localized(
				"sistemas financeiros em produção",
				"production financial systems",
			),
			domain.FactCategoryExperience,
			concepts(
				"financial-systems",
				"production",
				"transactional-consistency",
				"reliability",
			),
			context,
			localized(
				"Na uFind, João evolui sistemas financeiros em produção com validações e fluxos rastreáveis, focando consistência transacional e confiabilidade operacional.",
				"At uFind, João evolves production financial systems with validations and traceable flows, focusing on transactional consistency and operational reliability.",
			),
			1,
			period,
		),
		numberFact(
			"experience-current-financial-volume",
			EntityJoao,
			domain.RelationAutomated,
			1000000,
			"BRL/month",
			domain.NumberOperatorGreaterThan,
			domain.FactCategoryAchievement,
			concepts(
				"automation",
				"financial-systems",
				"operational-efficiency",
			),
			context,
			localized(
				"João automatizou fluxos financeiros responsáveis por mais de R$ 1 milhão por mês em movimentações.",
				"João automated financial flows responsible for more than BRL 1 million per month in transactions.",
			),
			1,
			period,
		),
		textFact(
			"experience-current-operation-time",
			EntityJoao,
			domain.RelationImproved,
			localized(
				"operações manuais de dias para minutos",
				"manual operations from days to minutes",
			),
			domain.FactCategoryAchievement,
			concepts(
				"automation",
				"operational-efficiency",
			),
			context,
			localized(
				"A automação reduziu operações manuais de dias para minutos.",
				"The automation reduced manual operations from days to minutes.",
			),
			0.95,
			period,
		),
		textFact(
			"experience-current-strategy-pattern",
			EntityJoao,
			domain.RelationImplemented,
			localized(
				"seleção dinâmica de estratégias de processamento utilizando Strategy Pattern",
				"dynamic processing strategy selection using the Strategy Pattern",
			),
			domain.FactCategoryExperience,
			concepts(
				"strategy-pattern",
				"software-design",
				"architecture",
			),
			context,
			localized(
				"João implementou seleção dinâmica de estratégias de processamento utilizando Strategy Pattern para adaptar fluxos às regras de cada operação.",
				"João implemented dynamic processing strategy selection using the Strategy Pattern to adapt flows to each operation's rules.",
			),
			0.9,
			period,
		),
		entityFact(
			"experience-current-mongodb",
			EntityJoao,
			domain.RelationUses,
			EntityMongoDB,
			domain.FactCategoryExperience,
			concepts("database", "data-modeling"),
			context,
			localized(
				"Na uFind, João modelou dados do mesmo domínio em schemas distintos no MongoDB, isolando estruturas e regras específicas de cada contexto.",
				"At uFind, João modeled data from the same domain in separate MongoDB schemas, isolating structures and rules for each context.",
			),
			0.9,
			period,
		),
		entityFact(
			"experience-current-aws",
			EntityJoao,
			domain.RelationUses,
			EntityAWS,
			domain.FactCategoryExperience,
			concepts("cloud", "infrastructure"),
			context,
			localized(
				"João contribui com a arquitetura AWS da uFind.",
				"João contributes to uFind's AWS architecture.",
			),
			0.9,
			period,
		),
		entityFact(
			"experience-current-ecs",
			EntityJoao,
			domain.RelationUses,
			EntityECS,
			domain.FactCategoryExperience,
			concepts("cloud", "containers"),
			context,
			localized(
				"Na arquitetura AWS da uFind, João atua com ECS.",
				"In uFind's AWS architecture, João works with ECS.",
			),
			0.8,
			period,
		),
		entityFact(
			"experience-current-s3",
			EntityJoao,
			domain.RelationUses,
			EntityS3,
			domain.FactCategoryExperience,
			concepts("cloud", "storage"),
			context,
			localized(
				"Na arquitetura AWS da uFind, João atua com S3.",
				"In uFind's AWS architecture, João works with S3.",
			),
			0.8,
			period,
		),
		entityFact(
			"experience-current-iam",
			EntityJoao,
			domain.RelationUses,
			EntityIAM,
			domain.FactCategoryExperience,
			concepts("cloud", "security"),
			context,
			localized(
				"Na arquitetura AWS da uFind, João atua com IAM.",
				"In uFind's AWS architecture, João works with IAM.",
			),
			0.8,
			period,
		),
		textFact(
			"experience-current-deployment",
			EntityJoao,
			domain.RelationHasExperience,
			localized(
				"manutenção de deploys automatizados",
				"maintenance of automated deployments",
			),
			domain.FactCategoryExperience,
			concepts(
				"deployment",
				"automation",
				"devops",
			),
			context,
			localized(
				"João atua na manutenção de deploys automatizados na uFind.",
				"João works on maintaining automated deployments at uFind.",
			),
			0.8,
			period,
		),
		entityFact(
			"experience-current-cloudflare",
			EntityJoao,
			domain.RelationUses,
			EntityCloudflare,
			domain.FactCategoryExperience,
			concepts(
				"dns",
				"proxy",
				"routing",
				"infrastructure",
			),
			context,
			localized(
				"João gerencia configurações de DNS, proxy e roteamento na Cloudflare para separar ambientes de testes e produção.",
				"João manages DNS, proxy and routing configurations in Cloudflare to separate testing and production environments.",
			),
			0.85,
			period,
		),
	}
}

func juniorExperienceFacts() []domain.Fact {
	period := closedPeriod(2024, 9, 2025, 5)
	context := []domain.EntityID{EntityRepresentaOnline}

	return []domain.Fact{
		entityFact(
			"experience-junior-company",
			EntityJoao,
			domain.RelationWorkedAt,
			EntityRepresentaOnline,
			domain.FactCategoryExperience,
			concepts("professional-experience"),
			context,
			localized(
				"João trabalhou como Desenvolvedor Júnior na Representa Online entre setembro de 2024 e maio de 2025.",
				"João worked as a Junior Software Developer at Representa Online from September 2024 to May 2025.",
			),
			1,
			period,
		),
		entityFact(
			"experience-junior-role",
			EntityJoao,
			domain.RelationHasRole,
			EntityRoleJunior,
			domain.FactCategoryExperience,
			concepts("professional-experience"),
			context,
			localized(
				"Na Representa Online, João atuou como Desenvolvedor Júnior.",
				"At Representa Online, João worked as a Junior Software Developer.",
			),
			0.95,
			period,
		),
		numberFact(
			"experience-junior-media-volume",
			EntityJoao,
			domain.RelationProcesses,
			16,
			"GB/flow",
			domain.NumberOperatorGreaterThan,
			domain.FactCategoryAchievement,
			concepts(
				"data-pipelines",
				"media-processing",
				"performance",
			),
			context,
			localized(
				"João desenvolveu pipelines de ingestão e processamento de mídia televisiva capazes de processar cargas superiores a 16 GB por fluxo.",
				"João developed television media ingestion and processing pipelines capable of handling workloads above 16 GB per flow.",
			),
			1,
			period,
		),
		entityFact(
			"experience-junior-node",
			EntityJoao,
			domain.RelationUses,
			EntityNodeJS,
			domain.FactCategoryExperience,
			concepts(
				"backend",
				"async-processing",
			),
			context,
			localized(
				"João estruturou fluxos assíncronos em Node.js.",
				"João structured asynchronous flows in Node.js.",
			),
			0.9,
			period,
		),
		entityFact(
			"experience-junior-worker-threads",
			EntityJoao,
			domain.RelationUses,
			EntityWorkerThread,
			domain.FactCategoryExperience,
			concepts(
				"concurrency",
				"async-processing",
				"fault-isolation",
			),
			context,
			localized(
				"João utilizou Worker Threads em fluxos assíncronos para melhorar uso de CPU, memória, isolamento de falhas e estabilidade operacional.",
				"João used Worker Threads in asynchronous flows to improve CPU and memory usage, fault isolation and operational stability.",
			),
			0.95,
			period,
		),
		textFact(
			"experience-junior-streams",
			EntityJoao,
			domain.RelationUses,
			localized(
				"Streams do Node.js",
				"Node.js Streams",
			),
			domain.FactCategoryExperience,
			concepts(
				"streams",
				"data-processing",
				"memory-efficiency",
			),
			context,
			localized(
				"João utilizou Streams do Node.js nos fluxos de processamento.",
				"João used Node.js Streams in processing flows.",
			),
			0.9,
			period,
		),
		textFact(
			"experience-junior-serverless-ai",
			EntityJoao,
			domain.RelationImplemented,
			localized(
				"Serverless Function stateless para integração entre serviço de IA e frontend Next.js",
				"stateless Serverless Function integrating an AI service with a Next.js frontend",
			),
			domain.FactCategoryExperience,
			concepts(
				"serverless",
				"stateless",
				"artificial-intelligence",
				"integration",
			),
			context,
			localized(
				"João implementou uma Serverless Function stateless como camada de integração entre um serviço de IA e um frontend em Next.js.",
				"João implemented a stateless Serverless Function as an integration layer between an AI service and a Next.js frontend.",
			),
			0.9,
			period,
		),
		entityFact(
			"experience-junior-base44",
			EntityJoao,
			domain.RelationUses,
			EntityBase44,
			domain.FactCategoryExperience,
			concepts("data-migration", "artificial-intelligence"),
			context,
			localized(
				"João migrou dados da Base44 para uma base de conhecimento.",
				"João migrated data from Base44 to a knowledge base.",
			),
			0.75,
			period,
		),
		entityFact(
			"experience-junior-vector-stores",
			EntityJoao,
			domain.RelationUses,
			EntityVectorStores,
			domain.FactCategoryExperience,
			concepts(
				"semantic-search",
				"retrieval",
				"artificial-intelligence",
			),
			context,
			localized(
				"Na migração da base de conhecimento, João utilizou OpenAI Vector Stores.",
				"During the knowledge-base migration, João used OpenAI Vector Stores.",
			),
			0.85,
			period,
		),
		entityFact(
			"experience-junior-file-search",
			EntityJoao,
			domain.RelationUses,
			EntityFileSearch,
			domain.FactCategoryExperience,
			concepts(
				"semantic-search",
				"retrieval",
				"artificial-intelligence",
			),
			context,
			localized(
				"João utilizou File Search para melhorar recuperação semântica e relevância das respostas.",
				"João used File Search to improve semantic retrieval and answer relevance.",
			),
			0.9,
			period,
		),
	}
}

func internshipFacts() []domain.Fact {
	period := closedPeriod(2024, 6, 2024, 8)
	context := []domain.EntityID{EntityRepresentaOnline}

	return []domain.Fact{
		entityFact(
			"experience-intern-company",
			EntityJoao,
			domain.RelationWorkedAt,
			EntityRepresentaOnline,
			domain.FactCategoryExperience,
			concepts("professional-experience"),
			context,
			localized(
				"João realizou estágio como Desenvolvedor de Sistemas na Representa Online entre junho e agosto de 2024.",
				"João worked as a Systems Developer Intern at Representa Online from June to August 2024.",
			),
			1,
			period,
		),
		entityFact(
			"experience-intern-role",
			EntityJoao,
			domain.RelationHasRole,
			EntityRoleIntern,
			domain.FactCategoryExperience,
			concepts("professional-experience"),
			context,
			localized(
				"João atuou como Desenvolvedor de Sistemas em estágio na Representa Online.",
				"João worked as a Systems Developer Intern at Representa Online.",
			),
			0.95,
			period,
		),
		entityFact(
			"experience-intern-postgresql",
			EntityJoao,
			domain.RelationUses,
			EntityPostgreSQL,
			domain.FactCategoryExperience,
			concepts(
				"database",
				"search",
				"full-text-search",
			),
			context,
			localized(
				"João implementou busca textual utilizando PostgreSQL Full Text Search.",
				"João implemented text search using PostgreSQL Full Text Search.",
			),
			0.9,
			period,
		),
		textFact(
			"experience-intern-haversine",
			EntityJoao,
			domain.RelationImplemented,
			localized(
				"busca geoespacial ordenada por proximidade utilizando a fórmula de Haversine",
				"geospatial search ordered by proximity using the Haversine formula",
			),
			domain.FactCategoryExperience,
			concepts(
				"geospatial-search",
				"search",
				"algorithms",
			),
			context,
			localized(
				"João desenvolveu busca geoespacial ordenada por proximidade utilizando a fórmula de Haversine.",
				"João developed geospatial search ordered by proximity using the Haversine formula.",
			),
			0.95,
			period,
		),
		entityFact(
			"experience-intern-jwt",
			EntityJoao,
			domain.RelationUses,
			EntityJWT,
			domain.FactCategoryExperience,
			concepts("authentication", "security"),
			context,
			localized(
				"João contribuiu com fluxos de autenticação JWT.",
				"João contributed to JWT authentication flows.",
			),
			0.8,
			period,
		),
		entityFact(
			"experience-intern-oauth2",
			EntityJoao,
			domain.RelationUses,
			EntityOAuth2,
			domain.FactCategoryExperience,
			concepts("authentication", "security"),
			context,
			localized(
				"João contribuiu com fluxos de autenticação OAuth2.",
				"João contributed to OAuth2 authentication flows.",
			),
			0.8,
			period,
		),
		entityFact(
			"experience-intern-angular",
			EntityJoao,
			domain.RelationUses,
			EntityAngular,
			domain.FactCategoryExperience,
			concepts(
				"frontend",
				"performance",
			),
			context,
			localized(
				"João realizou otimizações de performance no frontend em Angular.",
				"João performed frontend performance optimizations in Angular.",
			),
			0.85,
			period,
		),
		entityFact(
			"experience-intern-web-workers",
			EntityJoao,
			domain.RelationUses,
			EntityWebWorker,
			domain.FactCategoryExperience,
			concepts(
				"frontend",
				"concurrency",
				"performance",
			),
			context,
			localized(
				"João utilizou Web Workers na geração e download de comprovantes em PDF para evitar o bloqueio da interface.",
				"João used Web Workers for PDF receipt generation and download to avoid blocking the user interface.",
			),
			0.9,
			period,
		),
	}
}
