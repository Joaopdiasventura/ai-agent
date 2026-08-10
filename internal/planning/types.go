package planning

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/reasoning"
)

type PlanType string

const (
	PlanTypeUnknown         PlanType = "unknown"
	PlanTypeDirect          PlanType = "direct"
	PlanTypeOverview        PlanType = "overview"
	PlanTypeCapability      PlanType = "capability"
	PlanTypeExperience      PlanType = "experience"
	PlanTypeTechnologyUsage PlanType = "technology_usage"
	PlanTypeComparison      PlanType = "comparison"
	PlanTypeList            PlanType = "list"
)

type PlanStatus string

const (
	PlanStatusReady   PlanStatus = "ready"
	PlanStatusAbstain PlanStatus = "abstain"
)

type SectionKind string

const (
	SectionLead         SectionKind = "lead"
	SectionEvidence     SectionKind = "evidence"
	SectionDetails      SectionKind = "details"
	SectionAlternatives SectionKind = "alternatives"
	SectionList         SectionKind = "list"
	SectionAbstention   SectionKind = "abstention"
)

type ItemKind string

const (
	ItemFact                  ItemKind = "fact"
	ItemEntity                ItemKind = "entity"
	ItemSupport               ItemKind = "support"
	ItemComparisonWinner      ItemKind = "comparison_winner"
	ItemComparisonAlternative ItemKind = "comparison_alternative"
	ItemTechnologyUsage       ItemKind = "technology_usage"
)

type Item struct {
	Kind        ItemKind
	FactID      domain.FactID
	EntityID    domain.EntityID
	ConceptID   domain.ConceptID
	Support     reasoning.SupportStatus
	Score       float64
	Rank        int
	EvidenceIDs []domain.FactID
}

type Section struct {
	Kind  SectionKind
	Items []Item
}

type Plan struct {
	Type         PlanType
	Status       PlanStatus
	Language     domain.Language
	Intent       domain.Intent
	Target       domain.QueryTarget
	FocusEntity  domain.EntityID
	FocusConcept domain.ConceptID
	Sections     []Section
}

func (p Plan) Section(
	kind SectionKind,
) (Section, bool) {
	for _, section := range p.Sections {
		if section.Kind == kind {
			return section, true
		}
	}

	return Section{}, false
}

func (p Plan) FactIDs() []domain.FactID {
	seen := make(
		map[domain.FactID]struct{},
	)

	result := make(
		[]domain.FactID,
		0,
	)

	for _, section := range p.Sections {
		for _, item := range section.Items {
			if item.FactID != "" {
				if _, exists := seen[item.FactID]; !exists {
					seen[item.FactID] = struct{}{}

					result = append(
						result,
						item.FactID,
					)
				}
			}

			for _, factID := range item.EvidenceIDs {
				if factID == "" {
					continue
				}

				if _, exists :=
					seen[factID]; exists {
					continue
				}

				seen[factID] = struct{}{}

				result = append(
					result,
					factID,
				)
			}
		}
	}

	return result
}

func (p Plan) EntityIDs() []domain.EntityID {
	seen := make(
		map[domain.EntityID]struct{},
	)

	result := make(
		[]domain.EntityID,
		0,
	)

	if p.FocusEntity != "" {
		seen[p.FocusEntity] = struct{}{}

		result = append(
			result,
			p.FocusEntity,
		)
	}

	for _, section := range p.Sections {
		for _, item := range section.Items {
			if item.EntityID == "" {
				continue
			}

			if _, exists :=
				seen[item.EntityID]; exists {
				continue
			}

			seen[item.EntityID] = struct{}{}

			result = append(
				result,
				item.EntityID,
			)
		}
	}

	return result
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
