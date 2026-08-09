package knowledge

import "ai-agent/internal/domain"

func allFacts() []domain.Fact {
	facts := make([]domain.Fact, 0)

	facts = append(facts, profileFacts()...)
	facts = append(facts, skillFacts()...)
	facts = append(facts, experienceFacts()...)
	facts = append(facts, projectFacts()...)
	facts = append(facts, educationFacts()...)
	facts = append(facts, certificationFacts()...)
	facts = append(facts, languageFacts()...)

	return facts
}
