package ranking

import (
	"ai-agent/internal/domain"
	"testing"
)

func TestFuseCombinesDocumentsFromBothSources(t *testing.T) {
	vectorResults := []domain.SearchResult{
		rankingResult("a"),
		rankingResult("b"),
	}
	lexicalResults := []domain.SearchResult{
		rankingResult("b"),
		rankingResult("c"),
	}

	results := FuseWithK(vectorResults, lexicalResults, 3, 10)

	if len(results) != 3 {
		t.Fatalf("Fuse() returned %d results, want 3", len(results))
	}

	if results[0].Document.ID != "b" {
		t.Fatalf("Fuse() first result = %q, want b", results[0].Document.ID)
	}

	if results[0].VectorRank != 2 || results[0].LexicalRank != 1 {
		t.Fatalf("Fuse() ranks = vector %d lexical %d, want 2 and 1", results[0].VectorRank, results[0].LexicalRank)
	}
}

func TestFuseKeepsSingleSourceDocumentsAndLimit(t *testing.T) {
	results := FuseWithK([]domain.SearchResult{
		rankingResult("a"),
		rankingResult("b"),
	}, nil, 1, 10)

	if len(results) != 1 {
		t.Fatalf("Fuse() returned %d results, want 1", len(results))
	}

	if results[0].Document.ID != "a" || results[0].FusedRank != 1 {
		t.Fatalf("Fuse() returned %#v, want document a at fused rank 1", results[0])
	}
}

func TestFuseUsesDeterministicTieBreakAndHandlesEmptyLists(t *testing.T) {
	results := FuseWithK([]domain.SearchResult{
		rankingResult("b"),
		rankingResult("a"),
	}, []domain.SearchResult{
		rankingResult("a"),
		rankingResult("b"),
	}, 2, 10)

	if results[0].Document.ID != "a" {
		t.Fatalf("Fuse() first result = %q, want a on score tie", results[0].Document.ID)
	}

	if empty := FuseWithK(nil, nil, 5, 10); len(empty) != 0 {
		t.Fatalf("Fuse() returned %d results for empty rankings", len(empty))
	}
}

func rankingResult(id string) domain.SearchResult {
	return domain.SearchResult{
		Document: &domain.Document{
			ID:             id,
			Language:       "pt",
			Category:       "profile",
			Subject:        id,
			TemporalStatus: domain.TemporalTimeless,
			Keywords:       []string{id},
			Content:        id,
		},
	}
}
