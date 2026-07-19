package answer

import (
	"ai-agent/internal/nlp"
	"strings"
)

func RenderProjectRecommendation(plan Plan) string {
	projectName := plan.SelectedProject
	if projectName == "" {
		projectName = plan.Subject
	}

	if plan.Language == nlp.LanguageEnglish {
		return renderEnglishProjectRecommendation(plan, projectName)
	}

	return renderPortugueseProjectRecommendation(plan, projectName)
}

func renderPortugueseProjectRecommendation(plan Plan, projectName string) string {
	switch projectName {
	case "Auronix":
		return joinSentences([]string{
			portugueseProjectRecommendationOpening(plan.ProjectCriterion, projectName),
			findFact(plan.Facts, "plataforma financeira", "banco digital"),
			findFact(plan.Facts, "transferências confiáveis", "versionamento otimista", "ledger auditável"),
			"Ele demonstra capacidade de lidar com consistência, histórico verificável e infraestrutura cloud em um contexto financeiro.",
			technologyContext(plan, "Como contexto técnico, a solução usa %s."),
		})
	case "X Tube":
		return joinSentences([]string{
			"O X Tube é o projeto que melhor demonstra a capacidade de João Paulo de resolver problemas complexos.",
			findFact(plan.Facts, "equipe de três", "serviço de streaming"),
			findFact(plan.Facts, "liderança técnica", "upload", "processamento", "entrega"),
			findFact(plan.Facts, "receber eventos", "processar vídeos", "Amazon S3"),
			findFact(plan.Facts, "Amazon SQS", "Kafka", "Prometheus"),
			"Ele demonstra capacidade de coordenar várias etapas assíncronas, tornar um processo pesado mais acompanhável e dar visibilidade ao andamento da operação.",
			technologyContext(plan, "Como contexto técnico, a solução usa %s."),
		})
	case "GGCompress":
		return joinSentences([]string{
			portugueseProjectRecommendationOpening(plan.ProjectCriterion, projectName),
			findFact(plan.Facts, "compactar", "arquivar arquivos grandes"),
			findFact(plan.Facts, "chunks", "goroutines", "concorrente"),
			findFact(plan.Facts, "throughput", "1.23 GB/s", "9.77 GB"),
			"Ele demonstra cuidado com desempenho, concorrência e integridade de dados em arquivos grandes.",
			technologyContext(plan, "Como contexto técnico, a solução usa %s."),
		})
	case "Auditex":
		return joinSentences([]string{
			portugueseProjectRecommendationOpening(plan.ProjectCriterion, projectName),
			findFact(plan.Facts, "ledger financeiro", "blockchain centralizada"),
			findFact(plan.Facts, "detectar alterações", "validar a origem", "repetição indevida"),
			findFact(plan.Facts, "assinatura RSA", "nonce", "Merkle Root"),
			"Ele demonstra preocupação com auditabilidade, validação histórica e rastreabilidade de eventos financeiros.",
			technologyContext(plan, "Como contexto técnico, a solução usa %s."),
		})
	default:
		return FormatFacts(plan.Facts, plan.Language)
	}
}

