package agent

import (
	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
	"ai-agent/internal/evidence"
	"ai-agent/internal/queryanalysis"
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
	embedder         Embedder
	retriever        Retriever
	reranker         Reranker
	generator        Generator
	evidenceSelector evidence.Selector
	confidencePolicy confidence.Policy
}

var ErrMissingDependency = errors.New("missing agent dependency")

func NewService(embedder Embedder, retriever Retriever, reranker Reranker, generator Generator) (*Service, error) {
	if embedder == nil || retriever == nil || reranker == nil || generator == nil {
		return nil, ErrMissingDependency
	}

	return &Service{
		embedder:         embedder,
		retriever:        retriever,
		reranker:         reranker,
		generator:        generator,
		evidenceSelector: evidence.DefaultSelector(),
		confidencePolicy: confidence.DefaultPolicy(),
	}, nil
}

func (service *Service) Answer(ctx context.Context, question string, limit int) (string, bool, string, error) {
	query := queryanalysis.Analyze(question)

	vector, err := service.embedder.Embed(ctx, query.Text)
	if err != nil {
		return "", false, query.Language, err
	}

	query.Vector = vector

	results, err := service.retriever.Search(ctx, query, candidateLimit(limit))
	if err != nil {
		return "", false, query.Language, err
	}

	reranked := service.reranker.Rerank(query, results)
	selectedEvidence := service.evidenceSelector.Select(query, reranked)
	assessment := service.confidencePolicy.Assess(query, reranked, selectedEvidence)

	if assessment.Level == confidence.Low {
		return "", false, query.Language, nil
	}

	response, err := service.generator.Generate(ctx, query, selectedEvidence)
	if err != nil {
		return "", false, query.Language, err
	}

	return response, true, query.Language, nil
}

func candidateLimit(limit int) int {
	if limit <= 0 {
		return 0
	}

	return limit * 16
}
