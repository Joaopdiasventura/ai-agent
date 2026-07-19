package retrieval

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/vectorindex"
	"sort"
	"strings"
	"unicode"
)

func LexicalSearch(index vectorindex.Index, query domain.Query, limit int) []domain.SearchResult {
	if limit <= 0 || len(index.Entries) == 0 {
		return []domain.SearchResult{}
	}

	queryTokens := lexicalTokens(strings.Join(append([]string{query.Text}, query.ExactTerms...), " "))
	if len(queryTokens) == 0 {
		return []domain.SearchResult{}
	}

	results := make([]domain.SearchResult, 0)

	for _, entry := range index.Entries {
		score := lexicalScore(entry.Document, queryTokens)
		if score <= 0 {
			continue
		}

		document := entry.Document
		results = append(results, domain.SearchResult{
			Document:    &document,
			Score:       score,
			LexicalRank: 0,
			Sources:     []string{"lexical"},
		})
	}

	sort.Slice(results, func(firstIndex int, secondIndex int) bool {
		if results[firstIndex].Score == results[secondIndex].Score {
			return results[firstIndex].Document.ID < results[secondIndex].Document.ID
		}

		return results[firstIndex].Score > results[secondIndex].Score
	})

	if limit > len(results) {
		limit = len(results)
	}

	results = results[:limit]
	for index := range results {
		results[index].LexicalRank = index + 1
	}

	return results
}

func lexicalScore(document domain.Document, queryTokens []string) float64 {
	documentFields := []string{
		document.ID,
		document.Language,
		document.Category,
		document.Subject,
		document.Project,
		document.TemporalStatus,
		document.Content,
		strings.Join(document.Keywords, " "),
	}

	documentText := normalizeLexical(strings.Join(documentFields, " "))
	documentTokens := lexicalTokens(documentText)
	documentTokenSet := make(map[string]struct{}, len(documentTokens))

	for _, token := range documentTokens {
		documentTokenSet[token] = struct{}{}
	}

	score := 0.0

	for _, token := range queryTokens {
		if _, exists := documentTokenSet[token]; exists {
			score += exactLexicalWeight(token)
			continue
		}

		if len(token) < 4 {
			continue
		}

		for documentToken := range documentTokenSet {
			if len(documentToken) < 4 {
				continue
			}

			if strings.Contains(documentToken, token) || strings.Contains(token, documentToken) {
				score += 0.25
				break
			}
		}
	}

	return score
}

func exactLexicalWeight(token string) float64 {
	switch {
	case strings.Contains(token, "@"), strings.Contains(token, "+"), strings.Contains(token, "/"):
		return 6
	case hasDigit(token):
		return 4
	case len(token) <= 3:
		return 2
	default:
		return 1
	}
}

func lexicalTokens(text string) []string {
	text = normalizeLexical(text)
	fields := strings.FieldsFunc(text, func(value rune) bool {
		return !(unicode.IsLetter(value) || unicode.IsDigit(value) || value == '@' || value == '.' || value == '+' || value == '/' || value == '-' || value == '#')
	})

	tokens := make([]string, 0, len(fields))
	seen := make(map[string]struct{}, len(fields))

	for _, field := range fields {
		field = strings.Trim(field, ".-_")
		if field == "" {
			continue
		}

		if _, exists := seen[field]; exists {
			continue
		}

		seen[field] = struct{}{}
		tokens = append(tokens, field)
	}

	return tokens
}

func normalizeLexical(text string) string {
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "â", "a", "ã", "a", "ä", "a",
		"é", "e", "ê", "e", "è", "e", "ë", "e",
		"í", "i", "î", "i", "ì", "i", "ï", "i",
		"ó", "o", "ô", "o", "õ", "o", "ò", "o", "ö", "o",
		"ú", "u", "û", "u", "ù", "u", "ü", "u",
		"ç", "c",
	)

	return replacer.Replace(strings.ToLower(text))
}

func hasDigit(text string) bool {
	for _, value := range text {
		if unicode.IsDigit(value) {
			return true
		}
	}

	return false
}
