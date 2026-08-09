package ontology

import (
	"math"
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

func TestOntologyBuildsSuccessfully(t *testing.T) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if len(currentOntology.Concepts()) == 0 {
		t.Fatal("expected concepts")
	}
}

func TestOntologyCoversKnowledgeBase(t *testing.T) {
	base, err := knowledge.New()

	if err != nil {
		t.Fatal(err)
	}

	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if err := currentOntology.ValidateFacts(
		base.Facts(),
	); err != nil {
		t.Fatal(err)
	}
}

func TestResolvePortugueseConcurrencyAlias(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	matches := currentOntology.ResolveAlias(
		"paralelismo",
		domain.LanguagePortuguese,
	)

	if !containsConcept(
		matches,
		ConceptConcurrency,
	) {
		t.Fatalf(
			"expected concurrency, got %v",
			matches,
		)
	}
}

func TestResolveEnglishConcurrencyAlias(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	matches := currentOntology.ResolveAlias(
		"parallel processing",
		domain.LanguageEnglish,
	)

	if !containsConcept(
		matches,
		ConceptConcurrency,
	) {
		t.Fatalf(
			"expected concurrency, got %v",
			matches,
		)
	}
}

func TestResolveCrossLanguageAlias(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	matches := currentOntology.ResolveAlias(
		"concurrency",
		domain.LanguagePortuguese,
	)

	if !containsConcept(
		matches,
		ConceptConcurrency,
	) {
		t.Fatalf(
			"expected concurrency, got %v",
			matches,
		)
	}
}

func TestResolveAliasNormalizesAccents(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	matches := currentOntology.ResolveAlias(
		"CONCORRÊNCIA",
		domain.LanguagePortuguese,
	)

	if !containsConcept(
		matches,
		ConceptConcurrency,
	) {
		t.Fatalf(
			"expected concurrency, got %v",
			matches,
		)
	}
}

func TestGoroutinesExpandToConcurrency(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	expanded := currentOntology.Expand(
		ConceptGoroutines,
		1,
	)

	weight, found := conceptWeight(
		expanded,
		ConceptConcurrency,
	)

	if !found {
		t.Fatal(
			"expected concurrency expansion",
		)
	}

	if math.Abs(weight-0.85) > 0.000001 {
		t.Fatalf(
			"expected 0.85, got %f",
			weight,
		)
	}
}

func TestMessagingExpandsToEventDriven(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	expanded := currentOntology.Expand(
		ConceptMessaging,
		1,
	)

	weight, found := conceptWeight(
		expanded,
		ConceptEventDriven,
	)

	if !found {
		t.Fatal(
			"expected event-driven expansion",
		)
	}

	if math.Abs(weight-0.8) > 0.000001 {
		t.Fatalf(
			"expected 0.8, got %f",
			weight,
		)
	}
}

func TestKafkaBinding(t *testing.T) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	concepts := currentOntology.EntityConcepts(
		knowledge.EntityKafka,
	)

	weight, found := conceptWeight(
		concepts,
		ConceptEventDriven,
	)

	if !found {
		t.Fatal(
			"expected kafka to bind event-driven",
		)
	}

	if weight != 1 {
		t.Fatalf(
			"expected weight 1, got %f",
			weight,
		)
	}
}

func TestKubernetesBinding(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	concepts := currentOntology.EntityConcepts(
		knowledge.EntityKubernetes,
	)

	if _, found := conceptWeight(
		concepts,
		ConceptOrchestration,
	); !found {
		t.Fatal(
			"expected orchestration",
		)
	}

	if _, found := conceptWeight(
		concepts,
		ConceptDistributedSystems,
	); !found {
		t.Fatal(
			"expected distributed systems",
		)
	}
}

func TestUnknownAlias(t *testing.T) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	matches := currentOntology.ResolveAlias(
		"quantum computing",
		domain.LanguageEnglish,
	)

	if len(matches) != 0 {
		t.Fatalf(
			"expected no matches, got %v",
			matches,
		)
	}
}

func TestAliasesAreLongestFirst(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	aliases := currentOntology.Aliases(
		domain.LanguagePortuguese,
	)

	for i := 1; i < len(aliases); i++ {
		currentLength :=
			len([]rune(aliases[i].Normalized))

		previousLength :=
			len([]rune(aliases[i-1].Normalized))

		if currentLength > previousLength {
			t.Fatal(
				"expected aliases ordered by descending length",
			)
		}
	}
}

func TestValidateFactsRejectsUnknownConcept(
	t *testing.T,
) {
	currentOntology, err := New()

	if err != nil {
		t.Fatal(err)
	}

	facts := []domain.Fact{
		{
			ID: "invalid",
			Concepts: []domain.ConceptID{
				"does-not-exist",
			},
		},
	}

	if err := currentOntology.ValidateFacts(
		facts,
	); err == nil {
		t.Fatal(
			"expected validation error",
		)
	}
}

func containsConcept(
	values []domain.ConceptID,
	expected domain.ConceptID,
) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}

	return false
}

func conceptWeight(
	values []WeightedConcept,
	expected domain.ConceptID,
) (float64, bool) {
	for _, value := range values {
		if value.ID == expected {
			return value.Weight, true
		}
	}

	return 0, false
}
