package evaluation

import (
	"strings"
	"testing"

	"ai-agent/internal/agent"
	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

func createEvaluator(
	t *testing.T,
) *Evaluator {
	t.Helper()

	service, err :=
		agent.New()

	if err != nil {
		t.Fatal(err)
	}

	evaluator, err :=
		New(service)

	if err != nil {
		t.Fatal(err)
	}

	return evaluator
}

func TestRegressionCorpusPasses(
	t *testing.T,
) {
	evaluator :=
		createEvaluator(t)

	report :=
		evaluator.RunRegression()

	if report.Failed != 0 {
		t.Fatalf(
			"regression corpus failed\n%s\n\n%s",
			FormatMetrics(report),
			FormatFailures(report),
		)
	}

	if report.Accuracy != 1 {
		t.Fatalf(
			"expected 100%% accuracy, got %.2f%%",
			report.Accuracy*100,
		)
	}
}

func TestDefaultCorpusRuns(
	t *testing.T,
) {
	evaluator :=
		createEvaluator(t)

	report :=
		evaluator.RunDefault()

	if report.Total !=
		len(DefaultCases()) {
		t.Fatalf(
			"expected %d cases, got %d",
			len(DefaultCases()),
			report.Total,
		)
	}

	if report.Accuracy < 0 ||
		report.Accuracy > 1 {
		t.Fatalf(
			"invalid accuracy %f",
			report.Accuracy,
		)
	}

	t.Log(
		FormatMetrics(report),
	)

	if report.Failed > 0 {
		t.Log(
			FormatFailures(report),
		)
	}
}

func TestCorpusIDsAreUnique(
	t *testing.T,
) {
	seen := make(
		map[string]struct{},
	)

	for _, current := range DefaultCases() {
		if current.ID == "" {
			t.Fatal(
				"case id cannot be empty",
			)
		}

		if _, exists :=
			seen[current.ID]; exists {
			t.Fatalf(
				"duplicate case id %s",
				current.ID,
			)
		}

		seen[current.ID] =
			struct{}{}
	}
}

func TestCorpusQuestionsAreNotEmpty(
	t *testing.T,
) {
	for _, current := range DefaultCases() {
		if strings.TrimSpace(
			current.Question,
		) == "" {
			t.Fatalf(
				"case %s has empty question",
				current.ID,
			)
		}
	}
}

func TestEvaluatorDetectsWrongExpectation(
	t *testing.T,
) {
	evaluator :=
		createEvaluator(t)

	current := Case{
		ID:       "intentional-failure",
		Question: "Qual projeto melhor demonstra concorrência?",
		Category: CategoryComparison,
		Expectation: Expectation{
			Winner: knowledge.EntityVox,
		},
	}

	result :=
		evaluator.RunCase(
			current,
		)

	if result.Passed {
		t.Fatal(
			"expected evaluation failure",
		)
	}

	foundWinnerFailure := false

	for _, check := range result.Checks {
		if check.Name ==
			CheckWinner &&
			!check.Passed {
			foundWinnerFailure = true
		}
	}

	if !foundWinnerFailure {
		t.Fatal(
			"expected winner failure",
		)
	}
}

func TestEvaluatorChecksForbiddenFacts(
	t *testing.T,
) {
	evaluator :=
		createEvaluator(t)

	current := Case{
		ID:       "forbidden-fact",
		Question: "Ele sabe Go?",
		Category: CategoryCapability,
		Expectation: Expectation{
			ForbiddenFacts: []domain.FactID{
				"profile-email",
			},
		},
	}

	result :=
		evaluator.RunCase(
			current,
		)

	if !result.Passed {
		t.Fatalf(
			"unexpected failure\n%s",
			formatCaseFailure(
				result,
			),
		)
	}
}

func TestEvaluatorChecksConfidenceMode(
	t *testing.T,
) {
	evaluator :=
		createEvaluator(t)

	current := Case{
		ID:       "rust-mode",
		Question: "Ele sabe Rust?",
		Category: CategoryAbstention,
		Expectation: Expectation{
			HasResponse:    boolValue(false),
			ConfidenceMode: confidence.ModeAbstention,
		},
	}

	result :=
		evaluator.RunCase(
			current,
		)

	if !result.Passed {
		t.Fatalf(
			"unexpected failure\n%s",
			formatCaseFailure(
				result,
			),
		)
	}
}

func TestReportTotals(
	t *testing.T,
) {
	results := []CaseResult{
		{
			Case: Case{
				ID:       "a",
				Category: CategoryDirect,
				Expectation: Expectation{
					Language: domain.LanguagePortuguese,
				},
			},
			Passed: true,
			Checks: []Check{
				{
					Name:   CheckLanguage,
					Passed: true,
				},
			},
		},
		{
			Case: Case{
				ID:       "b",
				Category: CategoryDirect,
				Expectation: Expectation{
					Language: domain.LanguagePortuguese,
				},
			},
			Passed: false,
			Checks: []Check{
				{
					Name:   CheckLanguage,
					Passed: false,
				},
			},
		},
	}

	report :=
		BuildReport(results)

	if report.Total != 2 {
		t.Fatalf(
			"expected 2, got %d",
			report.Total,
		)
	}

	if report.Passed != 1 {
		t.Fatalf(
			"expected 1 passed, got %d",
			report.Passed,
		)
	}

	if report.Failed != 1 {
		t.Fatalf(
			"expected 1 failed, got %d",
			report.Failed,
		)
	}

	if report.Accuracy != 0.5 {
		t.Fatalf(
			"expected 0.5, got %f",
			report.Accuracy,
		)
	}

	category :=
		report.ByCategory[CategoryDirect]

	if category.Total != 2 ||
		category.Passed != 1 ||
		category.Failed != 1 {
		t.Fatalf(
			"invalid category counter %+v",
			category,
		)
	}

	language :=
		report.ByLanguage[domain.LanguagePortuguese]

	if language.Total != 2 {
		t.Fatalf(
			"expected 2 portuguese cases, got %d",
			language.Total,
		)
	}

	check :=
		report.ByCheck[CheckLanguage]

	if check.Total != 2 ||
		check.Passed != 1 {
		t.Fatalf(
			"invalid check counter %+v",
			check,
		)
	}
}

func TestCounterAccuracy(
	t *testing.T,
) {
	var counter Counter

	counter.Add(true)
	counter.Add(true)
	counter.Add(false)

	expected :=
		float64(2) / 3

	if counter.Accuracy() !=
		expected {
		t.Fatalf(
			"expected %f, got %f",
			expected,
			counter.Accuracy(),
		)
	}
}

func TestEmptyCounterAccuracy(
	t *testing.T,
) {
	var counter Counter

	if counter.Accuracy() != 0 {
		t.Fatalf(
			"expected zero, got %f",
			counter.Accuracy(),
		)
	}
}

func TestFailedCases(
	t *testing.T,
) {
	report :=
		BuildReport(
			[]CaseResult{
				{
					Case: Case{
						ID: "pass",
					},
					Passed: true,
				},
				{
					Case: Case{
						ID: "fail",
					},
					Passed: false,
				},
			},
		)

	failed :=
		report.FailedCases()

	if len(failed) != 1 {
		t.Fatalf(
			"expected one failure, got %d",
			len(failed),
		)
	}

	if failed[0].Case.ID !=
		"fail" {
		t.Fatalf(
			"expected fail, got %s",
			failed[0].Case.ID,
		)
	}
}

func TestFailureFormatting(
	t *testing.T,
) {
	report :=
		BuildReport(
			[]CaseResult{
				{
					Case: Case{
						ID:       "example",
						Question: "question",
					},
					Passed: false,
					Checks: []Check{
						{
							Name:     CheckWinner,
							Passed:   false,
							Expected: "ggcompress",
							Actual:   "auronix",
						},
					},
				},
			},
		)

	value :=
		FormatFailures(
			report,
		)

	if !strings.Contains(
		value,
		"example",
	) {
		t.Fatalf(
			"expected case id, got %q",
			value,
		)
	}

	if !strings.Contains(
		value,
		"ggcompress",
	) {
		t.Fatalf(
			"expected expected value, got %q",
			value,
		)
	}

	if !strings.Contains(
		value,
		"auronix",
	) {
		t.Fatalf(
			"expected actual value, got %q",
			value,
		)
	}
}

func formatCaseFailure(
	result CaseResult,
) string {
	report :=
		BuildReport(
			[]CaseResult{
				result,
			},
		)

	return FormatFailures(
		report,
	)
}
