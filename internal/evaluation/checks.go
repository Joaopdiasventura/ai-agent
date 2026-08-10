package evaluation

import (
	"fmt"
	"strconv"
	"strings"

	"ai-agent/internal/agent"
	"ai-agent/internal/domain"
)

func evaluateChecks(
	current Case,
	debug agent.DebugResult,
) []Check {
	result := make(
		[]Check,
		0,
	)

	expected :=
		current.Expectation

	if expected.HasResponse != nil {
		result = append(
			result,
			Check{
				Name: CheckHasResponse,
				Passed: debug.Result.HasResponse ==
					*expected.HasResponse,
				Expected: strconv.FormatBool(
					*expected.HasResponse,
				),
				Actual: strconv.FormatBool(
					debug.Result.HasResponse,
				),
			},
		)
	}

	if expected.Language != "" {
		result = append(
			result,
			Check{
				Name: CheckLanguage,
				Passed: debug.Result.Language ==
					expected.Language,
				Expected: string(
					expected.Language,
				),
				Actual: string(
					debug.Result.Language,
				),
			},
		)
	}

	if expected.Intent != "" {
		result = append(
			result,
			Check{
				Name: CheckIntent,
				Passed: debug.Query.Intent ==
					expected.Intent,
				Expected: string(
					expected.Intent,
				),
				Actual: string(
					debug.Query.Intent,
				),
			},
		)
	}

	if expected.Target != "" {
		result = append(
			result,
			Check{
				Name: CheckTarget,
				Passed: debug.Query.Target ==
					expected.Target,
				Expected: string(
					expected.Target,
				),
				Actual: string(
					debug.Query.Target,
				),
			},
		)
	}

	for _, entityID := range expected.Entities {
		result = append(
			result,
			Check{
				Name: CheckEntity,
				Passed: debug.Query.HasEntity(
					entityID,
				),
				Expected: string(entityID),
				Actual: queryEntities(
					debug.Query,
				),
			},
		)
	}

	for _, conceptID := range expected.Concepts {
		result = append(
			result,
			Check{
				Name: CheckConcept,
				Passed: debug.Query.HasConcept(
					conceptID,
				),
				Expected: string(conceptID),
				Actual: queryConcepts(
					debug.Query,
				),
			},
		)
	}

	if expected.Winner != "" {
		winner := domain.EntityID("")

		if group, found :=
			debug.Reasoning.TopGroup(); found {
			winner =
				group.EntityID
		}

		result = append(
			result,
			Check{
				Name: CheckWinner,
				Passed: winner ==
					expected.Winner,
				Expected: string(
					expected.Winner,
				),
				Actual: string(winner),
			},
		)
	}

	for _, factID := range expected.Facts {
		result = append(
			result,
			Check{
				Name: CheckFact,
				Passed: containsFactID(
					debug.Generation.FactIDs,
					factID,
				),
				Expected: string(factID),
				Actual: factIDsString(
					debug.Generation.FactIDs,
				),
			},
		)
	}

	for _, factID := range expected.ForbiddenFacts {
		result = append(
			result,
			Check{
				Name: CheckForbiddenFact,
				Passed: !containsFactID(
					debug.Generation.FactIDs,
					factID,
				),
				Expected: "not " +
					string(factID),
				Actual: factIDsString(
					debug.Generation.FactIDs,
				),
			},
		)
	}

	if expected.PlanStatus != "" {
		result = append(
			result,
			Check{
				Name: CheckPlanStatus,
				Passed: debug.Plan.Status ==
					expected.PlanStatus,
				Expected: string(
					expected.PlanStatus,
				),
				Actual: string(
					debug.Plan.Status,
				),
			},
		)
	}

	if expected.ConfidenceMode != "" {
		result = append(
			result,
			Check{
				Name: CheckConfidenceMode,
				Passed: debug.Confidence.Mode ==
					expected.ConfidenceMode,
				Expected: string(
					expected.ConfidenceMode,
				),
				Actual: string(
					debug.Confidence.Mode,
				),
			},
		)
	}

	if expected.MinConfidence != nil {
		result = append(
			result,
			Check{
				Name: CheckMinimumConfidence,
				Passed: debug.Confidence.Score >=
					*expected.MinConfidence,
				Expected: fmt.Sprintf(
					">= %.4f",
					*expected.MinConfidence,
				),
				Actual: fmt.Sprintf(
					"%.4f",
					debug.Confidence.Score,
				),
			},
		)
	}

	if expected.MaxConfidence != nil {
		result = append(
			result,
			Check{
				Name: CheckMaximumConfidence,
				Passed: debug.Confidence.Score <=
					*expected.MaxConfidence,
				Expected: fmt.Sprintf(
					"<= %.4f",
					*expected.MaxConfidence,
				),
				Actual: fmt.Sprintf(
					"%.4f",
					debug.Confidence.Score,
				),
			},
		)
	}

	for _, expectedText := range expected.ResponseContains {
		result = append(
			result,
			Check{
				Name: CheckResponseContains,
				Passed: containsText(
					debug.Result.Response,
					expectedText,
				),
				Expected: expectedText,
				Actual:   debug.Result.Response,
			},
		)
	}

	for _, expectedText := range expected.ResponseNotContains {
		result = append(
			result,
			Check{
				Name: CheckResponseNotContains,
				Passed: !containsText(
					debug.Result.Response,
					expectedText,
				),
				Expected: "not " +
					expectedText,
				Actual: debug.Result.Response,
			},
		)
	}

	for _, expectedText := range expected.GenerationContains {
		result = append(
			result,
			Check{
				Name: CheckGenerationContains,
				Passed: containsText(
					debug.Generation.Text,
					expectedText,
				),
				Expected: expectedText,
				Actual:   debug.Generation.Text,
			},
		)
	}

	return result
}

func checksPassed(
	values []Check,
) bool {
	for _, current := range values {
		if !current.Passed {
			return false
		}
	}

	return true
}

func containsText(
	value string,
	expected string,
) bool {
	return strings.Contains(
		strings.ToLower(value),
		strings.ToLower(expected),
	)
}

func containsFactID(
	values []domain.FactID,
	expected domain.FactID,
) bool {
	for _, current := range values {
		if current == expected {
			return true
		}
	}

	return false
}

func factIDsString(
	values []domain.FactID,
) string {
	result := make(
		[]string,
		0,
		len(values),
	)

	for _, value := range values {
		result = append(
			result,
			string(value),
		)
	}

	return strings.Join(
		result,
		", ",
	)
}

func queryEntities(
	currentQuery domain.Query,
) string {
	result := make(
		[]string,
		0,
		len(currentQuery.Entities),
	)

	for _, match := range currentQuery.Entities {
		result = append(
			result,
			fmt.Sprintf(
				"%s:%.3f",
				match.EntityID,
				match.Score,
			),
		)
	}

	return strings.Join(
		result,
		", ",
	)
}

func queryConcepts(
	currentQuery domain.Query,
) string {
	result := make(
		[]string,
		0,
		len(currentQuery.Concepts),
	)

	for _, match := range currentQuery.Concepts {
		result = append(
			result,
			fmt.Sprintf(
				"%s:%.3f",
				match.ConceptID,
				match.Score,
			),
		)
	}

	return strings.Join(
		result,
		", ",
	)
}
