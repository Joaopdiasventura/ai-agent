package answer

import "ai-agent/internal/nlp"

type DetailLevel string

const (
	DetailShort    DetailLevel = "short"
	DetailMedium   DetailLevel = "medium"
	DetailDetailed DetailLevel = "detailed"
)

var DetailKeys = map[string]DetailLevel{
	"brevemente":     DetailShort,
	"resuma":         DetailShort,
	"explique":       DetailMedium,
	"detalhes":       DetailDetailed,
	"detalhadamente": DetailDetailed,
}

func SelectDetailLevel(tokens []string) DetailLevel {
	detailLevel := DetailMedium

	for _, token := range tokens {
		newDetailLevel, exists := DetailKeys[token]

		if exists {
			detailLevel = newDetailLevel
		}
	}

	return detailLevel
}

func SelectIntentDetailLevel(intent nlp.Intent, detailLevel DetailLevel) DetailLevel {
	switch intent {
	case nlp.IntentVisitorSummary,
		nlp.IntentVisitorProjects,
		nlp.IntentVisitorServices,
		nlp.IntentHireReason:
		if detailLevel == DetailMedium {
			return DetailShort
		}
	}

	return detailLevel
}
