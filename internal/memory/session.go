package memory

import (
	"ai-agent/internal/nlp"
)

type Session struct {
	LastEntity        nlp.Entity
	HasLastentity     bool
	LastIntent        nlp.Intent
	LastTemplateIndex map[nlp.Intent]int
}

func NewSession() *Session {
	return &Session{
		LastTemplateIndex: make(map[nlp.Intent]int),
	}
}

func (session *Session) ResolveEntity(question string, detecetedEntity nlp.Entity, hasDetectedEntity bool) (nlp.Entity, bool) {
	if hasDetectedEntity {
		session.LastEntity = detecetedEntity
		session.HasLastentity = true

		return detecetedEntity, true
	}

	if !session.HasLastentity {
		return nlp.Entity{}, false
	}

	if referencesLastEntity(question) {
		return session.LastEntity, true
	}

	return nlp.Entity{}, false
}

func (session *Session) GetLastTemplateIndex(intent nlp.Intent) int {
	index, exists := session.LastTemplateIndex[intent]

	if !exists {
		return -1
	}

	return index
}

func (session *Session) SetLastTemplateIndex(intent nlp.Intent, index int) {
	session.LastTemplateIndex[intent] = index
	session.LastIntent = intent
}
