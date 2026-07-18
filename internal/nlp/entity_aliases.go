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
