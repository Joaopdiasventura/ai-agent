package reasoning

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

type CapabilityReasoner struct {
	base *knowledge.Knowledge
}

func NewCapabilityReasoner(
	base *knowledge.Knowledge,
) *CapabilityReasoner {
	return &CapabilityReasoner{
		base: base,
	}
}

func (r *CapabilityReasoner) Reason(
	currentQuery domain.Query,
	evidence []Evidence,
) Conclusion {
	relevant := make(
		[]Evidence,
		0,
	)

	for _, currentEvidence := range evidence {
		fact, found :=
			r.base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			continue
		}

		if !capabilityEvidenceRelevant(
			currentQuery,
			fact,
			r.base,
		) {
			continue
		}

		relevant = append(
			relevant,
			currentEvidence,
		)
	}

	sortEvidence(relevant)

	status :=
		capabilityStatus(
			currentQuery,
			relevant,
			r.base,
		)

	return Conclusion{
		Type:   ConclusionCapability,
		Status: status,
		FocusEntity: focusCapabilityEntity(
			currentQuery,
			r.base,
		),
		FocusConcept: focusConcept(
			currentQuery,
		),
		Evidence: relevant,
	}
}

func capabilityEvidenceRelevant(
	currentQuery domain.Query,
	fact domain.Fact,
	base *knowledge.Knowledge,
) bool {
	hasCapabilityTarget := false

	for _, entity := range currentQuery.Entities {
		if !entity.Explicit {
			continue
		}

		if !isCapabilityTargetEntity(
			entity.EntityID,
			base,
		) {
			continue
		}

		hasCapabilityTarget = true

		if factReferencesEntity(
			fact,
			entity.EntityID,
		) {
			return true
		}
	}

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		hasCapabilityTarget = true

		if factHasConcept(
			fact,
			concept.ConceptID,
		) {
			return true
		}
	}

	if !hasCapabilityTarget {
		return false
	}

	return false
}

func capabilityStatus(
	currentQuery domain.Query,
	evidence []Evidence,
	base *knowledge.Knowledge,
) SupportStatus {
	if len(evidence) == 0 {
		return SupportInsufficientEvidence
	}

	for _, currentEvidence := range evidence {
		if currentEvidence.Score < 0.35 {
			continue
		}

		fact, found :=
			base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			continue
		}

		if hasDirectCapabilityMatch(
			currentQuery,
			fact,
			base,
		) {
			return SupportSupported
		}
	}

	return SupportInsufficientEvidence
}

func hasDirectCapabilityMatch(
	currentQuery domain.Query,
	fact domain.Fact,
	base *knowledge.Knowledge,
) bool {
	for _, entity := range currentQuery.Entities {
		if !entity.Explicit {
			continue
		}

		if !isCapabilityTargetEntity(
			entity.EntityID,
			base,
		) {
			continue
		}

		if factReferencesEntity(
			fact,
			entity.EntityID,
		) {
			return true
		}
	}

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		if factHasConcept(
			fact,
			concept.ConceptID,
		) {
			return true
		}
	}

	return false
}

func isCapabilityTargetEntity(
	entityID domain.EntityID,
	base *knowledge.Knowledge,
) bool {
	entity, found :=
		base.Entity(entityID)

	if !found {
		return false
	}

	switch entity.Type {
	case domain.EntityTypePerson:
		return false

	default:
		return true
	}
}

func focusCapabilityEntity(
	currentQuery domain.Query,
	base *knowledge.Knowledge,
) domain.EntityID {
	bestScore := -1.0

	var best domain.EntityID

	for _, match := range currentQuery.Entities {
		if !match.Explicit {
			continue
		}

		if !isCapabilityTargetEntity(
			match.EntityID,
			base,
		) {
			continue
		}

		if match.Score <= bestScore {
			continue
		}

		bestScore =
			match.Score

		best =
			match.EntityID
	}

	return best
}
