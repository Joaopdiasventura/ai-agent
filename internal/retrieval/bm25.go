package retrieval

import (
	"math"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
)

type BM25FConfig struct {
	K1           float64
	FieldWeights map[searchindex.Field]float64
	FieldB       map[searchindex.Field]float64
}

func DefaultBM25FConfig() BM25FConfig {
	return BM25FConfig{
		K1: 1.2,
		FieldWeights: map[searchindex.Field]float64{
			searchindex.FieldStatement: 1,
			searchindex.FieldSubject:   6,
			searchindex.FieldObject:    4,
			searchindex.FieldConcept:   4.5,
			searchindex.FieldContext:   3.5,
			searchindex.FieldPredicate: 1.5,
			searchindex.FieldCategory:  0.75,
		},
		FieldB: map[searchindex.Field]float64{
			searchindex.FieldStatement: 0.75,
			searchindex.FieldSubject:   0.2,
			searchindex.FieldObject:    0.35,
			searchindex.FieldConcept:   0.25,
			searchindex.FieldContext:   0.3,
			searchindex.FieldPredicate: 0.1,
			searchindex.FieldCategory:  0.1,
		},
	}
}

type WeightedTerm struct {
	Value  string
	Weight float64
}

type BM25F struct {
	index  *searchindex.Index
	config BM25FConfig
}

func NewBM25F(
	currentIndex *searchindex.Index,
	config BM25FConfig,
) *BM25F {
	return &BM25F{
		index:  currentIndex,
		config: config,
	}
}

func (b *BM25F) Search(
	targetLanguage domain.Language,
	terms []WeightedTerm,
	limit int,
) []Candidate {
	if len(terms) == 0 {
		return nil
	}

	documentCount :=
		b.index.DocumentCount(
			targetLanguage,
		)

	if documentCount == 0 {
		return nil
	}

	type candidateState struct {
		score        float64
		matchedTerms []string
	}

	states := make(
		map[domain.FactID]*candidateState,
	)

	for _, term := range terms {
		if term.Value == "" ||
			term.Weight <= 0 {
			continue
		}

		postings :=
			b.index.Postings(
				targetLanguage,
				term.Value,
			)

		if len(postings) == 0 {
			continue
		}

		documentFrequency :=
			b.index.DocumentFrequency(
				targetLanguage,
				term.Value,
			)

		idf :=
			inverseDocumentFrequency(
				documentCount,
				documentFrequency,
			)

		fieldFrequencies := make(
			map[domain.FactID]map[searchindex.Field]int,
		)

		for _, posting := range postings {
			fields, exists :=
				fieldFrequencies[posting.FactID]

			if !exists {
				fields = make(
					map[searchindex.Field]int,
				)

				fieldFrequencies[posting.FactID] = fields
			}

			fields[posting.Field] +=
				posting.Frequency
		}

		for factID, fields := range fieldFrequencies {
			document, found :=
				b.index.Document(
					targetLanguage,
					factID,
				)

			if !found {
				continue
			}

			weightedFrequency :=
				b.weightedFrequency(
					document,
					targetLanguage,
					fields,
				)

			if weightedFrequency <= 0 {
				continue
			}

			score :=
				term.Weight *
					idf *
					(weightedFrequency * (b.config.K1 + 1)) /
					(weightedFrequency + b.config.K1)

			state, exists :=
				states[factID]

			if !exists {
				state =
					&candidateState{}

				states[factID] =
					state
			}

			state.score += score

			state.matchedTerms =
				appendUniqueString(
					state.matchedTerms,
					term.Value,
				)
		}
	}

	candidates := make(
		[]Candidate,
		0,
		len(states),
	)

	for factID, state := range states {
		candidates = append(
			candidates,
			Candidate{
				FactID:       factID,
				Language:     targetLanguage,
				Score:        state.score,
				Source:       SourceLexical,
				MatchedTerms: state.matchedTerms,
			},
		)
	}

	return limitCandidates(
		candidates,
		limit,
	)
}

func (b *BM25F) weightedFrequency(
	document searchindex.Document,
	targetLanguage domain.Language,
	fields map[searchindex.Field]int,
) float64 {
	total := 0.0

	for field, frequency := range fields {
		if frequency <= 0 {
			continue
		}

		weight :=
			b.fieldWeight(field)

		if weight <= 0 {
			continue
		}

		fieldData, found :=
			document.Field(field)

		if !found {
			continue
		}

		averageLength :=
			b.index.AverageFieldLength(
				targetLanguage,
				field,
			)

		normalization := 1.0

		if averageLength > 0 {
			bValue :=
				b.fieldB(field)

			normalization =
				(1 - bValue) +
					bValue*
						float64(fieldData.Length)/
						averageLength
		}

		if normalization <= 0 {
			normalization = 1
		}

		total +=
			weight *
				float64(frequency) /
				normalization
	}

	return total
}

func (b *BM25F) fieldWeight(
	field searchindex.Field,
) float64 {
	value, found :=
		b.config.FieldWeights[field]

	if !found {
		return 1
	}

	return value
}

func (b *BM25F) fieldB(
	field searchindex.Field,
) float64 {
	value, found :=
		b.config.FieldB[field]

	if !found {
		return 0.75
	}

	if value < 0 {
		return 0
	}

	if value > 1 {
		return 1
	}

	return value
}

func inverseDocumentFrequency(
	documentCount int,
	documentFrequency int,
) float64 {
	if documentCount <= 0 ||
		documentFrequency <= 0 {
		return 0
	}

	return math.Log(
		1 +
			(float64(documentCount)-float64(documentFrequency)+0.5)/
				(float64(documentFrequency)+0.5),
	)
}
