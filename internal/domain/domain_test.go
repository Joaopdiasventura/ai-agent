package domain

import "testing"

func TestLocalizedTextForPortuguese(t *testing.T) {
	text := LocalizedText{
		PT: "Olá",
		EN: "Hello",
	}

	if got := text.For(LanguagePortuguese); got != "Olá" {
		t.Fatalf("expected Olá, got %s", got)
	}
}

func TestLocalizedTextForEnglish(t *testing.T) {
	text := LocalizedText{
		PT: "Olá",
		EN: "Hello",
	}

	if got := text.For(LanguageEnglish); got != "Hello" {
		t.Fatalf("expected Hello, got %s", got)
	}
}

func TestLocalizedTextFallback(t *testing.T) {
	text := LocalizedText{
		PT: "Olá",
	}

	if got := text.For(LanguageEnglish); got != "Olá" {
		t.Fatalf("expected fallback Olá, got %s", got)
	}
}

func TestEntityAliases(t *testing.T) {
	entity := Entity{
		ID:   "ggcompress",
		Type: EntityTypeProject,
		Name: LocalizedText{
			PT: "GGCompress",
			EN: "GGCompress",
		},
		Aliases: map[Language][]string{
			LanguageUnknown: {
				"gg compress",
			},
			LanguagePortuguese: {
				"projeto de compressão",
			},
		},
	}

	aliases := entity.AllAliases(LanguagePortuguese)

	if len(aliases) != 3 {
		t.Fatalf("expected 3 aliases, got %d", len(aliases))
	}
}

func TestConceptAliases(t *testing.T) {
	concept := Concept{
		ID: "concurrency",
		Name: LocalizedText{
			PT: "concorrência",
			EN: "concurrency",
		},
		Aliases: map[Language][]string{
			LanguagePortuguese: {
				"paralelismo",
			},
			LanguageEnglish: {
				"parallelism",
			},
		},
	}

	aliases := concept.AllAliases(LanguageEnglish)

	if len(aliases) != 2 {
		t.Fatalf("expected 2 aliases, got %d", len(aliases))
	}
}

func TestValidFact(t *testing.T) {
	fact := Fact{
		ID:        "ggcompress-throughput",
		Subject:   "ggcompress",
		Predicate: RelationAchieved,
		Object:    NumberObject(1.23, "GB/s"),
		Category:  FactCategoryAchievement,
		Concepts: []ConceptID{
			"performance",
			"concurrency",
		},
		Statement: LocalizedText{
			PT: "O GGCompress alcançou throughput de até 1,23 GB/s.",
			EN: "GGCompress achieved throughput of up to 1.23 GB/s.",
		},
		Importance: 1,
	}

	if err := fact.Validate(); err != nil {
		t.Fatalf("expected valid fact, got %v", err)
	}
}

func TestFactRequiresSubject(t *testing.T) {
	fact := Fact{
		ID:        "invalid",
		Predicate: RelationUses,
		Object:    EntityObject("go"),
		Statement: LocalizedText{
			PT: "Texto",
		},
		Importance: 0.5,
	}

	if err := fact.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestFactRejectsInvalidImportance(t *testing.T) {
	fact := Fact{
		ID:        "invalid-importance",
		Subject:   "ggcompress",
		Predicate: RelationUses,
		Object:    EntityObject("go"),
		Statement: LocalizedText{
			PT: "O GGCompress utiliza Go.",
			EN: "GGCompress uses Go.",
		},
		Importance: 1.5,
	}

	if err := fact.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCurrentPeriodCannotHaveEnd(t *testing.T) {
	start := YearMonth{
		Year:  2025,
		Month: 6,
	}

	end := YearMonth{
		Year:  2026,
		Month: 8,
	}

	period := Period{
		Start:   &start,
		End:     &end,
		Current: true,
	}

	if err := period.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestQueryHasEntity(t *testing.T) {
	query := Query{
		Entities: []EntityMatch{
			{
				EntityID: "ggcompress",
				Score:    1,
				Explicit: true,
			},
		},
	}

	if !query.HasEntity("ggcompress") {
		t.Fatal("expected query to contain ggcompress")
	}

	if query.HasEntity("auronix") {
		t.Fatal("did not expect query to contain auronix")
	}
}

func TestQueryHasConcept(t *testing.T) {
	query := Query{
		Concepts: []ConceptMatch{
			{
				ConceptID: "concurrency",
				Score:     1,
			},
		},
	}

	if !query.HasConcept("concurrency") {
		t.Fatal("expected query to contain concurrency")
	}
}

func TestEvidenceSignal(t *testing.T) {
	evidence := Evidence{
		FactID: "ggcompress-throughput",
		Score:  0.9,
		Signals: []EvidenceSignal{
			{
				Name:  "concept",
				Score: 0.8,
			},
		},
	}

	score, found := evidence.Signal("concept")

	if !found {
		t.Fatal("expected signal to be found")
	}

	if score != 0.8 {
		t.Fatalf("expected 0.8, got %f", score)
	}
}
