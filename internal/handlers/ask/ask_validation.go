package handlers

import (
	"strings"
	"unicode"
)

const MaxQuestionLength = 300

type QuestionValidation struct {
	Valid   bool
	Message string
}

func ValidateQuestion(question string, lang Language) QuestionValidation {
	messages := messagesForLanguage(lang)

	if question == "" {
		return QuestionValidation{
			Valid:   false,
			Message: messages.Empty,
		}
	}

	if len([]rune(question)) > MaxQuestionLength {
		return QuestionValidation{
			Valid:   false,
			Message: messages.TooLong,
		}
	}

	if hasPromptInjectionIntent(question) {
		return QuestionValidation{
			Valid:   false,
			Message: messages.PromptInjection,
		}
	}

	if hasSensitivePersonalIntent(question) {
		return QuestionValidation{
			Valid:   false,
			Message: messages.Sensitive,
		}
	}

	if isLowQualityInput(question) {
		return QuestionValidation{
			Valid:   false,
			Message: messages.LowQuality,
		}
	}

	if !isPortfolioRelated(question) {
		return QuestionValidation{
			Valid:   false,
			Message: messages.OutOfScope,
		}
	}

	return QuestionValidation{
		Valid: true,
	}
}

func normalizeQuestion(question string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(question)), " ")
}

func hasPromptInjectionIntent(question string) bool {
	value := strings.ToLower(question)

	terms := []string{
		"ignore previous",
		"ignore all",
		"system prompt",
		"hidden prompt",
		"developer message",
		"internal rules",
		"reveal prompt",
		"show prompt",
		"bypass",
		"jailbreak",
		"forget instructions",
		"ignore as regras",
		"ignore as instruções",
		"prompt interno",
		"prompt escondido",
		"mensagem de sistema",
		"regras internas",
		"revele o prompt",
		"mostre o prompt",
	}

	return containsAny(value, terms)
}

func hasSensitivePersonalIntent(question string) bool {
	value := strings.ToLower(question)

	terms := []string{
		"cpf",
		"rg",
		"endereço",
		"endereco",
		"número pessoal",
		"numero pessoal",
		"namorada",
		"namorado",
		"religião",
		"religiao",
		"política",
		"politica",
		"partido",
		"salário",
		"salario",
		"quanto ganha",
		"vida pessoal",
		"personal life",
		"home address",
		"religion",
		"politics",
		"salary",
	}

	return containsAny(value, terms)
}

func isPortfolioRelated(question string) bool {
	value := strings.ToLower(question)

	terms := []string{
		"joão",
		"joao",
		"ele",
		"perfil",
		"portfolio",
		"portfólio",
		"currículo",
		"curriculo",
		"linkedin",
		"github",
		"contato",
		"contact",
		"email",
		"e-mail",
		"telefone",
		"phone",
		"whatsapp",
		"linkedin",
		"github",
		"experiência",
		"experiencia",
		"carreira",
		"trajetória",
		"trajetoria",
		"formação",
		"formacao",
		"educação",
		"educacao",
		"habilidade",
		"skill",
		"stack",
		"tecnologia",
		"projeto",
		"auronix",
		"modularis",
		"ggcompress",
		"votrix",
		"auditex",
		"backend",
		"frontend",
		"full stack",
		"fullstack",
		"node",
		"nestjs",
		"typescript",
		"go",
		"golang",
		"java",
		"spring",
		"angular",
		"postgresql",
		"mongodb",
		"redis",
		"rabbitmq",
		"aws",
		"docker",
		"microservices",
		"microsserviços",
		"event-driven",
		"sistemas distribuídos",
		"sistemas distribuidos",
		"ia",
		"ai",
		"openai",
		"gemini",
		"contratar",
		"hire",
		"recruiter",
		"recrutador",
		"vaga",
		"senioridade",
		"pleno",
		"junior",
		"júnior",
		"produção",
		"producao",
		"arquitetura",
		"performance",
		"pipeline",
		"streams",
		"sse",
		"api",
		"rest",
		"cloud",
	}

	return containsAny(value, terms)
}

func isLowQualityInput(question string) bool {
	runes := []rune(question)

	if len(runes) < 4 {
		return true
	}

	if hasTooManyRepeatedChars(runes) {
		return true
	}

	letters := 0
	for _, char := range runes {
		if unicode.IsLetter(char) {
			letters++
		}
	}

	return letters < 3
}

func hasTooManyRepeatedChars(chars []rune) bool {
	if len(chars) < 8 {
		return false
	}

	repeated := 1

	for i := 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			repeated++
			if repeated >= 6 {
				return true
			}
		} else {
			repeated = 1
		}
	}

	return false
}

func containsAny(value string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(value, term) {
			return true
		}
	}

	return false
}
