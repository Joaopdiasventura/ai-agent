package nlp

type AnswerMode string

const (
	AnswerModeDefault    AnswerMode = "default"
	AnswerModeAbout      AnswerMode = "about"
	AnswerModeTechnology AnswerMode = "technology"
	AnswerModeComparison AnswerMode = "comparison"
)

var answerModeToken = map[string]AnswerMode{
	"sobre":       AnswerModeAbout,
	"fale":        AnswerModeAbout,
	"explique":    AnswerModeAbout,
	"descreva":    AnswerModeAbout,
	"resuma":      AnswerModeAbout,
	"resumo":      AnswerModeAbout,
	"tecnologia":  AnswerModeTechnology,
	"tecnologias": AnswerModeTechnology,
	"stack":       AnswerModeTechnology,
	"usa":         AnswerModeTechnology,
	"utiliza":     AnswerModeTechnology,
	"utilizadas":  AnswerModeTechnology,
	"usadas":      AnswerModeTechnology,
	"comparar":    AnswerModeComparison,
	"comparação":  AnswerModeComparison,
	"comparacao":  AnswerModeComparison,
	"diferenca":   AnswerModeComparison,
	"diferença":   AnswerModeComparison,
	"melhor":      AnswerModeComparison,
	"principal":   AnswerModeComparison,
}

func DetectAnswerMode(tokens []string) AnswerMode {
	for _, token := range tokens {
		answerMode, exists := answerModeToken[token]

		if exists {
			return answerMode
		}
	}

	return AnswerModeDefault
}
