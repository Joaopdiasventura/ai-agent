package nlp

type Language string

const (
	LanguagePortuguese Language = "pt"
	LanguageEnglish    Language = "en"
)

func DetectLanguage(tokens []string) Language {
	portugueseScore := 0
	englishScore := 0

	for _, token := range tokens {
		portugueseScore += portugueseLanguageWeights[token]
		englishScore += englishLanguageWeights[token]
	}

	if englishScore > portugueseScore {
		return LanguageEnglish
	}

	return LanguagePortuguese
}

var portugueseLanguageWeights = map[string]int{
	"me":           1,
	"qual":         4,
	"quais":        4,
	"quem":         4,
	"onde":         4,
	"como":         4,
	"informação":   4,
	"informacao":   4,
	"profissional": 3,
	"perfil":       3,
	"ele":          3,
	"dele":         3,
	"seu":          3,
	"sua":          3,
	"fale":         4,
	"fala":         4,
	"falar":        4,
	"sobre":        4,
	"conte":        4,
	"explique":     4,
	"descreva":     4,
	"resumo":       3,
	"pergunta":     3,
	"relacao":      3,
	"relação":      3,
	"trabalho":     3,
	"trabalha":     3,
	"trabalhou":    3,
	"estuda":       3,
	"estudou":      3,
	"cursa":        3,
	"cursando":     3,
	"faculdade":    3,
	"formação":     3,
	"formacao":     3,
	"emprego":      3,
	"experiencia":  3,
	"experiência":  3,
	"projeto":      3,
	"projetos":     3,
	"portfolio":    1,
	"portfólio":    3,
	"contato":      3,
	"telefone":     3,
	"email":        1,
	"tecnologia":   3,
	"tecnologias":  3,
	"servico":      3,
	"serviço":      3,
	"servicos":     3,
	"serviços":     3,
	"contratar":    3,
	"contrato":     2,
	"auronix":      0,
	"joão":         0,
}

var englishLanguageWeights = map[string]int{
	"who":            4,
	"what":           4,
	"where":          4,
	"information":    4,
	"professional":   3,
	"profile":        3,
	"certification":  4,
	"certifications": 4,
	"why":            4,
	"how":            4,
	"which":          4,
	"tell":           4,
	"about":          4,
	"explain":        4,
	"describe":       4,
	"summary":        3,
	"question":       3,
	"unrelated":      3,
	"work":           3,
	"works":          3,
	"worked":         3,
	"job":            3,
	"experience":     3,
	"project":        3,
	"projects":       3,
	"portfolio":      2,
	"contact":        3,
	"phone":          3,
	"technologies":   3,
	"technology":     3,
	"stack":          3,
	"service":        3,
	"services":       3,
	"hire":           3,
	"hiring":         3,
	"does":           1,
	"his":            2,
	"him":            2,
	"he":             2,
	"should":         1,
	"can":            1,
	"email":          1,
	"linkedin":       0,
	"github":         0,
	"auronix":        0,
	"joão":           0,
}
