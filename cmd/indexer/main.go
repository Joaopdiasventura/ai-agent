package main

import (
	"ai-agent/internal/embedding"
	"ai-agent/internal/indexing"
	"ai-agent/internal/knowledge"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const outputDirectory = "data/generated"
const indexPath = "data/generated/knowledge.index.json"
const manifestPath = "data/generated/manifest.json"
const deterministicDimension = 128
const deterministicModel = "deterministic-hash-v1"

func main() {
	ctx := context.Background()

	embedder, model, err := configuredEmbedder()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to configure embedder: %v\n", err)
		os.Exit(1)
	}

	index, manifest, report, err := indexing.Build(ctx, knowledge.Documents(), embedder, model, time.Now())
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to build index: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(outputDirectory, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	if err := writeJSON(indexPath, index); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write index: %v\n", err)
		os.Exit(1)
	}

	if err := writeJSON(manifestPath, manifest); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write manifest: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("documents: %d\n", report.DocumentCount)
	fmt.Printf("model: %s\n", report.Model)
	fmt.Printf("dimension: %d\n", report.Dimension)
	fmt.Printf("base_hash: %s\n", report.BaseHash)
	fmt.Printf("index: %s\n", filepath.ToSlash(indexPath))
	fmt.Printf("manifest: %s\n", filepath.ToSlash(manifestPath))
}

func configuredEmbedder() (indexing.Embedder, string, error) {
	remoteConfig, err := embedding.RemoteConfigFromEnv()
	if err == nil {
		embedder, err := embedding.NewRemoteEmbedder(remoteConfig, &http.Client{Timeout: remoteConfig.Timeout})
		if err != nil {
			return nil, "", err
		}

		return embedder, remoteConfig.Model, nil
	}

	if err != embedding.ErrMissingRemoteConfig {
		return nil, "", err
	}

	embedder, err := embedding.NewDeterministicEmbedder(deterministicDimension)
	if err != nil {
		return nil, "", err
	}

	return embedder, deterministicModel, nil
}

func writeJSON(path string, value any) error {
	content, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}

	content = append(content, '\n')
	return os.WriteFile(path, content, 0644)
}
