package answer

import (
	"regexp"
	"strings"
)

func ExtractTechnologies(content string) []string {
	technologies := make([]string, 0)

	for _, technology := range knownTechnologies {
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
