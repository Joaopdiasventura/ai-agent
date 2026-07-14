package answer

import (
	"ai-agent/internal/nlp"
	"ai-agent/internal/search"
	"fmt"
	"strings"
)

var knownTechnologies = []string{
	"Node.js",
	"TypeScript",
	"Java",
	"Go",
	"Angular",
	"NestJS",
	"Spring Boot",
	"Fastify",
	"PostgreSQL",
	"MongoDB",
	"Redis",
	"RabbitMQ",
	"BullMQ",
	"SSE",
	"WebSocket",
	"Docker Compose",
	"nginx",
	"AsyncAPI",
	"Capacitor",
	"Tauri",
	"goroutines",
	"channels",
	"gzip",
	"SHA-256",
}

func Generate(result search.Result) string {
	switch result.Intent {
	case nlp.IntentTechnologies:
		documentContent := strings.ToLower(result.Document.Content)
		technologies := make([]string, 0)

		for _, technology := range knownTechnologies {
			if strings.Contains(documentContent, strings.ToLower(technology)) {
				technologies = append(technologies, technology)
			}
		}

		if len(technologies) == 0 {
			return result.Document.Content
		}

		subject := "João Paulo"

		if result.HasEntity {
			subject = result.Entity.Value
		}

		return fmt.Sprintf("As tecnologias relacionadas a %s são %s.", subject, strings.Join(technologies, ", "))

	case nlp.IntentCurrentJob:
		return result.Document.Content

	case nlp.IntentFirstJob:
		return result.Document.Content

	case nlp.IntentEducation:
		return result.Document.Content

	case nlp.IntentProject:
		return result.Document.Content

	default:
		return result.Document.Content
	}
}
