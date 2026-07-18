package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"ai-agent/internal/tfidf"
	"ai-agent/internal/tokenizer"
)

type Engine struct {
	Documents         []*domain.Document
	IDF               map[string]float64
	DocumentVectors   map[string]map[string]float64
	MinimumSimilarity float64
}

type SearchResult struct {
	Results  []Result
	Tokens   []string
	Intent   nlp.Intent
	Language nlp.Language
	Found    bool
}

func NewEngine(documents []*domain.Document, minimumSimilarity float64) *Engine {
	idf := tfidf.CalculateIDF(documents)
	documentsVectors := tfidf.CalculateDocumentVectors(documents, idf)

	return &Engine{
		Documents:         documents,
		IDF:               idf,
		DocumentVectors:   documentsVectors,
		MinimumSimilarity: minimumSimilarity,
	}
}

func (engine *Engine) Search(
	question string,
	limit int,
) *SearchResult {
	if limit <= 0 {
		return &SearchResult{
			Language: nlp.LanguagePortuguese,
			Found:    false,
		}
	}

	tokens := tokenizer.Tokenize(question)
	language := nlp.DetectLanguage(tokens)

	if len(tokens) == 0 {
		return &SearchResult{
			Language: language,
			Found:    false,
		}
	}

	expandedTokens := nlp.ExpandQuery(tokens)

	entity, hasEntity := nlp.DetectEntity(expandedTokens)

	analysis := nlp.AnalyzeQuery(expandedTokens, entity, hasEntity, language)

	candidates := FilterDocumentsByIntent(engine.Documents, analysis)

	if len(candidates) == 0 {
		return &SearchResult{
			Language: language,
			Found:    false,
		}
	}

	if hasEntity {
		entityTokens := tokenizer.Tokenize(entity.Value)
		expandedTokens = append(expandedTokens, entityTokens...)
	}

	questionVector := tfidf.CalculateTFIDF(expandedTokens, engine.IDF)

	if len(questionVector) == 0 {
		return &SearchResult{
			Language: language,
			Found:    false,
		}
	}

	searchLimit := 1

	if ShouldSearchMultipleDocuments(analysis.PrimaryIntent, hasEntity) {
		searchLimit = limit
	}

	results := FindTopDocuments(candidates, engine, questionVector, analysis, searchLimit)

	if len(results) == 0 {
		return &SearchResult{
			Language: language,
			Found:    false,
		}
	}

	for index := range results {
		results[index].Entity = entity
		results[index].HasEntity = hasEntity
	}

	return &SearchResult{
		Results:  results,
		Tokens:   expandedTokens,
		Intent:   analysis.PrimaryIntent,
		Language: language,
		Found:    true,
	}
}
