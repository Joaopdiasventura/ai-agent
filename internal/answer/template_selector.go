package answer

import (
	"ai-agent/internal/memory"
	"ai-agent/internal/nlp"
	"strings"
)

var templatesByIntent = map[nlp.Intent][]string{
	nlp.IntentCurrentJob: {
		"{fact}",
		"Atualmente, {subject} atua como desenvolvedor pleno na uFind Tecnologia.",
		"Hoje, {subject} trabalha como desenvolvedor pleno na uFind Tecnologia.",
	},
	nlp.IntentFirstJob: {
		"{fact}",
		"{subject} iniciou sua carreira como estagiário na Representa Online.",
		"A primeira experiência profissional de {subject} foi na Representa Online.",
	},
	nlp.IntentEducation: {
		"{fact}",
		"Atualmente, {subject} estuda Inteligência Artificial na FIAP.",
		"A formação atual de {subject} é em Inteligência Artificial na FIAP.",
	},
	nlp.IntentProject: {
		"{fact}",
		"Em resumo, {fact}",
		"De forma direta: {fact}",
	},
	nlp.IntentTechnologies: {
		"{subject} utiliza principalmente {technologies}.",
		"A stack de {subject} inclui {technologies}.",
		"As principais tecnologias relacionadas a {subject} são {technologies}.",
	},
	nlp.IntentVisitorSummary: {
		"{fact}",
		"Em resumo, {fact}",
	},
	nlp.IntentVisitorProjects: {
		"{fact}",
		"Entre os projetos, {fact}",
	},
	nlp.IntentVisitorServices: {
		"{fact}",
		"Na prática, {fact}",
	},
	nlp.IntentHireReason: {
		"{fact}",
		"Vale conversar com João porque {fact}",
	},
}

func SelectTemplateForPlan(plan Plan, session *memory.Session) string {
	templates, exists := templatesByIntent[plan.Intent]

	if !exists || len(templates) == 0 {
		return "{fact}"
	}

	lastIndex := session.GetLastTemplateIndex(plan.Intent)

	template, selectedIndex := SelectTemplate(templates, lastIndex)

	if plan.DetailLevel == DetailMedium ||
		plan.DetailLevel == DetailDetailed {
		template = strings.ReplaceAll(
			template,
			"{fact}",
			"{facts}",
		)
	}

	if selectedIndex >= 0 {
		session.SetLastTemplateIndex(plan.Intent, selectedIndex)
	}

	return template
}
