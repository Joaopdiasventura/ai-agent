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
	"React",
	"Next.js",
	"NestJS",
	"Spring Boot",
	"PostgreSQL",
	"MongoDB",
	"Redis",
	"RabbitMQ",
	"Kafka",
	"Amazon SQS",
	"AWS",
	"Docker",
	"Terraform",
	"Kubernetes",
	"Amazon EKS",
	"ECS",
	"S3",
	"IAM",
	"FFmpeg",
	"Prometheus",
	"IndexedDB",
	"goroutines",
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
