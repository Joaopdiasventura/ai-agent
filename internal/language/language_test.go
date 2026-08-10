package language

import (
	"math"
	"reflect"
	"testing"

	"ai-agent/internal/domain"
)

func TestNormalizePortuguese(t *testing.T) {
	got := Normalize(
		"  João já trabalhou com CONCORRÊNCIA?  ",
	)

	expected := "joao ja trabalhou com concorrencia"

	if got != expected {
		t.Fatalf(
			"expected %q, got %q",
			expected,
			got,
		)
	}
}

func TestNormalizeEnglish(t *testing.T) {
	got := Normalize(
		"Which PROJECT demonstrates concurrency?",
	)

	expected := "which project demonstrates concurrency"

	if got != expected {
		t.Fatalf(
			"expected %q, got %q",
			expected,
			got,
		)
	}
}

func TestNormalizeDecimalComma(t *testing.T) {
	got := Normalize(
		"Throughput de 1,23 GB/s.",
	)

	expected := "throughput de 1.23 gb s"

	if got != expected {
		t.Fatalf(
			"expected %q, got %q",
			expected,
			got,
		)
	}
}

func TestNormalizeDecimalPoint(t *testing.T) {
	got := Normalize(
		"1.23 GB/s",
	)

	expected := "1.23 gb s"

	if got != expected {
		t.Fatalf(
			"expected %q, got %q",
			expected,
			got,
		)
	}
}

func TestNormalizeTechnologyNames(t *testing.T) {
	got := Normalize(
		"Node.js, PostgreSQL, SHA-256 e X-Tube",
	)

	expected := "node js postgresql sha 256 e x tube"

	if got != expected {
		t.Fatalf(
			"expected %q, got %q",
			expected,
			got,
		)
	}
}

func TestTokenize(t *testing.T) {
	got := Tokenize(
		"Qual projeto usa Go?",
	)

	expected := []string{
		"qual",
		"projeto",
		"usa",
		"go",
	}

	if !reflect.DeepEqual(
		got,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			got,
		)
	}
}

func TestPortugueseContentTerms(t *testing.T) {
	tokens := Tokenize(
		"Qual projeto do João usa Go?",
	)

	got := ContentTerms(
		tokens,
		domain.LanguagePortuguese,
	)

	expected := []string{
		"projeto",
		"joao",
		"usa",
		"go",
	}

	if !reflect.DeepEqual(
		got,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			got,
		)
	}
}

func TestEnglishContentTerms(t *testing.T) {
	tokens := Tokenize(
		"Which project uses Go?",
	)

	got := ContentTerms(
		tokens,
		domain.LanguageEnglish,
	)

	expected := []string{
		"project",
		"uses",
		"go",
	}

	if !reflect.DeepEqual(
		got,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			got,
		)
	}
}

func TestGoIsNotEnglishStopWord(t *testing.T) {
	if IsStopWord(
		domain.LanguageEnglish,
		"go",
	) {
		t.Fatal(
			"go must not be treated as an english stop word",
		)
	}
}

func TestUniqueTokens(t *testing.T) {
	got := UniqueTokens(
		[]string{
			"go",
			"java",
			"go",
			"postgresql",
			"java",
		},
	)

	expected := []string{
		"go",
		"java",
		"postgresql",
	}

	if !reflect.DeepEqual(
		got,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			got,
		)
	}
}

func TestDetectPortuguese(t *testing.T) {
	detection := DetectLanguageDetailed(
		"Qual projeto demonstra melhor experiência com concorrência?",
	)

	if detection.Language != domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			detection.Language,
		)
	}

	if detection.Portuguese <= detection.English {
		t.Fatal(
			"expected portuguese score to be greater than english score",
		)
	}

	if detection.Confidence <= 0 {
		t.Fatal(
			"expected positive confidence",
		)
	}
}

func TestDetectEnglish(t *testing.T) {
	detection := DetectLanguageDetailed(
		"Which project best demonstrates experience with concurrency?",
	)

	if detection.Language != domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			detection.Language,
		)
	}

	if detection.English <= detection.Portuguese {
		t.Fatal(
			"expected english score to be greater than portuguese score",
		)
	}

	if detection.Confidence <= 0 {
		t.Fatal(
			"expected positive confidence",
		)
	}
}

func TestDetectPortugueseWithDiacritics(t *testing.T) {
	language := DetectLanguage(
		"experiência técnica",
	)

	if language != domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			language,
		)
	}
}

func TestUnknownLanguageForEntityOnly(t *testing.T) {
	language := DetectLanguage(
		"GGCompress",
	)

	if language != domain.LanguageUnknown {
		t.Fatalf(
			"expected unknown, got %s",
			language,
		)
	}
}

