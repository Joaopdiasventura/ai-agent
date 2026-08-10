package index

import (
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
)

func createIndex(
	t *testing.T,
) (*Index, *knowledge.Knowledge) {
	t.Helper()

	base, err := knowledge.New()

	if err != nil {
		t.Fatal(err)
	}

	currentOntology, err :=
		ontology.New()

	if err != nil {
		t.Fatal(err)
	}

	currentIndex, err :=
		New(
			base,
			currentOntology,
		)

	if err != nil {
		t.Fatal(err)
	}

	return currentIndex, base
}

func TestIndexBuildsSuccessfully(
	t *testing.T,
) {
	currentIndex, base :=
		createIndex(t)

	if currentIndex.DocumentCount(
		domain.LanguagePortuguese,
	) != len(base.Facts()) {
		t.Fatalf(
			"expected %d portuguese documents, got %d",
			len(base.Facts()),
			currentIndex.DocumentCount(
				domain.LanguagePortuguese,
			),
		)
	}

	if currentIndex.DocumentCount(
		domain.LanguageEnglish,
	) != len(base.Facts()) {
		t.Fatalf(
			"expected %d english documents, got %d",
			len(base.Facts()),
			currentIndex.DocumentCount(
				domain.LanguageEnglish,
			),
		)
	}
}

func TestGGCompressAppearsInPortugueseIndex(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	postings :=
		currentIndex.Postings(
			domain.LanguagePortuguese,
			"ggcompress",
		)

	if len(postings) == 0 {
		t.Fatal(
			"expected ggcompress postings",
		)
	}
}

func TestGGCompressAppearsInEnglishIndex(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	postings :=
		currentIndex.Postings(
			domain.LanguageEnglish,
			"ggcompress",
		)

	if len(postings) == 0 {
		t.Fatal(
			"expected ggcompress postings",
		)
	}
}

func TestConcurrencyPortugueseTerm(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	postings :=
		currentIndex.Postings(
			domain.LanguagePortuguese,
			"concorrencia",
		)

	if len(postings) == 0 {
		t.Fatal(
			"expected concurrency postings in portuguese",
		)
	}

	found := false

	for _, posting := range postings {
		if posting.FactID ==
			domain.FactID(
				"project-ggcompress-concurrency",
			) {
			found = true
			break
		}
	}

	if !found {
		t.Fatal(
			"expected ggcompress concurrency fact",
		)
	}
}

func TestConcurrencyEnglishTerm(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	postings :=
		currentIndex.Postings(
			domain.LanguageEnglish,
			"concurrency",
		)

	if len(postings) == 0 {
		t.Fatal(
			"expected concurrency postings in english",
		)
	}
}

func TestEntityIndexFindsGGCompress(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsByEntity(
			knowledge.EntityGGCompress,
		)

	if len(facts) == 0 {
		t.Fatal(
			"expected ggcompress facts",
		)
	}

	if !containsFactID(
		facts,
		"project-ggcompress-throughput",
	) {
		t.Fatal(
			"expected ggcompress throughput fact",
		)
	}
}

func TestEntityIndexIncludesObjectEntities(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsByEntity(
			knowledge.EntityKafka,
		)

	if !containsFactID(
		facts,
		"project-xtube-kafka",
	) {
		t.Fatal(
			"expected kafka x-tube fact",
		)
	}
}

func TestEntityIndexIncludesContextEntities(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsByEntity(
			knowledge.EntityXTube,
		)

	if !containsFactID(
		facts,
		"project-xtube-leadership",
	) {
		t.Fatal(
			"expected x-tube leadership fact",
		)
	}
}

func TestConceptIndexConcurrency(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsByConcept(
			ontology.ConceptConcurrency,
		)

	if len(facts) == 0 {
		t.Fatal(
			"expected concurrency facts",
		)
	}

	if !containsFactID(
		facts,
		"project-ggcompress-concurrency",
	) {
		t.Fatal(
			"expected ggcompress concurrency fact",
		)
	}
}

func TestCategoryIndex(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsByCategory(
			domain.FactCategoryProject,
		)

	if len(facts) == 0 {
		t.Fatal(
			"expected project facts",
		)
	}
}

func TestContextIndexUFind(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsByContext(
			knowledge.EntityUFind,
		)

	if len(facts) == 0 {
		t.Fatal(
			"expected ufind context facts",
		)
	}

	if !containsFactID(
		facts,
		"experience-current-financial-volume",
	) {
		t.Fatal(
			"expected financial volume fact",
		)
	}
}

func TestStatistics(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	stats :=
		currentIndex.Statistics(
			domain.LanguagePortuguese,
		)

	if stats.DocumentCount <= 0 {
		t.Fatal(
			"expected documents",
		)
	}

	if stats.AverageDocumentLength <= 0 {
		t.Fatal(
			"expected positive average document length",
		)
	}

	if stats.AverageFieldLength[FieldStatement] <= 0 {
		t.Fatal(
			"expected positive statement average length",
		)
	}
}

