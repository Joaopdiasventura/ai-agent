package nlp

type Language string

const (
	LanguagePortuguese Language = "pt"
	LanguageEnglish    Language = "en"
)

func DetectLanguage(tokens []string) Language {
	for _, token := range tokens {
		switch token {
		case "who", "what", "where", "why", "how", "which",
			"tell", "about", "me", "does", "should", "can",
			"question", "unrelated",
			"work", "works", "worked", "job", "projects", "portfolio",
			"contact", "email", "phone", "linkedin", "github",
			"technologies", "technology", "stack", "services", "hire":
			return LanguageEnglish
		}
	}

	return LanguagePortuguese
}
