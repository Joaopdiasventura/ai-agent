package ranking

import (
	"sort"

	"ai-agent/internal/domain"
	"ai-agent/internal/retrieval"
)

type RRFConfig struct {
	K             float64
	SourceWeights map[retrieval.Source]float64
}

func DefaultRRFConfig() RRFConfig {
	return RRFConfig{
		K: 60,
		SourceWeights: map[retrieval.Source]float64{
			retrieval.SourceLexical: 1,
			retrieval.SourceEntity:  1.15,
			retrieval.SourceConcept: 1.2,
			retrieval.SourceFuzzy:   0.7,
		},
	}
}

type RRFFusion struct {
	config RRFConfig
}

func NewRRFFusion(
	config RRFConfig,
) *RRFFusion {
	if config.K <= 0 {
		config.K = 60
	}

	return &RRFFusion{
		config: config,
	}
}

func (r *RRFFusion) Fuse(
	result retrieval.Result,
	limit int,
) []Candidate {
	type state struct {
		candidate Candidate
		rawScore  float64
	}

	states := make(
		map[domain.FactID]*state,
	)

	for _, ranking := range result.Rankings {
		sourceWeight :=
			r.sourceWeight(
				ranking.Source,
			)

		for index, sourceCandidate := range ranking.Candidates {
			position :=
				index + 1

			score :=
				sourceWeight /
					(r.config.K + float64(position))

			current, exists :=
				states[sourceCandidate.FactID]

			if !exists {
				current = &state{
					candidate: Candidate{
						FactID: sourceCandidate.FactID,
						SourceRanks: make(
							map[retrieval.Source]int,
						),
					},
				}

				states[sourceCandidate.FactID] = current
			}

			current.rawScore += score

			current.candidate.Sources =
				appendSource(
					current.candidate.Sources,
					ranking.Source,
				)

			current.candidate.SourceRanks[ranking.Source] = position

			for _, term := range sourceCandidate.MatchedTerms {
				current.candidate.MatchedTerms =
					appendString(
						current.candidate.MatchedTerms,
						term,
					)
			}

			for _, entity := range sourceCandidate.MatchedEntities {
				current.candidate.MatchedEntities =
					appendEntity(
						current.candidate.MatchedEntities,
						entity,
					)
			}

			for _, concept := range sourceCandidate.MatchedConcepts {
				current.candidate.MatchedConcepts =
					appendConcept(
						current.candidate.MatchedConcepts,
						concept,
					)
			}
		}
	}

	if len(states) == 0 {
		return nil
	}

	maximumRawScore := 0.0

	for _, current := range states {
		if current.rawScore >
			maximumRawScore {
			maximumRawScore =
				current.rawScore
		}
	}

	values := make(
		[]Candidate,
		0,
		len(states),
	)

	for _, current := range states {
		if maximumRawScore > 0 {
			current.candidate.FusionScore =
				current.rawScore /
					maximumRawScore
		}

		current.candidate.Signals =
			append(
				current.candidate.Signals,
				Signal{
					Name:   SignalFusion,
					Score:  current.candidate.FusionScore,
					Weight: 1,
				},
			)

		values = append(
			values,
			current.candidate,
		)
	}

	sort.SliceStable(
		values,
		func(left int, right int) bool {
			if values[left].FusionScore !=
				values[right].FusionScore {
				return values[left].FusionScore >
					values[right].FusionScore
			}

			return values[left].FactID <
				values[right].FactID
		},
	)

	if limit > 0 &&
		len(values) > limit {
		values =
			values[:limit]
	}

	return values
}

func (r *RRFFusion) sourceWeight(
	source retrieval.Source,
) float64 {
	weight, found :=
		r.config.SourceWeights[source]

	if !found || weight <= 0 {
		return 1
	}

	return weight
}
