package evaluation

import (
	"fmt"

	"ai-agent/internal/agent"
)

type Debugger interface {
	Debug(
		question string,
	) (agent.DebugResult, error)
}

type Evaluator struct {
	debugger Debugger
}

func New(
	debugger Debugger,
) (*Evaluator, error) {
	if debugger == nil {
		return nil,
			fmt.Errorf(
				"debugger is required",
			)
	}

	return &Evaluator{
		debugger: debugger,
	}, nil
}

func (e *Evaluator) RunCase(
	current Case,
) CaseResult {
	debug, err :=
		e.debugger.Debug(
			current.Question,
		)

	if err != nil {
		return CaseResult{
			Case:   current,
			Passed: false,
			Error:  err.Error(),
			Checks: []Check{
				{
					Name:     CheckExecution,
					Passed:   false,
					Expected: "successful execution",
					Actual:   err.Error(),
				},
			},
		}
	}

	checks :=
		evaluateChecks(
			current,
			debug,
		)

	return CaseResult{
		Case: current,
		Passed: checksPassed(
			checks,
		),
		Checks: checks,
	}
}

func (e *Evaluator) Run(
	cases []Case,
) Report {
	results := make(
		[]CaseResult,
		0,
		len(cases),
	)

	for _, current := range cases {
		results = append(
			results,
			e.RunCase(
				current,
			),
		)
	}

	return BuildReport(
		results,
	)
}

func (e *Evaluator) RunDefault() Report {
	return e.Run(
		DefaultCases(),
	)
}

func (e *Evaluator) RunRegression() Report {
	return e.Run(
		RegressionCases(),
	)
}
