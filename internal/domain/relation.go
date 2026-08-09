package domain

type Relation string

const (
	RelationUnknown        Relation = "unknown"
	RelationIs             Relation = "is"
	RelationWorksAt        Relation = "works_at"
	RelationWorkedAt       Relation = "worked_at"
	RelationHasRole        Relation = "has_role"
	RelationCreated        Relation = "created"
	RelationWorkedOn       Relation = "worked_on"
	RelationLed            Relation = "led"
	RelationUses           Relation = "uses"
	RelationBuiltWith      Relation = "built_with"
	RelationDemonstrates   Relation = "demonstrates"
	RelationImplemented    Relation = "implemented"
	RelationDeveloped      Relation = "developed"
	RelationDesigned       Relation = "designed"
	RelationProcesses      Relation = "processes"
	RelationStoresIn       Relation = "stores_in"
	RelationDeploysOn      Relation = "deploys_on"
	RelationIntegratesWith Relation = "integrates_with"
	RelationImproved       Relation = "improved"
	RelationAutomated      Relation = "automated"
	RelationAchieved       Relation = "achieved"
	RelationStudiedAt      Relation = "studied_at"
	RelationCertifiedIn    Relation = "certified_in"
	RelationSpeaks         Relation = "speaks"
	RelationLocatedIn      Relation = "located_in"
	RelationHasSkill       Relation = "has_skill"
	RelationHasExperience  Relation = "has_experience"
)
