package nlp

type Intent string

const (
	IntentUnknown      Intent = "unknown"
	IntentAbout        Intent = "about"
	IntentCurrentJob   Intent = "current_job"
	IntentFirstJob     Intent = "first_job"
	IntentEducation    Intent = "education"
	IntentProject      Intent = "project"
	IntentTechnologies Intent = "technologies"
	IntentContact      Intent = "contact"
)

type EntityType string

const (
	EntityProject     EntityType = "project"
	EntityCompany     EntityType = "company"
	EntityInstitution EntityType = "institution"
	EntityTechnology  EntityType = "technology"
)

type Entity struct {
	Type  EntityType
	Value string
}

type AnswerMode string

const (
	AnswerModeDefault    AnswerMode = "default"
	AnswerModeAbout      AnswerMode = "about"
	AnswerModeTechnology AnswerMode = "technology"
	AnswerModeComparison AnswerMode = "comparison"
)

type QueryAnalysis struct {
	PrimaryIntent Intent
	AnswerMode    AnswerMode
	Entity        Entity
	HasEntity     bool
}

func ExpandQuery(tokens []string) []string {
	expandedTokens := make([]string, 0, len(tokens))
	seenTokens := make(map[string]struct{})

	addToken := func(token string) {
		if token == "" {
			return
		}

		if _, exists := seenTokens[token]; exists {
			return
		}

		seenTokens[token] = struct{}{}
		expandedTokens = append(expandedTokens, token)
	}

	for _, token := range tokens {
		addToken(token)
	}

	for index := 0; index < len(expandedTokens); index++ {
		token := expandedTokens[index]

		expansions, exists := queryExpansions[token]
		if !exists {
			continue
		}

		for _, expansion := range expansions {
			addToken(expansion)
		}
	}

	return expandedTokens
}

func DetectEntity(tokens []string) (Entity, bool) {
	if len(tokens) == 0 {
		return Entity{}, false
	}

	for _, alias := range entityAliases {
		if containsTokenSequence(tokens, alias.Tokens) {
			return alias.Entity, true
		}
	}

	return Entity{}, false
}

func DetectIntent(tokens []string) Intent {
	scores := make(map[Intent]int)

	for _, token := range tokens {
		tokenScores, exists := intentKeywords[token]

		if !exists {
			continue
		}

		for intent, score := range tokenScores {
			scores[intent] += score
		}
	}

	bestIntent := IntentUnknown
	bestScore := 0

	for _, intent := range intentPriority {
		score := scores[intent]

		if score > bestScore {
			bestIntent = intent
			bestScore = score
		}
	}

	return bestIntent
}

func ResolveIntent(intent Intent, entity Entity, hasEntity bool) Intent {
	if !hasEntity {
		return intent
	}

	switch entity.Type {
	case EntityProject:
		if intent == IntentCurrentJob || intent == IntentUnknown {
			return IntentProject
		}

	case EntityCompany:
		if intent == IntentProject || intent == IntentUnknown {
			return IntentCurrentJob
		}

	case EntityInstitution:
		if intent == IntentUnknown {
			return IntentEducation
		}

	case EntityTechnology:
		if intent == IntentUnknown {
			return IntentTechnologies
		}
	}

	return intent
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

func DetectAnswerMode(tokens []string) AnswerMode {
	for _, token := range tokens {
		answerMode, exists := answerModeToken[token]

		if exists {
			return answerMode
		}
	}

	return AnswerModeDefault
}

func containsTokenSequence(tokens []string, sequence []string) bool {
	if len(sequence) == 0 || len(sequence) > len(tokens) {
		return false
	}

	maximumStartIndex := len(tokens) - len(sequence)

	for startIndex := 0; startIndex <= maximumStartIndex; startIndex++ {
		matches := true

		for sequenceIndex, expectedToken := range sequence {
			if tokens[startIndex+sequenceIndex] != expectedToken {
				matches = false
				break
			}
		}

		if matches {
			return true
		}
	}

	return false
}
