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

type RetrievalResult struct {
	Query      domain.Query
	Results    []domain.SearchResult
	Evidence   []domain.Evidence
	Confidence confidence.Assessment
}

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
	retrievalResult, err := service.Retrieve(ctx, question, limit)
	if err != nil {
		query := queryanalysis.Analyze(question)
		return "", false, query.Language, err
	}

	if retrievalResult.Confidence.Level == confidence.Low {
		return "", false, retrievalResult.Query.Language, nil
	}

	response, err := service.generator.Generate(ctx, retrievalResult.Query, retrievalResult.Evidence)
	if err != nil {
		return "", false, retrievalResult.Query.Language, err
	}

	return response, true, retrievalResult.Query.Language, nil
}

func (service *Service) Retrieve(ctx context.Context, question string, limit int) (RetrievalResult, error) {
	query := queryanalysis.Analyze(question)

	vector, err := service.embedder.Embed(ctx, query.Text)
	if err != nil {
		return RetrievalResult{Query: query}, err
	}

	query.Vector = vector

	results, err := service.retriever.Search(ctx, query, candidateLimit(limit))
	if err != nil {
		return RetrievalResult{Query: query}, err
	}

	reranked := service.reranker.Rerank(query, results)
	selectedEvidence := service.evidenceSelector.Select(query, reranked)
	assessment := service.confidencePolicy.Assess(query, reranked, selectedEvidence)

	return RetrievalResult{
		Query:      query,
		Results:    reranked,
		Evidence:   selectedEvidence,
		Confidence: assessment,
	}, nil
}

func candidateLimit(limit int) int {
	if limit <= 0 {
		return 0
	}

	return limit * 16
}
