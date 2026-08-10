package ranking

import (
	"sort"

	"ai-agent/internal/domain"
	"ai-agent/internal/retrieval"
)

type SignalName string

const (
	SignalFusion          SignalName = "fusion"
	SignalImportance      SignalName = "importance"
	SignalIntent          SignalName = "intent"
	SignalTarget          SignalName = "target"
	SignalTemporal        SignalName = "temporal"
	SignalEntityCoverage  SignalName = "entity_coverage"
	SignalConceptCoverage SignalName = "concept_coverage"
	SignalSourceDiversity SignalName = "source_diversity"
)

type Signal struct {
	Name   SignalName
	Score  float64
	Weight float64
}

type Candidate struct {
	FactID          domain.FactID
	Rank            int
	Score           float64
	FusionScore     float64
	FeatureScore    float64
	Sources         []retrieval.Source
	SourceRanks     map[retrieval.Source]int
	MatchedTerms    []string
	MatchedEntities []domain.EntityID
	MatchedConcepts []domain.ConceptID
	Signals         []Signal
}

func (c Candidate) HasSource(
	source retrieval.Source,
) bool {
	for _, current := range c.Sources {
		if current == source {
			return true
		}
	}

	return false
}

func (c Candidate) Signal(
	name SignalName,
) (Signal, bool) {
	for _, signal := range c.Signals {
		if signal.Name == name {
			return signal, true
		}
	}

	return Signal{}, false
}

type Result struct {
	Query      domain.Query
	Candidates []Candidate
}

func (r Result) Top(
	limit int,
) []Candidate {
	if limit <= 0 ||
		len(r.Candidates) == 0 {
		return nil
	}

	if limit > len(r.Candidates) {
		limit = len(r.Candidates)
	}

	result := make(
		[]Candidate,
		limit,
	)

	copy(
		result,
		r.Candidates[:limit],
	)

	return result
}

func (r Result) Candidate(
	factID domain.FactID,
) (Candidate, bool) {
	for _, candidate := range r.Candidates {
		if candidate.FactID == factID {
			return candidate, true
		}
	}

	return Candidate{}, false
}

func (r Result) FactIDs() []domain.FactID {
	result := make(
		[]domain.FactID,
		0,
		len(r.Candidates),
	)

	for _, candidate := range r.Candidates {
		result = append(
			result,
			candidate.FactID,
		)
	}

	return result
}

func sortFinalCandidates(
	values []Candidate,
) {
	sort.SliceStable(
		values,
		func(left int, right int) bool {
			if values[left].Score !=
				values[right].Score {
				return values[left].Score >
					values[right].Score
			}

			if values[left].FusionScore !=
				values[right].FusionScore {
				return values[left].FusionScore >
					values[right].FusionScore
			}

			return values[left].FactID <
				values[right].FactID
		},
	)

	for index := range values {
		values[index].Rank =
			index + 1
	}
}

func appendSource(
	values []retrieval.Source,
	value retrieval.Source,
) []retrieval.Source {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func appendString(
	values []string,
	value string,
) []string {
	if value == "" {
		return values
	}

	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func appendEntity(
	values []domain.EntityID,
	value domain.EntityID,
) []domain.EntityID {
	if value == "" {
		return values
	}

	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func appendConcept(
	values []domain.ConceptID,
	value domain.ConceptID,
) []domain.ConceptID {
	if value == "" {
		return values
	}

	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func clamp(
	value float64,
) float64 {
	if value < 0 {
		return 0
	}

	if value > 1 {
		return 1
	}

	return value
}
