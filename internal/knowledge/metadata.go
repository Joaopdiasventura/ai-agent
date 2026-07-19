package knowledge

import (
	"ai-agent/internal/domain"
	"regexp"
	"sort"
	"strings"
)

var keywordSeparator = regexp.MustCompile(`[^\p{L}\p{N}+#./-]+`)

func applyDocumentMetadata(document *domain.Document) {
	document.Subject = documentSubject(document.ID, document.Category)
	document.Project = documentProject(document.ID, document.Content)
	document.TemporalStatus = documentTemporalStatus(document.ID, document.Content)
	document.Keywords = documentKeywords(*document)
}

func documentSubject(id string, category string) string {
	id = strings.TrimSuffix(strings.TrimSuffix(id, "-pt"), "-en")
	parts := strings.Split(id, "-")

	if len(parts) <= 1 {
		return category
	}

	return strings.Join(parts[1:], " ")
}

func documentProject(id string, content string) string {
	normalized := strings.ToLower(id + " " + content)

	switch {
	case strings.Contains(normalized, "auronix"):
		return "auronix"
	case strings.Contains(normalized, "x tube"), strings.Contains(normalized, "xtube"):
		return "x-tube"
	case strings.Contains(normalized, "ggcompress"):
		return "ggcompress"
	case strings.Contains(normalized, "auditex"):
		return "auditex"
	default:
		return ""
	}
}

func documentTemporalStatus(id string, content string) string {
	normalizedID := strings.ToLower(id)
	normalized := strings.ToLower(id + " " + content)

	switch {
	case strings.Contains(normalizedID, "education-fiap"):
		return domain.TemporalFuture
	case strings.Contains(normalizedID, "career-current"):
		return domain.TemporalCurrent
	case strings.Contains(normalizedID, "career-junior"), strings.Contains(normalizedID, "career-intern"), strings.Contains(normalizedID, "education-etec"):
		return domain.TemporalPast
	case strings.Contains(normalized, "planned") && strings.Contains(normalized, "education"):
		return domain.TemporalFuture
	default:
		return domain.TemporalTimeless
	}
}

func documentKeywords(document domain.Document) []string {
	seen := make(map[string]struct{})
	keywords := make([]string, 0)

	add := func(value string) {
		value = strings.TrimSpace(strings.ToLower(value))
		if value == "" {
			return
		}

		if _, exists := seen[value]; exists {
			return
		}

		seen[value] = struct{}{}
		keywords = append(keywords, value)
	}

	add(document.ID)
	add(document.Language)
	add(document.Category)
	add(document.Subject)
	add(document.Project)
	add(document.TemporalStatus)

	for _, token := range keywordSeparator.Split(document.Content, -1) {
		if len([]rune(token)) < 3 {
			continue
		}

		add(token)
	}

	sort.Strings(keywords)

	return keywords
}
