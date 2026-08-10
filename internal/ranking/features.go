package ranking

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

type FeatureScorer struct {
	base *knowledge.Knowledge
}

func NewFeatureScorer(
	base *knowledge.Knowledge,
) *FeatureScorer {
	return &FeatureScorer{
		base: base,
	}
}

type featureAccumulator struct {
	total   float64
	weights float64
	signals []Signal
}

func (a *featureAccumulator) add(
	name SignalName,
	value float64,
	weight float64,
) {
	if weight <= 0 {
		return
	}

	value = clamp(value)

	a.total +=
		value * weight

	a.weights +=
		weight

	a.signals = append(
		a.signals,
		Signal{
			Name:   name,
			Score:  value,
			Weight: weight,
		},
	)
}

func (a featureAccumulator) score() float64 {
	if a.weights <= 0 {
		return 0
	}

	return clamp(
		a.total /
			a.weights,
	)
}

func (s *FeatureScorer) Score(
	currentQuery domain.Query,
	fact domain.Fact,
	candidate Candidate,
) (float64, []Signal) {
	var accumulator featureAccumulator

	accumulator.add(
		SignalImportance,
		fact.Importance,
		0.5,
	)

	accumulator.add(
		SignalIntent,
		intentCompatibility(
			currentQuery.Intent,
			fact,
		),
		1.3,
	)

	if targetScore, applicable :=
		s.targetCompatibility(
			currentQuery.Target,
			fact,
		); applicable {
		accumulator.add(
			SignalTarget,
			targetScore,
			1.4,
		)
	}

	if temporalScore, applicable :=
		temporalCompatibility(
			currentQuery.TemporalScope,
			fact,
		); applicable {
		accumulator.add(
			SignalTemporal,
			temporalScore,
			1.25,
		)
	}

	if entityScore, applicable :=
		s.entityCoverage(
			currentQuery,
			fact,
		); applicable {
		accumulator.add(
			SignalEntityCoverage,
			entityScore,
			1.45,
		)
	}

	if conceptScore, applicable :=
		conceptCoverage(
			currentQuery,
			fact,
		); applicable {
		accumulator.add(
			SignalConceptCoverage,
			conceptScore,
			1.55,
		)
	}

	accumulator.add(
		SignalSourceDiversity,
		sourceDiversity(candidate),
		0.45,
	)

	return accumulator.score(),
		accumulator.signals
}

func intentCompatibility(
	intent domain.Intent,
	fact domain.Fact,
) float64 {
	switch intent {
	case domain.IntentDirectFact:
		return 0.8

	case domain.IntentOverview:
		switch fact.Category {
		case domain.FactCategoryProfile:
			return 1

		case domain.FactCategoryProject:
			return 1

		case domain.FactCategoryExperience:
			return 0.8

		case domain.FactCategoryAchievement:
			return 0.75

		default:
			return 0.6
		}

	case domain.IntentCapability:
		switch fact.Category {
		case domain.FactCategorySkill:
			return 1

		case domain.FactCategoryExperience:
			return 0.95

		case domain.FactCategoryProject:
			return 0.9

		case domain.FactCategoryAchievement:
			return 0.85

		case domain.FactCategoryCertification:
			return 0.7

		default:
			return 0.45
		}

	case domain.IntentExperience:
		switch fact.Category {
		case domain.FactCategoryExperience:
			return 1

		case domain.FactCategoryAchievement:
			return 0.95

		case domain.FactCategoryProject:
			return 0.7

		case domain.FactCategorySkill:
			return 0.6

		default:
			return 0.4
		}

	case domain.IntentTechnologyUsage:
		switch fact.Predicate {
		case domain.RelationUses,
			domain.RelationBuiltWith,
			domain.RelationDeploysOn,
			domain.RelationIntegratesWith:
			return 1

		case domain.RelationImplemented,
			domain.RelationDeveloped:
			return 0.75

		case domain.RelationHasSkill:
			return 0.65

		default:
			return 0.45
		}

	case domain.IntentProject:
		if fact.Category ==
			domain.FactCategoryProject {
			return 1
		}

		if fact.Category ==
			domain.FactCategoryAchievement {
			return 0.8
		}

		return 0.4

	case domain.IntentComparison:
		switch fact.Category {
		case domain.FactCategoryProject:
			return 1

		case domain.FactCategoryAchievement:
			return 0.95

		case domain.FactCategoryExperience:
			return 0.7

		case domain.FactCategorySkill:
			return 0.5

		default:
			return 0.35
		}

	case domain.IntentList:
		return 0.8

	case domain.IntentContact:
		if fact.Category ==
			domain.FactCategoryContact {
			return 1
		}

		return 0.1

	case domain.IntentEducation:
		if fact.Category ==
			domain.FactCategoryEducation {
			return 1
		}

		return 0.1

	case domain.IntentCertification:
		if fact.Category ==
			domain.FactCategoryCertification {
			return 1
		}

		return 0.1

	default:
		return 0.5
	}
}

func temporalCompatibility(
	scope domain.TemporalScope,
	fact domain.Fact,
) (float64, bool) {
	if scope == domain.TemporalScopeAny {
		return 0, false
	}

	if fact.Period == nil {
		return 0.35, true
	}

	switch scope {
	case domain.TemporalScopeCurrent:
		if fact.Period.Current {
			return 1, true
		}

		return 0.05, true

	case domain.TemporalScopePast:
		if fact.Period.Current {
			return 0.05, true
		}

		if fact.Period.End != nil {
			return 1, true
		}

		return 0.5, true
	}

	return 0, false
}

