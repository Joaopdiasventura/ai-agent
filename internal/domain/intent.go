package domain

type Intent string

const (
	IntentUnknown         Intent = "unknown"
	IntentDirectFact      Intent = "direct_fact"
	IntentOverview        Intent = "overview"
	IntentCapability      Intent = "capability"
	IntentExperience      Intent = "experience"
	IntentTechnologyUsage Intent = "technology_usage"
	IntentProject         Intent = "project"
	IntentComparison      Intent = "comparison"
	IntentList            Intent = "list"
	IntentContact         Intent = "contact"
	IntentEducation       Intent = "education"
	IntentCertification   Intent = "certification"
)

type QueryTarget string

const (
	QueryTargetUnknown       QueryTarget = "unknown"
	QueryTargetAny           QueryTarget = "any"
	QueryTargetPerson        QueryTarget = "person"
	QueryTargetProject       QueryTarget = "project"
	QueryTargetTechnology    QueryTarget = "technology"
	QueryTargetCompany       QueryTarget = "company"
	QueryTargetExperience    QueryTarget = "experience"
	QueryTargetSkill         QueryTarget = "skill"
	QueryTargetEducation     QueryTarget = "education"
	QueryTargetCertification QueryTarget = "certification"
	QueryTargetContact       QueryTarget = "contact"
)

type TemporalScope string

const (
	TemporalScopeAny     TemporalScope = "any"
	TemporalScopeCurrent TemporalScope = "current"
	TemporalScopePast    TemporalScope = "past"
)
