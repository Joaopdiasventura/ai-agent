package nlp

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
		if intent == IntentProject {
			return IntentCurrentJob
		}

	case EntityInstitution:
		if intent == IntentUnknown {
			return IntentEducation
		}
	}

	return intent
}
