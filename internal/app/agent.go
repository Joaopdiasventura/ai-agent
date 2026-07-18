package app

import (
	"ai-agent/internal/answer"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/nlp"
	"ai-agent/internal/search"
)

var documents = knowledge.Documents()
var engine = search.NewEngine(documents, minimumSimilarity)

func AgentResponse(question string) (string, bool, nlp.Language) {
	searchResult := engine.Search(question, maximumSearchResults)

	if !searchResult.Found {
		return "", false, searchResult.Language
	}

	plan := answer.BuildPlan(searchResult)

	template := answer.SelectTemplateForPlan(plan)

	response := answer.RenderTemplate(template, plan)

	return response, true, searchResult.Language
}

func NotFoundMessage(language nlp.Language) string {
	if language == nlp.LanguageEnglish {
		return "I don't have that specific information, but I can talk about João's experience, projects, technologies, services, or contact details."
	}

	return "Não encontrei essa informação específica, mas posso falar sobre experiência, projetos, tecnologias, serviços ou contato do João."
}
