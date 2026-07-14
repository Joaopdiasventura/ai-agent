package memory

import (
	"ai-agent/internal/nlp"
	"strings"
)

type Session struct {
	LastEntity    nlp.Entity
	HasLastentity bool
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

	normalizedQuestion := strings.ToLower(question)

	replacer := strings.NewReplacer(
		"?", " ",
		"!", " ",
		".", " ",
		",", " ",
		";", " ",
		":", " ",
	)

	normalizedQuestion = replacer.Replace(normalizedQuestion)
	normalizedQuestion = " " + strings.Join(strings.Fields(normalizedQuestion), " ") + " "

	references := []string{
		" ele ",
		" ela ",
		" dele ",
		" dela ",
		" isso ",
		" esse ",
		" essa ",
		" este ",
		" esta ",
	}

	for _, reference := range references {
		if strings.Contains(normalizedQuestion, reference) {
			return session.LastEntity, true
		}
	}

	return nlp.Entity{}, false
}
