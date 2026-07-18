package nlp

import (
	"regexp"
	"strings"
)

var KnownTechnologies = []string{
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

func ExtractTechnologies(content string) []string {
	technologies := make([]string, 0)

	for _, technology := range KnownTechnologies {
		pattern := `(?i)(^|[^\p{L}\p{N}])` +
			regexp.QuoteMeta(technology) +
			`($|[^\p{L}\p{N}])`

		matched, err := regexp.MatchString(pattern, content)

		if err != nil || !matched {
			continue
		}

		technologies = append(technologies, strings.TrimSpace(technology))
	}

	return technologies
}
