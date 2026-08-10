package confidence

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/generation"
	"ai-agent/internal/planning"
	"ai-agent/internal/ranking"
	"ai-agent/internal/reasoning"
	"ai-agent/internal/retrieval"
)

type Mode string

const (
	ModeClaim      Mode = "claim"
	ModeAbstention Mode = "abstention"
)

type Level string

const (
	LevelLow    Level = "low"
	LevelMedium Level = "medium"
	LevelHigh   Level = "high"
)

type SignalName string

const (
	SignalQueryQuality       SignalName = "query_quality"
	SignalRetrievalAgreement SignalName = "retrieval_agreement"
	SignalSeparation         SignalName = "separation"
	SignalEvidenceStrength   SignalName = "evidence_strength"
	SignalEvidenceDirectness SignalName = "evidence_directness"
	SignalSemanticCoverage   SignalName = "semantic_coverage"
	SignalEvidenceAbsence    SignalName = "evidence_absence"
	SignalPlanGrounding      SignalName = "plan_grounding"
	SignalAnswerGrounding    SignalName = "answer_grounding"
)

type Signal struct {
	Name       SignalName
	Score      float64
	Weight     float64
	Applicable bool
}

type Input struct {
	Query     domain.Query
	Retrieval retrieval.Result
	Ranking   ranking.Result
	Reasoning reasoning.Result
	Plan      planning.Plan
	Answer    generation.Answer
}

type Result struct {
	Score   float64
	Mode    Mode
	Level   Level
	Signals []Signal
}

func (r Result) Signal(
	name SignalName,
) (Signal, bool) {
	for _, signal := range r.Signals {
		if signal.Name == name {
			return signal, true
		}
	}

	return Signal{}, false
}

func clamp(
	value float64,
) float64 {
	if value < 0 {
		return 0
	}

	if value > 1 {
		return 1
	}

	return value
}

func confidenceLevel(
	score float64,
) Level {
	switch {
	case score >= 0.8:
		return LevelHigh

	case score >= 0.6:
		return LevelMedium

	default:
		return LevelLow
	}
}

func weightedScore(
	signals []Signal,
) float64 {
	total := 0.0
	totalWeight := 0.0

	for _, signal := range signals {
		if !signal.Applicable ||
			signal.Weight <= 0 {
			continue
		}

		total +=
			clamp(signal.Score) *
				signal.Weight

		totalWeight +=
			signal.Weight
	}

	if totalWeight == 0 {
		return 0
	}

	return clamp(
		total / totalWeight,
	)
}
