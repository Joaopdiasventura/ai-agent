package retrieval

import (
	"errors"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
	"ai-agent/internal/knowledge"
)

type EntityRetriever struct {
	base  *knowledge.Knowledge
	index *searchindex.Index
}

func NewEntityRetriever(
	base *knowledge.Knowledge,
	currentIndex *searchindex.Index,
) (*EntityRetriever, error) {
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

	return &EntityRetriever{
		base:  base,
		index: currentIndex,
	}, nil
}

func (r *EntityRetriever) Search(
	currentQuery domain.Query,
	limit int,
) Ranking {
	if len(currentQuery.Entities) == 0 {
		return Ranking{
			Source: SourceEntity,
		}
	}

	type state struct {
		score    float64
		entities []domain.EntityID
	}

	states := make(
		map[domain.FactID]*state,
	)

	for _, match := range currentQuery.Entities {
		entity, found :=
			r.base.Entity(
				match.EntityID,
			)

		if !found {
			continue
		}

		typeWeight :=
			entityTypeWeight(
				entity.Type,
			)

		factIDs :=
			r.index.FactsByEntity(
				match.EntityID,
			)

		for _, factID := range factIDs {
			fact, found :=
				r.base.Fact(factID)

			if !found {
				continue
			}

			position :=
				entityPositionWeight(
					fact,
					match.EntityID,
				)

			importance :=
				0.7 +
					0.3*
						fact.Importance

			score :=
				match.Score *
					typeWeight *
					position *
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

			currentState.entities =
				appendUniqueEntity(
					currentState.entities,
					match.EntityID,
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
				Source:          SourceEntity,
				MatchedEntities: currentState.entities,
			},
		)
	}

	return Ranking{
		Source: SourceEntity,
		Candidates: limitCandidates(
			candidates,
			limit,
		),
	}
}

func entityTypeWeight(
	entityType domain.EntityType,
) float64 {
	switch entityType {
	case domain.EntityTypePerson:
		return 0.25

	case domain.EntityTypeProject:
		return 1

	case domain.EntityTypeTechnology:
		return 1

	case domain.EntityTypeCompany:
		return 0.9

	case domain.EntityTypeInstitution:
		return 0.9

	case domain.EntityTypeCertification:
		return 0.95

	case domain.EntityTypeRole:
		return 0.85

	case domain.EntityTypeLanguage:
		return 0.9

	case domain.EntityTypeLocation:
		return 0.8

	default:
		return 0.7
	}
}

func entityPositionWeight(
	fact domain.Fact,
	entityID domain.EntityID,
) float64 {
	if fact.Subject == entityID {
		return 1
	}

	if fact.Object.Kind ==
		domain.FactObjectEntity &&
		fact.Object.EntityID ==
			entityID {
		return 0.95
	}

	for _, contextID := range fact.Context {
		if contextID == entityID {
			return 0.85
		}
	}

	return 0.75
}
