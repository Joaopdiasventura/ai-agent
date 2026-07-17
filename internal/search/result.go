package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
)

type Result struct {
	Document   domain.Document
	Similarity float64
	Intent     nlp.Intent
	Entity     nlp.Entity
	HasEntity  bool
}
