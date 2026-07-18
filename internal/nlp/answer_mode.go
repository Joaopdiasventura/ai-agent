package nlp

type AnswerMode string

const (
	AnswerModeDefault    AnswerMode = "default"
	AnswerModeAbout      AnswerMode = "about"
	AnswerModeTechnology AnswerMode = "technology"
	AnswerModeComparison AnswerMode = "comparison"
)

var answerModeToken = map[string]AnswerMode{
	"sobre":        AnswerModeAbout,
	"about":        AnswerModeAbout,
	"fale":         AnswerModeAbout,
	"tell":         AnswerModeAbout,
	"explique":     AnswerModeAbout,
	"explain":      AnswerModeAbout,
	"descreva":     AnswerModeAbout,
	"describe":     AnswerModeAbout,
	"resuma":       AnswerModeAbout,
	"resumo":       AnswerModeAbout,
	"summary":      AnswerModeAbout,
	"tecnologia":   AnswerModeTechnology,
	"tecnologias":  AnswerModeTechnology,
	"technology":   AnswerModeTechnology,
	"technologies": AnswerModeTechnology,
	"stack":        AnswerModeTechnology,
	"usa":          AnswerModeTechnology,
	"use":          AnswerModeTechnology,
	"uses":         AnswerModeTechnology,
	"utiliza":      AnswerModeTechnology,
	"utilizadas":   AnswerModeTechnology,
	"usadas":       AnswerModeTechnology,
	"comparar":     AnswerModeComparison,
	"comparação":   AnswerModeComparison,
	"comparacao":   AnswerModeComparison,
	"diferenca":    AnswerModeComparison,
	"diferença":    AnswerModeComparison,
	"melhor":       AnswerModeComparison,
	"principal":    AnswerModeComparison,
	"compare":      AnswerModeComparison,
	"comparison":   AnswerModeComparison,
	"difference":   AnswerModeComparison,
	"best":         AnswerModeComparison,
	"main":         AnswerModeComparison,
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
