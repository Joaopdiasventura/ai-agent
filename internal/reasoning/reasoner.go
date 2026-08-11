package reasoning

import (
	"errors"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
	"ai-agent/internal/ranking"
)

type Reasoner struct {
	base       *knowledge.Knowledge
	evidence   *EvidenceBuilder
	grouper    *Grouper
	capability *CapabilityReasoner
	comparison *ComparisonReasoner
}

func New(
	base *knowledge.Knowledge,
) (*Reasoner, error) {
	if base == nil {
		return nil, errors.New(
			"knowledge base is required",
		)
	}

	grouper :=
		NewGrouper(base)

	return &Reasoner{
		base: base,
		evidence: NewEvidenceBuilder(
			base,
		),
		grouper: grouper,
		capability: NewCapabilityReasoner(
			base,
		),
		comparison: NewComparisonReasoner(
			base,
			grouper,
		),
	}, nil
}

func (r *Reasoner) Reason(
	currentQuery domain.Query,
	ranked ranking.Result,
) Result {
	evidence :=
		r.evidence.Build(
			currentQuery,
			ranked,
			40,
		)

	switch currentQuery.Intent {
	case domain.IntentComparison:
		return Result{
			Query: currentQuery,
			Conclusion: r.comparison.Reason(
				currentQuery,
				evidence,
			),
		}

	case domain.IntentCapability:
		return Result{
			Query: currentQuery,
			Conclusion: r.capability.Reason(
				currentQuery,
				evidence,
			),
		}

	case domain.IntentTechnologyUsage:
		return Result{
			Query: currentQuery,
			Conclusion: r.technologyUsage(
				currentQuery,
				evidence,
			),
		}

	case domain.IntentExperience:
		return Result{
			Query: currentQuery,
			Conclusion: r.generalConclusion(
				ConclusionExperience,
				currentQuery,
				evidence,
			),
		}

	case domain.IntentOverview:
		return Result{
			Query: currentQuery,
			Conclusion: r.overview(
				currentQuery,
				evidence,
			),
		}

	case domain.IntentList:
		return Result{
			Query: currentQuery,
			Conclusion: r.listConclusion(
				currentQuery,
				evidence,
			),
		}

	case domain.IntentContact,
		domain.IntentEducation,
		domain.IntentCertification,
		domain.IntentDirectFact,
		domain.IntentProject:
		return Result{
			Query: currentQuery,
			Conclusion: r.generalConclusion(
				ConclusionDirect,
				currentQuery,
				evidence,
			),
		}

	default:
		return Result{
			Query: currentQuery,
			Conclusion: Conclusion{
				Type:   ConclusionUnknown,
				Status: SupportInsufficientEvidence,
			},
		}
	}
}

func (r *Reasoner) generalConclusion(
	conclusionType ConclusionType,
	currentQuery domain.Query,
	evidence []Evidence,
) Conclusion {
	status :=
		SupportInsufficientEvidence

	filtered := r.directlyRelevantEvidence(
		currentQuery,
		evidence,
	)

	if len(filtered) > 0 {
		status =
			SupportSupported
	}

	return Conclusion{
		Type:   conclusionType,
		Status: status,
		FocusEntity: focusEntity(
			currentQuery,
		),
		FocusConcept: focusConcept(
			currentQuery,
		),
		Evidence: filtered,
	}
}

func (r *Reasoner) directlyRelevantEvidence(
	currentQuery domain.Query,
	evidence []Evidence,
) []Evidence {
	result := make([]Evidence, 0, len(evidence))

	for _, currentEvidence := range evidence {
		fact, found := r.base.Fact(currentEvidence.FactID)

		if !found {
			continue
		}

		if directFactRelevant(currentQuery, fact, r.base) {
			result = append(result, currentEvidence)
		}
	}

	sortEvidence(result)

	return result
}

func directFactRelevant(
	currentQuery domain.Query,
	fact domain.Fact,
	base *knowledge.Knowledge,
) bool {
	hasNonPersonEntity := false

	for _, entityMatch := range currentQuery.Entities {
		entity, found := base.Entity(entityMatch.EntityID)

		if !found || entity.Type == domain.EntityTypePerson {
			continue
		}

		hasNonPersonEntity = true

		if factReferencesEntity(fact, entityMatch.EntityID) {
			return true
		}
	}

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		if factHasConcept(fact, concept.ConceptID) {
			return true
		}
	}

	return !hasNonPersonEntity &&
		((currentQuery.Intent == domain.IntentExperience &&
			fact.Category == domain.FactCategoryExperience) ||
			(currentQuery.Target != domain.QueryTargetPerson &&
				currentQuery.Target != domain.QueryTargetAny &&
				evidenceAllowedByTarget(currentQuery.Target, fact)))
}

