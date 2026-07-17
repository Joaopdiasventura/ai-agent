package nlp

type entityAlias struct {
	Tokens []string
	Entity Entity
}

var entityAliases = []entityAlias{
	{
		Tokens: []string{"ufind", "tecnologia"},
		Entity: Entity{
			Type:  EntityCompany,
			Value: "uFind Tecnologia",
		},
	},
	{
		Tokens: []string{"representa", "online"},
		Entity: Entity{
			Type:  EntityCompany,
			Value: "Representa Online",
		},
	},
	{
		Tokens: []string{"etec", "guarulhos"},
		Entity: Entity{
			Type:  EntityInstitution,
			Value: "Etec Guarulhos",
		},
	},
	{
		Tokens: []string{"spring", "boot"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Spring Boot",
		},
	},
	{
		Tokens: []string{"docker", "compose"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Docker Compose",
		},
	},
	{
		Tokens: []string{"auronix"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "Auronix",
		},
	},
	{
		Tokens: []string{"modularis"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "Modularis",
		},
	},
	{
		Tokens: []string{"votrix"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "Votrix",
		},
	},
	{
		Tokens: []string{"ggcompress"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "GGCompress",
		},
	},
	{
		Tokens: []string{"ggc"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "GGCompress",
		},
	},
	{
		Tokens: []string{"vox"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "VOX",
		},
	},
	{
		Tokens: []string{"etecfy"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "Etecfy",
		},
	},
	{
		Tokens: []string{"ufind"},
		Entity: Entity{
			Type:  EntityCompany,
			Value: "uFind Tecnologia",
		},
	},
	{
		Tokens: []string{"representa"},
		Entity: Entity{
			Type:  EntityCompany,
			Value: "Representa Online",
		},
	},
	{
		Tokens: []string{"fiap"},
		Entity: Entity{
			Type:  EntityInstitution,
			Value: "FIAP",
		},
	},
	{
		Tokens: []string{"etec"},
		Entity: Entity{
			Type:  EntityInstitution,
			Value: "Etec Guarulhos",
		},
	},
	{
		Tokens: []string{"go"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Go",
		},
	},
	{
		Tokens: []string{"golang"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Go",
		},
	},
	{
		Tokens: []string{"java"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Java",
		},
	},
	{
		Tokens: []string{"typescript"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "TypeScript",
		},
	},
	{
		Tokens: []string{"angular"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Angular",
		},
	},
	{
		Tokens: []string{"nestjs"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "NestJS",
		},
	},
	{
		Tokens: []string{"fastify"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Fastify",
		},
	},
	{
		Tokens: []string{"redis"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Redis",
		},
	},
	{
		Tokens: []string{"rabbitmq"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "RabbitMQ",
		},
	},
	{
		Tokens: []string{"bullmq"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "BullMQ",
		},
	},
	{
		Tokens: []string{"postgresql"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "PostgreSQL",
		},
	},
	{
		Tokens: []string{"postgres"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "PostgreSQL",
		},
	},
	{
		Tokens: []string{"mongodb"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "MongoDB",
		},
	},
	{
		Tokens: []string{"capacitor"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Capacitor",
		},
	},
	{
		Tokens: []string{"tauri"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Tauri",
		},
	},
}

var intentPriority = []Intent{
	IntentContact,
	IntentProject,
	IntentTechnologies,
	IntentCurrentJob,
	IntentFirstJob,
	IntentEducation,
}

var intentKeywords = map[string]map[Intent]int{
	"faz": {
		IntentCurrentJob: 3,
	},
	"desenvolvido": {
		IntentTechnologies: 3,
	},
	"trabalha": {
		IntentCurrentJob: 3,
	},
	"trabalho": {
		IntentCurrentJob: 2,
	},
	"profissão": {
		IntentCurrentJob: 3,
	},
	"profissao": {
		IntentCurrentJob: 3,
	},
	"cargo": {
		IntentCurrentJob: 3,
	},
	"atualmente": {
		IntentCurrentJob: 2,
	},
	"primeiro": {
		IntentFirstJob: 3,
	},
	"emprego": {
		IntentFirstJob:   2,
		IntentCurrentJob: 1,
	},
	"estágio": {
		IntentFirstJob: 2,
	},
	"estagio": {
		IntentFirstJob: 2,
	},
	"estagiário": {
		IntentFirstJob: 2,
	},
	"estagiario": {
		IntentFirstJob: 2,
	},
	"estuda": {
		IntentEducation: 3,
	},
	"faculdade": {
		IntentEducation: 3,
	},
	"curso": {
		IntentEducation: 2,
	},
	"formação": {
		IntentEducation: 3,
	},
	"formacao": {
		IntentEducation: 3,
	},
	"projeto": {
		IntentProject: 3,
	},
	"projetos": {
		IntentProject: 3,
	},
	"auronix": {
		IntentProject: 3,
	},
	"modularis": {
		IntentProject: 3,
	},
	"votrix": {
		IntentProject: 3,
	},
	"ggcompress": {
		IntentProject: 3,
	},
	"vox": {
		IntentProject: 3,
	},
	"etecfy": {
		IntentProject: 3,
	},
	"usadas": {
		IntentTechnologies: 3,
	},
	"usados": {
		IntentTechnologies: 3,
	},
	"usada": {
		IntentTechnologies: 3,
	},
	"usado": {
		IntentTechnologies: 3,
	},
	"utiliza": {
		IntentTechnologies: 3,
	},
	"utilizam": {
		IntentTechnologies: 3,
	},
	"tecnologia": {
		IntentTechnologies: 3,
	},
	"tecnologias": {
		IntentTechnologies: 3,
	},
	"stack": {
		IntentTechnologies: 3,
	},
	"linguagem": {
		IntentTechnologies: 2,
	},
	"linguagens": {
		IntentTechnologies: 2,
	},
	"usa": {
		IntentTechnologies: 3,
	},
	"contato": {
		IntentContact: 3,
	},
	"email": {
		IntentContact: 3,
	},
	"telefone": {
		IntentContact: 3,
	},
	"linkedin": {
		IntentContact: 3,
	},
	"github": {
		IntentContact: 3,
	},
}

var queryExpansions = map[string][]string{
	"faz": {
		"trabalha",
		"cursa",
		"estuda",
		"atualmente",
		"desenvolvedor",
	},
	"trabalho": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"trabalha": {
		"atualmente",
		"desenvolvedor",
	},
	"profissão": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"profissao": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"ocupação": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"ocupacao": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"feito": {
		"desenvolvido",
		"criado",
		"como",
		"utilizou",
	},
	"utiliza": {
		"usa",
		"utilizou",
		"fez",
		"usou",
		"feito",
		"desenvolvido",
	},
	"cargo": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"faculdade": {
		"cursa",
		"fiap",
		"inteligência",
		"artificial",
	},
	"estuda": {
		"cursa",
		"fiap",
	},
	"curso": {
		"cursa",
		"inteligência",
		"artificial",
	},
	"banco": {
		"auronix",
		"digital",
		"transacional",
	},
	"compressão": {
		"ggcompress",
		"gzip",
		"concorrente",
	},
	"compressao": {
		"ggcompress",
		"gzip",
		"concorrente",
	},
}
