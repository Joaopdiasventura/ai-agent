package nlp

type QueryAnalysis struct {
	PrimaryIntent Intent
	AnswerMode    AnswerMode
	Entity        Entity
	HasEntity     bool
}

func AnalyzeQuery(tokens []string, entity Entity, hasEntity bool) *QueryAnalysis {
	intent := DetectIntent(tokens)
	intent = ResolveIntent(intent, entity, hasEntity)

	answerMode := DetectAnswerMode(tokens)

	return &QueryAnalysis{
		PrimaryIntent: intent,
		AnswerMode:    answerMode,
		Entity:        entity,
		HasEntity:     hasEntity,
	}
}
