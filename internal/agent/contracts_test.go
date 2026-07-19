package agent

import (
	"ai-agent/internal/domain"
	"context"
	"errors"
	"testing"
)

type testEmbedder struct{}

func (testEmbedder) Embed(context.Context, string) ([]float32, error) {
	return []float32{1}, nil
}

type testRetriever struct{}

func (testRetriever) Search(context.Context, domain.Query, int) ([]domain.SearchResult, error) {
	return nil, nil
}

type testReranker struct{}

func (testReranker) Rerank(domain.Query, []domain.SearchResult) []domain.SearchResult {
	return nil
}

type testGenerator struct{}

func (testGenerator) Generate(context.Context, domain.Query, []domain.Evidence) (string, error) {
	return "", nil
}

func TestNewServiceRequiresDependencies(t *testing.T) {
	_, err := NewService(nil, testRetriever{}, testReranker{}, testGenerator{})

	if !errors.Is(err, ErrMissingDependency) {
		t.Fatalf("NewService() error = %v, want %v", err, ErrMissingDependency)
	}
}

func TestNewServiceAcceptsCompleteDependencies(t *testing.T) {
	service, err := NewService(testEmbedder{}, testRetriever{}, testReranker{}, testGenerator{})

	if err != nil {
		t.Fatalf("NewService() returned error: %v", err)
	}

	if service == nil {
		t.Fatal("NewService() returned nil service")
	}
}
