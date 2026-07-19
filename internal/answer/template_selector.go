package answer

import (
	"ai-agent/internal/nlp"
	"strings"
)

var templatesByLanguageAndIntent = map[nlp.Language]map[nlp.Intent][]string{
	nlp.LanguagePortuguese: {
		nlp.IntentCurrentJob: {
			"{facts}",
		},
		nlp.IntentFirstJob: {
			"{facts}",
		},
		nlp.IntentEducation: {
			"{facts}",
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
			"Para quem está conhecendo o João agora: {fact}",
			"O essencial sobre João é isto: {fact}",
			"{fact}",
		},
		nlp.IntentVisitorProjects: {
			"{fact}",
		},
		nlp.IntentVisitorServices: {
			"Na prática, {fact}",
			"Para empresas e projetos, {fact}",
			"Em termos de entrega, {fact}",
			"Para transformar uma necessidade em software, {fact}",
			"{fact}",
		},
		nlp.IntentHireReason: {
			"Vale conversar com João porque {fact}",
			"Um bom motivo para falar com João é que {fact}",
			"Para uma conversa profissional, o ponto forte é que {fact}",
			"O diferencial aqui é que {fact}",
			"{fact}",
		},
	},
	nlp.LanguageEnglish: {
		nlp.IntentCurrentJob: {
			"{facts}",
		},
		nlp.IntentFirstJob: {
			"{facts}",
		},
		nlp.IntentEducation: {
			"{facts}",
		},
		nlp.IntentProject: {
			"In simple terms: {fact}",
			"About this project: {fact}",
			"{fact}",
		},
		nlp.IntentTechnologies: {
			"{subject} mainly works with {technologies}.",
			"In day-to-day work, {subject} uses technologies like {technologies}.",
			"The technologies most present in {subject}'s work are {technologies}.",
		},
		nlp.IntentVisitorSummary: {
			"In short, {fact}",
			"Simply put, {fact}",
			"If you are just getting to know João, {fact}",
			"The quick version is this: {fact}",
			"{fact}",
		},
		nlp.IntentVisitorProjects: {
			"{fact}",
		},
		nlp.IntentVisitorServices: {
			"In practice, {fact}",
			"For companies and projects, {fact}",
			"In terms of delivery, {fact}",
			"When a business need has to become software, {fact}",
			"{fact}",
		},
		nlp.IntentHireReason: {
			"João is worth talking to because {fact}",
			"A good reason to talk to João is that {fact}",
			"For a professional conversation, the strong point is that {fact}",
			"The useful signal here is that {fact}",
			"{fact}",
		},
	},
}

func SelectTemplateForPlan(plan Plan) string {
	templatesByIntent, exists := templatesByLanguageAndIntent[plan.Language]
	if !exists {
		templatesByIntent = templatesByLanguageAndIntent[nlp.LanguagePortuguese]
	}

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
