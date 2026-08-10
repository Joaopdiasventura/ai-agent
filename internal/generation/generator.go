package generation

import (
	"fmt"
	"strings"

	"ai-agent/internal/planning"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(
	plan planning.Plan,
	material Material,
) (Answer, error) {
	if err :=
		material.Validate(
			plan,
		); err != nil {
		return Answer{}, err
	}

	targetLanguage :=
		outputLanguage(
			plan.Language,
		)

	if plan.Status ==
		planning.PlanStatusAbstain {
		result :=
			realizeAbstention(
				targetLanguage,
			)

		return Answer{
			Text:     result.text,
			Language: targetLanguage,
		}, nil
	}

	var result realization

	switch plan.Type {
	case planning.PlanTypeDirect:
		result =
			realizeDirect(
				plan,
				material,
				targetLanguage,
			)

	case planning.PlanTypeOverview:
		result =
			realizeOverview(
				plan,
				material,
				targetLanguage,
			)

	case planning.PlanTypeCapability:
		result =
			realizeCapability(
				plan,
				material,
				targetLanguage,
			)

	case planning.PlanTypeExperience:
		result =
			realizeExperience(
				plan,
				material,
				targetLanguage,
			)

	case planning.PlanTypeTechnologyUsage:
		result =
			realizeTechnologyUsage(
				plan,
				material,
				targetLanguage,
			)

	case planning.PlanTypeComparison:
		result =
			realizeComparison(
				plan,
				material,
				targetLanguage,
			)

	case planning.PlanTypeList:
		result =
			realizeList(
				plan,
				material,
				targetLanguage,
			)

	default:
		return Answer{},
			fmt.Errorf(
				"unsupported plan type %s",
				plan.Type,
			)
	}

	result.text =
		strings.TrimSpace(
			result.text,
		)

	if result.text == "" {
		return Answer{},
			fmt.Errorf(
				"plan %s produced an empty answer",
				plan.Type,
			)
	}

	return Answer{
		Text:     result.text,
		Language: targetLanguage,
		FactIDs:  result.factIDs,
	}, nil
}
