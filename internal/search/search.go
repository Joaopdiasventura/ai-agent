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

func FindBestDocument(documents []domain.Document, documentsVectors map[string]map[string]float64, questionVector map[string]float64) (Result, bool) {
	if len(documents) == 0 || len(questionVector) == 0 {
		return Result{}, false
	}

	bestResult := Result{}
	found := false

	for _, document := range documents {
		documentVector, exists := documentsVectors[document.ID]

		if !exists {
			continue
		}

		similarity := CosineSimilarity(questionVector, documentVector)

		if !found || similarity > bestResult.Similarity {
			bestResult = Result{
				Document:   document,
				Similarity: similarity,
			}

			found = true
		}
	}

	return bestResult, found
}
