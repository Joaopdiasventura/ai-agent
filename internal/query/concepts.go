package query

import (
	"sort"

	"ai-agent/internal/domain"
	"ai-agent/internal/language"
	"ai-agent/internal/ontology"
)

type ConceptExtractor struct {
	currentOntology *ontology.Ontology
}

func NewConceptExtractor(
	currentOntology *ontology.Ontology,
) *ConceptExtractor {
	return &ConceptExtractor{
		currentOntology: currentOntology,
	}
}

func (c *ConceptExtractor) Extract(
	normalized string,
	terms []string,
	targetLanguage domain.Language,
) []domain.ConceptMatch {
	matches := make(
		map[domain.ConceptID]domain.ConceptMatch,
	)

	aliases :=
		c.currentOntology.Aliases(
			targetLanguage,
		)

	for _, alias := range aliases {
		if !containsPhrase(
			normalized,
			alias.Normalized,
		) {
			continue
		}

		score := alias.Weight

		aliasTokens :=
			language.TokenizeNormalized(
				alias.Normalized,
			)

		if len(aliasTokens) > 1 {
			score = maximum(
				score,
				0.98,
			)
		}

		updateConceptMatch(
			matches,
			domain.ConceptMatch{
				ConceptID:   alias.ConceptID,
				Score:       score,
				MatchedText: alias.Normalized,
			},
		)
	}

	for _, term := range terms {
		if runeLength(term) < 5 {
			continue
		}

		for _, alias := range aliases {
			if !isSingleToken(alias.Normalized) {
				continue
			}

			if runeLength(alias.Normalized) < 5 {
				continue
			}

			if term == alias.Normalized {
				continue
			}

			similarity :=
				language.FuzzySimilarity(
					term,
					alias.Normalized,
				)

			if similarity <
				conceptFuzzyThreshold(
					term,
					alias.Normalized,
				) {
				continue
			}

			score :=
				similarity *
					alias.Weight *
					0.82

			updateConceptMatch(
				matches,
				domain.ConceptMatch{
					ConceptID:   alias.ConceptID,
					Score:       score,
					MatchedText: term,
				},
			)
		}
	}

	return sortedConceptMatches(
		matches,
	)
}

func updateConceptMatch(
	values map[domain.ConceptID]domain.ConceptMatch,
	candidate domain.ConceptMatch,
) {
	existing, found := values[candidate.ConceptID]

	if found && existing.Score >= candidate.Score {
		return
	}

	values[candidate.ConceptID] = candidate
}

func sortedConceptMatches(
	values map[domain.ConceptID]domain.ConceptMatch,
) []domain.ConceptMatch {
	result := make(
		[]domain.ConceptMatch,
		0,
		len(values),
	)

	for _, match := range values {
		result = append(result, match)
	}

	sort.Slice(
		result,
		func(i int, j int) bool {
			if result[i].Score != result[j].Score {
				return result[i].Score >
					result[j].Score
			}

			return result[i].ConceptID <
				result[j].ConceptID
		},
	)

	return result
}

func conceptFuzzyThreshold(
	left string,
	right string,
) float64 {
	shortest := runeLength(left)

	if value := runeLength(right); value < shortest {
		shortest = value
	}

	switch {
	case shortest >= 10:
		return 0.78
	case shortest >= 8:
		return 0.80
	case shortest >= 6:
		return 0.84
	default:
		return 0.9
	}
}