func evidenceAllowedByTarget(
	target domain.QueryTarget,
	fact domain.Fact,
) bool {
	switch target {
	case domain.QueryTargetContact:
		return fact.Category == domain.FactCategoryContact
	case domain.QueryTargetEducation:
		return fact.Category == domain.FactCategoryEducation
	case domain.QueryTargetCertification:
		return fact.Category == domain.FactCategoryCertification
	case domain.QueryTargetExperience:
		return fact.Category == domain.FactCategoryExperience
	default:
		return false
	}
}

func (r *Reasoner) overview(
	currentQuery domain.Query,
	evidence []Evidence,
) Conclusion {
	focus :=
		focusEntity(
			currentQuery,
		)

	if focus == "" {
		return r.generalConclusion(
			ConclusionOverview,
			currentQuery,
			evidence,
		)
	}

	filtered := make(
		[]Evidence,
		0,
	)

	for _, currentEvidence := range evidence {
		fact, found :=
			r.base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			continue
		}

		if factReferencesEntity(
			fact,
			focus,
		) {
			filtered = append(
				filtered,
				currentEvidence,
			)
		}
	}

	sortEvidence(filtered)

	status :=
		SupportInsufficientEvidence

	if len(filtered) > 0 {
		status =
			SupportSupported
	}

	return Conclusion{
		Type:        ConclusionOverview,
		Status:      status,
		FocusEntity: focus,
		FocusConcept: focusConcept(
			currentQuery,
		),
		Evidence: filtered,
	}
}

func (r *Reasoner) technologyUsage(
	currentQuery domain.Query,
	evidence []Evidence,
) Conclusion {
	filtered := make(
		[]Evidence,
		0,
	)

	technology :=
		focusTechnology(
			currentQuery,
			r.base,
		)

	for _, currentEvidence := range evidence {
		fact, found :=
			r.base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			continue
		}

		if technology != "" &&
			!factReferencesEntity(
				fact,
				technology,
			) {
			continue
		}

		if !technologyUsageRelation(
			fact.Predicate,
		) {
			continue
		}

		filtered = append(
			filtered,
			currentEvidence,
		)
	}

	sortEvidence(filtered)

	status :=
		SupportInsufficientEvidence

	if len(filtered) > 0 {
		status =
			SupportSupported
	}

	groups :=
		r.grouper.Group(
			currentQuery,
			filtered,
			domain.EntityTypeProject,
		)

	return Conclusion{
		Type:        ConclusionTechnologyUsage,
		Status:      status,
		FocusEntity: technology,
		FocusConcept: focusConcept(
			currentQuery,
		),
		Evidence: filtered,
		Groups:   groups,
	}
}

func technologyUsageRelation(
	relation domain.Relation,
) bool {
	switch relation {
	case domain.RelationUses,
		domain.RelationBuiltWith,
		domain.RelationDeploysOn,
		domain.RelationIntegratesWith,
		domain.RelationImplemented,
		domain.RelationDeveloped:
		return true

	default:
		return false
	}
}

func (r *Reasoner) listConclusion(
	currentQuery domain.Query,
	evidence []Evidence,
) Conclusion {
	entityType, supported :=
		listEntityType(currentQuery)

	if !supported {
		return r.generalConclusion(
			ConclusionList,
			currentQuery,
			evidence,
		)
	}

	groups :=
		r.grouper.Group(
			currentQuery,
			evidence,
			entityType,
		)

	groups = r.filterListGroups(currentQuery, groups)

	status :=
		SupportInsufficientEvidence

	if len(groups) > 0 {
		status =
			SupportSupported
	}

	return Conclusion{
		Type:   ConclusionList,
		Status: status,
		FocusEntity: focusEntity(
			currentQuery,
		),
		FocusConcept: focusConcept(
			currentQuery,
		),
		Evidence: evidence,
		Groups:   groups,
	}
}

