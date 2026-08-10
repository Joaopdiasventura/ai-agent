package generation

import (
	"fmt"

	"ai-agent/internal/domain"
	"ai-agent/internal/planning"
)

type realization struct {
	text    string
	factIDs []domain.FactID
}

func realizeAbstention(
	targetLanguage domain.Language,
) realization {
	if targetLanguage ==
		domain.LanguagePortuguese {
		return realization{
			text: "Não encontrei evidência suficiente na base disponível para responder isso com segurança.",
		}
	}

	return realization{
		text: "I could not find enough evidence in the available knowledge base to answer that reliably.",
	}
}

func realizeDirect(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	statements := make(
		[]string,
		0,
	)

	factIDs := make(
		[]domain.FactID,
		0,
	)

	for _, section := range plan.Sections {
		for _, item := range section.Items {
			if item.FactID == "" {
				continue
			}

			statement :=
				factStatement(
					material,
					item.FactID,
					targetLanguage,
				)

			if statement == "" {
				continue
			}

			statements = append(
				statements,
				statement,
			)

			factIDs =
				appendFactID(
					factIDs,
					item.FactID,
				)
		}
	}

	return realization{
		text: joinSentences(
			statements,
		),
		factIDs: factIDs,
	}
}

func realizeOverview(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	label :=
		entityLabel(
			material,
			plan.FocusEntity,
			targetLanguage,
		)

	details, factIDs :=
		sectionStatements(
			plan,
			material,
			planning.SectionDetails,
			targetLanguage,
		)

	body :=
		joinSentences(details)

	if label == "" {
		return realization{
			text:    body,
			factIDs: factIDs,
		}
	}

	if targetLanguage ==
		domain.LanguagePortuguese {
		if body == "" {
			return realization{
				text: "Sobre " +
					label +
					".",
			}
		}

		return realization{
			text: "Sobre " +
				label +
				": " +
				body,
			factIDs: factIDs,
		}
	}

	if body == "" {
		return realization{
			text: "About " +
				label +
				".",
		}
	}

	return realization{
		text: "About " +
			label +
			": " +
			body,
		factIDs: factIDs,
	}
}

func realizeCapability(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	label :=
		entityLabel(
			material,
			plan.FocusEntity,
			targetLanguage,
		)

	evidence, factIDs :=
		sectionStatements(
			plan,
			material,
			planning.SectionEvidence,
			targetLanguage,
		)

	body :=
		joinSentences(evidence)

	var lead string

	if targetLanguage ==
		domain.LanguagePortuguese {
		if label != "" {
			lead =
				fmt.Sprintf(
					"Sim. Há evidências suficientes de experiência com %s.",
					label,
				)
		} else {
			lead =
				"Sim. Há evidências suficientes na base para sustentar essa capacidade."
		}
	} else {
		if label != "" {
			lead =
				fmt.Sprintf(
					"Yes. There is sufficient evidence of experience with %s.",
					label,
				)
		} else {
			lead =
				"Yes. There is sufficient evidence in the knowledge base to support that capability."
		}
	}

	if body == "" {
		return realization{
			text:    lead,
			factIDs: factIDs,
		}
	}

	return realization{
		text: lead +
			" " +
			body,
		factIDs: factIDs,
	}
}

func realizeComparison(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	leadSection, found :=
		plan.Section(
			planning.SectionLead,
		)

	if !found ||
		len(leadSection.Items) == 0 {
		return realization{}
	}

	winner :=
		leadSection.Items[0]

	winnerLabel :=
		entityLabel(
			material,
			winner.EntityID,
			targetLanguage,
		)

	factIDs :=
		appendFactIDs(
			nil,
			winner.EvidenceIDs,
		)

	evidence, evidenceFactIDs :=
		sectionStatements(
			plan,
			material,
			planning.SectionEvidence,
			targetLanguage,
		)

	factIDs =
		appendFactIDs(
			factIDs,
			evidenceFactIDs,
		)

	var lead string

	if targetLanguage ==
		domain.LanguagePortuguese {
		if winnerLabel != "" {
			lead =
				fmt.Sprintf(
					"Entre as evidências disponíveis, %s aparece com o suporte mais forte para essa comparação.",
					winnerLabel,
				)
		} else {
			lead =
				"As evidências disponíveis apontam uma opção com suporte mais forte para essa comparação."
		}
	} else {
		if winnerLabel != "" {
			lead =
				fmt.Sprintf(
					"Among the available evidence, %s has the strongest support in this comparison.",
					winnerLabel,
				)
		} else {
			lead =
				"The available evidence indicates one option with stronger support in this comparison."
		}
	}

	body :=
		joinSentences(evidence)

	alternativeLabels,
		alternativeFactIDs :=
		comparisonAlternatives(
			plan,
			material,
			targetLanguage,
		)

	factIDs =
		appendFactIDs(
			factIDs,
			alternativeFactIDs,
		)

	alternatives :=
		joinNatural(
			alternativeLabels,
			targetLanguage,
		)

	parts := []string{
		lead,
	}

	if body != "" {
		parts = append(
			parts,
			body,
		)
	}

	if alternatives != "" {
		if targetLanguage ==
			domain.LanguagePortuguese {
			parts = append(
				parts,
				"Outras opções com evidências relevantes incluem "+
					alternatives+
					".",
			)
		} else {
			parts = append(
				parts,
				"Other options with relevant evidence include "+
					alternatives+
					".",
			)
		}
	}

	return realization{
		text:    joinSentences(parts),
		factIDs: factIDs,
	}
}