func (s *FeatureScorer) targetCompatibility(
	target domain.QueryTarget,
	fact domain.Fact,
) (float64, bool) {
	switch target {
	case domain.QueryTargetUnknown,
		domain.QueryTargetAny:
		return 0, false

	case domain.QueryTargetPerson:
		if s.factReferencesType(
			fact,
			domain.EntityTypePerson,
		) {
			return 1, true
		}

		if fact.Category ==
			domain.FactCategoryProfile {
			return 0.9, true
		}

		return 0.3, true

	case domain.QueryTargetProject:
		if s.entityHasType(
			fact.Subject,
			domain.EntityTypeProject,
		) {
			return 1, true
		}

		if fact.Object.Kind ==
			domain.FactObjectEntity &&
			s.entityHasType(
				fact.Object.EntityID,
				domain.EntityTypeProject,
			) {
			return 0.95, true
		}

		for _, contextID := range fact.Context {
			if s.entityHasType(
				contextID,
				domain.EntityTypeProject,
			) {
				return 0.9, true
			}
		}

		if fact.Category ==
			domain.FactCategoryProject {
			return 0.8, true
		}

		return 0.25, true

	case domain.QueryTargetTechnology:
		if s.factReferencesType(
			fact,
			domain.EntityTypeTechnology,
		) {
			return 1, true
		}

		if fact.Category ==
			domain.FactCategorySkill {
			return 0.85, true
		}

		return 0.35, true

	case domain.QueryTargetCompany:
		if s.factReferencesType(
			fact,
			domain.EntityTypeCompany,
		) {
			return 1, true
		}

		return 0.25, true

	case domain.QueryTargetExperience:
		switch fact.Category {
		case domain.FactCategoryExperience:
			return 1, true

		case domain.FactCategoryAchievement:
			return 0.95, true

		case domain.FactCategorySkill:
			return 0.65, true

		case domain.FactCategoryProject:
			return 0.6, true

		default:
			return 0.3, true
		}

	case domain.QueryTargetSkill:
		switch fact.Category {
		case domain.FactCategorySkill:
			return 1, true

		case domain.FactCategoryExperience:
			return 0.9, true

		case domain.FactCategoryProject:
			return 0.85, true

		case domain.FactCategoryAchievement:
			return 0.75, true

		case domain.FactCategoryCertification:
			return 0.7, true

		default:
			return 0.3, true
		}

	case domain.QueryTargetEducation:
		if fact.Category ==
			domain.FactCategoryEducation {
			return 1, true
		}

		return 0.1, true

	case domain.QueryTargetCertification:
		if fact.Category ==
			domain.FactCategoryCertification {
			return 1, true
		}

		return 0.1, true

	case domain.QueryTargetContact:
		if fact.Category ==
			domain.FactCategoryContact {
			return 1, true
		}

		return 0.1, true
	}

	return 0, false
}

func (s *FeatureScorer) factReferencesType(
	fact domain.Fact,
	entityType domain.EntityType,
) bool {
	if s.entityHasType(
		fact.Subject,
		entityType,
	) {
		return true
	}

	if fact.Object.Kind ==
		domain.FactObjectEntity &&
		s.entityHasType(
			fact.Object.EntityID,
			entityType,
		) {
		return true
	}

	for _, contextID := range fact.Context {
		if s.entityHasType(
			contextID,
			entityType,
		) {
			return true
		}
	}

	return false
}

func (s *FeatureScorer) entityHasType(
	entityID domain.EntityID,
	entityType domain.EntityType,
) bool {
	entity, found :=
		s.base.Entity(entityID)

	if !found {
		return false
	}

	return entity.Type == entityType
}

func (s *FeatureScorer) entityCoverage(
	currentQuery domain.Query,
	fact domain.Fact,
) (float64, bool) {
	totalWeight := 0.0
	matchedWeight := 0.0

	for _, match := range currentQuery.Entities {
		entity, found :=
			s.base.Entity(
				match.EntityID,
			)

		if !found {
			continue
		}

		if entity.Type ==
			domain.EntityTypePerson {
			continue
		}

		weight :=
			clamp(match.Score)

		if weight <= 0 {
			continue
		}

		totalWeight += weight

		if factReferencesEntity(
			fact,
			match.EntityID,
		) {
			matchedWeight += weight
		}
	}

	if totalWeight == 0 {
		return 0, false
	}

	return matchedWeight /
		totalWeight, true
}

func conceptCoverage(
	currentQuery domain.Query,
	fact domain.Fact,
) (float64, bool) {
	if len(currentQuery.Concepts) == 0 {
		return 0, false
	}

	totalWeight := 0.0
	matchedWeight := 0.0

	for _, match := range currentQuery.Concepts {
		if match.Score < 0.15 {
			continue
		}

		weight :=
			clamp(match.Score)

		totalWeight += weight

		if factHasConcept(
			fact,
			match.ConceptID,
		) {
			matchedWeight += weight
		}
	}

	if totalWeight == 0 {
		return 0, false
	}

	return matchedWeight /
		totalWeight, true
}

func factReferencesEntity(
	fact domain.Fact,
	entityID domain.EntityID,
) bool {
	if fact.Subject == entityID {
		return true
	}

	if fact.Object.Kind ==
		domain.FactObjectEntity &&
		fact.Object.EntityID ==
			entityID {
		return true
	}

	for _, contextID := range fact.Context {
		if contextID == entityID {
			return true
		}
	}

	return false
}

func factHasConcept(
	fact domain.Fact,
	conceptID domain.ConceptID,
) bool {
	for _, current := range fact.Concepts {
		if current == conceptID {
			return true
		}
	}

	return false
}

func sourceDiversity(
	candidate Candidate,
) float64 {
	const totalSources = 4

	if len(candidate.Sources) == 0 {
		return 0
	}

	value :=
		float64(
			len(candidate.Sources),
		) /
			float64(totalSources)

	return clamp(value)
}
