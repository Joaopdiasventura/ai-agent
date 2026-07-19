package vectorindex

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"sync"
)

//go:embed generated/knowledge.index.json generated/manifest.json
var embeddedFiles embed.FS

var (
	loadedIndex Index
	loadErr     error
	loadOnce    sync.Once
)

var ErrInvalidIndex = errors.New("invalid embedded vector index")

func LoadEmbedded() (Index, error) {
	loadOnce.Do(func() {
		loadedIndex, loadErr = loadEmbedded()
	})

	if loadErr != nil {
		return Index{}, loadErr
	}

	return cloneIndex(loadedIndex), nil
}

func loadEmbedded() (Index, error) {
	indexContent, err := embeddedFiles.ReadFile("generated/knowledge.index.json")
	if err != nil {
		return Index{}, err
	}

	manifestContent, err := embeddedFiles.ReadFile("generated/manifest.json")
	if err != nil {
		return Index{}, err
	}

	var index Index
	if err := json.Unmarshal(indexContent, &index); err != nil {
		return Index{}, err
	}

	var manifest Manifest
	if err := json.Unmarshal(manifestContent, &manifest); err != nil {
		return Index{}, err
	}

	if err := Validate(index, manifest); err != nil {
		return Index{}, err
	}

	return index, nil
}

func Validate(index Index, manifest Manifest) error {
	if index.Version != Version || manifest.Version != Version {
		return fmt.Errorf("%w: unsupported version", ErrInvalidIndex)
	}

	if index.Dimension <= 0 || index.Dimension != manifest.Dimension {
		return fmt.Errorf("%w: invalid dimension", ErrInvalidIndex)
	}

	if index.Model == "" || index.Model != manifest.Model {
		return fmt.Errorf("%w: invalid model", ErrInvalidIndex)
	}

	if len(index.Entries) == 0 || len(index.Entries) != manifest.DocumentCount {
		return fmt.Errorf("%w: invalid document count", ErrInvalidIndex)
	}

	if index.BaseHash == "" || index.BaseHash != manifest.BaseHash {
		return fmt.Errorf("%w: invalid base hash", ErrInvalidIndex)
	}

	seenIDs := make(map[string]struct{}, len(index.Entries))

	for _, entry := range index.Entries {
		if entry.Document.ID == "" {
			return fmt.Errorf("%w: empty document id", ErrInvalidIndex)
		}

		if _, exists := seenIDs[entry.Document.ID]; exists {
			return fmt.Errorf("%w: duplicate document id %s", ErrInvalidIndex, entry.Document.ID)
		}

		seenIDs[entry.Document.ID] = struct{}{}

		if len(entry.Embedding) != index.Dimension {
			return fmt.Errorf("%w: embedding dimension mismatch for %s", ErrInvalidIndex, entry.Document.ID)
		}
	}

	expectedHash, err := baseHash(indexDocumentsFromKnowledge())
	if err != nil {
		return err
	}

	if expectedHash != index.BaseHash {
		return fmt.Errorf("%w: base hash mismatch", ErrInvalidIndex)
	}

	return nil
}

func indexDocumentsFromKnowledge() []domain.Document {
	documents := knowledge.Documents()
	copies := make([]domain.Document, 0, len(documents))

	for _, document := range documents {
		if document == nil {
			continue
		}

		copies = append(copies, *document)
	}

	sort.Slice(copies, func(firstIndex int, secondIndex int) bool {
		return copies[firstIndex].ID < copies[secondIndex].ID
	})

	return copies
}

func baseHash(documents []domain.Document) (string, error) {
	content, err := json.Marshal(documents)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:]), nil
}

func cloneIndex(index Index) Index {
	cloned := Index{
		Version:   index.Version,
		Model:     index.Model,
		Dimension: index.Dimension,
		BaseHash:  index.BaseHash,
		Entries:   make([]Entry, len(index.Entries)),
	}

	for indexPosition, entry := range index.Entries {
		cloned.Entries[indexPosition] = entry
		cloned.Entries[indexPosition].Embedding = append([]float32(nil), entry.Embedding...)
		cloned.Entries[indexPosition].Document.Keywords = append([]string(nil), entry.Document.Keywords...)
	}

	return cloned
}
