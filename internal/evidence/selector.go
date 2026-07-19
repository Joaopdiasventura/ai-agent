package evidence

import (
	"ai-agent/internal/domain"
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
	if !requiresSynthesis(query) {
		maxEvidence = 1
	}

	evidence := make([]domain.Evidence, 0, maxEvidence)
	seenIDs := make(map[string]struct{})
	selectedProject := query.Project

	for _, result := range results {
		if result.Document == nil {
			continue
		}

		document := result.Document

		if query.Language != "" && document.Language != query.Language {
			continue
		}

		if selectedProject != "" && document.Project != "" && document.Project != selectedProject {
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
