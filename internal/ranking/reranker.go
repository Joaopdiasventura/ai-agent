package ranking

import (
	"ai-agent/internal/domain"
	"sort"
)

type MetadataWeights struct {
	BaseScoreMultiplier         float64
	LanguageMatchBoost          float64
	LanguageMismatchPenalty     float64
	CategoryMatchBoost          float64
	TemporalMatchBoost          float64
	TemporalMismatchPenalty     float64
	ProjectMatchBoost           float64
	ProjectMismatchPenalty      float64
	SpecificDocumentBoost       float64
	GenericComparisonPenalty    float64
	BothRetrievalSourcesBoost   float64
	LexicalOnlyExactSignalBoost float64
}

type MetadataReranker struct {
	Weights MetadataWeights
}

func DefaultMetadataReranker() MetadataReranker {
	return MetadataReranker{
		Weights: MetadataWeights{
			BaseScoreMultiplier:         100,
			LanguageMatchBoost:          40,
			LanguageMismatchPenalty:     500,
			CategoryMatchBoost:          18,
			TemporalMatchBoost:          14,
			TemporalMismatchPenalty:     90,
			ProjectMatchBoost:           24,
			ProjectMismatchPenalty:      140,
			SpecificDocumentBoost:       8,
			GenericComparisonPenalty:    25,
			BothRetrievalSourcesBoost:   12,
			LexicalOnlyExactSignalBoost: 6,
		},
	}
}

func (reranker MetadataReranker) Rerank(query domain.Query, results []domain.SearchResult) []domain.SearchResult {
	reranked := make([]domain.SearchResult, 0, len(results))

	for _, result := range results {
		if result.Document == nil {
			continue
		}

		score := result.Score * reranker.Weights.BaseScoreMultiplier
		metadataScore := 0.0
		penalties := make([]string, 0)

		if query.Language != "" {
			if result.Document.Language == query.Language {
				metadataScore += reranker.Weights.LanguageMatchBoost
			} else {
				metadataScore -= reranker.Weights.LanguageMismatchPenalty
				penalties = append(penalties, "language_mismatch")
			}
		}

		if query.Category != "" && result.Document.Category == query.Category {
			metadataScore += reranker.Weights.CategoryMatchBoost
		}

		if query.TemporalStatus != "" && query.TemporalStatus != domain.TemporalTimeless {
			if result.Document.TemporalStatus == query.TemporalStatus {
				metadataScore += reranker.Weights.TemporalMatchBoost
			} else if result.Document.TemporalStatus != domain.TemporalTimeless {
				metadataScore -= reranker.Weights.TemporalMismatchPenalty
				penalties = append(penalties, "temporal_mismatch")
			}
		}

		if query.Project != "" {
			if result.Document.Project == query.Project {
				metadataScore += reranker.Weights.ProjectMatchBoost
			} else if result.Document.Project != "" {
				metadataScore -= reranker.Weights.ProjectMismatchPenalty
				penalties = append(penalties, "project_mismatch")
			}
		}

		if result.Document.Subject != "" && result.Document.Subject != result.Document.Category {
			metadataScore += reranker.Weights.SpecificDocumentBoost
		}

		if result.Document.Category == "comparison" && query.Category == "project" {
			metadataScore -= reranker.Weights.GenericComparisonPenalty
			penalties = append(penalties, "generic_comparison")
		}

		if hasSource(result.Sources, "vector") && hasSource(result.Sources, "lexical") {
			metadataScore += reranker.Weights.BothRetrievalSourcesBoost
		}

		if result.VectorRank == 0 && result.LexicalRank > 0 && len(query.ExactTerms) > 0 {
			metadataScore += reranker.Weights.LexicalOnlyExactSignalBoost
		}

		result.MetadataScore = metadataScore
		result.FinalScore = score + metadataScore
		result.PenaltyReasons = append(result.PenaltyReasons, penalties...)
		reranked = append(reranked, result)
	}

	sort.Slice(reranked, func(firstIndex int, secondIndex int) bool {
		if reranked[firstIndex].FinalScore == reranked[secondIndex].FinalScore {
			return reranked[firstIndex].Document.ID < reranked[secondIndex].Document.ID
		}

		return reranked[firstIndex].FinalScore > reranked[secondIndex].FinalScore
	})

	return reranked
}

func hasSource(sources []string, source string) bool {
	for _, existing := range sources {
		if existing == source {
			return true
		}
	}

	return false
}
