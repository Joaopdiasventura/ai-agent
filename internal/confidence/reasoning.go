package confidence

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/reasoning"
)

func reasoningEvidenceStrength(
	result reasoning.Result,
) (float64, bool) {
	evidence :=
		result.Conclusion.Evidence

	if len(evidence) == 0 {
		return 0, false
	}

	weights := []float64{
		1,
		0.7,
		0.5,
		0.35,
		0.25,
	}

	total := 0.0
	totalWeight := 0.0

	for index, current := range evidence {
		weight := 0.15

		if index < len(weights) {
			weight =
				weights[index]
		}

		value :=
			clamp(
				0.75*current.Score +
					0.25*current.Importance,
			)

		total +=
			value * weight

		totalWeight +=
			weight
	}

	if totalWeight == 0 {
		return 0, false
	}

	return clamp(
		total / totalWeight,
	), true
}

func reasoningDirectness(
	result reasoning.Result,
) (float64, bool) {
	evidence :=
		result.Conclusion.Evidence

	if len(evidence) == 0 {
		return 0, false
	}

	limit := 4

	if limit > len(evidence) {
		limit =
			len(evidence)
	}

	total := 0.0
	totalWeight := 0.0

	for index := 0; index < limit; index++ {
		weight :=
			1 /
				float64(index+1)

		total +=
			clamp(
				evidence[index].
					Directness,
			) *
				weight

		totalWeight +=
			weight
	}

	if totalWeight == 0 {
		return 0, false
	}

	return clamp(
		total / totalWeight,
	), true
}

func semanticCoverage(
	currentQuery domain.Query,
	result reasoning.Result,
) (float64, bool) {
	if len(
		result.Conclusion.Groups,
	) > 0 {
		top :=
			result.Conclusion.Groups[0]

		return clamp(
			top.ConceptCoverage,
		), true
	}

	targetEntities :=
		directQueryEntities(
			currentQuery,
		)

	targetConcepts :=
		directQueryConcepts(
			currentQuery,
		)

	if len(targetEntities) == 0 &&
		len(targetConcepts) == 0 {
		if len(
			result.Conclusion.Evidence,
		) > 0 {
			return 0.6, true
		}

		return 0, false
	}

	entityMatches := make(
		map[domain.EntityID]struct{},
	)

	conceptMatches := make(
		map[domain.ConceptID]struct{},
	)

	for _, evidence := range result.Conclusion.Evidence {
		for _, entityID := range evidence.MatchedEntities {
			entityMatches[entityID] = struct{}{}
		}

		for _, conceptID := range evidence.MatchedConcepts {
			conceptMatches[conceptID] = struct{}{}
		}
	}

	total := 0
	matched := 0

	for _, entityID := range targetEntities {
		total++

		if _, exists :=
			entityMatches[entityID]; exists {
			matched++
		}
	}

	for _, conceptID := range targetConcepts {
		total++

		if _, exists :=
			conceptMatches[conceptID]; exists {
			matched++
		}
	}

	if total == 0 {
		return 0, false
	}

	if matched == 0 &&
		result.Conclusion.Status ==
			reasoning.SupportSupported &&
		len(
			result.Conclusion.Evidence,
		) > 0 {
		return 0.55, true
	}

	return clamp(
		float64(matched) /
			float64(total),
	), true
}

func evidenceAbsence(
	result reasoning.Result,
) float64 {
	if result.Conclusion.Status !=
		reasoning.SupportInsufficientEvidence {
		return 0
	}

	if len(
		result.Conclusion.Evidence,
	) == 0 {
		return 1
	}

	strongest := 0.0

	for _, evidence := range result.Conclusion.Evidence {
		value :=
			clamp(
				0.65*evidence.Score +
					0.35*evidence.Directness,
			)

		if value > strongest {
			strongest =
				value
		}
	}

	return clamp(
		1 - strongest,
	)
}

func directQueryEntities(
	currentQuery domain.Query,
) []domain.EntityID {
	result := make(
		[]domain.EntityID,
		0,
	)

	seen := make(
		map[domain.EntityID]struct{},
	)

	for _, entity := range currentQuery.Entities {
		if !entity.Explicit {
			continue
		}

		if _, exists :=
			seen[entity.EntityID]; exists {
			continue
		}

		seen[entity.EntityID] =
			struct{}{}

		result = append(
			result,
			entity.EntityID,
		)
	}

	return result
}

func directQueryConcepts(
	currentQuery domain.Query,
) []domain.ConceptID {
	result := make(
		[]domain.ConceptID,
		0,
	)

	seen := make(
		map[domain.ConceptID]struct{},
	)

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		if _, exists :=
			seen[concept.ConceptID]; exists {
			continue
		}

		seen[concept.ConceptID] =
			struct{}{}

		result = append(
			result,
			concept.ConceptID,
		)
	}

	return result
}
