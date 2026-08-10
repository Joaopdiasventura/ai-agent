package evaluation

import (
	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
	"ai-agent/internal/planning"
)

type Category string

const (
	CategoryDirect          Category = "direct"
	CategoryOverview        Category = "overview"
	CategoryCapability      Category = "capability"
	CategoryTechnologyUsage Category = "technology_usage"
	CategoryComparison      Category = "comparison"
	CategoryAbstention      Category = "abstention"
	CategoryTypo            Category = "typo"
	CategoryUnknown         Category = "unknown"
	CategoryExperience      Category = "experience"
	CategoryEducation       Category = "education"
	CategoryCertification   Category = "certification"
)

type Expectation struct {
	HasResponse         *bool
	Language            domain.Language
	Intent              domain.Intent
	Target              domain.QueryTarget
	Entities            []domain.EntityID
	Concepts            []domain.ConceptID
	Winner              domain.EntityID
	Facts               []domain.FactID
	ForbiddenFacts      []domain.FactID
	PlanStatus          planning.PlanStatus
	ConfidenceMode      confidence.Mode
	MinConfidence       *float64
	MaxConfidence       *float64
	ResponseContains    []string
	ResponseNotContains []string
	GenerationContains  []string
}

type Case struct {
	ID          string
	Question    string
	Category    Category
	Expectation Expectation
}

type CheckName string

const (
	CheckExecution           CheckName = "execution"
	CheckHasResponse         CheckName = "has_response"
	CheckLanguage            CheckName = "language"
	CheckIntent              CheckName = "intent"
	CheckTarget              CheckName = "target"
	CheckEntity              CheckName = "entity"
	CheckConcept             CheckName = "concept"
	CheckWinner              CheckName = "winner"
	CheckFact                CheckName = "fact"
	CheckForbiddenFact       CheckName = "forbidden_fact"
	CheckPlanStatus          CheckName = "plan_status"
	CheckConfidenceMode      CheckName = "confidence_mode"
	CheckMinimumConfidence   CheckName = "minimum_confidence"
	CheckMaximumConfidence   CheckName = "maximum_confidence"
	CheckResponseContains    CheckName = "response_contains"
	CheckResponseNotContains CheckName = "response_not_contains"
	CheckGenerationContains  CheckName = "generation_contains"
)

type Check struct {
	Name     CheckName
	Passed   bool
	Expected string
	Actual   string
}

type CaseResult struct {
	Case   Case
	Passed bool
	Error  string
	Checks []Check
}
