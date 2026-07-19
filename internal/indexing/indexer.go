package indexing

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/embedding"
	"ai-agent/internal/vectorindex"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"sort"
	"time"
)

type Embedder interface {
	Embed(ctx context.Context, text string) ([]float32, error)
}

type Report struct {
	DocumentCount int
	Dimension     int
	Model         string
	BaseHash      string
}

var ErrNoDocuments = errors.New("cannot index empty document set")

func Build(ctx context.Context, documents []*domain.Document, embedder Embedder, model string, generatedAt time.Time) (vectorindex.Index, vectorindex.Manifest, Report, error) {
	if len(documents) == 0 {
		return vectorindex.Index{}, vectorindex.Manifest{}, Report{}, ErrNoDocuments
	}

	orderedDocuments := orderedDocumentCopies(documents)
	baseHash, err := BaseHash(orderedDocuments)
	if err != nil {
		return vectorindex.Index{}, vectorindex.Manifest{}, Report{}, err
	}

	entries := make([]vectorindex.Entry, 0, len(orderedDocuments))
	dimension := 0

	for _, document := range orderedDocuments {
		vector, err := embedder.Embed(ctx, document.Content)
		if err != nil {
			return vectorindex.Index{}, vectorindex.Manifest{}, Report{}, err
		}

		normalizedVector, err := embedding.Normalize(vector)
		if err != nil {
			return vectorindex.Index{}, vectorindex.Manifest{}, Report{}, err
		}

		if dimension == 0 {
			dimension = len(normalizedVector)
		}

		if len(normalizedVector) != dimension {
			return vectorindex.Index{}, vectorindex.Manifest{}, Report{}, embedding.ErrInvalidDimension
		}

		entries = append(entries, vectorindex.Entry{
			Document:    document,
			Embedding:   normalizedVector,
			ContentHash: ContentHash(document),
		})
	}

	index := vectorindex.Index{
		Version:   vectorindex.Version,
		Model:     model,
		Dimension: dimension,
		BaseHash:  baseHash,
		Entries:   entries,
	}

	manifest := vectorindex.Manifest{
		Version:       vectorindex.Version,
		Model:         model,
		Dimension:     dimension,
		DocumentCount: len(entries),
		BaseHash:      baseHash,
		GeneratedAt:   generatedAt.UTC().Format(time.RFC3339),
	}

	report := Report{
		DocumentCount: len(entries),
		Dimension:     dimension,
		Model:         model,
		BaseHash:      baseHash,
	}

	return index, manifest, report, nil
}

func BaseHash(documents []domain.Document) (string, error) {
	content, err := json.Marshal(documents)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(content)
	return hex.EncodeToString(sum[:]), nil
}

func ContentHash(document domain.Document) string {
	content := document.ID + "\n" + document.Language + "\n" + document.Category + "\n" + document.Subject + "\n" + document.Project + "\n" + document.TemporalStatus + "\n" + document.Content
	sum := sha256.Sum256([]byte(content))
	return hex.EncodeToString(sum[:])
}

func orderedDocumentCopies(documents []*domain.Document) []domain.Document {
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
