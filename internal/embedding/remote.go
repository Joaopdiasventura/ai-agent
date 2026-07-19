package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type RemoteConfig struct {
	URL          string
	Model        string
	APIKey       string
	Dimension    int
	Timeout      time.Duration
	MaxTextBytes int
}

type RemoteEmbedder struct {
	config RemoteConfig
	client *http.Client
}

type remoteRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type remoteResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

var ErrMissingRemoteConfig = errors.New("missing remote embedding configuration")

func RemoteConfigFromEnv() (RemoteConfig, error) {
	dimension, err := readPositiveInt("EMBEDDING_DIMENSION", 0)
	if err != nil {
		return RemoteConfig{}, err
	}

	timeoutMillis, err := readPositiveInt("EMBEDDING_TIMEOUT_MS", 5000)
	if err != nil {
		return RemoteConfig{}, err
	}

	maxTextBytes, err := readPositiveInt("EMBEDDING_MAX_TEXT_BYTES", 8192)
	if err != nil {
		return RemoteConfig{}, err
	}

	config := RemoteConfig{
		URL:          strings.TrimSpace(os.Getenv("EMBEDDING_URL")),
		Model:        strings.TrimSpace(os.Getenv("EMBEDDING_MODEL")),
		APIKey:       strings.TrimSpace(os.Getenv("EMBEDDING_API_KEY")),
		Dimension:    dimension,
		Timeout:      time.Duration(timeoutMillis) * time.Millisecond,
		MaxTextBytes: maxTextBytes,
	}

	if config.URL == "" || config.Model == "" || config.Dimension <= 0 {
		return RemoteConfig{}, ErrMissingRemoteConfig
	}

	return config, nil
}

func NewRemoteEmbedder(config RemoteConfig, client *http.Client) (*RemoteEmbedder, error) {
	if strings.TrimSpace(config.URL) == "" || strings.TrimSpace(config.Model) == "" || config.Dimension <= 0 {
		return nil, ErrMissingRemoteConfig
	}

	if config.Timeout <= 0 {
		config.Timeout = 5 * time.Second
	}

	if config.MaxTextBytes <= 0 {
		config.MaxTextBytes = 8192
	}

	if client == nil {
		client = &http.Client{Timeout: config.Timeout}
	}

	return &RemoteEmbedder{
		config: config,
		client: client,
	}, nil
}

func (embedder *RemoteEmbedder) Embed(ctx context.Context, text string) ([]float32, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, ErrEmptyText
	}

	if len([]byte(text)) > embedder.config.MaxTextBytes {
		return nil, fmt.Errorf("embedding text exceeds %d bytes", embedder.config.MaxTextBytes)
	}

	requestBody, err := json.Marshal(remoteRequest{
		Model: embedder.config.Model,
		Input: text,
	})
	if err != nil {
		return nil, err
	}

	requestContext, cancel := context.WithTimeout(ctx, embedder.config.Timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(requestContext, http.MethodPost, embedder.config.URL, bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	if embedder.config.APIKey != "" {
		request.Header.Set("Authorization", "Bearer "+embedder.config.APIKey)
	}

	response, err := embedder.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		_, _ = io.Copy(io.Discard, io.LimitReader(response.Body, 1024))
		return nil, fmt.Errorf("embedding provider returned status %d", response.StatusCode)
	}

	var payload remoteResponse
	decoder := json.NewDecoder(io.LimitReader(response.Body, 1<<20))
	if err := decoder.Decode(&payload); err != nil {
		return nil, err
	}

	if len(payload.Data) == 0 {
		return nil, errors.New("embedding provider returned no data")
	}

	vector := payload.Data[0].Embedding
	if len(vector) != embedder.config.Dimension {
		return nil, fmt.Errorf("embedding dimension = %d, want %d", len(vector), embedder.config.Dimension)
	}

	for _, value := range vector {
		if value != 0 {
			return vector, nil
		}
	}

	return nil, ErrEmptyVector
}

func readPositiveInt(name string, defaultValue int) (int, error) {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", name)
	}

	if parsed <= 0 {
		return 0, fmt.Errorf("%s must be positive", name)
	}

	return parsed, nil
}
