package indexing

import (
	"ai-agent/internal/domain"
	"context"
	"errors"
	"testing"
	"time"
)

type fixedEmbedder struct {
	vectors [][]float32
	index   int
}

func (embedder *fixedEmbedder) Embed(context.Context, string) ([]float32, error) {
	vector := embedder.vectors[embedder.index]
	embedder.index++
	return vector, nil
}

func TestBuildCreatesDeterministicIndex(t *testing.T) {
	documents := []*domain.Document{
		{ID: "b-pt", Language: "pt", Category: "profile", Subject: "b", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"b"}, Content: "second"},
		{ID: "a-pt", Language: "pt", Category: "profile", Subject: "a", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"a"}, Content: "first"},
	}

	embedder := &fixedEmbedder{
		vectors: [][]float32{
			{3, 0},
			{0, 4},
		},
	}

	index, manifest, report, err := Build(context.Background(), documents, embedder, "test-model", time.Date(2026, 7, 19, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("Build() returned error: %v", err)
	}

	if len(index.Entries) != 2 {
		t.Fatalf("Build() entries = %d, want 2", len(index.Entries))
	}

	if index.Entries[0].Document.ID != "a-pt" {
		t.Fatalf("first entry ID = %q, want a-pt", index.Entries[0].Document.ID)
	}

	if index.Entries[0].Embedding[0] != 1 || index.Entries[1].Embedding[1] != 1 {
		t.Fatalf("Build() did not normalize embeddings: %#v", index.Entries)
	}

	if manifest.DocumentCount != 2 || report.Dimension != 2 {
		t.Fatalf("manifest/report mismatch: %#v %#v", manifest, report)
	}
}

func TestBuildRejectsInvalidInputs(t *testing.T) {
	_, _, _, err := Build(context.Background(), nil, &fixedEmbedder{}, "test-model", time.Now())
	if !errors.Is(err, ErrNoDocuments) {
		t.Fatalf("Build() error = %v, want %v", err, ErrNoDocuments)
	}

	documents := []*domain.Document{
		{ID: "a-pt", Language: "pt", Category: "profile", Subject: "a", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"a"}, Content: "first"},
		{ID: "b-pt", Language: "pt", Category: "profile", Subject: "b", TemporalStatus: domain.TemporalTimeless, Keywords: []string{"b"}, Content: "second"},
	}

	embedder := &fixedEmbedder{
		vectors: [][]float32{
			{1, 0},
			{1, 0, 0},
		},
	}

	_, _, _, err = Build(context.Background(), documents, embedder, "test-model", time.Now())
	if err == nil {
		t.Fatal("Build() returned nil error for incompatible dimensions")
	}
}
