package retrieval

import (
	"errors"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
	"ai-agent/internal/knowledge"
)

type ConceptRetriever struct {
	base  *knowledge.Knowledge
	index *searchindex.Index
}

func NewConceptRetriever(
	base *knowledge.Knowledge,
	currentIndex *searchindex.Index,
) (*ConceptRetriever, error) {
	if base == nil {
		return nil, errors.New(
			"knowledge base is required",
		)
	}

	if currentIndex == nil {
		return nil, errors.New(
			"index is required",
		)
	}

	return &ConceptRetriever{
		base:  base,
		index: currentIndex,
	}, nil
}

func (r *ConceptRetriever) Search(
	currentQuery domain.Query,
	limit int,
) Ranking {
	if len(currentQuery.Concepts) == 0 {
		return Ranking{
			Source: SourceConcept,
		}
	}

	type state struct {
		score    float64
		concepts []domain.ConceptID
	}

	states := make(
		map[domain.FactID]*state,
	)

	for _, match := range currentQuery.Concepts {
		factIDs :=
			r.index.FactsByConcept(
				match.ConceptID,
			)

		if len(factIDs) == 0 {
			continue
		}

		explicitWeight := 0.72

		if match.MatchedText != "" {
			explicitWeight = 1
		}

		for _, factID := range factIDs {
			fact, found :=
				r.base.Fact(factID)

			if !found {
				continue
			}

			importance :=
				0.7 +
					0.3*
						fact.Importance

			score :=
				match.Score *
					explicitWeight *
					importance

			if score <= 0 {
				continue
			}

			currentState, exists :=
				states[factID]

			if !exists {
				currentState =
					&state{}

				states[factID] =
					currentState
			}

			currentState.score +=
				score

			currentState.concepts =
				appendUniqueConcept(
					currentState.concepts,
					match.ConceptID,
				)
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
				FactID:          factID,
				Language:        currentQuery.Language,
				Score:           currentState.score,
				Source:          SourceConcept,
				MatchedConcepts: currentState.concepts,
			},
		)
	}

	return Ranking{
		Source: SourceConcept,
		Candidates: limitCandidates(
			candidates,
			limit,
		),
	}
}
