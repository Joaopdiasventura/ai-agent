package confidence

import (
	"fmt"

	"ai-agent/internal/planning"
	"ai-agent/internal/reasoning"
)

type Calculator struct {
	config Config
}

func New() *Calculator {
	calculator, err :=
		NewWithConfig(
			DefaultConfig(),
		)

	if err != nil {
		panic(err)
	}

	return calculator
}

func NewWithConfig(
	config Config,
) (*Calculator, error) {
	if err :=
		config.Validate(); err != nil {
		return nil, err
	}

	return &Calculator{
		config: config,
	}, nil
}

func (c *Calculator) Assess(
	input Input,
) Result {
	mode :=
		confidenceMode(
			input,
		)

	if mode ==
		ModeAbstention {
		return c.assessAbstention(
			input,
		)
	}

	return c.assessClaim(
		input,
	)
}

func (c *Calculator) assessClaim(
	input Input,
) Result {
	agreement,
		agreementApplicable :=
		retrievalAgreement(
			input.Retrieval,
		)

	separation,
		separationApplicable :=
		rankingSeparation(
			input.Ranking,
			input.Reasoning,
		)

	strength,
		strengthApplicable :=
		reasoningEvidenceStrength(
			input.Reasoning,
		)

	directness,
		directnessApplicable :=
		reasoningDirectness(
			input.Reasoning,
		)

	coverage,
		coverageApplicable :=
		semanticCoverage(
			input.Query,
			input.Reasoning,
		)

	signals := []Signal{
		{
			Name: SignalQueryQuality,
			Score: queryQuality(
				input.Query,
			),
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalQueryQuality,
			),
			Applicable: true,
		},
		{
			Name:  SignalRetrievalAgreement,
			Score: agreement,
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalRetrievalAgreement,
			),
			Applicable: agreementApplicable,
		},
		{
			Name:  SignalSeparation,
			Score: separation,
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalSeparation,
			),
			Applicable: separationApplicable,
		},
		{
			Name:  SignalEvidenceStrength,
			Score: strength,
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalEvidenceStrength,
			),
			Applicable: strengthApplicable,
		},
		{
			Name:  SignalEvidenceDirectness,
			Score: directness,
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalEvidenceDirectness,
			),
			Applicable: directnessApplicable,
		},
		{
			Name:  SignalSemanticCoverage,
			Score: coverage,
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalSemanticCoverage,
			),
			Applicable: coverageApplicable,
		},
		{
			Name: SignalPlanGrounding,
			Score: planGrounding(
				input.Reasoning,
				input.Plan,
			),
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalPlanGrounding,
			),
			Applicable: true,
		},
		{
			Name: SignalAnswerGrounding,
			Score: answerGrounding(
				input.Plan,
				input.Answer,
			),
			Weight: signalWeight(
				c.config.ClaimWeights,
				SignalAnswerGrounding,
			),
			Applicable: true,
		},
	}

	score :=
		weightedScore(
			signals,
		)

	return Result{
		Score: score,
		Mode:  ModeClaim,
		Level: confidenceLevel(
			score,
		),
		Signals: signals,
	}
}

func (c *Calculator) assessAbstention(
	input Input,
) Result {
	signals := []Signal{
		{
			Name: SignalQueryQuality,
			Score: queryQuality(
				input.Query,
			),
			Weight: signalWeight(
				c.config.AbstentionWeights,
				SignalQueryQuality,
			),
			Applicable: true,
		},
		{
			Name: SignalEvidenceAbsence,
			Score: evidenceAbsence(
				input.Reasoning,
			),
			Weight: signalWeight(
				c.config.AbstentionWeights,
				SignalEvidenceAbsence,
			),
			Applicable: true,
		},
		{
			Name: SignalPlanGrounding,
			Score: planGrounding(
				input.Reasoning,
				input.Plan,
			),
			Weight: signalWeight(
				c.config.AbstentionWeights,
				SignalPlanGrounding,
			),
			Applicable: true,
		},
		{
			Name: SignalAnswerGrounding,
			Score: answerGrounding(
				input.Plan,
				input.Answer,
			),
			Weight: signalWeight(
				c.config.AbstentionWeights,
				SignalAnswerGrounding,
			),
			Applicable: true,
		},
	}

	score :=
		weightedScore(
			signals,
		)

	return Result{
		Score: score,
		Mode:  ModeAbstention,
		Level: confidenceLevel(
			score,
		),
		Signals: signals,
	}
}

func confidenceMode(
	input Input,
) Mode {
	if input.Plan.Status ==
		planning.PlanStatusAbstain {
		return ModeAbstention
	}

	if input.Reasoning.Conclusion.Status ==
		reasoning.SupportInsufficientEvidence {
		return ModeAbstention
	}

	return ModeClaim
}

func (c *Calculator) Validate(
	input Input,
) error {
	if input.Query.Original == "" {
		return fmt.Errorf(
			"query is required",
		)
	}

	if input.Plan.Language == "" {
		return fmt.Errorf(
			"plan language is required",
		)
	}

	return nil
}
