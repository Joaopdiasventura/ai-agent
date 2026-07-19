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
	Results          []Result
	Tokens           []string
	Intent           nlp.Intent
	Language         nlp.Language
	ProjectCriterion nlp.ProjectCriterion
	SelectedProject  string
	Found            bool
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

	analysisTokens := tokenizer.TokenizeForAnalysis(question)
	tokens := tokenizer.Tokenize(question)
	language := nlp.DetectLanguage(analysisTokens)

	if len(analysisTokens) == 0 && len(tokens) == 0 {
		return &SearchResult{
			Language: language,
			Found:    false,
		}
	}

	expandedTokens := nlp.ExpandQuery(tokens)
	expandedAnalysisTokens := nlp.ExpandQuery(mergeTokens(analysisTokens, expandedTokens))

	entity, hasEntity := nlp.DetectEntity(expandedAnalysisTokens)

	analysis := nlp.AnalyzeQuery(analysisTokens, entity, hasEntity, language)
	analysis.Question = question

	if analysis.PrimaryIntent == nlp.IntentProjectRecommendation {
		results, selectedProject := FindProjectRecommendationDocuments(engine.Documents, analysis, limit)

		if len(results) == 0 {
			return &SearchResult{
				Language: language,
				Found:    false,
			}
		}

		for index := range results {
			results[index].Entity = nlp.Entity{
				Type:  nlp.EntityProject,
				Value: selectedProject,
			}
			results[index].HasEntity = true
		}

		return &SearchResult{
			Results:          results,
			Tokens:           expandedAnalysisTokens,
			Intent:           analysis.PrimaryIntent,
			Language:         language,
			ProjectCriterion: analysis.ProjectCriterion,
			SelectedProject:  selectedProject,
			Found:            true,
		}
	}

	candidates := FilterDocumentsByIntent(engine.Documents, analysis)

	if len(candidates) == 0 {
		return &SearchResult{
			Language: language,
			Found:    false,
		}
	}

	if hasEntity && entity.Type != nlp.EntityPerson {
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
		results = FindTopDocuments(filterDocumentsByLanguage(engine.Documents, language), engine, questionVector, analysis, searchLimit)
	}

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
		Results:          results,
		Tokens:           expandedAnalysisTokens,
		Intent:           analysis.PrimaryIntent,
		Language:         language,
		ProjectCriterion: analysis.ProjectCriterion,
		Found:            true,
	}
}

func mergeTokens(tokenGroups ...[]string) []string {
	merged := make([]string, 0)
	seen := make(map[string]struct{})

	for _, tokens := range tokenGroups {
		for _, token := range tokens {
			if token == "" {
				continue
			}

			if _, exists := seen[token]; exists {
				continue
			}

			seen[token] = struct{}{}
			merged = append(merged, token)
		}
	}

	return merged
}
