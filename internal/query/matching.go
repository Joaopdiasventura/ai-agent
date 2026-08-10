package query

import (
	"strings"
	"unicode/utf8"
)

func containsPhrase(
	normalized string,
	phrase string,
) bool {
	if normalized == "" || phrase == "" {
		return false
	}

	if normalized == phrase {
		return true
	}

	paddedText := " " + normalized + " "
	paddedPhrase := " " + phrase + " "

	return strings.Contains(
		paddedText,
		paddedPhrase,
	)
}

func runeLength(value string) int {
	return utf8.RuneCountInString(value)
}

func isSingleToken(value string) bool {
	return !strings.ContainsRune(value, ' ')
}

func maximum(left, right float64) float64 {
	return max(left, right)
}
