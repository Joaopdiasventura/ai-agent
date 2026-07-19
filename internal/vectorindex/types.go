package vectorindex

import "ai-agent/internal/domain"

const Version = "knowledge-index-v1"

type Index struct {
	Version   string  `json:"version"`
	Model     string  `json:"model"`
	Dimension int     `json:"dimension"`
	BaseHash  string  `json:"base_hash"`
	Entries   []Entry `json:"entries"`
}

type Entry struct {
	Document    domain.Document `json:"document"`
	Embedding   []float32       `json:"embedding"`
	ContentHash string          `json:"content_hash"`
}

type Manifest struct {
	Version       string `json:"version"`
	Model         string `json:"model"`
	Dimension     int    `json:"dimension"`
	DocumentCount int    `json:"document_count"`
	BaseHash      string `json:"base_hash"`
	GeneratedAt   string `json:"generated_at"`
}
