package query

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/language"
	"ai-agent/internal/ontology"
	"errors"
)

type Analyzer struct {
	processor        *language.Processor
	entityExtractor  *EntityExtractor
	conceptExtractor *ConceptExtractor
	conceptExpander  *ConceptExpander
	intentDetector   *IntentDetector
	targetDetector   *TargetDetector
	temporalDetector *TemporalDetector
}

func NewAnalyzer(
	base *knowledge.Knowledge,
	currentOntology *ontology.Ontology,
) (*Analyzer, error) {
	if base == nil {
		return nil, errors.New("knowledge base is required")
	}

	if currentOntology == nil {
		return nil, errors.New("ontology is required")
	}

	if err := currentOntology.ValidateFacts(base.Facts()); err != nil {
		return nil, err
	}

	return &Analyzer{
		processor: language.NewProcessor(),
		entityExtractor: NewEntityExtractor(base),
		conceptExtractor: NewConceptExtractor(currentOntology),
		conceptExpander: NewconceptExpander(currentOntology),
		intentDetector: NewIntentDetector(),
		targetDetector: NewTargetDetector(base),
		temporalDetector: NewTemporalDetector(),
	}, nil
}

func (a *Analyzer) Analyze(value string) domain.Query {
	return a.AnalyzeWithLanguage(value, domain.LanguageUnknown)
}

func (a *Analyzer) AnalyzeWithLanguage(
	value string,
	fallbackLanguage domain.Language,
) domain.Query {
	analysis := a.processor.Analyze(value)

	queryLanguage := analysis.Language

	if queryLanguage == domain.LanguageUnknown && fallbackLanguage.IsSupported() {
		queryLanguage = fallbackLanguage
	}

	entities := a.entityExtractor.Extract(
		analysis.Original,
		analysis.Normalized,
		analysis.Terms,
		queryLanguage,
	)

	directConcepts := a.conceptExtractor.Extract(
		analysis.Normalized,
		analysis.Terms,
		queryLanguage,
	)

	expandedConcepts := a.conceptExpander.Expand(directConcepts, entities)

	intent := a.intentDetector.Detect(analysis.Normalized, entities, expandedConcepts)

	target := a.targetDetector.Detect(analysis.Normalized, intent, entities)

	temporalScope := a.temporalDetector.Detect(analysis.Normalized)

	return domain.Query{
		Original:      analysis.Original,
		Normalized:    analysis.Normalized,
		Language:      queryLanguage,
		Intent:        intent,
		Target:        target,
		TemporalScope: temporalScope,
		Tokens:        analysis.Tokens,
		Terms:         analysis.Terms,
		Entities:      entities,
		Concepts:      expandedConcepts,
	}
}