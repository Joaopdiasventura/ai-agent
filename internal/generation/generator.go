package generation

import (
	"ai-agent/internal/domain"
	"context"
	"errors"
	"sort"
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
	if projectName != "" {
		facts = selectedProjectFacts(evidence, projectName, "pt")
	}

	if projectName != "" && containsAny(strings.ToLower(query.Text), []string{"melhor", "maior", "mais", "demonstra", "impacto", "liderança", "desempenho", "auditabilidade"}) {
		return join([]string{
			portugueseOpening(query, projectName),
			facts,
		})
	}

	return facts
}

func (generator GroundedGenerator) generateEnglish(query domain.Query, evidence []domain.Evidence) string {
	projectName := displayProject(evidence)
	facts := selectedFacts(evidence)
	if projectName != "" {
		facts = selectedProjectFacts(evidence, projectName, "en")
	}

	if projectName != "" && containsAny(strings.ToLower(query.Text), []string{"best", "most", "demonstrates", "impact", "leadership", "performance", "auditability"}) {
		return join([]string{
			englishOpening(query, projectName),
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

func selectedProjectFacts(evidence []domain.Evidence, projectName string, language string) string {
	ordered := append([]domain.Evidence(nil), evidence...)

	sort.SliceStable(ordered, func(firstIndex int, secondIndex int) bool {
		firstPriority := projectFactPriority(ordered[firstIndex].DocumentID)
		secondPriority := projectFactPriority(ordered[secondIndex].DocumentID)
		if firstPriority == secondPriority {
			if ordered[firstIndex].Score == ordered[secondIndex].Score {
				return ordered[firstIndex].DocumentID < ordered[secondIndex].DocumentID
			}

			return ordered[firstIndex].Score > ordered[secondIndex].Score
		}

		return firstPriority < secondPriority
	})

	sentences := make([]string, 0, len(ordered))
	seen := make(map[string]struct{})

	for index, item := range ordered {
		content := strings.TrimSpace(item.Content)
		content = strings.TrimRight(content, ".")
		if content == "" {
			continue
		}

		content = polishProjectFact(content, projectName, language, item.DocumentID, index)
		normalized := strings.ToLower(content)
		if _, exists := seen[normalized]; exists {
			continue
		}

		seen[normalized] = struct{}{}
		sentences = append(sentences, content+".")
	}

	return strings.Join(sentences, " ")
}

func projectFactPriority(documentID string) int {
	switch {
	case strings.Contains(documentID, "description"):
		return 0
	case strings.Contains(documentID, "leadership"), strings.Contains(documentID, "consistency"), strings.Contains(documentID, "concurrency"), strings.Contains(documentID, "integrity"):
		return 10
	case strings.Contains(documentID, "processing"):
		return 11
	case strings.Contains(documentID, "architecture"), strings.Contains(documentID, "async-progress"), strings.Contains(documentID, "validation"), strings.Contains(documentID, "performance"):
		return 20
	case strings.Contains(documentID, "infrastructure"), strings.Contains(documentID, "wallet"):
		return 30
	case strings.Contains(documentID, "technologies"):
		return 40
	default:
		return 50
	}
}

func polishProjectFact(content string, projectName string, language string, documentID string, index int) string {
	if language == "en" {
		content = strings.Replace(content, "In "+projectName+", João Paulo", "João Paulo", 1)
		content = strings.Replace(content, "In "+projectName+", ", englishProjectContext(projectName, documentID, index), 1)
		return content
	}

	content = strings.Replace(content, "No "+projectName+", João Paulo", "João Paulo", 1)
	content = strings.Replace(content, "Na "+projectName+", João Paulo", "João Paulo", 1)
	content = strings.Replace(content, "No "+projectName+", ", portugueseProjectContext(projectName, documentID, index), 1)
	content = strings.Replace(content, "Na "+projectName+", ", portugueseProjectContext(projectName, documentID, index), 1)

	return content
}

func portugueseProjectContext(projectName string, documentID string, index int) string {
	if strings.Contains(documentID, "infrastructure") {
		return "Na infraestrutura, "
	}
	if strings.Contains(documentID, "async-progress") {
		return "Para o processamento em segundo plano, "
	}
	if index == 0 {
		return projectName + " "
	}

	return "Além disso, "
}

func englishProjectContext(projectName string, documentID string, index int) string {
	if strings.Contains(documentID, "infrastructure") {
		return "In the infrastructure, "
	}
	if strings.Contains(documentID, "async-progress") {
		return "For background processing, "
	}
	if index == 0 {
		return projectName + " "
	}

	return "Additionally, "
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

func portugueseOpening(query domain.Query, projectName string) string {
	text := strings.ToLower(query.Text)

	switch {
	case containsAny(text, []string{"capacidade técnica", "capacidade tecnica"}):
		return "O " + projectName + " é o projeto mais alinhado ao critério de capacidade técnica."
	case containsAny(text, []string{"complexo", "complexos", "difícil", "dificil", "difíceis", "dificeis", "desafio"}):
		return "O " + projectName + " é o projeto que melhor demonstra capacidade de resolver problemas complexos."
	case containsAny(text, []string{"auditabilidade"}):
		return "O " + projectName + " é o projeto mais alinhado ao critério de auditabilidade e validação histórica."
	case containsAny(text, []string{"desempenho", "concorrência", "concorrencia"}):
		return "O " + projectName + " é o projeto mais alinhado ao critério de desempenho e concorrência."
	case containsAny(text, []string{"liderança", "lideranca"}):
		return "O " + projectName + " é o projeto mais alinhado ao critério de liderança técnica."
	default:
		return "O " + projectName + " é o projeto mais alinhado a esse critério."
	}
}

func englishOpening(query domain.Query, projectName string) string {
	text := strings.ToLower(query.Text)

	switch {
	case containsAny(text, []string{"technical capability"}):
		return projectName + " is the project that best matches the technical capability criterion."
	case containsAny(text, []string{"complex", "difficult", "challenge"}):
		return projectName + " is the project that best demonstrates the ability to solve complex problems."
	case containsAny(text, []string{"auditability"}):
		return projectName + " is the project that best matches the auditability and historical validation criterion."
	case containsAny(text, []string{"performance", "concurrency"}):
		return projectName + " is the project that best matches the performance and concurrency criterion."
	case containsAny(text, []string{"leadership"}):
		return projectName + " is the project that best matches the technical leadership criterion."
	default:
		return projectName + " is the project that best matches this criterion."
	}
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
