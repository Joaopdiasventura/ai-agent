package answer

import (
	"ai-agent/internal/nlp"
	"strings"
)

var templatesByIntent = map[nlp.Intent][]string{
	nlp.IntentCurrentJob: {
		"Hoje, {subject} trabalha como desenvolvedor pleno na uFind Tecnologia.",
		"{subject} está hoje na uFind Tecnologia, onde atua como desenvolvedor pleno.",
		"No momento, {subject} trabalha na uFind Tecnologia como desenvolvedor pleno.",
		"{fact}",
	},
	nlp.IntentFirstJob: {
		"{subject} começou a carreira como estagiário na Representa Online.",
		"A primeira experiência profissional de {subject} foi na Representa Online, como estagiário.",
		"No início da carreira, {subject} trabalhou como estagiário na Representa Online.",
		"{fact}",
	},
	nlp.IntentEducation: {
		"{subject} estuda Inteligência Artificial na FIAP.",
		"Na formação, {subject} segue estudando Inteligência Artificial na FIAP.",
		"A formação atual de {subject} é voltada para Inteligência Artificial na FIAP.",
		"{fact}",
	},
	nlp.IntentProject: {
		"De forma simples: {fact}",
		"Sobre esse projeto: {fact}",
		"{fact}",
	},
	nlp.IntentTechnologies: {
		"{subject} trabalha principalmente com {technologies}.",
		"No dia a dia, {subject} usa tecnologias como {technologies}.",
		"As tecnologias mais presentes no trabalho de {subject} são {technologies}.",
	},
	nlp.IntentVisitorSummary: {
		"Em poucas palavras, {fact}",
		"De forma simples, {fact}",
		"{fact}",
	},
	nlp.IntentVisitorProjects: {
		"Entre os projetos do portfólio, {fact}",
		"Um bom resumo dos projetos é: {fact}",
		"{fact}",
	},
	nlp.IntentVisitorServices: {
		"Na prática, {fact}",
		"Para empresas e projetos, {fact}",
		"{fact}",
	},
	nlp.IntentHireReason: {
		"Vale conversar com João porque {fact}",
		"Um bom motivo para falar com João é que {fact}",
		"{fact}",
	},
}

func SelectTemplateForPlan(plan Plan) string {
	templates, exists := templatesByIntent[plan.Intent]

	if !exists || len(templates) == 0 {
		return "{fact}"
	}

	template := SelectTemplate(templates)

	if plan.DetailLevel == DetailMedium ||
		plan.DetailLevel == DetailDetailed {
		template = strings.ReplaceAll(
			template,
			"{fact}",
			"{facts}",
		)
	}

	return template
}
