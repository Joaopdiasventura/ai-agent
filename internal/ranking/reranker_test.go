package ranking

import (
	"ai-agent/internal/domain"
	"testing"
)

func TestMetadataRerankerPrioritizesCompatibleMetadata(t *testing.T) {
	reranker := DefaultMetadataReranker()
	query := domain.Query{
		Language:       "pt",
		Category:       "education",
		TemporalStatus: domain.TemporalPast,
	}

	results := []domain.SearchResult{
		metadataResult("identity-basic-pt", "pt", "identity", "", domain.TemporalTimeless, 1),
		metadataResult("education-etec-pt", "pt", "education", "", domain.TemporalPast, 0.8),
	}

	reranked := reranker.Rerank(query, results)

	if reranked[0].Document.ID != "education-etec-pt" {
		t.Fatalf("Rerank() first result = %q, want education-etec-pt", reranked[0].Document.ID)
	}
}

func TestMetadataRerankerPenalizesLanguageAndProjectMismatch(t *testing.T) {
	reranker := DefaultMetadataReranker()
	query := domain.Query{
		Language: "pt",
		Category: "project",
		Project:  "auronix",
	}

	results := []domain.SearchResult{
		metadataResult("project-xtube-pt", "pt", "project", "x-tube", domain.TemporalTimeless, 1),
		metadataResult("project-auronix-en", "en", "project", "auronix", domain.TemporalTimeless, 1),
		metadataResult("project-auronix-pt", "pt", "project", "auronix", domain.TemporalTimeless, 0.8),
	}

	reranked := reranker.Rerank(query, results)

	if reranked[0].Document.ID != "project-auronix-pt" {
		t.Fatalf("Rerank() first result = %q, want project-auronix-pt", reranked[0].Document.ID)
	}
}

func TestMetadataRerankerPenalizesGenericComparisonForProjectQuestion(t *testing.T) {
	reranker := DefaultMetadataReranker()
	query := domain.Query{
		Language: "pt",
		Category: "project",
	}

	results := []domain.SearchResult{
		metadataResult("project-comparison-best-pt", "pt", "comparison", "", domain.TemporalTimeless, 1),
		metadataResult("project-auronix-description-pt", "pt", "project", "auronix", domain.TemporalTimeless, 0.9),
	}

	reranked := reranker.Rerank(query, results)

	if reranked[0].Document.ID != "project-auronix-description-pt" {
		t.Fatalf("Rerank() first result = %q, want project-auronix-description-pt", reranked[0].Document.ID)
	}
}

func metadataResult(id string, language string, category string, project string, temporalStatus string, score float64) domain.SearchResult {
	return domain.SearchResult{
		Document: &domain.Document{
			ID:             id,
			Language:       language,
			Category:       category,
			Subject:        id,
			Project:        project,
			TemporalStatus: temporalStatus,
			Keywords:       []string{id},
			Content:        id,
		},
		Score:   score,
		Sources: []string{"vector", "lexical"},
	}
}
