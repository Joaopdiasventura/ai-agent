package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"sort"
	"strings"
)

type projectCandidate struct {
	Name  string
	Score float64
	Docs  []*domain.Document
}

var projectNamesByPrefix = map[string]string{
	"project-auronix":    "Auronix",
	"project-xtube":      "X Tube",
	"project-ggcompress": "GGCompress",
	"project-auditex":    "Auditex",
}

var projectBaseScoreByCriterion = map[nlp.ProjectCriterion]map[string]float64{
	nlp.ProjectCriterionComplexProblem: {
		"X Tube":     8,
		"Auronix":    5,
		"GGCompress": 4,
		"Auditex":    4,
	},
	nlp.ProjectCriterionTechnicalCapability: {
		"Auronix":    8,
		"X Tube":     6,
		"GGCompress": 5,
		"Auditex":    4,
	},
	nlp.ProjectCriterionFinancialSystems: {
		"Auronix": 8,
		"Auditex": 5,
	},
	nlp.ProjectCriterionTechnicalLeadership: {
		"X Tube": 8,
	},
	nlp.ProjectCriterionGoPerformance: {
		"GGCompress": 8,
		"X Tube":     5,
	},
	nlp.ProjectCriterionAuditability: {
		"Auditex": 8,
		"Auronix": 3,
	},
	nlp.ProjectCriterionAsyncProcessing: {
		"X Tube": 8,
	},
	nlp.ProjectCriterionGeneralRecommendation: {
		"Auronix":    12,
		"X Tube":     6,
		"GGCompress": 5,
		"Auditex":    4,
	},
}

var projectCriterionTerms = map[nlp.ProjectCriterion][]string{
	nlp.ProjectCriterionComplexProblem: {
		"liderança", "lideranca", "leadership", "processamento", "processing",
		"assíncrono", "assincrono", "asynchronous", "upload", "entrega", "delivery",
		"nuvem", "cloud", "progresso", "progress", "observabilidade", "observability",
		"consistência", "consistency", "concorrente", "concurrent", "integridade", "integrity",
	},
	nlp.ProjectCriterionTechnicalCapability: {
		"full stack", "financeira", "financial", "banco digital", "digital banking",
		"consistência", "consistency", "ledger", "auditável", "auditable",
		"stateless", "kubernetes", "terraform", "aws eks", "cloud", "infraestrutura",
		"infrastructure", "spring boot", "angular", "postgresql", "rabbitmq", "redis",
	},
	nlp.ProjectCriterionFinancialSystems: {
		"financeira", "financeiro", "financial", "banco", "banking", "transferências",
		"transfers", "consistência", "consistency", "ledger", "cloud", "transacional",
		"transactional",
	},
	nlp.ProjectCriterionTechnicalLeadership: {
		"liderança", "lideranca", "leadership", "equipe de três", "team of three",
		"upload", "processamento", "processing", "entrega", "delivery",
	},
	nlp.ProjectCriterionGoPerformance: {
		"go", "concorrência", "concorrencia", "concurrency", "desempenho", "performance",
		"throughput", "benchmark", "9.77 gb", "1.23 gb/s", "integridade", "integrity",
	},
	nlp.ProjectCriterionAuditability: {
		"auditabilidade", "auditability", "criptografia", "cryptography", "histórico",
		"historical", "rsa", "nonce", "merkle", "assinatura", "signature",
		"alterações", "changes",
	},
	nlp.ProjectCriterionAsyncProcessing: {
		"assíncrono", "assincrono", "asynchronous", "streaming", "upload",
		"processamento", "processing", "sqs", "kafka", "prometheus", "progresso",
		"progress", "eventos", "events",
	},
	nlp.ProjectCriterionGeneralRecommendation: {
		"full stack", "financeira", "financial", "banco digital", "digital banking",
		"consistência", "consistency", "confiáveis", "reliable", "ledger", "auditável",
		"auditable", "stateless", "distribuída", "distributed", "kubernetes", "terraform",
		"aws eks", "cloud", "infraestrutura", "infrastructure",
	},
}

var preferredEvidenceByProjectAndCriterion = map[string]map[nlp.ProjectCriterion][]string{
	"X Tube": {
		nlp.ProjectCriterionComplexProblem:        {"description", "leadership", "processing", "async-progress"},
		nlp.ProjectCriterionTechnicalLeadership:   {"leadership", "description", "processing"},
		nlp.ProjectCriterionAsyncProcessing:       {"async-progress", "processing", "leadership"},
		nlp.ProjectCriterionGeneralRecommendation: {"leadership", "processing", "async-progress"},
	},
	"Auronix": {
		nlp.ProjectCriterionFinancialSystems:      {"description", "consistency", "architecture", "infrastructure"},
		nlp.ProjectCriterionComplexProblem:        {"description", "consistency", "architecture", "infrastructure"},
		nlp.ProjectCriterionTechnicalCapability:   {"description", "consistency", "architecture", "infrastructure"},
		nlp.ProjectCriterionGeneralRecommendation: {"description", "consistency", "architecture", "infrastructure"},
	},
	"GGCompress": {
		nlp.ProjectCriterionGoPerformance:  {"description", "concurrency", "integrity", "performance"},
		nlp.ProjectCriterionComplexProblem: {"description", "concurrency", "integrity", "performance"},
	},
	"Auditex": {
		nlp.ProjectCriterionAuditability:   {"description", "integrity", "validation", "wallet"},
		nlp.ProjectCriterionComplexProblem: {"description", "integrity", "validation"},
	},
}

