package domain

type Language string

const (
	LanguageUnknown    Language = "unknown"
	LanguagePortuguese Language = "pt"
	LanguageEnglish    Language = "en"
)

func (l Language) IsSupported() bool {
	return l == LanguagePortuguese || l == LanguageEnglish
}

type LocalizedText struct {
	PT string
	EN string
}

func (t LocalizedText) For(language Language) string {
	switch language {
	case LanguageEnglish:
		if t.EN != "" {
			return t.EN
		}
		return t.PT

	default:
		if t.PT != "" {
			return t.PT
		}
		return t.EN
	}
}

func (t LocalizedText) Empty() bool {
    return t.PT == "" && t.EN == ""
}