func renderEnglishProjectRecommendation(plan Plan, projectName string) string {
	switch projectName {
	case "Auronix":
		return joinSentences([]string{
			englishProjectRecommendationOpening(plan.ProjectCriterion, projectName),
			findFact(plan.Facts, "financial platform", "digital banking"),
			findFact(plan.Facts, "reliable transfers", "optimistic versioning", "auditable ledger"),
			"It shows the ability to handle consistency, verifiable history, and cloud infrastructure in a financial context.",
			technologyContext(plan, "As technical context, the solution uses %s."),
		})
	case "X Tube":
		return joinSentences([]string{
			"X Tube is the project that best demonstrates João Paulo's ability to solve complex problems.",
			findFact(plan.Facts, "team of three", "streaming service"),
			findFact(plan.Facts, "technical leadership", "upload", "processing", "delivery"),
			findFact(plan.Facts, "receiving events", "processing videos", "Amazon S3"),
			findFact(plan.Facts, "Amazon SQS", "Kafka", "Prometheus"),
			"It shows the ability to coordinate several asynchronous steps, make a heavy process easier to follow, and provide visibility into the operation.",
			technologyContext(plan, "As technical context, the solution uses %s."),
		})
	case "GGCompress":
		return joinSentences([]string{
			englishProjectRecommendationOpening(plan.ProjectCriterion, projectName),
			findFact(plan.Facts, "compress", "archive large files"),
			findFact(plan.Facts, "chunks", "goroutines", "concurrently"),
			findFact(plan.Facts, "throughput", "1.23 GB/s", "9.77 GB"),
			"It shows attention to performance, concurrency, and data integrity when handling large files.",
			technologyContext(plan, "As technical context, the solution uses %s."),
		})
	case "Auditex":
		return joinSentences([]string{
			englishProjectRecommendationOpening(plan.ProjectCriterion, projectName),
			findFact(plan.Facts, "financial ledger", "centralized blockchain"),
			findFact(plan.Facts, "detect changes", "validate the origin", "transaction replay"),
			findFact(plan.Facts, "RSA signatures", "nonces", "Merkle Root"),
			"It shows concern for auditability, historical validation, and traceability of financial events.",
			technologyContext(plan, "As technical context, the solution uses %s."),
		})
	default:
		return FormatFacts(plan.Facts, plan.Language)
	}
}

func portugueseProjectRecommendationOpening(criterion nlp.ProjectCriterion, projectName string) string {
	switch criterion {
	case nlp.ProjectCriterionTechnicalCapability:
		return "O " + projectName + " é o projeto que melhor demonstra capacidade técnica de ponta a ponta"
	case nlp.ProjectCriterionGeneralRecommendation:
		return "O " + projectName + " é o projeto que eu destacaria para um recrutador"
	case nlp.ProjectCriterionFinancialSystems:
		return "O " + projectName + " é o projeto que melhor demonstra experiência com sistemas financeiros"
	case nlp.ProjectCriterionGoPerformance:
		return "O " + projectName + " é o projeto que melhor demonstra experiência com Go, desempenho e concorrência"
	case nlp.ProjectCriterionAuditability:
		return "O " + projectName + " é o projeto que melhor demonstra preocupação com auditabilidade"
	default:
		return "O " + projectName + " é o projeto que melhor responde a esse critério"
	}
}

func englishProjectRecommendationOpening(criterion nlp.ProjectCriterion, projectName string) string {
	switch criterion {
	case nlp.ProjectCriterionTechnicalCapability:
		return projectName + " is the project that best demonstrates end-to-end technical capability"
	case nlp.ProjectCriterionGeneralRecommendation:
		return projectName + " is the project I would highlight for a recruiter"
	case nlp.ProjectCriterionFinancialSystems:
		return projectName + " is the project that best demonstrates experience with financial systems"
	case nlp.ProjectCriterionGoPerformance:
		return projectName + " is the project that best demonstrates experience with Go, performance, and concurrency"
	case nlp.ProjectCriterionAuditability:
		return projectName + " is the project that best demonstrates concern for auditability"
	default:
		return projectName + " is the project that best matches this criterion"
	}
}

func findFact(facts []string, terms ...string) string {
	for _, fact := range facts {
		normalizedFact := strings.ToLower(fact)
		for _, term := range terms {
			if strings.Contains(normalizedFact, strings.ToLower(term)) {
				return fact
			}
		}
	}

	return ""
}

func technologyContext(plan Plan, template string) string {
	if len(plan.Technologies) == 0 {
		return ""
	}

	return strings.TrimSpace(strings.Replace(template, "%s", formatList(plan.Technologies, plan.Language), 1))
}

func joinSentences(sentences []string) string {
	selected := make([]string, 0, len(sentences))
	seen := make(map[string]struct{})

	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		sentence = strings.TrimRight(sentence, ".")

		if sentence == "" {
			continue
		}

		normalizedSentence := strings.ToLower(sentence)
		if _, exists := seen[normalizedSentence]; exists {
			continue
		}

		seen[normalizedSentence] = struct{}{}
		selected = append(selected, sentence+".")
	}

	return strings.Join(selected, " ")
}
