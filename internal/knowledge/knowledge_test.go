package knowledge

import (
	"testing"

	"ai-agent/internal/domain"
)

func TestKnowledgeBuildsSuccessfully(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if len(base.Entities()) == 0 {
		t.Fatal("expected entities")
	}

	if len(base.Facts()) == 0 {
		t.Fatal("expected facts")
	}
}

func TestKnowledgeContainsJoao(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	entity, found := base.Entity(EntityJoao)

	if !found {
		t.Fatal("expected joao entity")
	}

	if entity.Type != domain.EntityTypePerson {
		t.Fatalf(
			"expected person entity, got %s",
			entity.Type,
		)
	}
}

func TestKnowledgeContainsProjects(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	expected := []domain.EntityID{
		EntityAuronix,
		EntityXTube,
		EntityGGCompress,
		EntityVox,
	}

	for _, id := range expected {
		entity, found := base.Entity(id)

		if !found {
			t.Fatalf(
				"expected project %s",
				id,
			)
		}

		if entity.Type != domain.EntityTypeProject {
			t.Fatalf(
				"expected %s to be project",
				id,
			)
		}
	}
}

func TestGGCompressHasFacts(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	facts := base.FactsBySubject(EntityGGCompress)

	if len(facts) == 0 {
		t.Fatal("expected ggcompress facts")
	}
}

func TestConcurrencyFactsExist(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	facts := base.FactsByConcept(
		domain.ConceptID("concurrency"),
	)

	if len(facts) == 0 {
		t.Fatal("expected concurrency facts")
	}
}

func TestFinancialVolumeUsesGreaterThan(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	fact, found := base.Fact(
		domain.FactID(
			"experience-current-financial-volume",
		),
	)

	if !found {
		t.Fatal("expected financial volume fact")
	}

	if fact.Object.Kind != domain.FactObjectNumber {
		t.Fatal("expected numeric fact")
	}

	if fact.Object.Number != 1000000 {
		t.Fatalf(
			"expected 1000000, got %f",
			fact.Object.Number,
		)
	}

	if fact.Object.Operator != domain.NumberOperatorGreaterThan {
		t.Fatalf(
			"expected greater_than, got %s",
			fact.Object.Operator,
		)
	}
}

func TestGGCompressThroughput(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	fact, found := base.Fact(
		domain.FactID(
			"project-ggcompress-throughput",
		),
	)

	if !found {
		t.Fatal("expected ggcompress throughput fact")
	}

	if fact.Object.Number != 1.23 {
		t.Fatalf(
			"expected 1.23, got %f",
			fact.Object.Number,
		)
	}

	if fact.Object.Unit != "GB/s" {
		t.Fatalf(
			"expected GB/s, got %s",
			fact.Object.Unit,
		)
	}

	if fact.Object.Operator != domain.NumberOperatorUpTo {
		t.Fatalf(
			"expected up_to, got %s",
			fact.Object.Operator,
		)
	}
}

func TestVoxVoterCount(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	fact, found := base.Fact(
		domain.FactID("project-vox-voters"),
	)

	if !found {
		t.Fatal("expected vox voter fact")
	}

	if fact.Object.Number != 500 {
		t.Fatalf(
			"expected 500, got %f",
			fact.Object.Number,
		)
	}

	if fact.Object.Operator != domain.NumberOperatorGreaterThan {
		t.Fatalf(
			"expected greater_than, got %s",
			fact.Object.Operator,
		)
	}
}

func TestAllFactsHaveBothLanguages(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	for _, fact := range base.Facts() {
		if fact.Statement.PT == "" {
			t.Fatalf(
				"fact %s has no portuguese statement",
				fact.ID,
			)
		}

		if fact.Statement.EN == "" {
			t.Fatalf(
				"fact %s has no english statement",
				fact.ID,
			)
		}
	}
}

func TestFactsByContext(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	facts := base.FactsByContext(EntityUFind)

	if len(facts) == 0 {
		t.Fatal("expected ufind context facts")
	}

	for _, fact := range facts {
		found := false

		for _, context := range fact.Context {
			if context == EntityUFind {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf(
				"fact %s should contain ufind context",
				fact.ID,
			)
		}
	}
}

func TestFactsAreOrderedByImportance(t *testing.T) {
	base, err := New()

	if err != nil {
		t.Fatal(err)
	}

	facts := base.FactsBySubject(EntityGGCompress)

	for i := 1; i < len(facts); i++ {
		if facts[i].Importance > facts[i-1].Importance {
			t.Fatalf(
				"facts are not ordered by importance",
			)
		}
	}
}
