package retrieval

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/embedding"
	"ai-agent/internal/ranking"
	"ai-agent/internal/vectorindex"
	"context"
	"errors"
)

type HybridRetriever struct {
	index vectorindex.Index
}

var ErrMissingQueryVector = errors.New("query embedding vector is required")

func NewHybridRetriever(index vectorindex.Index) HybridRetriever {
	return HybridRetriever{index: index}
}

func (retriever HybridRetriever) Search(ctx context.Context, query domain.Query, limit int) ([]domain.SearchResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if len(query.Vector) == 0 {
		return nil, ErrMissingQueryVector
	}

	normalizedVector, err := embedding.Normalize(query.Vector)
	if err != nil {
		return nil, err
	}

	vectorResults, err := VectorSearch(retriever.index, normalizedVector, expandedLimit(limit))
	if err != nil {
		return nil, err
	}

	lexicalResults := LexicalSearch(retriever.index, query, expandedLimit(limit))
	return ranking.Fuse(vectorResults, lexicalResults, limit), nil
}

func expandedLimit(limit int) int {
	if limit <= 0 {
		return 0
	}

	return limit * 4
}
