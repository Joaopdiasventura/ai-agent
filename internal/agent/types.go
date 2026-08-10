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

type Result struct {
	Response        string           `json:"response"`
	HasResponse     bool             `json:"has_response"`
	Language        domain.Language  `json:"language"`
	Confidence      float64          `json:"confidence"`
	ConfidenceLevel confidence.Level `json:"confidence_level"`
}

type execution struct {
	question   string
	query      domain.Query
	retrieval  retrieval.Result
	ranking    ranking.Result
	reasoning  reasoning.Result
	plan       planning.Plan
	answer     generation.Answer
	confidence confidence.Result
}
