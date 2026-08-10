package reasoning

import (
	"sort"

	"ai-agent/internal/domain"
)

type ConclusionType string

const (
	ConclusionUnknown         ConclusionType = "unknown"
	ConclusionDirect          ConclusionType = "direct"
	ConclusionOverview        ConclusionType = "overview"
	ConclusionCapability      ConclusionType = "capability"
	ConclusionExperience      ConclusionType = "experience"
	ConclusionTechnologyUsage ConclusionType = "technology_usage"
	ConclusionComparison      ConclusionType = "comparison"
	ConclusionList            ConclusionType = "list"
)

type SupportStatus string

const (
	SupportUnknown              SupportStatus = "unknown"
	SupportSupported            SupportStatus = "supported"
	SupportInsufficientEvidence SupportStatus = "insufficient_evidence"
)

type Evidence struct {
	FactID          domain.FactID
	Rank            int
	Score           float64
	Importance      float64
	Directness      float64
	MatchedEntities []domain.EntityID
	MatchedConcepts []domain.ConceptID
}

type EntityGroup struct {
	EntityID         domain.EntityID
	Score            float64
	Rank             int
	Evidence         []Evidence
	EvidenceStrength float64
	ConceptCoverage  float64
	Diversity        float64
	Quantity         float64
}

func (g EntityGroup) TopEvidence(
	limit int,
) []Evidence {
	if limit <= 0 ||
		len(g.Evidence) == 0 {
		return nil
	}

	if limit > len(g.Evidence) {
		limit = len(g.Evidence)
	}

	result := make(
		[]Evidence,
		limit,
	)

	copy(
		result,
		g.Evidence[:limit],
	)

	return result
}

type Conclusion struct {
	Type         ConclusionType
	Status       SupportStatus
	FocusEntity  domain.EntityID
	FocusConcept domain.ConceptID
	Evidence     []Evidence
	Groups       []EntityGroup
}

type Result struct {
	Query      domain.Query
	Conclusion Conclusion
}

func (r Result) TopEvidence(
	limit int,
) []Evidence {
	if limit <= 0 ||
		len(r.Conclusion.Evidence) == 0 {
		return nil
	}

	if limit >
		len(r.Conclusion.Evidence) {
		limit =
			len(r.Conclusion.Evidence)
	}

	result := make(
		[]Evidence,
		limit,
	)

	copy(
		result,
		r.Conclusion.Evidence[:limit],
	)

	return result
}

func (r Result) TopGroup() (
	EntityGroup,
	bool,
) {
	if len(r.Conclusion.Groups) == 0 {
		return EntityGroup{}, false
	}

	return r.Conclusion.Groups[0], true
}

func sortEvidence(
	values []Evidence,
) {
	sort.SliceStable(
		values,
		func(left int, right int) bool {
			if values[left].Score !=
				values[right].Score {
				return values[left].Score >
					values[right].Score
			}

			if values[left].Directness !=
				values[right].Directness {
				return values[left].Directness >
					values[right].Directness
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

func sortGroups(
	values []EntityGroup,
) {
	sort.SliceStable(
		values,
		func(left int, right int) bool {
			if values[left].Score !=
				values[right].Score {
				return values[left].Score >
					values[right].Score
			}

			if values[left].EvidenceStrength !=
				values[right].EvidenceStrength {
				return values[left].EvidenceStrength >
					values[right].EvidenceStrength
			}

			return values[left].EntityID <
				values[right].EntityID
		},
	)

	for index := range values {
		values[index].Rank =
			index + 1
	}
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
