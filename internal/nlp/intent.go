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
		if intent == IntentAbout ||
			intent == IntentCurrentJob ||
			intent == IntentUnknown {
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
