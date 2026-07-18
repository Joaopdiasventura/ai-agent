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

func (engine *Engine) SearchResults(
	question string,
	session *memory.Session,
	limit int,
) ([]Result, []string, bool) {
	if limit <= 0 {
		return nil, nil, false
	}

	tokens := tokenizer.Tokenize(question)

	if len(tokens) == 0 {
		return nil, nil, false
	}

	expandedTokens := nlp.ExpandQuery(tokens)

	detectedEntity, hasDetectedEntity := nlp.DetectEntity(expandedTokens)

	entity, hasEntity := session.ResolveEntity(question, detectedEntity, hasDetectedEntity)

	intent := nlp.DetectIntent(expandedTokens)
	intent = nlp.ResolveIntent(intent, entity, hasEntity)

	candidates := FilterDocumentsByIntent(engine.Documents, intent)

	if len(candidates) == 0 {
		return nil, nil, false
	}

	if hasEntity {
		entityTokens := tokenizer.Tokenize(entity.Value)
		expandedTokens = append(expandedTokens, entityTokens...)
	}

	questionVector := tfidf.CalculateTFIDF(expandedTokens, engine.IDF)

	if len(questionVector) == 0 {
		return nil, nil, false
	}

	searchLimit := 1

	if ShouldSearchMultipleDocuments(intent, hasEntity) {
		searchLimit = limit
	}

	results := FindTopDocuments(candidates, engine.DocumentVectors, questionVector, searchLimit)

	results = FilterRelevantResults(results, engine.MinimumSimilarity)

	if len(results) == 0 {
		return nil, nil, false
	}

	for index := range results {
		results[index].Intent = intent
		results[index].Entity = entity
		results[index].HasEntity = hasEntity
	}

	return results, expandedTokens, true
}
