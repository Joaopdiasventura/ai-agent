package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/memory"
	"ai-agent/internal/nlp"
	"ai-agent/internal/tfidf"
	"ai-agent/internal/tokenizer"
)

type Engine struct {
	Documents         []domain.Document
	IDF               map[string]float64
	DocumentVectors   map[string]map[string]float64
	MinimumSimilarity float64
}

func NewEngine(documents []domain.Document, minimumSimilarity float64) *Engine {
	idf := tfidf.CalculateIDF(documents)
	documentsVectors := tfidf.CalculateDocumentVectors(documents, idf)

	return &Engine{
		Documents:         documents,
		IDF:               idf,
		DocumentVectors:   documentsVectors,
		MinimumSimilarity: minimumSimilarity,
	}
}

func (engine *Engine) Search(question string, session *memory.Session) (Result, bool) {
	tokens := tokenizer.Tokenize(question)

	if len(tokens) == 0 {
		return Result{}, false
	}

	detectedEntity, hasDetectedEntity := nlp.DetecEntity(tokens)

	entity, hasEntity := session.ResolveEntity(question, detectedEntity, hasDetectedEntity)

	intent := nlp.DetectIntent(tokens)
	intent = nlp.ResolveIntent(intent, entity, hasEntity)

	candidates := FilterDocumentsByIntent(engine.Documents, intent)

	if len(candidates) == 0 {
		return Result{}, false
	}

	expandedTokens := nlp.ExpandQuery(tokens)

	if hasEntity {
		entityTokens := tokenizer.Tokenize(entity.Value)
		expandedTokens = append(expandedTokens, entityTokens...)
	}

	questionVector := tfidf.CalculateTFIDF(expandedTokens, engine.IDF)

	if len(questionVector) == 0 {
		return Result{}, false
	}

	result, found := FindBestDocument(candidates, engine.DocumentVectors, questionVector)

	if !found {
		return Result{}, false
	}

	if !IsRelevant(result.Similarity, engine.MinimumSimilarity) {
		return Result{}, false
	}

	result.Intent = intent
	result.Entity = entity
	result.HasEntity = hasEntity

	return result, true
}
