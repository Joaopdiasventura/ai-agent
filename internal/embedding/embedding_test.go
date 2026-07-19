package embedding

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDeterministicEmbedderReturnsStableVector(t *testing.T) {
	embedder, err := NewDeterministicEmbedder(16)
	if err != nil {
		t.Fatalf("NewDeterministicEmbedder() returned error: %v", err)
	}

	first, err := embedder.Embed(context.Background(), "Auronix financial platform")
	if err != nil {
		t.Fatalf("Embed() returned error: %v", err)
	}

	second, err := embedder.Embed(context.Background(), "Auronix financial platform")
	if err != nil {
		t.Fatalf("Embed() returned error: %v", err)
	}

	if len(first) != 16 {
		t.Fatalf("Embed() length = %d, want 16", len(first))
	}

	for index := range first {
		if first[index] != second[index] {
			t.Fatalf("Embed() is not deterministic at index %d", index)
		}
	}
}

func TestDeterministicEmbedderRejectsInvalidInput(t *testing.T) {
	if _, err := NewDeterministicEmbedder(0); !errors.Is(err, ErrInvalidDimension) {
		t.Fatalf("NewDeterministicEmbedder() error = %v, want %v", err, ErrInvalidDimension)
	}

	embedder, err := NewDeterministicEmbedder(8)
	if err != nil {
		t.Fatalf("NewDeterministicEmbedder() returned error: %v", err)
	}

	if _, err := embedder.Embed(context.Background(), " "); !errors.Is(err, ErrEmptyText) {
		t.Fatalf("Embed() error = %v, want %v", err, ErrEmptyText)
	}
}

func TestNormalizeRejectsEmptyVector(t *testing.T) {
	if _, err := Normalize([]float32{0, 0}); !errors.Is(err, ErrEmptyVector) {
		t.Fatalf("Normalize() error = %v, want %v", err, ErrEmptyVector)
	}
}

func TestRemoteEmbedderEmbedsOpenAICompatibleResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Fatalf("Authorization header = %q", r.Header.Get("Authorization"))
		}

		_ = json.NewEncoder(w).Encode(remoteResponse{
			Data: []struct {
				Embedding []float32 `json:"embedding"`
			}{
				{Embedding: []float32{1, 2, 3}},
			},
		})
	}))
	defer server.Close()

	embedder, err := NewRemoteEmbedder(RemoteConfig{
		URL:          server.URL,
		Model:        "test-model",
		APIKey:       "test-key",
		Dimension:    3,
		Timeout:      time.Second,
		MaxTextBytes: 128,
	}, server.Client())
	if err != nil {
		t.Fatalf("NewRemoteEmbedder() returned error: %v", err)
	}

	vector, err := embedder.Embed(context.Background(), "hello")
	if err != nil {
		t.Fatalf("Embed() returned error: %v", err)
	}

	if len(vector) != 3 {
		t.Fatalf("Embed() length = %d, want 3", len(vector))
	}
}

func TestRemoteEmbedderValidatesStatusAndDimension(t *testing.T) {
	statusServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "failed", http.StatusBadGateway)
	}))
	defer statusServer.Close()

	embedder, err := NewRemoteEmbedder(RemoteConfig{
		URL:          statusServer.URL,
		Model:        "test-model",
		Dimension:    2,
		Timeout:      time.Second,
		MaxTextBytes: 128,
	}, statusServer.Client())
	if err != nil {
		t.Fatalf("NewRemoteEmbedder() returned error: %v", err)
	}

	if _, err := embedder.Embed(context.Background(), "hello"); err == nil {
		t.Fatal("Embed() returned nil error for bad status")
	}

	dimensionServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(remoteResponse{
			Data: []struct {
				Embedding []float32 `json:"embedding"`
			}{
				{Embedding: []float32{1}},
			},
		})
	}))
	defer dimensionServer.Close()

	embedder, err = NewRemoteEmbedder(RemoteConfig{
		URL:          dimensionServer.URL,
		Model:        "test-model",
		Dimension:    2,
		Timeout:      time.Second,
		MaxTextBytes: 128,
	}, dimensionServer.Client())
	if err != nil {
		t.Fatalf("NewRemoteEmbedder() returned error: %v", err)
	}

	if _, err := embedder.Embed(context.Background(), "hello"); err == nil {
		t.Fatal("Embed() returned nil error for invalid dimension")
	}
}
