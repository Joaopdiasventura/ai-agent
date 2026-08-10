package cli

import (
	"strings"

	"ai-agent/internal/domain"
)

func NotFoundMessage(
	language domain.Language,
) string {
	if language ==
		domain.LanguageEnglish {
		return "I don't have that specific information, but I can talk about João's experience, projects, technologies, services, or contact details."
	}

	return "Não encontrei essa informação específica, mas posso falar sobre experiência, projetos, tecnologias, serviços ou contato do João."
}

func ShouldExit(
	input string,
) bool {
	value :=
		strings.ToLower(
			strings.TrimSpace(
				input,
			),
		)

	switch value {
	case "sair",
		"exit",
		"quit",
		"encerrar":
		return true

	default:
		return false
	}
}
