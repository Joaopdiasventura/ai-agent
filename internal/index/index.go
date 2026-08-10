package index

import (
	"errors"
	"sort"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/language"
	"ai-agent/internal/ontology"
)

var indexedLanguages = []domain.Language{
	domain.LanguagePortuguese,
	domain.LanguageEnglish,
}

type Index struct {
	documents map[domain.Language]map[domain.FactID]Document

	inverted map[domain.Language]map[string][]Posting

	documentFrequency map[domain.Language]map[string]int

	documentLengthTotals map[domain.Language]int

	fieldLengthTotals map[domain.Language]map[Field]int

	statistics map[domain.Language]Statistics

	vocabulary map[domain.Language][]string

	ngramTerms map[domain.Language]map[string][]string

	subjectFacts map[domain.EntityID][]domain.FactID

	entityFacts map[domain.EntityID][]domain.FactID

	conceptFacts map[domain.ConceptID][]domain.FactID

	categoryFacts map[domain.FactCategory][]domain.FactID

	contextFacts map[domain.EntityID][]domain.FactID
}

func New(
	base *knowledge.Knowledge,
	currentOntology *ontology.Ontology,
) (*Index, error) {
	if base == nil {
		return nil, errors.New(
			"knowledge base is required",
		)
	}

	if currentOntology == nil {
		return nil, errors.New(
			"ontology is required",
		)
	}

	if err := currentOntology.ValidateFacts(
		base.Facts(),
	); err != nil {
		return nil, err
	}

	result := createEmptyIndex()

	for _, fact := range base.Facts() {
		result.indexFactStructure(
			fact,
		)

		for _, targetLanguage :=
			range indexedLanguages {
			document, err :=
				buildDocument(
					fact,
					targetLanguage,
					base,
					currentOntology,
				)

			if err != nil {
				return nil, err
			}

			result.indexDocument(
				document,
			)
		}
	}

	sortStructuralIndexes(
		result,
	)

	finalizeLexicalIndex(
		result,
	)

	buildNGramIndexes(
		result,
	)

	return result, nil
}

func createEmptyIndex() *Index {
	result := &Index{
		documents: make(
			map[domain.Language]map[domain.FactID]Document,
		),
		inverted: make(
			map[domain.Language]map[string][]Posting,
		),
		documentFrequency: make(
			map[domain.Language]map[string]int,
		),
		documentLengthTotals: make(
			map[domain.Language]int,
		),
		fieldLengthTotals: make(
			map[domain.Language]map[Field]int,
		),
		statistics: make(
			map[domain.Language]Statistics,
		),
		vocabulary: make(
			map[domain.Language][]string,
		),
		ngramTerms: make(
			map[domain.Language]map[string][]string,
		),
		subjectFacts: make(
			map[domain.EntityID][]domain.FactID,
		),
		entityFacts: make(
			map[domain.EntityID][]domain.FactID,
		),
		conceptFacts: make(
			map[domain.ConceptID][]domain.FactID,
		),
		categoryFacts: make(
			map[domain.FactCategory][]domain.FactID,
		),
		contextFacts: make(
			map[domain.EntityID][]domain.FactID,
		),
	}

	for _, targetLanguage :=
		range indexedLanguages {
		result.documents[
			targetLanguage,
		] = make(
			map[domain.FactID]Document,
		)

		result.inverted[
			targetLanguage,
		] = make(
			map[string][]Posting,
		)

		result.documentFrequency[
			targetLanguage,
		] = make(
			map[string]int,
		)

		result.fieldLengthTotals[
			targetLanguage,
		] = make(
			map[Field]int,
		)

		result.ngramTerms[
			targetLanguage,
		] = make(
			map[string][]string,
		)
	}

	return result
}

func (i *Index) Document(
	targetLanguage domain.Language,
	factID domain.FactID,
) (Document, bool) {
	documents, exists :=
		i.documents[targetLanguage]

	if !exists {
		return Document{}, false
	}

	document, found :=
		documents[factID]

	return document, found
}

func (i *Index) Documents(
	targetLanguage domain.Language,
) []Document {
	documents :=
		i.documents[targetLanguage]

	result := make(
		[]Document,
		0,
		len(documents),
	)

	for _, document := range documents {
		result = append(
			result,
			document,
		)
	}

	sort.Slice(
		result,
		func(left int, right int) bool {
			return result[left].FactID <
				result[right].FactID
		},
	)

	return result
}

func (i *Index) DocumentCount(
	targetLanguage domain.Language,
) int {
	return len(
		i.documents[targetLanguage],
	)
}

func (i *Index) Postings(
	targetLanguage domain.Language,
	term string,
) []Posting {
	normalizedTokens :=
		language.Tokenize(term)

	if len(normalizedTokens) != 1 {
		return nil
	}

	return copyPostings(
		i.inverted[
			targetLanguage,
		][normalizedTokens[0]],
	)
}

func (i *Index) DocumentFrequency(
	targetLanguage domain.Language,
	term string,
) int {
	normalizedTokens :=
		language.Tokenize(term)

	if len(normalizedTokens) != 1 {
		return 0
	}

	return i.documentFrequency[
		targetLanguage,
	][normalizedTokens[0]]
}

func (i *Index) Statistics(
	targetLanguage domain.Language,
) Statistics {
	stats :=
		i.statistics[targetLanguage]

	fieldLengths := make(
		map[Field]float64,
		len(
			stats.AverageFieldLength,
		),
	)

	for field, value :=
		range stats.AverageFieldLength {
		fieldLengths[field] = value
	}

	return Statistics{
		DocumentCount:
			stats.DocumentCount,
		AverageDocumentLength:
			stats.AverageDocumentLength,
		AverageFieldLength:
			fieldLengths,
	}
}

func (i *Index) AverageFieldLength(
	targetLanguage domain.Language,
	field Field,
) float64 {
	return i.statistics[
		targetLanguage,
	].AverageFieldLength[field]
}

func (i *Index) Vocabulary(
	targetLanguage domain.Language,
) []string {
	values :=
		i.vocabulary[targetLanguage]

	result := make(
		[]string,
		len(values),
	)

	copy(result, values)

	return result
}

func (i *Index) TermsForNGram(
	targetLanguage domain.Language,
	gram string,
) []string {
	normalized :=
		language.Normalize(gram)

	values :=
		i.ngramTerms[
			targetLanguage,
		][normalized]

	result := make(
		[]string,
		len(values),
	)

	copy(result, values)

	return result
}

func (i *Index) FuzzyTermCandidates(
	targetLanguage domain.Language,
	term string,
) []string {
	grams :=
		language.CharacterNGrams(
			term,
			3,
			4,
		)

	return uniqueCandidateTerms(
		i,
		targetLanguage,
		grams,
	)
}

func (i *Index) FactsBySubject(
	entityID domain.EntityID,
) []domain.FactID {
	return copyFactIDs(
		i.subjectFacts[entityID],
	)
}

func (i *Index) FactsByEntity(
	entityID domain.EntityID,
) []domain.FactID {
	return copyFactIDs(
		i.entityFacts[entityID],
	)
}

func (i *Index) FactsByConcept(
	conceptID domain.ConceptID,
) []domain.FactID {
	return copyFactIDs(
		i.conceptFacts[conceptID],
	)
}

func (i *Index) FactsByCategory(
	category domain.FactCategory,
) []domain.FactID {
	return copyFactIDs(
		i.categoryFacts[category],
	)
}

func (i *Index) FactsByContext(
	entityID domain.EntityID,
) []domain.FactID {
	return copyFactIDs(
		i.contextFacts[entityID],
	)
}