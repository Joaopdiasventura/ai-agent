package retrieval

import (
	"sort"

	"ai-agent/internal/domain"
)

type Source string

const (
	SourceLexical Source = "lexical"
	SourceEntity  Source = "entity"
	SourceConcept Source = "concept"
	SourceFuzzy   Source = "fuzzy"
)

type Candidate struct {
	FactID          domain.FactID
	Language        domain.Language
	Score           float64
	Source          Source
	MatchedTerms    []string
	MatchedEntities []domain.EntityID
	MatchedConcepts []domain.ConceptID
}

type Ranking struct {
	Source     Source
	Candidates []Candidate
}

func (r Ranking) Top(limit int) []Candidate {
	if limit <= 0 || len(r.Candidates) == 0 {
		return nil
	}

	if limit > len(r.Candidates) {
		limit = len(r.Candidates)
	}

	result := make(
		[]Candidate,
		limit,
	)

	copy(
		result,
		r.Candidates[:limit],
	)

	return result
}

type Result struct {
	Query    domain.Query
	Rankings []Ranking
}

func (r Result) Ranking(
	source Source,
) (Ranking, bool) {
	for _, ranking := range r.Rankings {
		if ranking.Source == source {
			return ranking, true
		}
	}

	return Ranking{}, false
}

func (r Result) Empty() bool {
	for _, ranking := range r.Rankings {
		if len(ranking.Candidates) > 0 {
			return false
		}
	}

	return true
}

func (r Result) FactIDs() []domain.FactID {
	seen := make(
		map[domain.FactID]struct{},
	)

	result := make(
		[]domain.FactID,
		0,
	)

	for _, ranking := range r.Rankings {
		for _, candidate := range ranking.Candidates {
			if _, exists := seen[candidate.FactID]; exists {
				continue
			}

			seen[candidate.FactID] = struct{}{}

			result = append(
				result,
				candidate.FactID,
			)
		}
	}

	sort.Slice(
		result,
		func(left int, right int) bool {
			return result[left] <
				result[right]
		},
	)

	return result
}

func sortCandidates(
	values []Candidate,
) {
	sort.SliceStable(
		values,
		func(left int, right int) bool {
			if values[left].Score !=
				values[right].Score {
				return values[left].Score >
					values[right].Score
			}

			return values[left].FactID <
				values[right].FactID
		},
	)
}

func limitCandidates(
	values []Candidate,
	limit int,
) []Candidate {
	if limit <= 0 || len(values) == 0 {
		return nil
	}

	sortCandidates(values)

	if len(values) > limit {
		values = values[:limit]
	}

	result := make(
		[]Candidate,
		len(values),
	)

	copy(result, values)

	return result
}

func appendUniqueString(
	values []string,
	value string,
) []string {
	if value == "" {
		return values
	}

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

func appendUniqueEntity(
	values []domain.EntityID,
	value domain.EntityID,
) []domain.EntityID {
	if value == "" {
		return values
	}

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

func appendUniqueConcept(
	values []domain.ConceptID,
	value domain.ConceptID,
) []domain.ConceptID {
	if value == "" {
		return values
	}

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

func mergeCandidateMaximum(
	values map[domain.FactID]Candidate,
	candidate Candidate,
) {
	existing, found :=
		values[candidate.FactID]

	if !found {
		values[candidate.FactID] =
			candidate

		return
	}

	if candidate.Score > existing.Score {
		existing.Score =
			candidate.Score

		existing.Language =
			candidate.Language
	}

	for _, term := range candidate.MatchedTerms {
		existing.MatchedTerms =
			appendUniqueString(
				existing.MatchedTerms,
				term,
			)
	}

	for _, entity := range candidate.MatchedEntities {
		existing.MatchedEntities =
			appendUniqueEntity(
				existing.MatchedEntities,
				entity,
			)
	}

	for _, concept := range candidate.MatchedConcepts {
		existing.MatchedConcepts =
			appendUniqueConcept(
				existing.MatchedConcepts,
				concept,
			)
	}

	values[candidate.FactID] =
		existing
}

func candidatesFromMap(
	values map[domain.FactID]Candidate,
	limit int,
) []Candidate {
	result := make(
		[]Candidate,
		0,
		len(values),
	)

	for _, candidate := range values {
		result = append(
			result,
			candidate,
		)
	}

	return limitCandidates(
		result,
		limit,
	)
}

func retrievalLanguages(
	value domain.Language,
) []domain.Language {
	if value == domain.LanguagePortuguese {
		return []domain.Language{
			domain.LanguagePortuguese,
		}
	}

	if value == domain.LanguageEnglish {
		return []domain.Language{
			domain.LanguageEnglish,
		}
	}

	return []domain.Language{
		domain.LanguagePortuguese,
		domain.LanguageEnglish,
	}
}
