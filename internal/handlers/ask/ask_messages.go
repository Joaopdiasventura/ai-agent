package handlers

type Language string

const (
	LanguagePT Language = "pt"
	LanguageEN Language = "en"
)

type ValidationMessages struct {
	InvalidBody     string
	Empty           string
	TooLong         string
	PromptInjection string
	Sensitive       string
	LowQuality      string
	OutOfScope      string
	ServiceError    string
}

var validationMessages = map[Language]ValidationMessages{
	LanguagePT: {
		InvalidBody:     "Body da requisição inválido.",
		Empty:           "Digite uma pergunta sobre experiência, projetos, habilidades ou trajetória profissional.",
		TooLong:         "Sua pergunta está muito longa. Faça uma pergunta mais objetiva sobre experiência, projetos, habilidades ou trajetória profissional.",
		PromptInjection: "Não posso responder solicitações para alterar regras internas, revelar instruções ou ignorar o contexto do portfólio.",
		Sensitive:       "Essa informação não está disponível no contexto público do portfólio.",
		LowQuality:      "Faça uma pergunta mais clara sobre experiência, projetos, habilidades, formação ou trajetória profissional.",
		OutOfScope:      "Posso responder apenas perguntas relacionadas ao portfólio: experiência, projetos, habilidades técnicas, formação, trajetória profissional ou contratação.",
		ServiceError:    "Serviço temporariamente indisponível.",
	},
	LanguageEN: {
		InvalidBody:     "Invalid request body.",
		Empty:           "Ask a question about experience, projects, skills, or professional background.",
		TooLong:         "Your question is too long. Ask a more objective question about experience, projects, skills, or professional background.",
		PromptInjection: "I cannot answer requests to change internal rules, reveal instructions, or ignore the portfolio context.",
		Sensitive:       "This information is not available in the public portfolio context.",
		LowQuality:      "Ask a clearer question about experience, projects, skills, education, or professional background.",
		OutOfScope:      "I can only answer questions related to the portfolio: experience, projects, technical skills, education, professional background, or hiring.",
		ServiceError:    "Service temporarily unavailable.",
	},
}

func messagesForLanguage(lang Language) ValidationMessages {
	messages, ok := validationMessages[lang]
	if !ok {
		return validationMessages[LanguagePT]
	}

	return messages
}
