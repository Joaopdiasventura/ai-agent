package nlp

type CategoryHint string

const (
	CategoryHintUnknown     CategoryHint = ""
	CategoryHintContact     CategoryHint = "contact"
	CategoryHintEducation   CategoryHint = "education"
	CategoryHintCareer      CategoryHint = "career"
	CategoryHintTechnology  CategoryHint = "technology"
	CategoryHintProject     CategoryHint = "project"
	CategoryHintImpact      CategoryHint = "impact"
	CategoryHintProfile     CategoryHint = "profile"
	CategoryHintCertificate CategoryHint = "certificate"
)

var categoryKeywordWeights = map[string]map[CategoryHint]int{
	"email":          {CategoryHintContact: 5},
	"e-mail":         {CategoryHintContact: 5},
	"telefone":       {CategoryHintContact: 5},
	"phone":          {CategoryHintContact: 5},
	"contato":        {CategoryHintContact: 4},
	"contact":        {CategoryHintContact: 4},
	"linkedin":       {CategoryHintContact: 4},
	"github":         {CategoryHintContact: 4},
	"estuda":         {CategoryHintEducation: 4},
	"estudou":        {CategoryHintEducation: 4},
	"estudando":      {CategoryHintEducation: 4},
	"study":          {CategoryHintEducation: 4},
	"studies":        {CategoryHintEducation: 4},
	"studied":        {CategoryHintEducation: 4},
	"cursa":          {CategoryHintEducation: 4},
	"cursando":       {CategoryHintEducation: 4},
	"curso":          {CategoryHintEducation: 3},
	"faculdade":      {CategoryHintEducation: 4},
	"formação":       {CategoryHintEducation: 4},
	"formacao":       {CategoryHintEducation: 4},
	"formou":         {CategoryHintEducation: 4},
	"concluiu":       {CategoryHintEducation: 4},
	"completed":      {CategoryHintEducation: 4},
	"graduated":      {CategoryHintEducation: 4},
	"fiap":           {CategoryHintEducation: 3},
	"etec":           {CategoryHintEducation: 3},
	"trabalha":       {CategoryHintCareer: 4},
	"trabalhou":      {CategoryHintCareer: 4},
	"trabalho":       {CategoryHintCareer: 3},
	"work":           {CategoryHintCareer: 3},
	"works":          {CategoryHintCareer: 4},
	"worked":         {CategoryHintCareer: 4},
	"job":            {CategoryHintCareer: 4},
	"emprego":        {CategoryHintCareer: 4},
	"cargo":          {CategoryHintCareer: 4},
	"carreira":       {CategoryHintCareer: 3},
	"career":         {CategoryHintCareer: 3},
	"ufind":          {CategoryHintCareer: 3},
	"representa":     {CategoryHintCareer: 3},
	"tecnologia":     {CategoryHintTechnology: 4},
	"tecnologias":    {CategoryHintTechnology: 4},
	"technology":     {CategoryHintTechnology: 4},
	"technologies":   {CategoryHintTechnology: 4},
	"stack":          {CategoryHintTechnology: 4},
	"usa":            {CategoryHintTechnology: 3},
	"uses":           {CategoryHintTechnology: 3},
	"utiliza":        {CategoryHintTechnology: 3},
	"projeto":        {CategoryHintProject: 4},
	"projetos":       {CategoryHintProject: 4},
	"project":        {CategoryHintProject: 4},
	"projects":       {CategoryHintProject: 4},
	"auronix":        {CategoryHintProject: 4},
	"xtube":          {CategoryHintProject: 4},
	"tube":           {CategoryHintProject: 2},
	"ggcompress":     {CategoryHintProject: 4},
	"auditex":        {CategoryHintProject: 4},
	"resultado":      {CategoryHintImpact: 4},
	"result":         {CategoryHintImpact: 4},
	"impacto":        {CategoryHintImpact: 4},
	"impact":         {CategoryHintImpact: 4},
	"reduziu":        {CategoryHintImpact: 4},
	"reduced":        {CategoryHintImpact: 4},
	"automação":      {CategoryHintImpact: 3},
	"automacao":      {CategoryHintImpact: 3},
	"automation":     {CategoryHintImpact: 3},
	"perfil":         {CategoryHintProfile: 4},
	"profile":        {CategoryHintProfile: 4},
	"idioma":         {CategoryHintProfile: 4},
	"idiomas":        {CategoryHintProfile: 4},
	"language":       {CategoryHintProfile: 4},
	"languages":      {CategoryHintProfile: 4},
	"autonomia":      {CategoryHintProfile: 3},
	"autonomy":       {CategoryHintProfile: 3},
	"certificação":   {CategoryHintCertificate: 5},
	"certificacao":   {CategoryHintCertificate: 5},
	"certificações":  {CategoryHintCertificate: 5},
	"certificacoes":  {CategoryHintCertificate: 5},
	"certification":  {CategoryHintCertificate: 5},
	"certifications": {CategoryHintCertificate: 5},
	"aws":            {CategoryHintCertificate: 1, CategoryHintTechnology: 1},
	"edb":            {CategoryHintCertificate: 2},
}

var categoryPriority = []CategoryHint{
	CategoryHintContact,
	CategoryHintEducation,
	CategoryHintCareer,
	CategoryHintTechnology,
	CategoryHintProject,
	CategoryHintImpact,
	CategoryHintProfile,
	CategoryHintCertificate,
}

func DetectCategoryHint(tokens []string, intent Intent) CategoryHint {
	scores := make(map[CategoryHint]int)

	for _, token := range tokens {
		tokenScores, exists := categoryKeywordWeights[token]
		if !exists {
			continue
		}

		for category, score := range tokenScores {
			scores[category] += score
		}
	}

	switch intent {
	case IntentContact:
		scores[CategoryHintContact] += 3
	case IntentEducation:
		scores[CategoryHintEducation] += 3
	case IntentCurrentJob, IntentFirstJob:
		scores[CategoryHintCareer] += 3
	case IntentTechnologies:
		scores[CategoryHintTechnology] += 3
	case IntentProject, IntentProjectRecommendation, IntentVisitorProjects:
		scores[CategoryHintProject] += 3
	case IntentHireReason:
		scores[CategoryHintImpact] += 1
	}

	bestCategory := CategoryHintUnknown
	bestScore := 0

	for _, category := range categoryPriority {
		score := scores[category]
		if score > bestScore {
			bestCategory = category
			bestScore = score
		}
	}

	return bestCategory
}
