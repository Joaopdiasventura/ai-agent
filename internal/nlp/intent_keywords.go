package nlp

var intentPriority = []Intent{
	IntentContact,
	IntentHireReason,
	IntentVisitorServices,
	IntentVisitorProjects,
	IntentProject,
	IntentTechnologies,
	IntentVisitorSummary,
	IntentAbout,
	IntentCurrentJob,
	IntentFirstJob,
	IntentEducation,
}

var intentKeywords = map[string]map[Intent]int{
	"faz": {
		IntentVisitorSummary: 3,
		IntentCurrentJob:     2,
		IntentAbout:          1,
	},
	"joão": {
		IntentVisitorSummary: 2,
		IntentAbout:          1,
	},
	"joao": {
		IntentVisitorSummary: 2,
		IntentAbout:          1,
	},
	"quem": {
		IntentVisitorSummary: 3,
	},
	"who": {
		IntentVisitorSummary: 3,
	},
	"what": {
		IntentVisitorSummary: 1,
	},
	"resumo": {
		IntentVisitorSummary: 3,
	},
	"summary": {
		IntentVisitorSummary: 3,
	},
	"perfil": {
		IntentVisitorSummary: 3,
	},
	"profile": {
		IntentVisitorSummary: 3,
	},
	"apresentação": {
		IntentVisitorSummary: 3,
	},
	"apresentacao": {
		IntentVisitorSummary: 3,
	},
	"desenvolvido": {
		IntentTechnologies: 3,
	},
	"trabalha": {
		IntentCurrentJob: 3,
	},
	"work": {
		IntentCurrentJob:      2,
		IntentVisitorServices: 1,
	},
	"works": {
		IntentCurrentJob: 3,
	},
	"working": {
		IntentCurrentJob: 2,
	},
	"trabalho": {
		IntentCurrentJob: 2,
	},
	"job": {
		IntentCurrentJob: 3,
	},
	"current": {
		IntentCurrentJob: 2,
	},
	"currently": {
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
	"first": {
		IntentFirstJob: 3,
	},
	"emprego": {
		IntentFirstJob:   2,
		IntentCurrentJob: 1,
	},
	"intern": {
		IntentFirstJob: 2,
	},
	"internship": {
		IntentFirstJob: 2,
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
	"studies": {
		IntentEducation: 3,
	},
	"study": {
		IntentEducation: 3,
	},
	"faculdade": {
		IntentEducation: 3,
	},
	"college": {
		IntentEducation: 3,
	},
	"education": {
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
		IntentVisitorProjects: 3,
		IntentProject:         2,
	},
	"project": {
		IntentVisitorProjects: 3,
		IntentProject:         2,
	},
	"projetos": {
		IntentVisitorProjects: 3,
		IntentProject:         2,
	},
	"projects": {
		IntentVisitorProjects: 3,
		IntentProject:         2,
	},
	"portfólio": {
		IntentVisitorProjects: 3,
	},
	"portfolio": {
		IntentVisitorProjects: 3,
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
	"technology": {
		IntentTechnologies: 3,
	},
	"tecnologias": {
		IntentTechnologies: 3,
	},
	"technologies": {
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
	"contact": {
		IntentContact: 3,
	},
	"contatar": {
		IntentContact: 3,
	},
	"conversar": {
		IntentContact:    2,
		IntentHireReason: 1,
	},
	"email": {
		IntentContact: 3,
	},
	"telefone": {
		IntentContact: 3,
	},
	"phone": {
		IntentContact: 3,
	},
	"linkedin": {
		IntentContact: 3,
	},
	"github": {
		IntentContact: 3,
	},
	"serviço": {
		IntentVisitorServices: 10,
	},
	"service": {
		IntentVisitorServices: 10,
	},
	"servico": {
		IntentVisitorServices: 10,
	},
	"serviços": {
		IntentVisitorServices: 10,
	},
	"services": {
		IntentVisitorServices: 10,
	},
	"servicos": {
		IntentVisitorServices: 10,
	},
	"pode": {
		IntentVisitorServices: 2,
	},
	"desenvolve": {
		IntentVisitorServices: 3,
	},
	"develop": {
		IntentVisitorServices: 3,
	},
	"develops": {
		IntentVisitorServices: 3,
	},
	"build": {
		IntentVisitorServices: 2,
		IntentVisitorSummary:  1,
	},
	"builds": {
		IntentVisitorServices: 2,
		IntentVisitorSummary:  1,
	},
	"cria": {
		IntentVisitorServices: 2,
		IntentVisitorSummary:  1,
	},
	"ajudar": {
		IntentVisitorServices: 2,
		IntentHireReason:      2,
	},
	"contratar": {
		IntentHireReason: 3,
	},
	"hire": {
		IntentHireReason: 3,
	},
	"hiring": {
		IntentHireReason: 3,
	},
	"contrataria": {
		IntentHireReason: 3,
	},
	"contrato": {
		IntentHireReason: 2,
	},
	"vale": {
		IntentHireReason: 2,
	},
	"motivo": {
		IntentHireReason: 2,
	},
	"motivos": {
		IntentHireReason: 2,
	},
	"diferencial": {
		IntentHireReason: 3,
	},
	"diferenciais": {
		IntentHireReason: 3,
	},
	"reason": {
		IntentHireReason: 2,
	},
	"reasons": {
		IntentHireReason: 2,
	},
	"differential": {
		IntentHireReason: 3,
	},
	"why": {
		IntentHireReason: 2,
	},
	"sobre": {
		IntentAbout: 3,
	},
	"about": {
		IntentAbout: 3,
	},
	"fale": {
		IntentAbout: 2,
	},
	"tell": {
		IntentAbout: 2,
	},
	"explique": {
		IntentAbout: 3,
	},
	"explain": {
		IntentAbout: 3,
	},
	"descreva": {
		IntentAbout: 3,
	},
	"é": {
		IntentAbout: 2,
	},
}
