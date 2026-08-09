package language

import "ai-agent/internal/domain"

func Tokenize(value string) []string {
	return TokenizeNormalized(Normalize(value))
}

func TokenizeNormalized(value string) []string {
	if value == "" {
		return nil
	}

	tokens := make([]string, 0)
	start := -1

	for index, current := range value {
		if current == ' ' {
			if start >= 0 {
				tokens = append(tokens, value[start:index])
				start = -1
			}

			continue
		}

		if start < 0 {
			start = index
		}
	}

	if start >= 0 {
		tokens = append(tokens, value[start:])
	}

	return tokens
}

func ContentTerms(
	tokens []string,
	language domain.Language,
) []string {
	if len(tokens) == 0 {
		return nil
	}

	terms := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if token == "" {
			continue
		}

		if IsStopWord(language, token) {
			continue
		}

		terms = append(terms, token)
	}

	return terms
}

func UniqueTokens(tokens []string) []string {
	if len(tokens) == 0 {
		return nil
	}

	result := make([]string, 0, len(tokens))
	seen := make(map[string]struct{}, len(tokens))

	for _, token := range tokens {
		if token == "" {
			continue
		}

		if _, exists := seen[token]; exists {
			continue
		}

		seen[token] = struct{}{}
		result = append(result, token)
	}

	return result
}
