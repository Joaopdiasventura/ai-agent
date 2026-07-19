package retrieval

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/vectorindex"
	"errors"
	"testing"
)

func TestVectorSearchRanksByDotProduct(t *testing.T) {
	index := testVectorIndex([]vectorindex.Entry{
		testEntry("same", []float32{1, 0}),
		testEntry("opposite", []float32{-1, 0}),
		testEntry("near", []float32{0.8, 0.2}),
	})

	results, err := VectorSearch(index, []float32{1, 0}, 3)
	if err != nil {
		t.Fatalf("VectorSearch() returned error: %v", err)
	}

	expectedIDs := []string{"same", "near", "opposite"}
	for index, expectedID := range expectedIDs {
		if results[index].Document.ID != expectedID {
			t.Fatalf("result[%d] = %q, want %q", index, results[index].Document.ID, expectedID)
		}
	}
}

func TestVectorSearchUsesDeterministicTieBreak(t *testing.T) {
	index := testVectorIndex([]vectorindex.Entry{
		testEntry("b", []float32{1, 0}),
		testEntry("a", []float32{1, 0}),
	})

	results, err := VectorSearch(index, []float32{1, 0}, 2)
	if err != nil {
		t.Fatalf("VectorSearch() returned error: %v", err)
	}

	if results[0].Document.ID != "a" || results[1].Document.ID != "b" {
		t.Fatalf("VectorSearch() results = %q, %q; want a, b", results[0].Document.ID, results[1].Document.ID)
	}
}

func TestVectorSearchHandlesLimitAndEmptyIndex(t *testing.T) {
	index := testVectorIndex([]vectorindex.Entry{
		testEntry("a", []float32{1, 0}),
		testEntry("b", []float32{0, 1}),
	})

	results, err := VectorSearch(index, []float32{1, 0}, 1)
	if err != nil {
		t.Fatalf("VectorSearch() returned error: %v", err)
	}

	if len(results) != 1 || results[0].VectorRank != 1 {
		t.Fatalf("VectorSearch() returned %#v, want one ranked result", results)
	}

	results, err = VectorSearch(vectorindex.Index{Dimension: 2}, []float32{1, 0}, 5)
	if err != nil {
		t.Fatalf("VectorSearch() returned error for empty index: %v", err)
	}

	if len(results) != 0 {
		t.Fatalf("VectorSearch() returned %d results for empty index", len(results))
	}
}

func TestVectorSearchRejectsInvalidDimension(t *testing.T) {
	index := testVectorIndex([]vectorindex.Entry{
		testEntry("a", []float32{1, 0}),
	})

	if _, err := VectorSearch(index, []float32{1}, 1); !errors.Is(err, ErrDimensionMismatch) {
		t.Fatalf("VectorSearch() error = %v, want %v", err, ErrDimensionMismatch)
	}
}

func testVectorIndex(entries []vectorindex.Entry) vectorindex.Index {
	return vectorindex.Index{
		Version:   vectorindex.Version,
		Model:     "test",
		Dimension: 2,
		Entries:   entries,
	}
}

func testEntry(id string, vector []float32) vectorindex.Entry {
	return vectorindex.Entry{
		Document: domain.Document{
			ID:             id,
			Language:       "pt",
			Category:       "profile",
			Subject:        id,
			TemporalStatus: domain.TemporalTimeless,
			Keywords:       []string{id},
			Content:        id,
		},
		Embedding: vector,
	}
}
