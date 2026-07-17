package search

import "ai-agent/internal/nlp"

func ShouldSearchMultipleDocuments(intent nlp.Intent, hasEntity bool) bool {
	if hasEntity {
		return false
	}

	switch intent {
	case nlp.IntentTechnologies:
		return true
	case nlp.IntentProject:
		return true
	default:
		return false
	}
}
