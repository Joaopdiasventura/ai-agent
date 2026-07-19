package agent

import (
	"ai-agent/internal/domain"
	"context"
	"errors"
)

type Embedder interface {
	Embed(ctx context.Context, text string) ([]float32, error)
}

type Retriever interface {
	Search(ctx context.Context, query domain.Query, limit int) ([]domain.SearchResult, error)
}

type Reranker interface {
	Rerank(query domain.Query, results []domain.SearchResult) []domain.SearchResult
}

type Generator interface {
	Generate(ctx context.Context, query domain.Query, evidence []domain.Evidence) (string, error)
}

type Service struct {
	embedder  Embedder
	retriever Retriever
	reranker  Reranker
	generator Generator
}

var ErrMissingDependency = errors.New("missing agent dependency")

func NewService(embedder Embedder, retriever Retriever, reranker Reranker, generator Generator) (*Service, error) {
	if embedder == nil || retriever == nil || reranker == nil || generator == nil {
		return nil, ErrMissingDependency
	}

	return &Service{
		embedder:  embedder,
		retriever: retriever,
		reranker:  reranker,
		generator: generator,
	}, nil
}
