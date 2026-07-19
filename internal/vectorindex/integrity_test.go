package vectorindex

import (
	"bytes"
	"os"
	"testing"
)

func TestGeneratedIndexArtifactsAreSynchronized(t *testing.T) {
	publicIndex, err := os.ReadFile("../../data/generated/knowledge.index.json")
	if err != nil {
		t.Fatalf("failed to read public index: %v", err)
	}

	embeddedIndex, err := embeddedFiles.ReadFile("generated/knowledge.index.json")
	if err != nil {
		t.Fatalf("failed to read embedded index: %v", err)
	}

	if !bytes.Equal(publicIndex, embeddedIndex) {
		t.Fatal("public generated index and embedded index are not synchronized")
	}

	publicManifest, err := os.ReadFile("../../data/generated/manifest.json")
	if err != nil {
		t.Fatalf("failed to read public manifest: %v", err)
	}

	embeddedManifest, err := embeddedFiles.ReadFile("generated/manifest.json")
	if err != nil {
		t.Fatalf("failed to read embedded manifest: %v", err)
	}

	if !bytes.Equal(publicManifest, embeddedManifest) {
		t.Fatal("public generated manifest and embedded manifest are not synchronized")
	}
}
