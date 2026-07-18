package nlp

var intentPriority = []Intent{
	IntentContact,
	IntentProject,
	IntentTechnologies,
	IntentAbout,
	IntentCurrentJob,
	IntentFirstJob,
	IntentEducation,
}

var intentKeywords = map[string]map[Intent]int{
	"faz": {
		IntentCurrentJob: 3,
		IntentAbout:      1,
	},
	"joão": {
		IntentAbout: 1,
	},
	"joao": {
		IntentAbout: 1,
	},
	"desenvolvido": {
		IntentTechnologies: 3,
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
	"modularis": {
		IntentProject: 3,
	},
	"votrix": {
		IntentProject: 3,
	},
	"ggcompress": {
		IntentProject: 3,
	},
	"vox": {
		IntentProject: 3,
	},
	"etecfy": {
		IntentProject: 3,
	},
	"usadas": {
		IntentTechnologies: 3,
	},
	"usados": {
		IntentTechnologies: 3,
	},
	"usada": {
		IntentTechnologies: 3,
	},
	"usado": {
		IntentTechnologies: 3,
	},
	"utiliza": {
		IntentTechnologies: 3,
	},
	"utilizam": {
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
	"usa": {
		IntentTechnologies: 3,
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
	"sobre": {
		IntentAbout: 3,
	},
	"fale": {
		IntentAbout: 2,
	},
	"explique": {
		IntentAbout: 3,
	},
	"descreva": {
		IntentAbout: 3,
	},
	"é": {
		IntentAbout: 2,
	},
}
