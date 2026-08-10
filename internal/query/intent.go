package query

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/language"
)

type IntentDetector struct{}

func NewIntentDetector() *IntentDetector {
	return &IntentDetector{}
}

func (d *IntentDetector) Detect(
	normalized string,
	entities []domain.EntityMatch,
	concepts []domain.ConceptMatch,
) domain.Intent {
	if hasAnyPhrase(
		normalized,
		comparisonMarkers,
	) {
		return domain.IntentComparison
	}

	if hasAnyPhrase(
		normalized,
		contactMarkers,
	) {
		return domain.IntentContact
	}

	if hasAnyPhrase(
		normalized,
		certificationMarkers,
	) {
		return domain.IntentCertification
	}

	if hasAnyPhrase(
		normalized,
		educationMarkers,
	) {
		return domain.IntentEducation
	}

	if hasTechnologyEntity(entities) &&
		hasAnyPhrase(
			normalized,
			technologyUsageMarkers,
		) {
		return domain.IntentTechnologyUsage
	}

	if hasAnyPhrase(
		normalized,
		capabilityMarkers,
	) {
		return domain.IntentCapability
	}

	if hasAnyPhrase(
		normalized,
		listMarkers,
	) {
		return domain.IntentList
	}

	if hasAnyPhrase(
		normalized,
		experienceMarkers,
	) {
		return domain.IntentExperience
	}

	if hasAnyPhrase(
		normalized,
		overviewMarkers,
	) {
		return domain.IntentOverview
	}

	if hasProjectEntity(entities) {
		return domain.IntentOverview
	}

	if len(entities) > 0 ||
		len(concepts) > 0 {
		return domain.IntentDirectFact
	}

	return domain.IntentUnknown
}

var comparisonMarkers = []string{
	"melhor",
	"mais demonstra",
	"mais representa",
	"mais forte",
	"qual projeto mais",
	"qual deles",
	"comparar",
	"compare",
	"comparacao",
	"comparação",
	"entre os projetos",
	"best",
	"most demonstrates",
	"strongest",
	"which project best",
	"which one",
	"comparison",
	"between the projects",
}

var contactMarkers = []string{
	"email",
	"e mail",
	"telefone",
	"celular",
	"contato",
	"contact",
	"phone",
	"phone number",
	"email address",
}

var certificationMarkers = []string{
	"certificacao",
	"certificacoes",
	"certificado",
	"certificados",
	"certification",
	"certifications",
	"certificate",
	"certificates",
}

var educationMarkers = []string{
	"educacao",
	"formacao",
	"estuda",
	"estudou",
	"faculdade",
	"curso",
	"escola",
	"fiap",
	"etec",
	"education",
	"studies",
	"studied",
	"college",
	"university",
	"course",
	"school",
	"academic",
}

var technologyUsageMarkers = []string{
	"onde usou",
	"onde utiliza",
	"onde utilizou",
	"onde trabalha com",
	"onde trabalhou com",
	"em qual projeto usa",
	"em qual projeto usou",
	"em quais projetos usa",
	"em quais projetos usou",
	"como usou",
	"como utiliza",
	"how did",
	"how does",
	"where did",
	"where does",
	"where has",
	"which project uses",
	"which projects use",
	"what project uses",
	"what projects use",
	"used in",
	"uses in",
}

var capabilityMarkers = []string{
	"sabe",
	"conhece",
	"domina",
	"tem experiencia",
	"tem experiência",
	"possui experiencia",
	"possui experiência",
	"ja trabalhou com",
	"já trabalhou com",
	"capacidade",
	"habilidade",
	"does he know",
	"does joao know",
	"does joão know",
	"has experience with",
	"have experience with",
	"experienced with",
	"can he",
	"can joao",
	"can joão",
	"skilled in",
	"proficient in",
}

var listMarkers = []string{
	"liste",
	"lista",
	"quais",
	"quais sao",
	"quais são",
	"todos os",
	"todas as",
	"mostre os",
	"mostre as",
	"list",
	"what are",
	"which technologies",
	"which skills",
	"all projects",
	"all technologies",
	"all skills",
}

var experienceMarkers = []string{
	"experiencia profissional",
	"experiência profissional",
	"carreira",
	"trajetoria",
	"trajetória",
	"historico profissional",
	"histórico profissional",
	"empregos",
	"professional experience",
	"work experience",
	"career",
	"employment history",
	"previous roles",
}

var overviewMarkers = []string{
	"me fale sobre",
	"fale sobre",
	"conte sobre",
	"explique",
	"o que e",
	"o que é",
	"quem e",
	"quem é",
	"tell me about",
	"tell me more about",
	"explain",
	"what is",
	"who is",
	"give me an overview",
	"overview",
}

func hasAnyPhrase(
	normalized string,
	markers []string,
) bool {
	for _, marker := range markers {
		if containsPhrase(
			normalized,
			normalizeMarker(marker),
		) {
			return true
		}
	}

	return false
}

func normalizeMarker(value string) string {
	return language.Normalize(value)
}

func hasTechnologyEntity(
	entities []domain.EntityMatch,
) bool {
	for _, entity := range entities {
		switch entity.EntityID {
		case knowledge.EntityAngular,
			knowledge.EntityReact,
			knowledge.EntityNextJS,
			knowledge.EntityJava,
			knowledge.EntitySpringBoot,
			knowledge.EntityGo,
			knowledge.EntityNodeJS,
			knowledge.EntityNestJS,
			knowledge.EntityPostgreSQL,
			knowledge.EntityMongoDB,
			knowledge.EntityRedis,
			knowledge.EntityRabbitMQ,
			knowledge.EntityKafka,
			knowledge.EntitySQS,
			knowledge.EntityDocker,
			knowledge.EntityTerraform,
			knowledge.EntityKubernetes,
			knowledge.EntityAWS,
			knowledge.EntityECS,
			knowledge.EntityEKS,
			knowledge.EntityS3,
			knowledge.EntityIAM,
			knowledge.EntityCloudflare,
			knowledge.EntityFFmpeg,
			knowledge.EntityPrometheus,
			knowledge.EntityTauri,
			knowledge.EntityJWT,
			knowledge.EntityOAuth2,
			knowledge.EntityWorkerThread,
			knowledge.EntityWebWorker,
			knowledge.EntitySHA256,
			knowledge.EntityBase44,
			knowledge.EntityVectorStores,
			knowledge.EntityFileSearch:
			return true
		}
	}

	return false
}

func hasProjectEntity(
	entities []domain.EntityMatch,
) bool {
	for _, entity := range entities {
		switch entity.EntityID {
		case knowledge.EntityAuronix,
			knowledge.EntityXTube,
			knowledge.EntityGGCompress,
			knowledge.EntityVox:
			return true
		}
	}

	return false
}
