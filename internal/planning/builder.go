package planning

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/reasoning"
)

func evidenceSection(
	kind SectionKind,
	evidence []reasoning.Evidence,
) Section {
	items := make(
		[]Item,
		0,
		len(evidence),
	)

	for index, current := range evidence {
		items = append(
			items,
			Item{
				Kind:   ItemFact,
				FactID: current.FactID,
				Score: clamp(
					current.Score,
				),
				Rank: index + 1,
			},
		)
	}

	return Section{
		Kind:  kind,
		Items: items,
	}
}

func supportSection(
	entityID domain.EntityID,
	conceptID domain.ConceptID,
	status reasoning.SupportStatus,
) Section {
	return Section{
		Kind: SectionLead,
		Items: []Item{
			{
				Kind:      ItemSupport,
				EntityID:  entityID,
				ConceptID: conceptID,
				Support:   status,
				Rank:      1,
			},
		},
	}
}

func winnerSection(
	group reasoning.EntityGroup,
	evidence []reasoning.Evidence,
) Section {
	return Section{
		Kind: SectionLead,
		Items: []Item{
			{
				Kind:     ItemComparisonWinner,
				EntityID: group.EntityID,
				Score: clamp(
					group.Score,
				),
				Rank: 1,
				EvidenceIDs: evidenceIDs(
					evidence,
				),
			},
		},
	}
}

func alternativesSection(
	groups []reasoning.EntityGroup,
	selector *EvidenceSelector,
	limit int,
) Section {
	if limit <= 0 {
		return Section{
			Kind: SectionAlternatives,
		}
	}

	if limit > len(groups) {
		limit = len(groups)
	}

	items := make(
		[]Item,
		0,
		limit,
	)

	for index := 0; index < limit; index++ {
		group :=
			groups[index]

		evidence :=
			selector.Select(
				group.Evidence,
				2,
			)

		items = append(
			items,
			Item{
				Kind:     ItemComparisonAlternative,
				EntityID: group.EntityID,
				Score: clamp(
					group.Score,
				),
				Rank: index + 1,
				EvidenceIDs: evidenceIDs(
					evidence,
				),
			},
		)
	}

	return Section{
		Kind:  SectionAlternatives,
		Items: items,
	}
}

func entityListSection(
	groups []reasoning.EntityGroup,
	selector *EvidenceSelector,
	limit int,
	itemKind ItemKind,
) Section {
	if limit <= 0 {
		return Section{
			Kind: SectionList,
		}
	}

	if limit > len(groups) {
		limit = len(groups)
	}

	items := make(
		[]Item,
		0,
		limit,
	)

	for index := 0; index < limit; index++ {
		group :=
			groups[index]

		evidence :=
			selector.Select(
				group.Evidence,
				3,
			)

		items = append(
			items,
			Item{
				Kind:     itemKind,
				EntityID: group.EntityID,
				Score: clamp(
					group.Score,
				),
				Rank: index + 1,
				EvidenceIDs: evidenceIDs(
					evidence,
				),
			},
		)
	}

	return Section{
		Kind:  SectionList,
		Items: items,
	}
}

func abstentionSection(
	entityID domain.EntityID,
	conceptID domain.ConceptID,
) Section {
	return Section{
		Kind: SectionAbstention,
		Items: []Item{
			{
				Kind:      ItemSupport,
				EntityID:  entityID,
				ConceptID: conceptID,
				Support:   reasoning.SupportInsufficientEvidence,
				Rank:      1,
			},
		},
	}
}
