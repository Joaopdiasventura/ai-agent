package generation

import (
	"strings"

	"ai-agent/internal/domain"
)

func outputLanguage(
	value domain.Language,
) domain.Language {
	switch value {
	case domain.LanguagePortuguese:
		return domain.LanguagePortuguese

	case domain.LanguageEnglish:
		return domain.LanguageEnglish

	default:
		return domain.LanguageEnglish
	}
}

func localizedText(
	value domain.LocalizedText,
	targetLanguage domain.Language,
) string {
	result :=
		strings.TrimSpace(
			value.For(
				targetLanguage,
			),
		)

	if result != "" {
		return result
	}

	switch targetLanguage {
	case domain.LanguagePortuguese:
		if value.PT != "" {
			return strings.TrimSpace(
				value.PT,
			)
		}

		return strings.TrimSpace(
			value.EN,
		)

	case domain.LanguageEnglish:
		if value.EN != "" {
			return strings.TrimSpace(
				value.EN,
			)
		}

		return strings.TrimSpace(
			value.PT,
		)

	default:
		if value.EN != "" {
			return strings.TrimSpace(
				value.EN,
			)
		}

		return strings.TrimSpace(
			value.PT,
		)
	}
}

func entityLabel(
	material Material,
	id domain.EntityID,
	targetLanguage domain.Language,
) string {
	if id == "" {
		return ""
	}

	entity, found :=
		material.Entity(id)

	if !found {
		return ""
	}

	return localizedText(
		entity.Name,
		targetLanguage,
	)
}

func factStatement(
	material Material,
	id domain.FactID,
	targetLanguage domain.Language,
) string {
	fact, found :=
		material.Fact(id)

	if !found {
		return ""
	}

	return normalizeSentence(
		localizedText(
			fact.Statement,
			targetLanguage,
		),
	)
}

func normalizeSentence(
	value string,
) string {
	value =
		strings.TrimSpace(value)

	if value == "" {
		return ""
	}

	last :=
		value[len(value)-1]

	switch last {
	case '.',
		'!',
		'?':
		return value

	default:
		return value + "."
	}
}

func joinSentences(
	values []string,
) string {
	filtered := make(
		[]string,
		0,
		len(values),
	)

	seen := make(
		map[string]struct{},
	)

	for _, value := range values {
		value =
			normalizeSentence(value)

		if value == "" {
			continue
		}

		if _, exists :=
			seen[value]; exists {
			continue
		}

		seen[value] =
			struct{}{}

		filtered = append(
			filtered,
			value,
		)
	}

	return strings.Join(
		filtered,
		" ",
	)
}

func joinNatural(
	values []string,
	targetLanguage domain.Language,
) string {
	filtered := make(
		[]string,
		0,
		len(values),
	)

	seen := make(
		map[string]struct{},
	)

	for _, value := range values {
		value =
			strings.TrimSpace(value)

		if value == "" {
			continue
		}

		if _, exists :=
			seen[value]; exists {
			continue
		}

		seen[value] =
			struct{}{}

		filtered = append(
			filtered,
			value,
		)
	}

	switch len(filtered) {
	case 0:
		return ""

	case 1:
		return filtered[0]

	case 2:
		if targetLanguage ==
			domain.LanguagePortuguese {
			return filtered[0] +
				" e " +
				filtered[1]
		}

		return filtered[0] +
			" and " +
			filtered[1]
	}

	if targetLanguage ==
		domain.LanguagePortuguese {
		return strings.Join(
			filtered[:len(filtered)-1],
			", ",
		) +
			" e " +
			filtered[len(filtered)-1]
	}

	return strings.Join(
		filtered[:len(filtered)-1],
		", ",
	) +
		", and " +
		filtered[len(filtered)-1]
}