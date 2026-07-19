package generation

import (
	"ai-agent/internal/domain"
	"context"
	"errors"
	"strings"
)

type GroundedGenerator struct{}

var ErrNoEvidence = errors.New("no evidence available for generation")

func NewGroundedGenerator() GroundedGenerator {
	return GroundedGenerator{}
}

func (generator GroundedGenerator) Generate(ctx context.Context, query domain.Query, evidence []domain.Evidence) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	if len(evidence) == 0 {
		return "", ErrNoEvidence
	}

	if isDirectAnswer(query, evidence) {
		return evidence[0].Content, nil
	}

	if query.Language == "en" {
		return generator.generateEnglish(query, evidence), nil
	}

	return generator.generatePortuguese(query, evidence), nil
}

func isDirectAnswer(query domain.Query, evidence []domain.Evidence) bool {
	if len(evidence) != 1 {
		return false
	}

	if query.Category == "contact" || query.Category == "education" || query.Category == "career" || query.Category == "certificate" {
		return true
	}

	text := strings.ToLower(query.Text)
	return !containsAny(text, []string{"melhor", "maior", "mais", "compar", "impacto", "why", "best", "most", "compare", "impact"})
}

func (generator GroundedGenerator) generatePortuguese(query domain.Query, evidence []domain.Evidence) string {
	projectName := displayProject(evidence)
	facts := selectedFacts(evidence)

	if projectName != "" && containsAny(strings.ToLower(query.Text), []string{"melhor", "maior", "mais", "demonstra", "impacto", "liderança", "desempenho", "auditabilidade"}) {
		return join([]string{
			"O " + projectName + " é o projeto mais alinhado a esse critério.",
			facts,
		})
	}

	return facts
}

func (generator GroundedGenerator) generateEnglish(query domain.Query, evidence []domain.Evidence) string {
	projectName := displayProject(evidence)
	facts := selectedFacts(evidence)

	if projectName != "" && containsAny(strings.ToLower(query.Text), []string{"best", "most", "demonstrates", "impact", "leadership", "performance", "auditability"}) {
		return join([]string{
			projectName + " is the project that best matches this criterion.",
			facts,
		})
	}

	return facts
}

func selectedFacts(evidence []domain.Evidence) string {
	sentences := make([]string, 0, len(evidence))
	seen := make(map[string]struct{})

	for _, item := range evidence {
		content := strings.TrimSpace(item.Content)
		content = strings.TrimRight(content, ".")
		if content == "" {
			continue
		}

		normalized := strings.ToLower(content)
		if _, exists := seen[normalized]; exists {
			continue
		}

		seen[normalized] = struct{}{}
		sentences = append(sentences, content+".")
	}

	return strings.Join(sentences, " ")
}

func displayProject(evidence []domain.Evidence) string {
	for _, item := range evidence {
		switch item.Project {
		case "auronix":
			return "Auronix"
		case "x-tube":
			return "X Tube"
		case "ggcompress":
			return "GGCompress"
		case "auditex":
			return "Auditex"
		}
	}

	return ""
}

func join(parts []string) string {
	selected := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		selected = append(selected, part)
	}

	return strings.Join(selected, " ")
}

func containsAny(text string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(text, term) {
			return true
		}
	}

	return false
}
