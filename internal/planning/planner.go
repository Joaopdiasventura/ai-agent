package planning

import (
	"errors"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/reasoning"
)

type Planner struct {
	base     *knowledge.Knowledge
	selector *EvidenceSelector
}

func New(
	base *knowledge.Knowledge,
) (*Planner, error) {
	if base == nil {
		return nil, errors.New(
			"knowledge base is required",
		)
	}

	return &Planner{
		base: base,
		selector: NewEvidenceSelector(
			base,
		),
	}, nil
}

func (p *Planner) Plan(
	result reasoning.Result,
) Plan {
	plan := Plan{
		Type: planType(
			result.Conclusion.Type,
		),
		Language:     result.Query.Language,
		Intent:       result.Query.Intent,
		Target:       result.Query.Target,
		FocusEntity:  result.Conclusion.FocusEntity,
		FocusConcept: result.Conclusion.FocusConcept,
	}

	if result.Conclusion.Status !=
		reasoning.SupportSupported {
		plan.Status =
			PlanStatusAbstain

		plan.Sections =
			[]Section{
				abstentionSection(
					result.Conclusion.FocusEntity,
					result.Conclusion.FocusConcept,
				),
			}

		return plan
	}

	plan.Status =
		PlanStatusReady

	switch result.Conclusion.Type {
	case reasoning.ConclusionComparison:
		plan.Sections =
			p.planComparison(
				result,
			)

	case reasoning.ConclusionCapability:
		plan.Sections =
			p.planCapability(
				result,
			)

	case reasoning.ConclusionTechnologyUsage:
		plan.Sections =
			p.planTechnologyUsage(
				result,
			)

	case reasoning.ConclusionOverview:
		plan.Sections =
			p.planOverview(
				result,
			)

	case reasoning.ConclusionExperience:
		plan.Sections =
			p.planExperience(
				result,
			)

	case reasoning.ConclusionList:
		plan.Sections =
			p.planList(
				result,
			)

	case reasoning.ConclusionDirect:
		plan.Sections =
			p.planDirect(
				result,
			)

	default:
		plan.Status =
			PlanStatusAbstain

		plan.Sections =
			[]Section{
				abstentionSection(
					result.Conclusion.FocusEntity,
					result.Conclusion.FocusConcept,
				),
			}
	}

	return plan
}

func (p *Planner) planComparison(
	result reasoning.Result,
) []Section {
	if len(
		result.Conclusion.Groups,
	) == 0 {
		return []Section{
			abstentionSection(
				result.Conclusion.FocusEntity,
				result.Conclusion.FocusConcept,
			),
		}
	}

	winner :=
		result.Conclusion.Groups[0]

	winnerEvidence :=
		p.selector.Select(
			winner.Evidence,
			4,
		)

	sections := []Section{
		winnerSection(
			winner,
			winnerEvidence,
		),
		evidenceSection(
			SectionEvidence,
			winnerEvidence,
		),
	}

	if len(
		result.Conclusion.Groups,
	) > 1 {
		sections = append(
			sections,
			alternativesSection(
				result.Conclusion.Groups[1:],
				p.selector,
				2,
			),
		)
	}

	return sections
}

func (p *Planner) planCapability(
	result reasoning.Result,
) []Section {
	evidence :=
		p.selector.Select(
			result.Conclusion.Evidence,
			4,
		)

	sections := []Section{
		supportSection(
			result.Conclusion.FocusEntity,
			result.Conclusion.FocusConcept,
			result.Conclusion.Status,
		),
	}

	if len(evidence) > 0 {
		sections = append(
			sections,
			evidenceSection(
				SectionEvidence,
				evidence,
			),
		)
	}

	return sections
}

func (p *Planner) planTechnologyUsage(
	result reasoning.Result,
) []Section {
	evidence :=
		p.selector.Select(
			result.Conclusion.Evidence,
			5,
		)

	sections := []Section{
		supportSection(
			result.Conclusion.FocusEntity,
			result.Conclusion.FocusConcept,
			result.Conclusion.Status,
		),
	}

	if len(
		result.Conclusion.Groups,
	) > 0 {
		sections = append(
			sections,
			entityListSection(
				result.Conclusion.Groups,
				p.selector,
				5,
				ItemTechnologyUsage,
			),
		)
	}

	if len(evidence) > 0 {
		sections = append(
			sections,
			evidenceSection(
				SectionEvidence,
				evidence,
			),
		)
	}

	return sections
}

func (p *Planner) planOverview(
	result reasoning.Result,
) []Section {
	evidence :=
		p.selector.Select(
			result.Conclusion.Evidence,
			6,
		)

	return []Section{
		{
			Kind: SectionLead,
			Items: []Item{
				{
					Kind:     ItemEntity,
					EntityID: result.Conclusion.FocusEntity,
					Rank:     1,
				},
			},
		},
		evidenceSection(
			SectionDetails,
			evidence,
		),
	}
}

func (p *Planner) planExperience(
	result reasoning.Result,
) []Section {
	evidence :=
		p.selector.Select(
			result.Conclusion.Evidence,
			7,
		)

	return []Section{
		evidenceSection(
			SectionDetails,
			evidence,
		),
	}
}

func (p *Planner) planList(
	result reasoning.Result,
) []Section {
	if len(
		result.Conclusion.Groups,
	) > 0 {
		return []Section{
			entityListSection(
				result.Conclusion.Groups,
				p.selector,
				10,
				ItemEntity,
			),
		}
	}

	evidence :=
		p.selector.Select(
			result.Conclusion.Evidence,
			10,
		)

	return []Section{
		evidenceSection(
			SectionList,
			evidence,
		),
	}
}

func (p *Planner) planDirect(
	result reasoning.Result,
) []Section {
	limit :=
		directEvidenceLimit(
			result.Query.Intent,
		)

	evidence :=
		p.selector.Select(
			result.Conclusion.Evidence,
			limit,
		)

	return []Section{
		evidenceSection(
			SectionDetails,
			evidence,
		),
	}
}

func directEvidenceLimit(
	intent domain.Intent,
) int {
	switch intent {
	case domain.IntentContact:
		return 1

	case domain.IntentEducation:
		return 4

	case domain.IntentCertification:
		return 8

	default:
		return 3
	}
}

func planType(
	conclusionType reasoning.ConclusionType,
) PlanType {
	switch conclusionType {
	case reasoning.ConclusionDirect:
		return PlanTypeDirect

	case reasoning.ConclusionOverview:
		return PlanTypeOverview

	case reasoning.ConclusionCapability:
		return PlanTypeCapability

	case reasoning.ConclusionExperience:
		return PlanTypeExperience

	case reasoning.ConclusionTechnologyUsage:
		return PlanTypeTechnologyUsage

	case reasoning.ConclusionComparison:
		return PlanTypeComparison

	case reasoning.ConclusionList:
		return PlanTypeList

	default:
		return PlanTypeUnknown
	}
}
