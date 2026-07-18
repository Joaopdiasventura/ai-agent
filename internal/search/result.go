package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
)

type Result struct {
	Document   domain.Document
	Similarity float64
	Entity     nlp.Entity
	HasEntity  bool
}
