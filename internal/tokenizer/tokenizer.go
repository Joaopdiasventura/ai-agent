package tokenizer

import (
	"regexp"
	"strings"
)

var stopWords = map[string]bool{
	"a": true, "o": true, "as": true, "os": true,
	"de": true, "da": true, "do": true, "das": true, "dos": true,
	"e": true, "é": true, "em": true, "um": true, "uma": true,
	"para": true, "com": true, "que": true, "por": true,
	"qual": true, "quais": true, "quem": true, "quando": true,
	"onde": true, "como": true, "porque": true,
	"foi": true, "era": true, "ser": true, "são": true,
	"ele": true, "ela": true, "dele": true, "dela": true,
	"seu": true, "sua": true, "seus": true, "suas": true,
}

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