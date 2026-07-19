package vectorindex

import (
	"errors"
	"testing"
)

func TestLoadEmbeddedLoadsImmutableIndex(t *testing.T) {
	index, err := LoadEmbedded()
	if err != nil {
		t.Fatalf("LoadEmbedded() returned error: %v", err)
	}

	if index.Version != Version {
		t.Fatalf("index version = %q, want %q", index.Version, Version)
	}

	if len(index.Entries) != 134 {
		t.Fatalf("index entries = %d, want 134", len(index.Entries))
	}

	original := index.Entries[0].Embedding[0]
	index.Entries[0].Embedding[0] = 999

	reloaded, err := LoadEmbedded()
	if err != nil {
		t.Fatalf("LoadEmbedded() returned error: %v", err)
	}

	if reloaded.Entries[0].Embedding[0] != original {
		t.Fatal("LoadEmbedded() returned mutable shared embedding data")
	}
}

func TestValidateRejectsInvalidManifest(t *testing.T) {
	index, err := LoadEmbedded()
	if err != nil {
		t.Fatalf("LoadEmbedded() returned error: %v", err)
	}

	manifest := Manifest{
		Version:       Version,
		Model:         index.Model,
		Dimension:     index.Dimension + 1,
		DocumentCount: len(index.Entries),
		BaseHash:      index.BaseHash,
	}

	if err := Validate(index, manifest); !errors.Is(err, ErrInvalidIndex) {
		t.Fatalf("Validate() error = %v, want %v", err, ErrInvalidIndex)
	}
}
