package app

import (
	"ai-agent/internal/answer"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/search"
)

var documents = knowledge.Documents()
var engine = search.NewEngine(documents, minimumSimilarity)

func AgentResponse(question string) (string, bool) {
	searchResult := engine.Search(question, maximumSearchResults)

	if !searchResult.Found {
		return "", false
	}

	plan := answer.BuildPlan(searchResult.Tokens, searchResult.Intent, searchResult.Results)

	template := answer.SelectTemplateForPlan(plan)

	response := answer.RenderTemplate(template, plan)

	return response, true
}
