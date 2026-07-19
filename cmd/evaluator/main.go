package main

import (
	"ai-agent/internal/evaluation"
	"fmt"
	"os"
)

const evaluationPath = "data/evaluation.json"

func main() {
	cases, err := evaluation.LoadCases(evaluationPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load evaluation cases: %v\n", err)
		os.Exit(1)
	}

	metrics := evaluation.Run(cases)
	evaluation.Print(metrics)
}
