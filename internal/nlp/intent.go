package nlp

type Intent string

const (
	IntentUnknown      Intent = "unknown"
	IntentCurrentJob   Intent = "current_job"
	IntentFirstJob     Intent = "first_job"
	IntentEducation    Intent = "education"
	IntentProject      Intent = "project"
	IntentTechnologies Intent = "technologies"
	IntentContact      Intent = "contact"
)

var keywords = map[string]map[Intent]int{
	"faz": {
		IntentCurrentJob: 3,
	},
	"trabalha": {
		IntentCurrentJob: 3,
	},
	"trabalho": {
		IntentCurrentJob: 2,
	},
	"profissão": {
		IntentCurrentJob: 3,
	},
	"profissao": {
		IntentCurrentJob: 3,
	},
	"cargo": {
		IntentCurrentJob: 3,
	},
	"atualmente": {
		IntentCurrentJob: 2,
	},
	"primeiro": {
		IntentFirstJob: 3,
	},
	"emprego": {
		IntentFirstJob:   2,
		IntentCurrentJob: 1,
	},
	"estágio": {
		IntentFirstJob: 2,
	},
	"estagio": {
		IntentFirstJob: 2,
	},
	"estagiário": {
		IntentFirstJob: 2,
	},
	"estagiario": {
		IntentFirstJob: 2,
	},
	"estuda": {
		IntentEducation: 3,
	},
	"faculdade": {
		IntentEducation: 3,
	},
	"curso": {
		IntentEducation: 2,
	},
	"formação": {
		IntentEducation: 3,
	},
	"formacao": {
		IntentEducation: 3,
	},
	"projeto": {
		IntentProject: 3,
	},
	"projetos": {
		IntentProject: 3,
	},
	"auronix": {
		IntentProject: 3,
	},
	"ggcompress": {
		IntentProject: 3,
	},
	"usadas": {
		IntentTechnologies: 3,
	},
	"usado": {
		IntentTechnologies: 3,
	},
	"tecnologia": {
		IntentTechnologies: 3,
	},
	"tecnologias": {
		IntentTechnologies: 3,
	},
	"stack": {
		IntentTechnologies: 3,
	},
	"linguagem": {
		IntentTechnologies: 2,
	},
	"linguagens": {
		IntentTechnologies: 2,
	},
	"contato": {
		IntentContact: 3,
	},
	"email": {
		IntentContact: 3,
	},
	"telefone": {
		IntentContact: 3,
	},
	"linkedin": {
		IntentContact: 3,
	},
	"github": {
		IntentContact: 3,
	},
}

func DetectIntent(tokens []string) Intent {
	scores := map[Intent]int{
		IntentCurrentJob:   0,
		IntentFirstJob:     0,
		IntentEducation:    0,
		IntentProject:      0,
		IntentTechnologies: 0,
		IntentContact:      0,
	}

	for _, token := range tokens {
		tokenScores, exists := keywords[token]

		if !exists {
			continue
		}

		for intent, score := range tokenScores {
			scores[intent] += score
		}
	}

	bestIntent := IntentUnknown
	bestScore := 0

	for intent, score := range scores {
		if score > bestScore {
			bestIntent = intent
			bestScore = score
		}
	}

	return bestIntent
}
