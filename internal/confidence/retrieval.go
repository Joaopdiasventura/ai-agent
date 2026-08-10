package confidence

import (
	"math"

	"ai-agent/internal/domain"
	"ai-agent/internal/ranking"
	"ai-agent/internal/reasoning"
	"ai-agent/internal/retrieval"
)

const agreementDepth = 5

func retrievalAgreement(
	result retrieval.Result,
) (float64, bool) {
	active := make(
		[]retrieval.Ranking,
		0,
	)

	for _, current := range result.Rankings {
		if len(
			current.Candidates,
		) == 0 {
			continue
		}

		active = append(
			active,
			current,
		)
	}

	if len(active) == 0 {
		return 0, false
	}

	if len(active) == 1 {
		return 0.5, true
	}

	topConsensus :=
		retrievalTopConsensus(
			active,
		)

	overlap :=
		retrievalPairwiseOverlap(
			active,
		)

	return clamp(
		0.65*topConsensus +
			0.35*overlap,
	), true
}

func retrievalTopConsensus(
	rankings []retrieval.Ranking,
) float64 {
	counts := make(
		map[domain.FactID]int,
	)

	maximum := 0

	for _, current := range rankings {
		limit :=
			agreementDepth

		if limit >
			len(current.Candidates) {
			limit =
				len(current.Candidates)
		}

		seen := make(
			map[domain.FactID]struct{},
		)

		for index := 0; index < limit; index++ {
			factID :=
				current.Candidates[index].FactID

			if _, exists :=
				seen[factID]; exists {
				continue
			}

			seen[factID] =
				struct{}{}

			counts[factID]++

			if counts[factID] >
				maximum {
				maximum =
					counts[factID]
			}
		}
	}

	if maximum == 0 {
		return 0
	}

	return clamp(
		float64(maximum) /
			float64(len(rankings)),
	)
}

func retrievalPairwiseOverlap(
	rankings []retrieval.Ranking,
) float64 {
	if len(rankings) < 2 {
		return 0.5
	}

	total := 0.0
	pairs := 0

	for left := 0; left < len(rankings); left++ {
		for right := left + 1; right < len(rankings); right++ {
			total +=
				rankingOverlap(
					rankings[left],
					rankings[right],
				)

			pairs++
		}
	}

	if pairs == 0 {
		return 0.5
	}

	return clamp(
		total /
			float64(pairs),
	)
}

func rankingOverlap(
	left retrieval.Ranking,
	right retrieval.Ranking,
) float64 {
	leftSet :=
		rankingFactSet(
			left,
			agreementDepth,
		)

	rightSet :=
		rankingFactSet(
			right,
			agreementDepth,
		)

	if len(leftSet) == 0 &&
		len(rightSet) == 0 {
		return 1
	}

	union := make(
		map[domain.FactID]struct{},
	)

	intersection := 0

	for factID := range leftSet {
		union[factID] =
			struct{}{}

		if _, exists :=
			rightSet[factID]; exists {
			intersection++
		}
	}

	for factID := range rightSet {
		union[factID] =
			struct{}{}
	}

	if len(union) == 0 {
		return 0
	}

	return clamp(
		float64(intersection) /
			float64(len(union)),
	)
}

func rankingFactSet(
	value retrieval.Ranking,
	limit int,
) map[domain.FactID]struct{} {
	result := make(
		map[domain.FactID]struct{},
	)

	if limit >
		len(value.Candidates) {
		limit =
			len(value.Candidates)
	}

	for index := 0; index < limit; index++ {
		result[value.Candidates[index].FactID] = struct{}{}
	}

	return result
}

func rankingSeparation(
	ranked ranking.Result,
	reasoned reasoning.Result,
) (float64, bool) {
	if len(
		reasoned.Conclusion.Groups,
	) > 0 {
		return groupSeparation(
			reasoned.Conclusion.Groups,
		)
	}

	return candidateSeparation(
		ranked.Candidates,
	)
}

func groupSeparation(
	groups []reasoning.EntityGroup,
) (float64, bool) {
	if len(groups) == 0 {
		return 0, false
	}

	if len(groups) == 1 {
		return 0.9, true
	}

	first :=
		groups[0].Score

	second :=
		groups[1].Score

	return normalizedSeparation(
		first,
		second,
	), true
}

func candidateSeparation(
	candidates []ranking.Candidate,
) (float64, bool) {
	if len(candidates) == 0 {
		return 0, false
	}

	if len(candidates) == 1 {
		return 0.9, true
	}

	return normalizedSeparation(
		candidates[0].Score,
		candidates[1].Score,
	), true
}

func normalizedSeparation(
	first float64,
	second float64,
) float64 {
	if first <= 0 {
		return 0
	}

	margin :=
		(first - second) /
			math.Max(
				first,
				0.000001,
			)

	if margin <= 0 {
		return 0
	}

	const strongMargin = 0.25

	return clamp(
		margin / strongMargin,
	)
}
