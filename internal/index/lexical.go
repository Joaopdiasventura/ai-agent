package index

import (
	"sort"

	"ai-agent/internal/domain"
)

func (i *Index) indexDocument(
	document Document,
) {
	languageDocuments :=
		i.documents[document.Language]

	languageDocuments[document.FactID] =
		document

	seenDocumentTerms :=
		make(map[string]struct{})

	for fieldName, field := range document.Fields {
		i.fieldLengthTotals[document.Language][fieldName] += field.Length

		for term, frequency := range field.Terms {
			posting := Posting{
				FactID:    document.FactID,
				Field:     fieldName,
				Frequency: frequency,
			}

			i.inverted[document.Language][term] = append(
				i.inverted[document.Language][term],
				posting,
			)

			seenDocumentTerms[term] =
				struct{}{}
		}
	}

	for term := range seenDocumentTerms {
		i.documentFrequency[document.Language][term]++
	}

	i.documentLengthTotals[document.Language] += document.Length
}

func finalizeLexicalIndex(
	currentIndex *Index,
) {
	for _, targetLanguage := range indexedLanguages {
		documents :=
			currentIndex.documents[targetLanguage]

		documentCount :=
			len(documents)

		if documentCount > 0 {
			currentIndex.statistics[targetLanguage] = Statistics{
				DocumentCount: documentCount,
				AverageDocumentLength: float64(
					currentIndex.documentLengthTotals[targetLanguage],
				) /
					float64(documentCount),
				AverageFieldLength: make(
					map[Field]float64,
				),
			}
		} else {
			currentIndex.statistics[targetLanguage] = Statistics{
				AverageFieldLength: make(
					map[Field]float64,
				),
			}
		}

		stats :=
			currentIndex.statistics[targetLanguage]

		for _, field := range AllFields {
			if documentCount == 0 {
				stats.AverageFieldLength[field] = 0
				continue
			}

			stats.AverageFieldLength[field] =
				float64(
					currentIndex.fieldLengthTotals[targetLanguage][field],
				) /
					float64(documentCount)
		}

		currentIndex.statistics[targetLanguage] = stats

		for term, postings := range currentIndex.inverted[targetLanguage] {
			sort.Slice(
				postings,
				func(left int, right int) bool {
					if postings[left].FactID !=
						postings[right].FactID {
						return postings[left].FactID <
							postings[right].FactID
					}

					return postings[left].Field <
						postings[right].Field
				},
			)

			currentIndex.inverted[targetLanguage][term] = postings
		}
	}
}

func copyPostings(
	values []Posting,
) []Posting {
	result := make(
		[]Posting,
		len(values),
	)

	copy(result, values)

	return result
}

func copyFactIDs(
	values []domain.FactID,
) []domain.FactID {
	result := make(
		[]domain.FactID,
		len(values),
	)

	copy(result, values)

	return result
}
