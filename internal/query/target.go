package query

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

type TargetDetector struct {
	base *knowledge.Knowledge
}

func NewTargetDetector(
	base *knowledge.Knowledge,
) *TargetDetector {
	return &TargetDetector{
		base: base,
	}
}

func (d *TargetDetector) Detect(
	normalized string,
	intent domain.Intent,
	entities []domain.EntityMatch,
) domain.QueryTarget {
	switch intent {
	case domain.IntentContact:
		return domain.QueryTargetContact
	case domain.IntentEducation:
		return domain.QueryTargetEducation
	case domain.IntentCertification:
		return domain.QueryTargetCertification
	}

	if hasAnyPhrase(
		normalized,
		projectTargetMarkers,
	) {
		return domain.QueryTargetProject
	}

	if hasAnyPhrase(
		normalized,
		companyTargetMarkers,
	) {
		return domain.QueryTargetCompany
	}

	if hasAnyPhrase(
		normalized,
		technologyTargetMarkers,
	) {
		return domain.QueryTargetTechnology
	}

	if hasAnyPhrase(
		normalized,
		skillTargetMarkers,
	) {
		return domain.QueryTargetSkill
	}

	if hasAnyPhrase(
		normalized,
		experienceTargetMarkers,
	) {
		return domain.QueryTargetExperience
	}

	if intent == domain.IntentTechnologyUsage {
		if hasAnyPhrase(
			normalized,
			projectTargetMarkers,
		) {
			return domain.QueryTargetProject
		}

		return domain.QueryTargetExperience
	}

	if intent == domain.IntentCapability {
		return domain.QueryTargetSkill
	}

	if d.containsEntityType(
		entities,
		domain.EntityTypeProject,
	) {
		return domain.QueryTargetProject
	}

	if d.containsEntityType(
		entities,
		domain.EntityTypeTechnology,
	) {
		return domain.QueryTargetTechnology
	}

	if d.containsEntityType(
		entities,
		domain.EntityTypeCompany,
	) {
		return domain.QueryTargetCompany
	}

	if d.containsEntityType(
		entities,
		domain.EntityTypePerson,
	) {
		return domain.QueryTargetPerson
	}

	return domain.QueryTargetAny
}

var projectTargetMarkers = []string{
	"projeto",
	"projetos",
	"case de estudo",
	"cases de estudo",
	"project",
	"projects",
	"case study",
	"case studies",
}

var companyTargetMarkers = []string{
	"empresa",
	"empresas",
	"companhia",
	"company",
	"companies",
	"employer",
	"employers",
}

var technologyTargetMarkers = []string{
	"tecnologia",
	"tecnologias",
	"stack",
	"framework",
	"frameworks",
	"linguagem",
	"linguagens",
	"technology",
	"technologies",
	"tech stack",
	"language",
	"languages",
}

var skillTargetMarkers = []string{
	"habilidade",
	"habilidades",
	"competencia",
	"competencias",
	"conhecimento",
	"capacidade",
	"skill",
	"skills",
	"capability",
	"capabilities",
	"knowledge",
}

var experienceTargetMarkers = []string{
	"experiencia",
	"experiencias",
	"carreira",
	"trabalho",
	"trabalhou",
	"trabalha",
	"experience",
	"career",
	"work",
	"worked",
	"works",
}

func (d *TargetDetector) containsEntityType(
	entities []domain.EntityMatch,
	entityType domain.EntityType,
) bool {
	for _, match := range entities {
		entity, found :=
			d.base.Entity(match.EntityID)

		if !found {
			continue
		}

		if entity.Type == entityType {
			return true
		}
	}

	return false
}
