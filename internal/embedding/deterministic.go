package embedding

import (
	"context"
	"encoding/binary"
	"errors"
	"hash/fnv"
	"math"
	"strings"
)

type DeterministicEmbedder struct {
	dimension int
}

func NewDeterministicEmbedder(dimension int) (*DeterministicEmbedder, error) {
	if dimension <= 0 {
		return nil, ErrInvalidDimension
	}

	return &DeterministicEmbedder{dimension: dimension}, nil
}

func (embedder *DeterministicEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if strings.TrimSpace(text) == "" {
		return nil, ErrEmptyText
	}

	vector := make([]float32, embedder.dimension)
	tokens := tokenFeatures(text)

	for _, token := range tokens {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		index, sign := featurePosition(token, embedder.dimension)
		vector[index] += sign
	}

	return vector, nil
}

func tokenFeatures(text string) []string {
	fields := strings.Fields(strings.ToLower(text))
	features := make([]string, 0, len(fields)*2)

	for _, field := range fields {
		field = strings.Trim(field, ".,;:!?()[]{}\"'")
		if field == "" {
			continue
		}

		features = append(features, field)
		if len([]rune(field)) > 4 {
			features = append(features, field[:4])
		}
	}

	if len(features) == 0 {
		return []string{text}
	}

	return features
}

func featurePosition(feature string, dimension int) (int, float32) {
	hasher := fnv.New64a()
	_, _ = hasher.Write([]byte(feature))
	sum := hasher.Sum64()

	index := int(sum % uint64(dimension))
	sign := float32(1)
	if binary.LittleEndian.Uint16([]byte{byte(sum >> 8), byte(sum >> 16)})%2 == 1 {
		sign = -1
	}

	return index, sign
}

func Normalize(vector []float32) ([]float32, error) {
	if len(vector) == 0 {
		return nil, ErrInvalidDimension
	}

	var squared float64
	for _, value := range vector {
		squared += float64(value * value)
	}

	if squared == 0 {
		return nil, ErrEmptyVector
	}

	norm := float32(math.Sqrt(squared))
	normalized := make([]float32, len(vector))

	for index, value := range vector {
		normalized[index] = value / norm
	}

	return normalized, nil
}

var ErrInvalidDimension = errors.New("invalid embedding dimension")
var ErrEmptyText = errors.New("empty text cannot be embedded")
var ErrEmptyVector = errors.New("embedding vector is empty")
