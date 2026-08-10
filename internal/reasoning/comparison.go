package reasoning

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

type ComparisonReasoner struct {
	base    *knowledge.Knowledge
	grouper *Grouper
}

func NewComparisonReasoner(
	base *knowledge.Knowledge,
	grouper *Grouper,
) *ComparisonReasoner {
	return &ComparisonReasoner{
		base:    base,
		grouper: grouper,
	}
}

func (r *ComparisonReasoner) Reason(
	currentQuery domain.Query,
	evidence []Evidence,
) Conclusion {
	entityType :=
		comparisonEntityType(
			currentQuery.Target,
		)

	groups :=
		r.grouper.Group(
			currentQuery,
			evidence,
			entityType,
		)

	status :=
		SupportInsufficientEvidence

	if len(groups) > 0 {
		status =
			SupportSupported
	}

	resultEvidence :=
		flattenGroupEvidence(
			groups,
			20,
		)

	return Conclusion{
		Type:   ConclusionComparison,
		Status: status,
		FocusEntity: focusEntity(
			currentQuery,
		),
		FocusConcept: focusConcept(
			currentQuery,
		),
		Evidence: resultEvidence,
		Groups:   groups,
	}
}

func comparisonEntityType(
	target domain.QueryTarget,
) domain.EntityType {
	switch target {
	case domain.QueryTargetTechnology:
		return domain.EntityTypeTechnology

	case domain.QueryTargetCompany:
		return domain.EntityTypeCompany

	case domain.QueryTargetPerson:
		return domain.EntityTypePerson

	case domain.QueryTargetEducation:
		return domain.EntityTypeInstitution

	case domain.QueryTargetCertification:
		return domain.EntityTypeCertification

	default:
		return domain.EntityTypeProject
	}
}

func flattenGroupEvidence(
	groups []EntityGroup,
	limit int,
) []Evidence {
	result := make(
		[]Evidence,
		0,
	)

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, group := range groups {
		for _, currentEvidence := range group.Evidence {
			if _, exists :=
				seen[currentEvidence.FactID]; exists {
				continue
			}

			seen[currentEvidence.FactID] =
				struct{}{}

			result = append(
				result,
				currentEvidence,
			)
		}
	}

	sortEvidence(result)

	if limit > 0 &&
		len(result) > limit {
		result =
			result[:limit]

		for index := range result {
			result[index].Rank =
				index + 1
		}
	}

	return result
}
