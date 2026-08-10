package retrieval

import (
	"errors"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
	"ai-agent/internal/knowledge"
)

const DefaultLimit = 20

type Retriever interface {
	Search(
		currentQuery domain.Query,
		limit int,
	) Ranking
}

type HybridRetriever struct {
	lexical *LexicalRetriever
	entity  *EntityRetriever
	concept *ConceptRetriever
	fuzzy   *FuzzyRetriever
}

func NewHybridRetriever(
	base *knowledge.Knowledge,
	currentIndex *searchindex.Index,
) (*HybridRetriever, error) {
	if base == nil {
		return nil, errors.New(
			"knowledge base is required",
		)
	}

	if currentIndex == nil {
		return nil, errors.New(
			"index is required",
		)
	}

	entityRetriever, err :=
		NewEntityRetriever(
			base,
			currentIndex,
		)

	if err != nil {
		return nil, err
	}

	conceptRetriever, err :=
		NewConceptRetriever(
			base,
			currentIndex,
		)

	if err != nil {
		return nil, err
	}

	return &HybridRetriever{
		lexical: NewLexicalRetriever(
			currentIndex,
		),
		entity:  entityRetriever,
		concept: conceptRetriever,
		fuzzy: NewFuzzyRetriever(
			currentIndex,
		),
	}, nil
}

func (r *HybridRetriever) Search(
	currentQuery domain.Query,
	limit int,
) Result {
	if limit <= 0 {
		limit = DefaultLimit
	}

	return Result{
		Query: currentQuery,
		Rankings: []Ranking{
			r.lexical.Search(
				currentQuery,
				limit,
			),
			r.entity.Search(
				currentQuery,
				limit,
			),
			r.concept.Search(
				currentQuery,
				limit,
			),
			r.fuzzy.Search(
				currentQuery,
				limit,
			),
		},
	}
}
