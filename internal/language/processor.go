package language

import "ai-agent/internal/domain"

type TextAnalysis struct {
	Original           string
	Normalized         string
	Language           domain.Language
	LanguageConfidence float64
	Tokens             []string
	Terms              []string
}

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) Analyze(value string) TextAnalysis {
	normalized := Normalize(value)

	tokens := TokenizeNormalized(normalized)

	detection := DetectLanguageDetailed(value)

	terms := ContentTerms(tokens, detection.Language)

	return TextAnalysis{
		Original: value,
		Normalized: normalized,
		Language: detection.Language,
		LanguageConfidence: detection.Confidence,
		Tokens: tokens,
		Terms: terms,
	}
}
