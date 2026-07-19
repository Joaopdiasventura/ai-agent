package evidence

import (
	"ai-agent/internal/domain"
	"sort"
	"strings"
)

type Selector struct {
	MaxEvidence int
}

func DefaultSelector() Selector {
	return Selector{MaxEvidence: 4}
}

func (selector Selector) Select(query domain.Query, results []domain.SearchResult) []domain.Evidence {
	if selector.MaxEvidence <= 0 || len(results) == 0 {
		return []domain.Evidence{}
	}

	maxEvidence := selector.MaxEvidence
	synthesis := requiresSynthesis(query)
	if !synthesis {
		maxEvidence = 1
	}

	if synthesis && query.Project != "" {
		results = prioritizeProjectEvidence(results)
	}

	evidence := make([]domain.Evidence, 0, maxEvidence)
	seenIDs := make(map[string]struct{})
	selectedProject := query.Project

	for _, result := range results {
		if result.Document == nil {
			continue
		}

		document := result.Document

		if strings.Contains(document.ID, "project-comparison") {
			continue
		}

		if query.Language != "" && document.Language != query.Language {
			continue
		}

		if selectedProject != "" && document.Project != "" && document.Project != selectedProject {
			continue
		}

		if selectedProject != "" && requiresSynthesis(query) && query.Category == "project" && document.Project == "" {
			continue
		}

		if selectedProject == "" && len(evidence) > 0 && evidence[0].Project != "" && document.Project != "" && document.Project != evidence[0].Project {
			continue
		}

		if _, exists := seenIDs[document.ID]; exists {
			continue
		}

		if selectedProject == "" && document.Project != "" {
			selectedProject = document.Project
		}

		seenIDs[document.ID] = struct{}{}
		evidence = append(evidence, domain.Evidence{
			DocumentID:     document.ID,
			Language:       document.Language,
			Category:       document.Category,
			Project:        document.Project,
			TemporalStatus: document.TemporalStatus,
			Content:        document.Content,
			Score:          result.FinalScore,
			Sources:        append([]string(nil), result.Sources...),
		})

		if len(evidence) >= maxEvidence {
			break
		}
	}

	return evidence
}

func requiresSynthesis(query domain.Query) bool {
	text := strings.ToLower(query.Text)
	synthesisSignals := []string{
		"compar", "melhor", "maior", "mais", "demonstra", "impacto", "porque", "por que",
		"complex", "liderança", "lideranca", "desempenho", "concorr", "auditabilidade",
		"compare", "best", "most", "why", "impact", "leadership", "performance", "auditability",
	}

	for _, signal := range synthesisSignals {
		if strings.Contains(text, signal) {
			return true
		}
	}

	return false
}

func prioritizeProjectEvidence(results []domain.SearchResult) []domain.SearchResult {
	prioritized := append([]domain.SearchResult(nil), results...)

	sort.SliceStable(prioritized, func(firstIndex int, secondIndex int) bool {
		firstPriority := evidencePriority(prioritized[firstIndex])
		secondPriority := evidencePriority(prioritized[secondIndex])

		if firstPriority == secondPriority {
			return prioritized[firstIndex].FinalScore > prioritized[secondIndex].FinalScore
		}

		return firstPriority < secondPriority
	})

	return prioritized
}

func evidencePriority(result domain.SearchResult) int {
	if result.Document == nil {
		return 99
	}

	switch result.Document.Category {
	case "impact":
		return 0
	case "project":
		return 1
	case "technology":
		return 2
	default:
		return 3
	}
}
