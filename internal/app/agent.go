package app

import (
	"ai-agent/internal/agent"
	"ai-agent/internal/embedding"
	"ai-agent/internal/generation"
	"ai-agent/internal/ranking"
	"ai-agent/internal/retrieval"
	"ai-agent/internal/vectorindex"
	"context"
)

var agentIndex = mustLoadAgentIndex()
var service = mustCreateAgentService(agentIndex)

func AgentResponse(question string) (string, bool, string) {
	response, hasResponse, language, err := service.Answer(context.Background(), question, maximumSearchResults)
	if err != nil {
		return "", false, "pt"
	}

	return response, hasResponse, language
}

func NotFoundMessage(language string) string {
	if language == "en" {
		return "I don't have that specific information, but I can talk about João's experience, projects, technologies, services, or contact details."
	}

	return "Não encontrei essa informação específica, mas posso falar sobre experiência, projetos, tecnologias, serviços ou contato do João."
}

func DocumentCount() int {
	return len(agentIndex.Entries)
}

func mustLoadAgentIndex() vectorindex.Index {
	index, err := vectorindex.LoadEmbedded()
	if err != nil {
		panic(err)
	}

	return index
}

func mustCreateAgentService(index vectorindex.Index) *agent.Service {
	embedder, err := embedding.NewDeterministicEmbedder(index.Dimension)
	if err != nil {
		panic(err)
	}

	service, err := agent.NewService(
		embedder,
		retrieval.NewHybridRetriever(index),
		ranking.DefaultMetadataReranker(),
		generation.NewGroundedGenerator(),
	)
	if err != nil {
		panic(err)
	}

	return service
}
