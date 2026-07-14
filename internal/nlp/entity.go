package nlp

import "strings"

type EntityType string

const (
	EntityProject     EntityType = "project"
	EntityCompany     EntityType = "company"
	EntityInstitution EntityType = "institution"
	EntityTechnology  EntityType = "technology"
)

type Entity struct {
	Type  EntityType
	Value string
}

var entityAliases = map[string]Entity{
	"auronix": {
		Type:  EntityProject,
		Value: "Auronix",
	},
	"modularis": {
		Type:  EntityProject,
		Value: "Modularis",
	},
	"votrix": {
		Type:  EntityProject,
		Value: "Votrix",
	},
	"ggcompress": {
		Type:  EntityProject,
		Value: "GGCompress",
	},
	"ggc": {
		Type:  EntityProject,
		Value: "GGCompress",
	},
	"vox": {
		Type:  EntityProject,
		Value: "VOX",
	},
	"etecfy": {
		Type:  EntityProject,
		Value: "Etecfy",
	},
	"ufind": {
		Type:  EntityCompany,
		Value: "uFind Tecnologia",
	},
	"ufind tecnologia": {
		Type:  EntityCompany,
		Value: "uFind Tecnologia",
	},
	"representa": {
		Type:  EntityCompany,
		Value: "Representa Online",
	},
	"representa online": {
		Type:  EntityCompany,
		Value: "Representa Online",
	},
	"fiap": {
		Type:  EntityInstitution,
		Value: "FIAP",
	},
	"etec": {
		Type:  EntityInstitution,
		Value: "Etec Guarulhos",
	},
	"etec guarulhos": {
		Type:  EntityInstitution,
		Value: "Etec Guarulhos",
	},
	"go": {
		Type:  EntityTechnology,
		Value: "Go",
	},
	"golang": {
		Type:  EntityTechnology,
		Value: "Go",
	},
	"java": {
		Type:  EntityTechnology,
		Value: "Java",
	},
	"angular": {
		Type:  EntityTechnology,
		Value: "Angular",
	},
	"nestjs": {
		Type:  EntityTechnology,
		Value: "NestJS",
	},
	"redis": {
		Type:  EntityTechnology,
		Value: "Redis",
	},
	"rabbitmq": {
		Type:  EntityTechnology,
		Value: "RabbitMQ",
	},
	"postgresql": {
		Type:  EntityTechnology,
		Value: "PostgreSQL",
	},
	"postgres": {
		Type:  EntityTechnology,
		Value: "PostgreSQL",
	},
}

func DetecEntity(tokens []string) (Entity, bool) {
	if len(tokens) == 0 {
		return Entity{}, false
	}

	text := strings.Join(tokens, " ")

	for alias, entity := range entityAliases {
		if strings.Contains(text, alias) {
			return entity, true
		}
	}

	return Entity{}, false
}
