package confidence

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/generation"
	"ai-agent/internal/planning"
	"ai-agent/internal/reasoning"
)

func planGrounding(
	reasoned reasoning.Result,
	plan planning.Plan,
) float64 {
	if reasoned.Conclusion.Status ==
		reasoning.SupportInsufficientEvidence {
		if plan.Status !=
			planning.PlanStatusAbstain {
			return 0
		}

		if len(
			plan.FactIDs(),
		) != 0 {
			return 0
		}

		return 1
	}

	if reasoned.Conclusion.Status !=
		reasoning.SupportSupported {
		return 0.25
	}

	if plan.Status !=
		planning.PlanStatusReady {
		return 0
	}

	allowed :=
		reasoningFactSet(
			reasoned,
		)

	planned :=
		plan.FactIDs()

	if len(planned) == 0 {
		return 0.35
	}

	matched := 0

	for _, factID := range planned {
		if _, exists :=
			allowed[factID]; exists {
			matched++
		}
	}

	return clamp(
		float64(matched) /
			float64(len(planned)),
	)
}

func answerGrounding(
	plan planning.Plan,
	answer generation.Answer,
) float64 {
	if answer.Empty() {
		return 0
	}

	planned :=
		factSet(
			plan.FactIDs(),
		)

	if plan.Status ==
		planning.PlanStatusAbstain {
		if len(
			answer.FactIDs,
		) != 0 {
			return 0
		}

		return 1
	}

	if len(answer.FactIDs) == 0 {
		if len(planned) == 0 {
			return 0.5
		}

		return 0.2
	}

	matched := 0

	for _, factID := range answer.FactIDs {
		if _, exists :=
			planned[factID]; exists {
			matched++
		}
	}

	return clamp(
		float64(matched) /
			float64(len(answer.FactIDs)),
	)
}

func reasoningFactSet(
	result reasoning.Result,
) map[domain.FactID]struct{} {
	values := make(
		map[domain.FactID]struct{},
	)

	for _, evidence := range result.Conclusion.Evidence {
		values[evidence.FactID] = struct{}{}
	}

	for _, group := range result.Conclusion.Groups {
		for _, evidence := range group.Evidence {
			values[evidence.FactID] = struct{}{}
		}
	}

	return values
}

func factSet(
	values []domain.FactID,
) map[domain.FactID]struct{} {
	result := make(
		map[domain.FactID]struct{},
	)

	for _, value := range values {
		if value == "" {
			continue
		}

		result[value] =
			struct{}{}
	}

	return result
}
