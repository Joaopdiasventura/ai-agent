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

func TestGroundedGeneratorOrdersProjectEvidenceNarratively(t *testing.T) {
	generator := NewGroundedGenerator()
	response, err := generator.Generate(context.Background(), domain.Query{
		Text:     "Me fale sobre o Auronix",
		Language: "pt",
		Category: "project",
		Project:  "auronix",
	}, []domain.Evidence{
		{DocumentID: "project-auronix-consistency-pt", Language: "pt", Category: "impact", Project: "auronix", Content: "No Auronix, João Paulo trabalhou com transferências confiáveis, versionamento otimista para evitar confirmações incompatíveis ao mesmo tempo e um ledger auditável para histórico verificável das movimentações.", Score: 10},
		{DocumentID: "project-auronix-infrastructure-pt", Language: "pt", Category: "project", Project: "auronix", Content: "No Auronix, Docker empacota os serviços, Kubernetes organiza sua execução e Terraform descreve a infraestrutura AWS EKS por arquivos versionados.", Score: 9},
		{DocumentID: "project-auronix-description-pt", Language: "pt", Category: "project", Project: "auronix", Content: "Auronix é uma plataforma financeira full stack criada para representar operações de um banco digital.", Score: 8},
	})

	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}

	if !strings.HasPrefix(response, "Auronix é uma plataforma financeira full stack") {
		t.Fatalf("Generate() = %q, want description first", response)
	}

	if strings.Contains(response, "No Auronix") {
		t.Fatalf("Generate() = %q, repeats project prefix", response)
	}
}

func TestGroundedGeneratorAvoidsRepeatedProjectPrefixes(t *testing.T) {
	generator := NewGroundedGenerator()
	response, err := generator.Generate(context.Background(), domain.Query{
		Text:     "Me fale sobre o X Tube",
		Language: "pt",
		Category: "project",
		Project:  "x-tube",
	}, []domain.Evidence{
		{DocumentID: "project-xtube-leadership-pt", Language: "pt", Category: "impact", Project: "x-tube", Content: "No X Tube, João Paulo teve liderança técnica no desenho do fluxo de upload, processamento e entrega de vídeos.", Score: 10},
		{DocumentID: "project-xtube-processing-pt", Language: "pt", Category: "project", Project: "x-tube", Content: "João Paulo desenvolveu em Go o serviço do X Tube responsável por receber eventos, processar vídeos com FFmpeg e salvar os artefatos gerados no Amazon S3.", Score: 9},
		{DocumentID: "project-xtube-description-pt", Language: "pt", Category: "project", Project: "x-tube", Content: "X Tube é um serviço de streaming desenvolvido por uma equipe de três pessoas.", Score: 8},
	})

	if err != nil {
		t.Fatalf("Generate() returned error: %v", err)
	}

	if !strings.HasPrefix(response, "X Tube é um serviço de streaming") {
		t.Fatalf("Generate() = %q, want description first", response)
	}

	if strings.Contains(response, "No X Tube") {
		t.Fatalf("Generate() = %q, repeats project prefix", response)
	}
}

func TestGroundedGeneratorRequiresEvidence(t *testing.T) {
	_, err := NewGroundedGenerator().Generate(context.Background(), domain.Query{}, nil)

	if !errors.Is(err, ErrNoEvidence) {
		t.Fatalf("Generate() error = %v, want %v", err, ErrNoEvidence)
	}
}
