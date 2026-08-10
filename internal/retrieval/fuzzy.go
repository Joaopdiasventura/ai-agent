package retrieval

import (
	"math"
	"sort"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
	"ai-agent/internal/language"
)

type FuzzyRetriever struct {
	index *searchindex.Index
}

type fuzzyAlternative struct {
	Term       string
	Similarity float64
}

func NewFuzzyRetriever(
	currentIndex *searchindex.Index,
) *FuzzyRetriever {
	return &FuzzyRetriever{
		index: currentIndex,
	}
}

func (r *FuzzyRetriever) Search(
	currentQuery domain.Query,
	limit int,
) Ranking {
	if len(currentQuery.Terms) == 0 {
		return Ranking{
			Source: SourceFuzzy,
		}
	}

	merged := make(
		map[domain.FactID]Candidate,
	)

	for _, targetLanguage := range retrievalLanguages(
		currentQuery.Language,
	) {
		candidates :=
			r.searchLanguage(
				currentQuery,
				targetLanguage,
				limit*2,
			)

		for _, candidate := range candidates {
			mergeCandidateMaximum(
				merged,
				candidate,
			)
		}
	}

	return Ranking{
		Source: SourceFuzzy,
		Candidates: candidatesFromMap(
			merged,
			limit,
		),
	}
}

func (r *FuzzyRetriever) searchLanguage(
	currentQuery domain.Query,
	targetLanguage domain.Language,
	limit int,
) []Candidate {
	type state struct {
		score        float64
		matchedTerms []string
	}

	states := make(
		map[domain.FactID]*state,
	)

	documentCount :=
		r.index.DocumentCount(
			targetLanguage,
		)

	if documentCount == 0 {
		return nil
	}

	seenQueryTerms :=
		make(map[string]struct{})

	for _, queryTerm := range currentQuery.Terms {
		if queryTerm == "" {
			continue
		}

		if _, exists :=
			seenQueryTerms[queryTerm]; exists {
			continue
		}

		seenQueryTerms[queryTerm] =
			struct{}{}

		if len([]rune(queryTerm)) < 4 {
			continue
		}

		if len(
			r.index.Postings(
				targetLanguage,
				queryTerm,
			),
		) > 0 {
			continue
		}

		alternatives :=
			r.alternatives(
				targetLanguage,
				queryTerm,
			)

		for _, alternative := range alternatives {
			documentFrequency :=
				r.index.DocumentFrequency(
					targetLanguage,
					alternative.Term,
				)

			idf :=
				inverseDocumentFrequency(
					documentCount,
					documentFrequency,
				)

			postings :=
				r.index.Postings(
					targetLanguage,
					alternative.Term,
				)

			for _, posting := range postings {
				fieldWeight :=
					fuzzyFieldWeight(
						posting.Field,
					)

				frequencyWeight :=
					1 +
						math.Log(
							float64(
								posting.Frequency,
							),
						)

				score :=
					alternative.Similarity *
						idf *
						fieldWeight *
						frequencyWeight

				if score <= 0 {
					continue
				}

				currentState, exists :=
					states[posting.FactID]

				if !exists {
					currentState =
						&state{}

					states[posting.FactID] = currentState
				}

				currentState.score +=
					score

				currentState.matchedTerms =
					appendUniqueString(
						currentState.matchedTerms,
						queryTerm+
							"→"+
							alternative.Term,
					)
			}
		}
	}

	candidates := make(
		[]Candidate,
		0,
		len(states),
	)

	for factID, currentState := range states {
		candidates = append(
			candidates,
			Candidate{
				FactID:       factID,
				Language:     targetLanguage,
				Score:        currentState.score,
				Source:       SourceFuzzy,
				MatchedTerms: currentState.matchedTerms,
			},
		)
	}

	return limitCandidates(
		candidates,
		limit,
	)
}

func (r *FuzzyRetriever) alternatives(
	targetLanguage domain.Language,
	queryTerm string,
) []fuzzyAlternative {
	candidates :=
		r.index.FuzzyTermCandidates(
			targetLanguage,
			queryTerm,
		)

	result := make(
		[]fuzzyAlternative,
		0,
	)

	for _, candidate := range candidates {
		if candidate == queryTerm {
			continue
		}

		similarity :=
			language.FuzzySimilarity(
				queryTerm,
				candidate,
			)

		if similarity <
			fuzzyThreshold(
				queryTerm,
				candidate,
			) {
			continue
		}

		result = append(
			result,
			fuzzyAlternative{
				Term:       candidate,
				Similarity: similarity,
			},
		)
	}

	sort.Slice(
		result,
		func(left int, right int) bool {
			if result[left].Similarity !=
				result[right].Similarity {
				return result[left].Similarity >
					result[right].Similarity
			}

			return result[left].Term <
				result[right].Term
		},
	)

	const maximumAlternatives = 8

	if len(result) >
		maximumAlternatives {
		result =
			result[:maximumAlternatives]
	}

	return result
}

func fuzzyThreshold(
	left string,
	right string,
) float64 {
	shortest := len([]rune(left))

	if value :=
		len([]rune(right)); value < shortest {
		shortest = value
	}

	switch {
	case shortest >= 10:
		return 0.78

	case shortest >= 8:
		return 0.8

	case shortest >= 6:
		return 0.84

	default:
		return 0.9
	}
}

func fuzzyFieldWeight(
	field searchindex.Field,
) float64 {
	switch field {
	case searchindex.FieldSubject:
		return 1.35

	case searchindex.FieldObject:
		return 1.25

	case searchindex.FieldConcept:
		return 1.25

	case searchindex.FieldContext:
		return 1.15

	case searchindex.FieldStatement:
		return 1

	case searchindex.FieldPredicate:
		return 0.8

	case searchindex.FieldCategory:
		return 0.7

	default:
		return 1
	}
}