func realizeTechnologyUsage(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	technology :=
		entityLabel(
			material,
			plan.FocusEntity,
			targetLanguage,
		)

	entities, entityFactIDs :=
		entityItems(
			plan,
			material,
			planning.SectionList,
			targetLanguage,
		)

	places :=
		joinNatural(
			entities,
			targetLanguage,
		)

	var lead string

	if targetLanguage ==
		domain.LanguagePortuguese {
		switch {
		case technology != "" &&
			places != "":
			lead =
				fmt.Sprintf(
					"As evidências mostram uso de %s em %s.",
					technology,
					places,
				)

		case technology != "":
			lead =
				fmt.Sprintf(
					"Há evidências de uso de %s.",
					technology,
				)

		default:
			lead =
				"Há evidências de uso dessa tecnologia."
		}
	} else {
		switch {
		case technology != "" &&
			places != "":
			lead =
				fmt.Sprintf(
					"The evidence shows %s being used in %s.",
					technology,
					places,
				)

		case technology != "":
			lead =
				fmt.Sprintf(
					"There is evidence of %s being used.",
					technology,
				)

		default:
			lead =
				"There is evidence of that technology being used."
		}
	}

	evidence, evidenceFactIDs :=
		sectionStatements(
			plan,
			material,
			planning.SectionEvidence,
			targetLanguage,
		)

	factIDs :=
		appendFactIDs(
			entityFactIDs,
			evidenceFactIDs,
		)

	body :=
		joinSentences(evidence)

	if body == "" {
		return realization{
			text:    lead,
			factIDs: factIDs,
		}
	}

	return realization{
		text: lead +
			" " +
			body,
		factIDs: factIDs,
	}
}

func realizeExperience(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	details, factIDs :=
		sectionStatements(
			plan,
			material,
			planning.SectionDetails,
			targetLanguage,
		)

	body :=
		joinSentences(details)

	if body == "" {
		return realization{}
	}

	if targetLanguage ==
		domain.LanguagePortuguese {
		return realization{
			text: "As evidências profissionais mais relevantes são: " +
				body,
			factIDs: factIDs,
		}
	}

	return realization{
		text: "The most relevant professional evidence is: " +
			body,
		factIDs: factIDs,
	}
}

func realizeList(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) realization {
	entities, entityFactIDs :=
		entityItems(
			plan,
			material,
			planning.SectionList,
			targetLanguage,
		)

	if len(entities) > 0 {
		values :=
			joinNatural(
				entities,
				targetLanguage,
			)

		if targetLanguage ==
			domain.LanguagePortuguese {
			return realization{
				text: "Os principais itens encontrados são " +
					values +
					".",
				factIDs: entityFactIDs,
			}
		}

		return realization{
			text: "The main items found are " +
				values +
				".",
			factIDs: entityFactIDs,
		}
	}

	statements, factIDs :=
		sectionStatements(
			plan,
			material,
			planning.SectionList,
			targetLanguage,
		)

	return realization{
		text: joinSentences(
			statements,
		),
		factIDs: factIDs,
	}
}

func sectionStatements(
	plan planning.Plan,
	material Material,
	kind planning.SectionKind,
	targetLanguage domain.Language,
) ([]string, []domain.FactID) {
	section, found :=
		plan.Section(kind)

	if !found {
		return nil, nil
	}

	statements := make(
		[]string,
		0,
	)

	factIDs := make(
		[]domain.FactID,
		0,
	)

	for _, item := range section.Items {
		if item.FactID == "" {
			continue
		}

		statement :=
			factStatement(
				material,
				item.FactID,
				targetLanguage,
			)

		if statement == "" {
			continue
		}

		statements = append(
			statements,
			statement,
		)

		factIDs =
			appendFactID(
				factIDs,
				item.FactID,
			)
	}

	return statements,
		factIDs
}

func entityItems(
	plan planning.Plan,
	material Material,
	kind planning.SectionKind,
	targetLanguage domain.Language,
) ([]string, []domain.FactID) {
	section, found :=
		plan.Section(kind)

	if !found {
		return nil, nil
	}

	values := make(
		[]string,
		0,
	)

	factIDs := make(
		[]domain.FactID,
		0,
	)

	for _, item := range section.Items {
		if item.EntityID == "" {
			continue
		}

		label :=
			entityLabel(
				material,
				item.EntityID,
				targetLanguage,
			)

		if label == "" {
			continue
		}

		values = append(
			values,
			label,
		)

		factIDs =
			appendFactIDs(
				factIDs,
				item.EvidenceIDs,
			)
	}

	return values,
		factIDs
}

func comparisonAlternatives(
	plan planning.Plan,
	material Material,
	targetLanguage domain.Language,
) ([]string, []domain.FactID) {
	return entityItems(
		plan,
		material,
		planning.SectionAlternatives,
		targetLanguage,
	)
}
