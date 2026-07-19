package generation

import (
	"ai-agent/internal/domain"
	"context"
	"errors"
	"strings"
	"testing"
)

func TestGroundedGeneratorReturnsDirectFactForSimpleQuestion(t *testing.T) {
	generator := NewGroundedGenerator()
	response, err := generator.Generate(context.Background(), domain.Query{
		Text:     "Qual é o email dele?",
		Language: "pt",
		Category: "contact",
	}, []domain.Evidence{
		{DocumentID: "contact-email-pt", Language: "pt", Category: "contact", Content: "Para entrar em contato com João Paulo por email, use joaopdias.dev@gmail.com."},
	})

	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}

	if response != "Para entrar em contato com João Paulo por email, use joaopdias.dev@gmail.com." {
		t.Fatalf("Generate() = %q", response)
	}
}

func TestGroundedGeneratorSynthesizesComparativeAnswer(t *testing.T) {
	generator := NewGroundedGenerator()
	response, err := generator.Generate(context.Background(), domain.Query{
		Text:     "Qual projeto demonstra liderança?",
		Language: "pt",
		Category: "project",
	}, []domain.Evidence{
		{DocumentID: "project-xtube-leadership-pt", Language: "pt", Category: "impact", Project: "x-tube", Content: "No X Tube, João Paulo teve liderança técnica no desenho do fluxo de upload, processamento e entrega de vídeos."},
		{DocumentID: "project-xtube-processing-pt", Language: "pt", Category: "project", Project: "x-tube", Content: "João Paulo desenvolveu em Go o serviço do X Tube responsável por receber eventos, processar vídeos com FFmpeg e salvar os artefatos gerados no Amazon S3."},
	})

	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}

	expectedTerms := []string{"X Tube", "liderança técnica", "upload", "Amazon S3"}
	for _, term := range expectedTerms {
		if !strings.Contains(response, term) {
			t.Fatalf("Generate() = %q, want term %q", response, term)
		}
	}

	if strings.Contains(response, "DocumentID") || strings.Contains(response, "score") {
		t.Fatalf("Generate() exposed internal data: %q", response)
	}
}

func TestGroundedGeneratorRequiresEvidence(t *testing.T) {
	_, err := NewGroundedGenerator().Generate(context.Background(), domain.Query{}, nil)

	if !errors.Is(err, ErrNoEvidence) {
		t.Fatalf("Generate() error = %v, want %v", err, ErrNoEvidence)
	}
}
