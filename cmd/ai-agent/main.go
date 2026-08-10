package main

import (
	"fmt"
	"os"

	"ai-agent/internal/agent"
	"ai-agent/internal/cli"
)

func main() {
	service, err :=
		agent.New()

	if err != nil {
		fail(err)
	}

	runner, err :=
		cli.New(
			service,
			os.Stdin,
			os.Stdout,
		)

	if err != nil {
		fail(err)
	}

	if err :=
		runner.Run(); err != nil {
		fail(err)
	}
}

func fail(
	err error,
) {
	fmt.Fprintln(
		os.Stderr,
		err,
	)

	os.Exit(1)
}