func TestCharacterNGrams(t *testing.T) {
	grams := CharacterNGrams(
		"go",
		3,
		4,
	)

	expected := []string{
		"go",
	}

	if !reflect.DeepEqual(
		grams,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			grams,
		)
	}
}

func TestCharacterNGramSimilarityExact(t *testing.T) {
	score := CharacterNGramSimilarity(
		"kubernetes",
		"kubernetes",
		3,
		4,
	)

	if math.Abs(score-1) > 0.000001 {
		t.Fatalf(
			"expected 1, got %f",
			score,
		)
	}
}

func TestCharacterNGramSimilarityTypo(t *testing.T) {
	score := CharacterNGramSimilarity(
		"kubernets",
		"kubernetes",
		3,
		4,
	)

	if score <= 0.4 {
		t.Fatalf(
			"expected typo similarity above 0.4, got %f",
			score,
		)
	}
}

func TestCharacterNGramSimilarityUnrelated(t *testing.T) {
	score := CharacterNGramSimilarity(
		"kubernetes",
		"postgresql",
		3,
		4,
	)

	if score >= 0.3 {
		t.Fatalf(
			"expected low similarity, got %f",
			score,
		)
	}
}

func TestProcessorPortuguese(t *testing.T) {
	processor := NewProcessor()

	result := processor.Analyze(
		"Qual projeto do João demonstra concorrência?",
	)

	if result.Language != domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			result.Language,
		)
	}

	expectedNormalized :=
		"qual projeto do joao demonstra concorrencia"

	if result.Normalized != expectedNormalized {
		t.Fatalf(
			"expected %q, got %q",
			expectedNormalized,
			result.Normalized,
		)
	}

	expectedTerms := []string{
		"projeto",
		"joao",
		"demonstra",
		"concorrencia",
	}

	if !reflect.DeepEqual(
		result.Terms,
		expectedTerms,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expectedTerms,
			result.Terms,
		)
	}
}

func TestProcessorEnglish(t *testing.T) {
	processor := NewProcessor()

	result := processor.Analyze(
		"Which project demonstrates concurrency?",
	)

	if result.Language != domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			result.Language,
		)
	}

	expectedTerms := []string{
		"project",
		"demonstrates",
		"concurrency",
	}

	if !reflect.DeepEqual(
		result.Terms,
		expectedTerms,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expectedTerms,
			result.Terms,
		)
	}
}

func TestProcessorUnknownLanguage(t *testing.T) {
	processor := NewProcessor()

	result := processor.Analyze(
		"GGCompress",
	)

	if result.Language != domain.LanguageUnknown {
		t.Fatalf(
			"expected unknown, got %s",
			result.Language,
		)
	}

	expected := []string{
		"ggcompress",
	}

	if !reflect.DeepEqual(
		result.Terms,
		expected,
	) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			result.Terms,
		)
	}
}

func TestEmptyInput(t *testing.T) {
	processor := NewProcessor()

	result := processor.Analyze("")

	if result.Normalized != "" {
		t.Fatalf(
			"expected empty normalized string, got %q",
			result.Normalized,
		)
	}

	if result.Language != domain.LanguageUnknown {
		t.Fatalf(
			"expected unknown language, got %s",
			result.Language,
		)
	}

	if len(result.Tokens) != 0 {
		t.Fatalf(
			"expected no tokens, got %v",
			result.Tokens,
		)
	}

	if len(result.Terms) != 0 {
		t.Fatalf(
			"expected no terms, got %v",
			result.Terms,
		)
	}
}

func TestEditSimilarityKubernetesTypo(
	t *testing.T,
) {
	score := EditSimilarity(
		"kubernets",
		"kubernetes",
	)

	if score < 0.85 {
		t.Fatalf(
			"expected high similarity, got %f",
			score,
		)
	}
}

func TestEditSimilarityConcurrencyTypo(
	t *testing.T,
) {
	score := EditSimilarity(
		"concorencia",
		"concorrencia",
	)

	if score < 0.85 {
		t.Fatalf(
			"expected high similarity, got %f",
			score,
		)
	}
}

func TestFuzzySimilarityPostgresTypo(
	t *testing.T,
) {
	score := FuzzySimilarity(
		"postgress",
		"postgresql",
	)

	if score < 0.75 {
		t.Fatalf(
			"expected high similarity, got %f",
			score,
		)
	}
}

func TestFuzzySimilarityRejectsUnrelatedWords(
	t *testing.T,
) {
	score := FuzzySimilarity(
		"kubernetes",
		"rabbitmq",
	)

	if score >= 0.5 {
		t.Fatalf(
			"expected low similarity, got %f",
			score,
		)
	}
}
