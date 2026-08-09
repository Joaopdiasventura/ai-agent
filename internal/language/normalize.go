package language

import (
	"strings"
	"unicode"
)

func Normalize(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))

	if value == "" {
		return ""
	}

	runes := []rune(value)

	var builder strings.Builder

	lastSeparator := true

	for index, current := range runes {
		current = foldRune(current)

		if unicode.IsLetter(current) || unicode.IsDigit(current) {
			builder.WriteRune(current)
			lastSeparator = false
			continue
		}

		if isDecimalSeparator(runes, index) {
			builder.WriteRune('.')
			lastSeparator = false
			continue
		}

		if !lastSeparator && builder.Len() > 0 {
			builder.WriteRune(' ')
			lastSeparator = true
		}
	}

	return strings.TrimSpace(builder.String())
}

func foldRune(value rune) rune {
	switch value {
	case 'á', 'à', 'â', 'ã', 'ä':
		return 'a'
	case 'é', 'è', 'ê', 'ë':
		return 'e'
	case 'í', 'ì', 'î', 'ï':
		return 'i'
	case 'ó', 'ò', 'ô', 'õ', 'ö':
		return 'o'
	case 'ú', 'ù', 'û', 'ü':
		return 'u'
	case 'ç':
		return 'c'
	case 'ñ':
		return 'n'
	default:
		return value
	}
}

func isDecimalSeparator(runes []rune, index int) bool {
	if index <= 0 || index >= len(runes)-1 {
		return false
	}

	current := runes[index]

	if current != '.' && current != ',' {
		return false
	}

	previous := runes[index-1]
	next := runes[index+1]

	return unicode.IsDigit(previous) && unicode.IsDigit(next)
}
