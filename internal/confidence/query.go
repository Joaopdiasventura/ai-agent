package confidence

import (
	"ai-agent/internal/domain"
)

func queryQuality(
	currentQuery domain.Query,
) float64 {
	languageScore :=
		queryLanguageScore(
			currentQuery.Language,
		)

	intentScore :=
		queryIntentScore(
			currentQuery.Intent,
		)

	targetScore :=
		queryTargetScore(
			currentQuery.Target,
		)

	semanticScore :=
		querySemanticScore(
			currentQuery,
		)

	lexicalScore :=
		queryLexicalScore(
			currentQuery,
		)

	return clamp(
		0.15*languageScore +
			0.30*intentScore +
			0.15*targetScore +
			0.30*semanticScore +
			0.10*lexicalScore,
	)
}

func queryLanguageScore(
	value domain.Language,
) float64 {
	switch value {
	case domain.LanguagePortuguese,
		domain.LanguageEnglish:
		return 1

	default:
		return 0.35
	}
}

func queryIntentScore(
	value domain.Intent,
) float64 {
	if value == domain.IntentUnknown {
		return 0
	}

	return 1
}

func queryTargetScore(
	value domain.QueryTarget,
) float64 {
	switch value {
	case domain.QueryTargetUnknown:
		return 0

	case domain.QueryTargetAny:
		return 0.55

	default:
		return 1
	}
}

func querySemanticScore(
	currentQuery domain.Query,
) float64 {
	bestEntity := 0.0
	bestConcept := 0.0

	for _, entity := range currentQuery.Entities {
		if !entity.Explicit {
			continue
		}

		if entity.Score > bestEntity {
			bestEntity =
				entity.Score
		}
	}

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		if concept.Score >
			bestConcept {
			bestConcept =
				concept.Score
		}
	}

	switch {
	case bestEntity > 0 &&
		bestConcept > 0:
		return clamp(
			0.5*bestEntity +
				0.5*bestConcept,
		)

	case bestEntity > 0:
		return clamp(
			bestEntity,
		)

	case bestConcept > 0:
		return clamp(
			bestConcept,
		)

	default:
		if len(
			currentQuery.Terms,
		) >= 2 {
			return 0.45
		}

		return 0.2
	}
}

func queryLexicalScore(
	currentQuery domain.Query,
) float64 {
	termCount :=
		len(currentQuery.Terms)

	switch {
	case termCount >= 3:
		return 1

	case termCount == 2:
		return 0.85

	case termCount == 1:
		return 0.65

	default:
		return 0
	}
}
