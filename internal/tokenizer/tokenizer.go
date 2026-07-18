package tokenizer

import (
	"regexp"
	"strings"
)

var invalidCharactersRegex = regexp.MustCompile(`[^\p{L}\p{N}]+`)

func IsStopWord(word string) bool {
	return stopWords[word]
}

func Tokenize(text string) []string {
	text = strings.ToLower(text)
	text = invalidCharactersRegex.ReplaceAllString(text, " ")

	words := strings.Fields(text)
	tokens := make([]string, 0, len(words))

	for _, word := range words {
		if !IsStopWord(word) {
			tokens = append(tokens, word)
		}
	}

	return tokens
}