func listEntityType(currentQuery domain.Query) (domain.EntityType, bool) {
	if currentQuery.HasConcept(ontology.ConceptProgrammingLanguage) ||
		currentQuery.HasConcept(ontology.ConceptFramework) ||
		currentQuery.HasConcept(ontology.ConceptRuntime) ||
		currentQuery.HasConcept(ontology.ConceptDatabase) ||
		currentQuery.HasConcept(ontology.ConceptMessaging) ||
		currentQuery.HasConcept(ontology.ConceptCloud) ||
		currentQuery.HasConcept(ontology.ConceptDevOps) ||
		currentQuery.HasConcept(ontology.ConceptInfrastructure) {
		return domain.EntityTypeTechnology, true
	}

	if currentQuery.HasConcept(ontology.ConceptLanguage) {
		return domain.EntityTypeLanguage, true
	}

	return targetEntityType(currentQuery.Target)
}

func (r *Reasoner) filterListGroups(
	currentQuery domain.Query,
	groups []EntityGroup,
) []EntityGroup {
	focus := focusConcept(currentQuery)

	if focus == "" {
		return groups
	}

	result := make([]EntityGroup, 0, len(groups))

	for _, group := range groups {
		if !entityMatchesListConcept(group.EntityID, focus) {
			continue
		}

		result = append(result, group)
	}

	sortGroups(result)

	return result
}

func entityMatchesListConcept(
	entityID domain.EntityID,
	conceptID domain.ConceptID,
) bool {
	concepts := map[domain.ConceptID][]domain.EntityID{
		ontology.ConceptProgrammingLanguage: {
			knowledge.EntityJavaScript,
			knowledge.EntityTypeScript,
			knowledge.EntityJava,
			knowledge.EntityGo,
		},
		ontology.ConceptFramework: {
			knowledge.EntityAngular,
			knowledge.EntityReact,
			knowledge.EntityNextJS,
			knowledge.EntitySpringBoot,
			knowledge.EntityNestJS,
		},
		ontology.ConceptRuntime: {
			knowledge.EntityNodeJS,
		},
	}

	allowed, constrained := concepts[conceptID]

	if !constrained {
		return true
	}

	for _, current := range allowed {
		if current == entityID {
			return true
		}
	}

	return false
}

func targetEntityType(
	target domain.QueryTarget,
) (
	domain.EntityType,
	bool,
) {
	switch target {
	case domain.QueryTargetProject:
		return domain.EntityTypeProject,
			true

	case domain.QueryTargetTechnology:
		return domain.EntityTypeTechnology,
			true

	case domain.QueryTargetCompany:
		return domain.EntityTypeCompany,
			true

	case domain.QueryTargetPerson:
		return domain.EntityTypePerson,
			true

	case domain.QueryTargetEducation:
		return domain.EntityTypeInstitution,
			true

	case domain.QueryTargetCertification:
		return domain.EntityTypeCertification,
			true

	default:
		return domain.EntityTypeUnknown,
			false
	}
}

func focusEntity(
	currentQuery domain.Query,
) domain.EntityID {
	bestScore := -1.0
	var best domain.EntityID

	for _, entity := range currentQuery.Entities {
		if !entity.Explicit {
			continue
		}

		if entity.Score > bestScore {
			bestScore =
				entity.Score

			best =
				entity.EntityID
		}
	}

	return best
}

func focusTechnology(
	currentQuery domain.Query,
	base *knowledge.Knowledge,
) domain.EntityID {
	bestScore := -1.0
	var best domain.EntityID

	for _, match := range currentQuery.Entities {
		entity, found :=
			base.Entity(
				match.EntityID,
			)

		if !found ||
			entity.Type !=
				domain.EntityTypeTechnology {
			continue
		}

		if match.Score > bestScore {
			bestScore =
				match.Score

			best =
				match.EntityID
		}
	}

	return best
}

func focusConcept(
	currentQuery domain.Query,
) domain.ConceptID {
	bestScore := -1.0
	var best domain.ConceptID

	for _, concept := range currentQuery.Concepts {
		if concept.MatchedText == "" {
			continue
		}

		if concept.Score > bestScore {
			bestScore =
				concept.Score

			best =
				concept.ConceptID
		}
	}

	if best != "" {
		return best
	}

	for _, concept := range currentQuery.Concepts {
		if concept.Score > bestScore {
			bestScore =
				concept.Score

			best =
				concept.ConceptID
		}
	}

	return best
}
