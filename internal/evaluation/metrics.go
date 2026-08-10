package evaluation

import (
	"ai-agent/internal/domain"
)

type Counter struct {
	Total  int
	Passed int
	Failed int
}

func (c *Counter) Add(
	passed bool,
) {
	c.Total++

	if passed {
		c.Passed++
		return
	}

	c.Failed++
}

func (c Counter) Accuracy() float64 {
	if c.Total == 0 {
		return 0
	}

	return float64(c.Passed) /
		float64(c.Total)
}

type Report struct {
	Total      int
	Passed     int
	Failed     int
	Accuracy   float64
	ByCategory map[Category]Counter
	ByLanguage map[domain.Language]Counter
	ByCheck    map[CheckName]Counter
	Cases      []CaseResult
}

func BuildReport(
	results []CaseResult,
) Report {
	report := Report{
		ByCategory:
			make(
				map[Category]Counter,
			),
		ByLanguage:
			make(
				map[domain.Language]Counter,
			),
		ByCheck:
			make(
				map[CheckName]Counter,
			),
		Cases:
			make(
				[]CaseResult,
				len(results),
			),
	}

	copy(
		report.Cases,
		results,
	)

	for _, current := range results {
		report.Total++

		if current.Passed {
			report.Passed++
		} else {
			report.Failed++
		}

		categoryCounter :=
			report.ByCategory[
				current.Case.Category,
			]

		categoryCounter.Add(
			current.Passed,
		)

		report.ByCategory[
			current.Case.Category,
		] = categoryCounter

		language :=
			current.Case.
				Expectation.
				Language

		if language != "" {
			languageCounter :=
				report.ByLanguage[
					language,
				]

			languageCounter.Add(
				current.Passed,
			)

			report.ByLanguage[
				language,
			] = languageCounter
		}

		for _, check :=
			range current.Checks {
			checkCounter :=
				report.ByCheck[
					check.Name,
				]

			checkCounter.Add(
				check.Passed,
			)

			report.ByCheck[
				check.Name,
			] = checkCounter
		}
	}

	if report.Total > 0 {
		report.Accuracy =
			float64(
				report.Passed,
			) /
				float64(
					report.Total,
				)
	}

	return report
}

func (r Report) FailedCases() []CaseResult {
	result := make(
		[]CaseResult,
		0,
	)

	for _, current :=
		range r.Cases {
		if current.Passed {
			continue
		}

		result = append(
			result,
			current,
		)
	}

	return result
}