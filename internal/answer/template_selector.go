package answer

import (
	"ai-agent/internal/nlp"
	"strings"
)

var templatesByLanguageAndIntent = map[nlp.Language]map[nlp.Intent][]string{
	nlp.LanguagePortuguese: {
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
	},
	nlp.LanguageEnglish: {
		nlp.IntentCurrentJob: {
			"Today, {subject} works as a mid-level developer at uFind Tecnologia.",
			"{subject} currently works at uFind Tecnologia as a mid-level developer.",
			"Right now, {subject} works as a mid-level developer at uFind Tecnologia.",
			"{fact}",
		},
		nlp.IntentFirstJob: {
			"{subject} started his career as a systems development intern at Representa Online.",
			"{subject}'s first professional experience was at Representa Online as an intern.",
			"Early in his career, {subject} worked as an intern at Representa Online.",
			"{fact}",
		},
		nlp.IntentEducation: {
			"{subject} studies Artificial Intelligence at FIAP.",
			"{subject}'s current education is focused on Artificial Intelligence at FIAP.",
			"{fact}",
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
			"{fact}",
		},
		nlp.IntentVisitorProjects: {
			"Among the portfolio projects, {fact}",
			"A good project summary is: {fact}",
			"{fact}",
		},
		nlp.IntentVisitorServices: {
			"In practice, {fact}",
			"For companies and projects, {fact}",
			"{fact}",
		},
		nlp.IntentHireReason: {
			"João is worth talking to because {fact}",
			"A good reason to talk to João is that {fact}",
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
