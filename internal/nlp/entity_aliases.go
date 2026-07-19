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
			Value: "Etec de Guarulhos",
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
		Tokens: []string{"amazon", "sqs"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Amazon SQS",
		},
	},
	{
		Tokens: []string{"amazon", "eks"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Amazon EKS",
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
		Tokens: []string{"x", "tube"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "X Tube",
		},
	},
	{
		Tokens: []string{"xtube"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "X Tube",
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
		Tokens: []string{"auditex"},
		Entity: Entity{
			Type:  EntityProject,
			Value: "Auditex",
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
			Value: "Etec de Guarulhos",
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
		Tokens: []string{"react"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "React",
		},
	},
	{
		Tokens: []string{"next", "js"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Next.js",
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
		Tokens: []string{"kafka"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Kafka",
		},
	},
	{
		Tokens: []string{"sqs"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Amazon SQS",
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
		Tokens: []string{"docker"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Docker",
		},
	},
	{
		Tokens: []string{"terraform"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Terraform",
		},
	},
	{
		Tokens: []string{"kubernetes"},
		Entity: Entity{
			Type:  EntityTechnology,
			Value: "Kubernetes",
		},
	},
}
