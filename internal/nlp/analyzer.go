package nlp

type QueryAnalysis struct {
	PrimaryIntent   Intent
	AnswerMode      AnswerMode
	Entity          Entity
	HasEntity       bool
	Language        Language
	TemporalContext TemporalContext
	CategoryHint    CategoryHint
	Question        string
}

func AnalyzeQuery(tokens []string, entity Entity, hasEntity bool, language Language) *QueryAnalysis {
	intent := DetectIntent(tokens)
	intent = ResolveIntent(intent, entity, hasEntity)

	answerMode := DetectAnswerMode(tokens)
	temporalContext := DetectTemporalContext(tokens, intent)
	categoryHint := DetectCategoryHint(tokens, intent)

	return &QueryAnalysis{
		PrimaryIntent:   intent,
		AnswerMode:      answerMode,
		Entity:          entity,
		HasEntity:       hasEntity,
		Language:        language,
		TemporalContext: temporalContext,
		CategoryHint:    categoryHint,
	}
}
