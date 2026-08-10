package query

import (
	"sort"
	"strings"
	"unicode"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/language"
)

type entityAlias struct {
	EntityID   domain.EntityID
	EntityType domain.EntityType
	Language   domain.Language
	Normalized string
	Weight     float64
}

type EntityExtractor struct {
	aliases []entityAlias
}

func NewEntityExtractor(
	base *knowledge.Knowledge,
) *EntityExtractor {
	aliases := make([]entityAlias, 0)

	for _, entity := range base.Entities() {
		aliases = appendEntityAlias(
			aliases,
			entity,
			domain.LanguagePortuguese,
			entity.Name.PT,
			1,
		)

		aliases = appendEntityAlias(
			aliases,
			entity,
			domain.LanguageEnglish,
			entity.Name.EN,
			1,
		)

		for aliasLanguage, values := range entity.Aliases {
			for _, value := range values {
				weight := 0.95

				if aliasLanguage == domain.LanguageUnknown {
					weight = 0.97
				}

				aliases = appendEntityAlias(
					aliases,
					entity,
					aliasLanguage,
					value,
					weight,
				)
			}
		}
	}

	sort.Slice(
		aliases,
		func(i int, j int) bool {
			leftLength :=
				runeLength(aliases[i].Normalized)

			rightLength :=
				runeLength(aliases[j].Normalized)

			if leftLength != rightLength {
				return leftLength > rightLength
			}

			if aliases[i].Weight != aliases[j].Weight {
				return aliases[i].Weight >
					aliases[j].Weight
			}

			if aliases[i].Normalized !=
				aliases[j].Normalized {
				return aliases[i].Normalized <
					aliases[j].Normalized
			}

			return aliases[i].EntityID <
				aliases[j].EntityID
		},
	)

	return &EntityExtractor{
		aliases: aliases,
	}
}

func (e *EntityExtractor) Extract(
	original string,
	normalized string,
	terms []string,
	targetLanguage domain.Language,
) []domain.EntityMatch {
	matches := make(
		map[domain.EntityID]domain.EntityMatch,
	)

	for _, alias := range e.aliases {
		if !languageCompatible(
			alias.Language,
			targetLanguage,
		) {
			continue
		}

		if !containsPhrase(
			normalized,
			alias.Normalized,
		) {
			continue
		}

		if shouldRejectAmbiguousGo(
			original,
			normalized,
			alias,
			targetLanguage,
		) {
			continue
		}

		score := alias.Weight

		if alias.Language != domain.LanguageUnknown &&
			targetLanguage != domain.LanguageUnknown &&
			alias.Language != targetLanguage {
			score *= 0.85
		}

		updateEntityMatch(
			matches,
			domain.EntityMatch{
				EntityID:    alias.EntityID,
				Score:       score,
				Explicit:    true,
				MatchedText: alias.Normalized,
			},
		)
	}

	for _, term := range terms {
		if runeLength(term) < 4 {
			continue
		}

		for _, alias := range e.aliases {
			if !languageCompatible(
				alias.Language,
				targetLanguage,
			) {
				continue
			}

			if !isSingleToken(alias.Normalized) {
				continue
			}

			if runeLength(alias.Normalized) < 4 {
				continue
			}

			if term == alias.Normalized {
				continue
			}

			similarity :=
				language.FuzzySimilarity(
					term,
					alias.Normalized,
				)

			threshold :=
				entityFuzzyThreshold(
					term,
					alias.Normalized,
				)

			if similarity < threshold {
				continue
			}

			score :=
				similarity *
					alias.Weight *
					0.88

			updateEntityMatch(
				matches,
				domain.EntityMatch{
					EntityID:    alias.EntityID,
					Score:       score,
					Explicit:    true,
					MatchedText: term,
				},
			)
		}
	}

	result := make(
		[]domain.EntityMatch,
		0,
		len(matches),
	)

	for _, match := range matches {
		result = append(result, match)
	}

	sort.Slice(
		result,
		func(i int, j int) bool {
			if result[i].Score != result[j].Score {
				return result[i].Score >
					result[j].Score
			}

			return result[i].EntityID <
				result[j].EntityID
		},
	)

	return result
}

func appendEntityAlias(
	aliases []entityAlias,
	entity domain.Entity,
	aliasLanguage domain.Language,
	value string,
	weight float64,
) []entityAlias {
	normalized := language.Normalize(value)

	if normalized == "" {
		return aliases
	}

	for _, existing := range aliases {
		if existing.EntityID == entity.ID &&
			existing.Language == aliasLanguage &&
			existing.Normalized == normalized {
			return aliases
		}
	}

	return append(
		aliases,
		entityAlias{
			EntityID:   entity.ID,
			EntityType: entity.Type,
			Language:   aliasLanguage,
			Normalized: normalized,
			Weight:     weight,
		},
	)
}

func updateEntityMatch(
	values map[domain.EntityID]domain.EntityMatch,
	candidate domain.EntityMatch,
) {
	existing, found := values[candidate.EntityID]

	if found && existing.Score >= candidate.Score {
		return
	}

	values[candidate.EntityID] = candidate
}

func languageCompatible(
	aliasLanguage domain.Language,
	queryLanguage domain.Language,
) bool {
	if aliasLanguage == domain.LanguageUnknown {
		return true
	}

	if queryLanguage == domain.LanguageUnknown {
		return true
	}

	return true
}

func entityFuzzyThreshold(
	left string,
	right string,
) float64 {
	shortest := runeLength(left)

	if value := runeLength(right); value < shortest {
		shortest = value
	}

	switch {
	case shortest >= 9:
		return 0.78
	case shortest >= 7:
		return 0.80
	case shortest >= 5:
		return 0.84
	default:
		return 0.9
	}
}

func shouldRejectAmbiguousGo(
	original string,
	normalized string,
	alias entityAlias,
	targetLanguage domain.Language,
) bool {
	if alias.EntityID != knowledge.EntityGo {
		return false
	}

	if alias.Normalized != "go" {
		return false
	}

	if targetLanguage != domain.LanguageEnglish {
		return false
	}

	if containsCapitalizedGo(original) {
		return false
	}

	technicalMarkers := []string{
		"use",
		"uses",
		"used",
		"using",
		"know",
		"knows",
		"experience",
		"experienced",
		"language",
		"programming",
		"technology",
		"technologies",
		"stack",
		"backend",
		"project",
		"projects",
		"develop",
		"developed",
		"code",
		"coding",
	}

	for _, marker := range technicalMarkers {
		if containsPhrase(
			normalized,
			marker,
		) {
			return false
		}
	}

	return true
}

func containsCapitalizedGo(value string) bool {
	fields := strings.FieldsFunc(
		value,
		func(current rune) bool {
			return !unicode.IsLetter(current)
		},
	)

	for _, field := range fields {
		if field == "Go" {
			return true
		}
	}

	return false
}
