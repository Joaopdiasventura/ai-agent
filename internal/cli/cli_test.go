package cli

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"ai-agent/internal/agent"
	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
)

type fakeService struct {
	results map[string]agent.Result
	err     error
	calls   []string
}

func (f *fakeService) Answer(
	question string,
) (agent.Result, error) {
	f.calls = append(
		f.calls,
		question,
	)

	if f.err != nil {
		return agent.Result{},
			f.err
	}

	result, found :=
		f.results[question]

	if !found {
		return agent.Result{
			HasResponse: false,
			Language:    domain.LanguagePortuguese,
		}, nil
	}

	return result, nil
}

func (f *fakeService) FactCount() int {
	return 42
}

func TestShouldExit(
	t *testing.T,
) {
	values := []string{
		"sair",
		"SAIR",
		" sair ",
		"exit",
		"EXIT",
		"quit",
		"QUIT",
		"encerrar",
		"ENCERRAR",
	}

	for _, value := range values {
		if !ShouldExit(value) {
			t.Fatalf(
				"expected %q to exit",
				value,
			)
		}
	}
}

func TestShouldNotExit(
	t *testing.T,
) {
	values := []string{
		"",
		"hello",
		"go",
		"me fale sobre o sair",
		"exit agora",
	}

	for _, value := range values {
		if ShouldExit(value) {
			t.Fatalf(
				"expected %q not to exit",
				value,
			)
		}
	}
}

func TestPortugueseConversation(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: map[string]agent.Result{
				"Qual o email do João?": {
					Response:        "O email do João é joaopdias.dev@gmail.com.",
					HasResponse:     true,
					Language:        domain.LanguagePortuguese,
					Confidence:      0.9,
					ConfidenceLevel: confidence.LevelHigh,
				},
			},
		}

	input :=
		strings.NewReader(
			"Qual o email do João?\nsair\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	value :=
		output.String()

	expected := []string{
		"Chatbot iniciado.",
		"Digite 'sair' para encerrar.",
		"Você:",
		"Bot: O email do João é joaopdias.dev@gmail.com.",
		"Chatbot encerrado.",
	}

	for _, current := range expected {
		if !strings.Contains(
			value,
			current,
		) {
			t.Fatalf(
				"expected %q in output:\n%s",
				current,
				value,
			)
		}
	}

	if len(service.calls) != 1 {
		t.Fatalf(
			"expected one agent call, got %d",
			len(service.calls),
		)
	}
}

func TestRunnerDoesNotPrintFactCount(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: make(
				map[string]agent.Result,
			),
		}

	input :=
		strings.NewReader(
			"sair\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	value :=
		output.String()

	if strings.Contains(
		strings.ToLower(value),
		"fatos carregados",
	) {
		t.Fatalf(
			"runner must not print fact count:\n%s",
			value,
		)
	}
}

func TestEnglishNotFoundMessage(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: map[string]agent.Result{
				"Does João know Rust?": {
					HasResponse:     false,
					Language:        domain.LanguageEnglish,
					Confidence:      0.95,
					ConfidenceLevel: confidence.LevelHigh,
				},
			},
		}

	input :=
		strings.NewReader(
			"Does João know Rust?\nexit\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	value :=
		output.String()

	expected :=
		"I don't have that specific information"

	if !strings.Contains(
		value,
		expected,
	) {
		t.Fatalf(
			"expected english fallback, got:\n%s",
			value,
		)
	}
}

func TestPortugueseNotFoundMessage(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: map[string]agent.Result{
				"Ele sabe Rust?": {
					HasResponse: false,
					Language:    domain.LanguagePortuguese,
				},
			},
		}

	input :=
		strings.NewReader(
			"Ele sabe Rust?\nsair\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	value :=
		output.String()

	expected :=
		"Não encontrei essa informação específica"

	if !strings.Contains(
		value,
		expected,
	) {
		t.Fatalf(
			"expected portuguese fallback, got:\n%s",
			value,
		)
	}
}

func TestEmptyQuestionIsIgnored(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: map[string]agent.Result{
				"Ele sabe Go?": {
					Response:    "Sim.",
					HasResponse: true,
					Language:    domain.LanguagePortuguese,
				},
			},
		}

	input :=
		strings.NewReader(
			"\n   \nEle sabe Go?\nsair\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	if len(service.calls) != 1 {
		t.Fatalf(
			"expected one call, got %d",
			len(service.calls),
		)
	}

	if service.calls[0] !=
		"Ele sabe Go?" {
		t.Fatalf(
			"unexpected call %q",
			service.calls[0],
		)
	}
}

func TestExitDoesNotCallAgent(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: make(
				map[string]agent.Result,
			),
		}

	input :=
		strings.NewReader(
			"sair\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	if len(service.calls) != 0 {
		t.Fatalf(
			"expected zero calls, got %d",
			len(service.calls),
		)
	}
}

func TestAnswerErrorIsReturned(
	t *testing.T,
) {
	expected :=
		errors.New(
			"agent failure",
		)

	service :=
		&fakeService{
			err: expected,
		}

	input :=
		strings.NewReader(
			"Ele sabe Go?\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	err =
		runner.Run()

	if !errors.Is(
		err,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			err,
		)
	}
}

func TestNewRequiresService(
	t *testing.T,
) {
	var output bytes.Buffer

	_, err :=
		New(
			nil,
			strings.NewReader(""),
			&output,
		)

	if err == nil {
		t.Fatal(
			"expected service error",
		)
	}
}

func TestNewRequiresInput(
	t *testing.T,
) {
	service :=
		&fakeService{}

	var output bytes.Buffer

	_, err :=
		New(
			service,
			nil,
			&output,
		)

	if err == nil {
		t.Fatal(
			"expected input error",
		)
	}
}

func TestNewRequiresOutput(
	t *testing.T,
) {
	service :=
		&fakeService{}

	_, err :=
		New(
			service,
			strings.NewReader(""),
			nil,
		)

	if err == nil {
		t.Fatal(
			"expected output error",
		)
	}
}

func TestMultipleQuestions(
	t *testing.T,
) {
	service :=
		&fakeService{
			results: map[string]agent.Result{
				"Ele sabe Go?": {
					Response:    "Sim, há evidências de experiência com Go.",
					HasResponse: true,
					Language:    domain.LanguagePortuguese,
				},
				"Me fale sobre o GGCompress": {
					Response:    "Sobre GGCompress: projeto em Go.",
					HasResponse: true,
					Language:    domain.LanguagePortuguese,
				},
			},
		}

	input :=
		strings.NewReader(
			"Ele sabe Go?\nMe fale sobre o GGCompress\nquit\n",
		)

	var output bytes.Buffer

	runner, err :=
		New(
			service,
			input,
			&output,
		)

	if err != nil {
		t.Fatal(err)
	}

	if err :=
		runner.Run(); err != nil {
		t.Fatal(err)
	}

	if len(service.calls) != 2 {
		t.Fatalf(
			"expected two calls, got %d",
			len(service.calls),
		)
	}

	value :=
		output.String()

	if !strings.Contains(
		value,
		"experiência com Go",
	) {
		t.Fatalf(
			"missing first response:\n%s",
			value,
		)
	}

	if !strings.Contains(
		value,
		"GGCompress",
	) {
		t.Fatalf(
			"missing second response:\n%s",
			value,
		)
	}
}
