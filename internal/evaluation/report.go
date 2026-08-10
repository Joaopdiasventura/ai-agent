package evaluation

import (
	"fmt"
	"sort"
	"strings"
)

func Summary(
	report Report,
) string {
	return fmt.Sprintf(
		"total=%d passed=%d failed=%d accuracy=%.2f%%",
		report.Total,
		report.Passed,
		report.Failed,
		report.Accuracy*100,
	)
}

func FormatFailures(
	report Report,
) string {
	failed :=
		report.FailedCases()

	if len(failed) == 0 {
		return ""
	}

	var builder strings.Builder

	for index, current := range failed {
		if index > 0 {
			builder.WriteString(
				"\n\n",
			)
		}

		builder.WriteString(
			fmt.Sprintf(
				"[%s] %s\n",
				current.Case.ID,
				current.Case.Question,
			),
		)

		if current.Error != "" {
			builder.WriteString(
				"  error: ",
			)

			builder.WriteString(
				current.Error,
			)

			builder.WriteString(
				"\n",
			)
		}

		for _, check := range current.Checks {
			if check.Passed {
				continue
			}

			builder.WriteString(
				fmt.Sprintf(
					"  %s\n    expected: %s\n    actual:   %s\n",
					check.Name,
					check.Expected,
					check.Actual,
				),
			)
		}
	}

	return strings.TrimSpace(
		builder.String(),
	)
}

func FormatMetrics(
	report Report,
) string {
	var builder strings.Builder

	builder.WriteString(
		Summary(report),
	)

	categories := make(
		[]string,
		0,
		len(report.ByCategory),
	)

	for category := range report.ByCategory {
		categories = append(
			categories,
			string(category),
		)
	}

	sort.Strings(categories)

	for _, categoryValue := range categories {
		category :=
			Category(
				categoryValue,
			)

		counter :=
			report.ByCategory[category]

		builder.WriteString(
			fmt.Sprintf(
				"\ncategory.%s=%d/%d %.2f%%",
				category,
				counter.Passed,
				counter.Total,
				counter.Accuracy()*100,
			),
		)
	}

	checks := make(
		[]string,
		0,
		len(report.ByCheck),
	)

	for check := range report.ByCheck {
		checks = append(
			checks,
			string(check),
		)
	}

	sort.Strings(checks)

	for _, checkValue := range checks {
		check :=
			CheckName(
				checkValue,
			)

		counter :=
			report.ByCheck[check]

		builder.WriteString(
			fmt.Sprintf(
				"\ncheck.%s=%d/%d %.2f%%",
				check,
				counter.Passed,
				counter.Total,
				counter.Accuracy()*100,
			),
		)
	}

	return builder.String()
}