func FindProjectRecommendationDocuments(
	documents []*domain.Document,
	analysis *nlp.QueryAnalysis,
	limit int,
) ([]Result, string) {
	if limit <= 0 {
		return []Result{}, ""
	}

	candidates := projectCandidates(documents, analysis)

	if len(candidates) == 0 {
		return []Result{}, ""
	}

	sort.Slice(candidates, func(firstIndex int, secondIndex int) bool {
		if candidates[firstIndex].Score == candidates[secondIndex].Score {
			return candidates[firstIndex].Name < candidates[secondIndex].Name
		}

		return candidates[firstIndex].Score > candidates[secondIndex].Score
	})

	selectedCandidate := candidates[0]
	selectedDocs := selectProjectEvidence(selectedCandidate, analysis.ProjectCriterion, limit)
	results := make([]Result, 0, len(selectedDocs))

	for index, document := range selectedDocs {
		results = append(results, Result{
			Document:   document,
			Similarity: selectedCandidate.Score - float64(index)*0.01,
		})
	}

	return results, selectedCandidate.Name
}

func projectCandidates(documents []*domain.Document, analysis *nlp.QueryAnalysis) []projectCandidate {
	candidatesByProject := make(map[string]*projectCandidate)
	criterionTerms := projectCriterionTerms[analysis.ProjectCriterion]

	for _, document := range documents {
		if document.Language != string(analysis.Language) {
			continue
		}

		projectName, ok := projectNameFromDocumentID(document.ID)
		if !ok {
			continue
		}

		candidate := candidatesByProject[projectName]
		if candidate == nil {
			candidate = &projectCandidate{Name: projectName}
			candidatesByProject[projectName] = candidate
		}

		candidate.Docs = append(candidate.Docs, document)
		candidate.Score += documentEvidenceScore(document, criterionTerms)
	}

	for projectName, score := range projectBaseScoreByCriterion[analysis.ProjectCriterion] {
		candidate := candidatesByProject[projectName]
		if candidate != nil {
			candidate.Score += score
		}
	}

	candidates := make([]projectCandidate, 0, len(candidatesByProject))
	for _, candidate := range candidatesByProject {
		candidates = append(candidates, *candidate)
	}

	return candidates
}

func documentEvidenceScore(document *domain.Document, criterionTerms []string) float64 {
	content := normalizeForBoost(document.Content)
	score := 0.0

	for _, term := range criterionTerms {
		if strings.Contains(content, normalizeForBoost(term)) {
			score += 1
		}
	}

	switch document.Category {
	case "impact":
		score += 1.5
	case "project":
		score += 1
	case "technology":
		score += 0.5
	}

	if strings.Contains(document.ID, "project-comparison-best") {
		score -= 8
	}

	return score
}

func selectProjectEvidence(candidate projectCandidate, criterion nlp.ProjectCriterion, limit int) []*domain.Document {
	preferredIDs := preferredEvidenceByProjectAndCriterion[candidate.Name][criterion]
	if len(preferredIDs) == 0 {
		preferredIDs = preferredEvidenceByProjectAndCriterion[candidate.Name][nlp.ProjectCriterionComplexProblem]
	}

	selected := make([]*domain.Document, 0, limit)
	seen := make(map[string]struct{})

	for _, idPart := range preferredIDs {
		for _, document := range candidate.Docs {
			if len(selected) >= limit {
				return selected
			}

			if _, exists := seen[document.ID]; exists {
				continue
			}

			if strings.Contains(document.ID, idPart) && !strings.Contains(document.ID, "project-comparison-best") {
				seen[document.ID] = struct{}{}
				selected = append(selected, document)
				break
			}
		}
	}

	sort.Slice(candidate.Docs, func(firstIndex int, secondIndex int) bool {
		return candidate.Docs[firstIndex].ID < candidate.Docs[secondIndex].ID
	})

	for _, document := range candidate.Docs {
		if len(selected) >= limit {
			break
		}

		if _, exists := seen[document.ID]; exists {
			continue
		}

		if strings.Contains(document.ID, "project-comparison-best") {
			continue
		}

		seen[document.ID] = struct{}{}
		selected = append(selected, document)
	}

	return selected
}

func projectNameFromDocumentID(documentID string) (string, bool) {
	for prefix, projectName := range projectNamesByPrefix {
		if strings.HasPrefix(documentID, prefix) {
			return projectName, true
		}
	}

	switch {
	case strings.Contains(documentID, "comparison-financial"):
		return "Auronix", true
	case strings.Contains(documentID, "comparison-streaming"):
		return "X Tube", true
	case strings.Contains(documentID, "comparison-go-performance"):
		return "GGCompress", true
	case strings.Contains(documentID, "comparison-auditability"):
		return "Auditex", true
	default:
		return "", false
	}
}
