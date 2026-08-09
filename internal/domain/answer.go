package domain

type AnswerPlanType string

const (
	AnswerPlanUnknown         AnswerPlanType = "unknown"
	AnswerPlanDirect          AnswerPlanType = "direct"
	AnswerPlanOverview        AnswerPlanType = "overview"
	AnswerPlanCapability      AnswerPlanType = "capability"
	AnswerPlanExperience      AnswerPlanType = "experience"
	AnswerPlanTechnologyUsage AnswerPlanType = "technology_usage"
	AnswerPlanComparison      AnswerPlanType = "comparison"
	AnswerPlanList            AnswerPlanType = "list"
	AnswerPlanAbstention      AnswerPlanType = "abstention"
)

type AnswerSectionRole string

const (
	AnswerSectionLead       AnswerSectionRole = "lead"
	AnswerSectionContext    AnswerSectionRole = "context"
	AnswerSectionSupport    AnswerSectionRole = "support"
	AnswerSectionProof      AnswerSectionRole = "proof"
	AnswerSectionComparison AnswerSectionRole = "comparison"
	AnswerSectionCaveat     AnswerSectionRole = "caveat"
)

type AnswerSection struct {
	Role    AnswerSectionRole
	Label   LocalizedText
	FactIDs []FactID
}

type AnswerPlan struct {
	Type          AnswerPlanType
	Language      Language
	FocusEntity   EntityID
	FocusConcept  ConceptID
	Sections      []AnswerSection
	Confidence    float64
	Abstain       bool
	AbstainReason LocalizedText
}

type Answer struct {
	Text       string
	Language   Language
	Confidence float64
	Evidence   []Evidence
	Abstained  bool
}