func TestDocumentFrequency(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	frequency :=
		currentIndex.DocumentFrequency(
			domain.LanguagePortuguese,
			"go",
		)

	if frequency <= 0 {
		t.Fatal(
			"expected go document frequency",
		)
	}

	if frequency >
		currentIndex.DocumentCount(
			domain.LanguagePortuguese,
		) {
		t.Fatal(
			"document frequency cannot exceed document count",
		)
	}
}

func TestDocumentContainsFields(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	document, found :=
		currentIndex.Document(
			domain.LanguagePortuguese,
			domain.FactID(
				"project-ggcompress-throughput",
			),
		)

	if !found {
		t.Fatal(
			"expected document",
		)
	}

	if document.Length <= 0 {
		t.Fatal(
			"expected positive document length",
		)
	}

	if document.TermFrequency(
		FieldSubject,
		"ggcompress",
	) <= 0 {
		t.Fatal(
			"expected ggcompress in subject field",
		)
	}

	if document.TermFrequency(
		FieldConcept,
		"performance",
	) <= 0 {
		t.Fatal(
			"expected performance in concept field",
		)
	}
}

func TestNumericFactIsIndexed(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	document, found :=
		currentIndex.Document(
			domain.LanguageEnglish,
			domain.FactID(
				"project-ggcompress-throughput",
			),
		)

	if !found {
		t.Fatal(
			"expected document",
		)
	}

	if document.TermFrequency(
		FieldObject,
		"1.23",
	) <= 0 {
		t.Fatal(
			"expected numeric value",
		)
	}

	if document.TermFrequency(
		FieldObject,
		"gb",
	) <= 0 {
		t.Fatal(
			"expected unit gb",
		)
	}
}

func TestVocabularyContainsKubernetes(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	vocabulary :=
		currentIndex.Vocabulary(
			domain.LanguagePortuguese,
		)

	if !containsString(
		vocabulary,
		"kubernetes",
	) {
		t.Fatal(
			"expected kubernetes in vocabulary",
		)
	}
}

func TestNGramIndexKubernetes(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	values :=
		currentIndex.TermsForNGram(
			domain.LanguagePortuguese,
			"kub",
		)

	if !containsString(
		values,
		"kubernetes",
	) {
		t.Fatalf(
			"expected kubernetes, got %v",
			values,
		)
	}
}

func TestFuzzyCandidateKubernets(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	values :=
		currentIndex.FuzzyTermCandidates(
			domain.LanguagePortuguese,
			"kubernets",
		)

	if !containsString(
		values,
		"kubernetes",
	) {
		t.Fatalf(
			"expected kubernetes candidate, got %v",
			values,
		)
	}
}

func TestFuzzyCandidateConcurrencyTypo(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	values :=
		currentIndex.FuzzyTermCandidates(
			domain.LanguagePortuguese,
			"concorencia",
		)

	if !containsString(
		values,
		"concorrencia",
	) {
		t.Fatalf(
			"expected concorrencia candidate, got %v",
			values,
		)
	}
}

func TestFactsBySubject(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	facts :=
		currentIndex.FactsBySubject(
			knowledge.EntityGGCompress,
		)

	if len(facts) == 0 {
		t.Fatal(
			"expected ggcompress subject facts",
		)
	}

	for _, factID := range facts {
		_, found :=
			currentIndex.Document(
				domain.LanguagePortuguese,
				factID,
			)

		if !found {
			t.Fatalf(
				"expected document for %s",
				factID,
			)
		}
	}
}

func TestPostingsHaveFrequency(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	postings :=
		currentIndex.Postings(
			domain.LanguagePortuguese,
			"ggcompress",
		)

	for _, posting := range postings {
		if posting.Frequency <= 0 {
			t.Fatalf(
				"expected positive frequency for %s",
				posting.FactID,
			)
		}
	}
}

func TestUnknownTerm(
	t *testing.T,
) {
	currentIndex, _ :=
		createIndex(t)

	postings :=
		currentIndex.Postings(
			domain.LanguageEnglish,
			"doesnotexistxyz",
		)

	if len(postings) != 0 {
		t.Fatalf(
			"expected no postings, got %v",
			postings,
		)
	}

	frequency :=
		currentIndex.DocumentFrequency(
			domain.LanguageEnglish,
			"doesnotexistxyz",
		)

	if frequency != 0 {
		t.Fatalf(
			"expected zero document frequency, got %d",
			frequency,
		)
	}
}

func containsFactID(
	values []domain.FactID,
	expected domain.FactID,
) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}

	return false
}

func containsString(
	values []string,
	expected string,
) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}

	return false
}
