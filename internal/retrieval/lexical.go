package retrieval

import (
	"strconv"

	"ai-agent/internal/domain"
	"ai-agent/internal/index"
)

type LexicalRetriever struct {
	bm25 *BM25F
}

func NewLexicalRetriever(
	currentIndex *index.Index,
) *LexicalRetriever {
	return &LexicalRetriever{
		bm25: NewBM25F(
			currentIndex,
			DefaultBM25FConfig(),
		),
	}
}

func (r *LexicalRetriever) Search(
	currentQuery domain.Query,
	limit int,
) Ranking {
	terms :=
		buildLexicalTerms(
			currentQuery,
		)

	if len(terms) == 0 {
		return Ranking{
			Source: SourceLexical,
		}
	}

	merged := make(
		map[domain.FactID]Candidate,
	)

	for _, targetLanguage := range retrievalLanguages(
		currentQuery.Language,
	) {
		candidates :=
			r.bm25.Search(
				targetLanguage,
				terms,
				limit*2,
			)

		for _, candidate := range candidates {
			mergeCandidateMaximum(
				merged,
				candidate,
			)
		}
	}

	return Ranking{
		Source: SourceLexical,
		Candidates: candidatesFromMap(
			merged,
			limit,
		),
	}
}

func buildLexicalTerms(
	currentQuery domain.Query,
) []WeightedTerm {
	result := make(
		[]WeightedTerm,
		0,
		len(currentQuery.Terms),
	)

	seen := make(
		map[string]struct{},
	)

	for _, term := range currentQuery.Terms {
		if term == "" {
			continue
		}

		if _, exists := seen[term]; exists {
			continue
		}

		seen[term] = struct{}{}

		result = append(
			result,
			WeightedTerm{
				Value: term,
				Weight: lexicalTermWeight(
					term,
				),
			},
		)
	}

	return result
}

func lexicalTermWeight(
	term string,
) float64 {
	if _, err :=
		strconv.ParseFloat(
			term,
			64,
		); err == nil {
		return 1.35
	}

	length := len([]rune(term))

	if length <= 2 {
		return 1.15
	}

	return 1
}
