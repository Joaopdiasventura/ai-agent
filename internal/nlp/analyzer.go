package nlp

type QueryAnalysis struct {
	PrimaryIntent Intent
	AnswerMode    AnswerMode
	Entity        Entity
	HasEntity     bool
	Language      Language
}

func AnalyzeQuery(tokens []string, entity Entity, hasEntity bool, language Language) *QueryAnalysis {
	intent := DetectIntent(tokens)
	intent = ResolveIntent(intent, entity, hasEntity)

	answerMode := DetectAnswerMode(tokens)

	return &QueryAnalysis{
		PrimaryIntent: intent,
		AnswerMode:    answerMode,
		Entity:        entity,
		HasEntity:     hasEntity,
		Language:      language,
	}
}
