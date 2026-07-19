package nlp

type ProjectCriterion string

const (
	ProjectCriterionUnknown               ProjectCriterion = ""
	ProjectCriterionComplexProblem        ProjectCriterion = "complex_problem"
	ProjectCriterionTechnicalCapability   ProjectCriterion = "technical_capability"
	ProjectCriterionFinancialSystems      ProjectCriterion = "financial_systems"
	ProjectCriterionTechnicalLeadership   ProjectCriterion = "technical_leadership"
	ProjectCriterionGoPerformance         ProjectCriterion = "go_performance"
	ProjectCriterionAuditability          ProjectCriterion = "auditability"
	ProjectCriterionAsyncProcessing       ProjectCriterion = "async_processing"
	ProjectCriterionGeneralRecommendation ProjectCriterion = "general_recommendation"
)

var projectCriterionWeights = map[string]map[ProjectCriterion]int{
	"financeiro":     {ProjectCriterionFinancialSystems: 5},
	"financeiros":    {ProjectCriterionFinancialSystems: 5},
	"financial":      {ProjectCriterionFinancialSystems: 5},
	"banco":          {ProjectCriterionFinancialSystems: 4},
	"banking":        {ProjectCriterionFinancialSystems: 4},
	"consistência":   {ProjectCriterionFinancialSystems: 4},
	"consistencia":   {ProjectCriterionFinancialSystems: 4},
	"consistency":    {ProjectCriterionFinancialSystems: 4},
	"transacional":   {ProjectCriterionFinancialSystems: 4},
	"transactional":  {ProjectCriterionFinancialSystems: 4},
	"transferências": {ProjectCriterionFinancialSystems: 4},
	"transferencias": {ProjectCriterionFinancialSystems: 4},
	"transfers":      {ProjectCriterionFinancialSystems: 4},
	"liderança":      {ProjectCriterionTechnicalLeadership: 5},
	"lideranca":      {ProjectCriterionTechnicalLeadership: 5},
	"leadership":     {ProjectCriterionTechnicalLeadership: 5},
	"liderou":        {ProjectCriterionTechnicalLeadership: 4},
	"led":            {ProjectCriterionTechnicalLeadership: 4},
	"recrutador":     {ProjectCriterionGeneralRecommendation: 5},
	"recruiter":      {ProjectCriterionGeneralRecommendation: 5},
	"go":             {ProjectCriterionGoPerformance: 5},
	"golang":         {ProjectCriterionGoPerformance: 5},
	"concorrência":   {ProjectCriterionGoPerformance: 5},
	"concorrencia":   {ProjectCriterionGoPerformance: 5},
	"concurrency":    {ProjectCriterionGoPerformance: 5},
	"desempenho":     {ProjectCriterionGoPerformance: 5},
	"performance":    {ProjectCriterionGoPerformance: 5},
	"throughput":     {ProjectCriterionGoPerformance: 5},
	"benchmark":      {ProjectCriterionGoPerformance: 5},
	"integridade":    {ProjectCriterionGoPerformance: 3},
	"auditabilidade": {ProjectCriterionAuditability: 5},
	"auditability":   {ProjectCriterionAuditability: 5},
	"criptografia":   {ProjectCriterionAuditability: 4},
	"cryptography":   {ProjectCriterionAuditability: 4},
	"blockchain":     {ProjectCriterionAuditability: 4},
	"histórica":      {ProjectCriterionAuditability: 4},
	"historica":      {ProjectCriterionAuditability: 4},
	"historical":     {ProjectCriterionAuditability: 4},
	"validação":      {ProjectCriterionAuditability: 3},
	"validacao":      {ProjectCriterionAuditability: 3},
	"validation":     {ProjectCriterionAuditability: 3},
	"assinados":      {ProjectCriterionAuditability: 3},
	"signed":         {ProjectCriterionAuditability: 3},
	"assíncrono":     {ProjectCriterionAsyncProcessing: 5},
	"assincrono":     {ProjectCriterionAsyncProcessing: 5},
	"assíncrona":     {ProjectCriterionAsyncProcessing: 5},
	"assincrona":     {ProjectCriterionAsyncProcessing: 5},
	"asynchronous":   {ProjectCriterionAsyncProcessing: 5},
	"streaming":      {ProjectCriterionAsyncProcessing: 4},
	"upload":         {ProjectCriterionAsyncProcessing: 4},
	"processamento":  {ProjectCriterionAsyncProcessing: 4},
	"processing":     {ProjectCriterionAsyncProcessing: 4},
	"eventos":        {ProjectCriterionAsyncProcessing: 3},
	"events":         {ProjectCriterionAsyncProcessing: 3},
	"complexo":       {ProjectCriterionComplexProblem: 5},
	"complexos":      {ProjectCriterionComplexProblem: 5},
	"complex":        {ProjectCriterionComplexProblem: 5},
	"difíceis":       {ProjectCriterionComplexProblem: 5},
	"dificeis":       {ProjectCriterionComplexProblem: 5},
	"difficult":      {ProjectCriterionComplexProblem: 5},
	"desafio":        {ProjectCriterionComplexProblem: 5},
	"challenge":      {ProjectCriterionComplexProblem: 5},
	"capacidade":     {ProjectCriterionTechnicalCapability: 3},
	"ability":        {ProjectCriterionTechnicalCapability: 3},
	"técnica":        {ProjectCriterionTechnicalCapability: 4},
	"tecnica":        {ProjectCriterionTechnicalCapability: 4},
	"technical":      {ProjectCriterionTechnicalCapability: 4},
	"destacaria":     {ProjectCriterionGeneralRecommendation: 5},
	"destacar":       {ProjectCriterionGeneralRecommendation: 5},
	"highlight":      {ProjectCriterionGeneralRecommendation: 5},
	"relevante":      {ProjectCriterionGeneralRecommendation: 4},
	"relevant":       {ProjectCriterionGeneralRecommendation: 4},
}

var projectCriterionPriority = []ProjectCriterion{
	ProjectCriterionAuditability,
	ProjectCriterionGoPerformance,
	ProjectCriterionTechnicalLeadership,
	ProjectCriterionFinancialSystems,
	ProjectCriterionAsyncProcessing,
	ProjectCriterionComplexProblem,
	ProjectCriterionTechnicalCapability,
	ProjectCriterionGeneralRecommendation,
}

func DetectProjectCriterion(tokens []string, intent Intent) ProjectCriterion {
	if intent != IntentProjectRecommendation {
		return ProjectCriterionUnknown
	}

	scores := make(map[ProjectCriterion]int)

	for _, token := range tokens {
		tokenScores, exists := projectCriterionWeights[token]
		if !exists {
			continue
		}

		for criterion, score := range tokenScores {
			scores[criterion] += score
		}
	}

	bestCriterion := ProjectCriterionUnknown
	bestScore := 0

	for _, criterion := range projectCriterionPriority {
		score := scores[criterion]
		if score > bestScore {
			bestCriterion = criterion
			bestScore = score
		}
	}

	if bestCriterion == ProjectCriterionUnknown {
		return ProjectCriterionComplexProblem
	}

	return bestCriterion
}
