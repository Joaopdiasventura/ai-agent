package domain

type EntityMatch struct {
	EntityID    EntityID
	Score       float64
	Explicit    bool
	MatchedText string
}

type ConceptMatch struct {
	ConceptID   ConceptID
	Score       float64
	MatchedText string
}

type Query struct {
	Original      string
	Normalized    string
	Language      Language
	Intent        Intent
	Target        QueryTarget
	TemporalScope TemporalScope
	Tokens        []string
	Terms         []string
	Entities      []EntityMatch
	Concepts      []ConceptMatch
}

func (q Query) HasEntity(entityID EntityID) bool {
	for _, entity := range q.Entities {
		if entity.EntityID == entityID {
			return true
		}
	}

	return false
}

func (q Query) HasConcept(conceptID ConceptID) bool {
	for _, concept := range q.Concepts {
		if concept.ConceptID == conceptID {
			return true
		}
	}

	return false
}
