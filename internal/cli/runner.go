package cli

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"ai-agent/internal/agent"
)

type Service interface {
	Answer(
		question string,
	) (agent.Result, error)
}

type Runner struct {
	service Service
	input   io.Reader
	output  io.Writer
}

func New(
	service Service,
	input io.Reader,
	output io.Writer,
) (*Runner, error) {
	if service == nil {
		return nil,
			fmt.Errorf(
				"agent service is required",
			)
	}

	if input == nil {
		return nil,
			fmt.Errorf(
				"input is required",
			)
	}

	if output == nil {
		return nil,
			fmt.Errorf(
				"output is required",
			)
	}

	return &Runner{
		service: service,
		input:   input,
		output:  output,
	}, nil
}

func (r *Runner) Run() error {
	scanner :=
		bufio.NewScanner(
			r.input,
		)

	fmt.Fprintln(
		r.output,
		"Chatbot iniciado.",
	)

	fmt.Fprintln(
		r.output,
		"Digite 'sair' para encerrar.",
	)

	for {
		fmt.Fprint(
			r.output,
			"\nVocê: ",
		)

		if !scanner.Scan() {
			break
		}

		question :=
			strings.TrimSpace(
				scanner.Text(),
			)

		if question == "" {
			continue
		}

		if ShouldExit(question) {
			fmt.Fprintln(
				r.output,
				"Chatbot encerrado.",
			)

			break
		}

		result, err :=
			r.service.Answer(
				question,
			)

		if err != nil {
			return err
		}

		if !result.HasResponse {
			fmt.Fprintf(
				r.output,
				"Bot: %s\n",
				NotFoundMessage(
					result.Language,
				),
			)

			continue
		}

		fmt.Fprintf(
			r.output,
			"Bot: %s\n",
			result.Response,
		)
	}

	if err :=
		scanner.Err(); err != nil {
		return fmt.Errorf(
			"read terminal input: %w",
			err,
		)
	}

	return nil
}
