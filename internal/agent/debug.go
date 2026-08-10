package agent

import (
	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
	"ai-agent/internal/generation"
	"ai-agent/internal/planning"
	"ai-agent/internal/ranking"
	"ai-agent/internal/reasoning"
	"ai-agent/internal/retrieval"
)

type DebugResult struct {
	Question   string            `json:"question"`
	Result     Result            `json:"result"`
	Query      domain.Query      `json:"query"`
	Retrieval  retrieval.Result  `json:"retrieval"`
	Ranking    ranking.Result    `json:"ranking"`
	Reasoning  reasoning.Result  `json:"reasoning"`
	Plan       planning.Plan     `json:"plan"`
	Generation generation.Answer `json:"generation"`
	Confidence confidence.Result `json:"confidence"`
}

func (s *Service) Debug(
	question string,
) (DebugResult, error) {
	currentExecution, err :=
		s.execute(question)

	if err != nil {
		return DebugResult{}, err
	}

	hasResponse :=
		executionHasResponse(
			currentExecution,
		)

	response := ""

	if hasResponse {
		response =
			currentExecution.answer.Text
	}

	return DebugResult{
		Question: currentExecution.question,
		Result: Result{
			Response:        response,
			HasResponse:     hasResponse,
			Language:        currentExecution.answer.Language,
			Confidence:      currentExecution.confidence.Score,
			ConfidenceLevel: currentExecution.confidence.Level,
		},
		Query:      currentExecution.query,
		Retrieval:  currentExecution.retrieval,
		Ranking:    currentExecution.ranking,
		Reasoning:  currentExecution.reasoning,
		Plan:       currentExecution.plan,
		Generation: currentExecution.answer,
		Confidence: currentExecution.confidence,
	}, nil
}
