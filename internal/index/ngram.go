package index

import (
	"sort"

	"ai-agent/internal/domain"
	"ai-agent/internal/language"
)

func buildNGramIndexes(
	currentIndex *Index,
) {
	for _, targetLanguage := range indexedLanguages {
		terms := make(
			[]string,
			0,
			len(
				currentIndex.inverted[targetLanguage],
			),
		)

		for term := range currentIndex.inverted[targetLanguage] {
			terms = append(
				terms,
				term,
			)
		}

		sort.Strings(terms)

		currentIndex.vocabulary[targetLanguage] = terms

		for _, term := range terms {
			grams :=
				language.CharacterNGrams(
					term,
					3,
					4,
				)

			for _, gram := range grams {
				currentIndex.ngramTerms[targetLanguage][gram] =
					appendUniqueString(
						currentIndex.ngramTerms[targetLanguage][gram],
						term,
					)
			}
		}

		for gram, candidates := range currentIndex.ngramTerms[targetLanguage] {
			sort.Strings(candidates)

			currentIndex.ngramTerms[targetLanguage][gram] = candidates
		}
	}
}

func appendUniqueString(
	values []string,
	value string,
) []string {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}

	return append(
		values,
		value,
	)
}

func uniqueCandidateTerms(
	currentIndex *Index,
	targetLanguage domain.Language,
	grams []string,
) []string {
	seen := make(
		map[string]struct{},
	)

	result := make(
		[]string,
		0,
	)

	for _, gram := range grams {
		for _, term := range currentIndex.ngramTerms[targetLanguage][gram] {
			if _, exists := seen[term]; exists {
				continue
			}

			seen[term] = struct{}{}

			result = append(
				result,
				term,
			)
		}
	}

	sort.Strings(result)

	return result
}
