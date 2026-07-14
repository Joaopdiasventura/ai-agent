package tfidf

func CalculateTFIDF(
	tokens []string,
	idf map[string]float64,
) map[string]float64 {
	vector := make(map[string]float64)
	tf := CalculateTF(tokens)

	for term, frequency := range tf {
		idfValue, exists := idf[term]

		if !exists {
			continue
		}

		vector[term] = frequency * idfValue
	}

	return vector
}
