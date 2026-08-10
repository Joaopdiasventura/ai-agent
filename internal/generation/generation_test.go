package generation

import (
	"strings"
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/index"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
	"ai-agent/internal/planning"
	"ai-agent/internal/query"
	"ai-agent/internal/ranking"
	"ai-agent/internal/reasoning"
	"ai-agent/internal/retrieval"
)

type testEnvironment struct {
	base      *knowledge.Knowledge
	analyzer  *query.Analyzer
	retriever *retrieval.HybridRetriever
	ranker    *ranking.Ranker
	reasoner  *reasoning.Reasoner
	planner   *planning.Planner
	generator *Generator
}

func createEnvironment(
	t *testing.T,
) testEnvironment {
	t.Helper()

	base, err :=
		knowledge.New()

	if err != nil {
		t.Fatal(err)
	}

	currentOntology, err :=
		ontology.New()

	if err != nil {
		t.Fatal(err)
	}

	currentIndex, err :=
		index.New(
			base,
			currentOntology,
		)

	if err != nil {
		t.Fatal(err)
	}

	analyzer, err :=
		query.NewAnalyzer(
			base,
			currentOntology,
		)

	if err != nil {
		t.Fatal(err)
	}

	retriever, err :=
		retrieval.NewHybridRetriever(
			base,
			currentIndex,
		)

	if err != nil {
		t.Fatal(err)
	}

	ranker, err :=
		ranking.New(base)

	if err != nil {
		t.Fatal(err)
	}

	reasoner, err :=
		reasoning.New(base)

	if err != nil {
		t.Fatal(err)
	}

	planner, err :=
		planning.New(base)

	if err != nil {
		t.Fatal(err)
	}

	return testEnvironment{
		base:      base,
		analyzer:  analyzer,
		retriever: retriever,
		ranker:    ranker,
		reasoner:  reasoner,
		planner:   planner,
		generator: New(),
	}
}

func buildPlan(
	t *testing.T,
	environment testEnvironment,
	value string,
) planning.Plan {
	t.Helper()

	currentQuery :=
		environment.analyzer.Analyze(
			value,
		)

	retrieved :=
		environment.retriever.Search(
			currentQuery,
			80,
		)

	ranked :=
		environment.ranker.Rank(
			retrieved,
			40,
		)

	reasoned :=
		environment.reasoner.Reason(
			currentQuery,
			ranked,
		)

	return environment.planner.Plan(
		reasoned,
	)
}

func generateQuestion(
	t *testing.T,
	environment testEnvironment,
	value string,
) (
	planning.Plan,
	Material,
	Answer,
) {
	t.Helper()

	plan :=
		buildPlan(
			t,
			environment,
			value,
		)

	material, err :=
		Materialize(
			plan,
			environment.base,
		)

	if err != nil {
		t.Fatal(err)
	}

	answer, err :=
		environment.generator.Generate(
			plan,
			material,
		)

	if err != nil {
		t.Fatal(err)
	}

	return plan,
		material,
		answer
}

func TestGeneratorBuilds(
	t *testing.T,
) {
	generator :=
		New()

	if generator == nil {
		t.Fatal(
			"expected generator",
		)
	}
}

func TestMaterialContainsOnlyPlannedFacts(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		buildPlan(
			t,
			environment,
			"Ele sabe Go?",
		)

	material, err :=
		Materialize(
			plan,
			environment.base,
		)

	if err != nil {
		t.Fatal(err)
	}

	if len(material.Facts) !=
		len(plan.FactIDs()) {
		t.Fatalf(
			"expected %d facts, got %d",
			len(plan.FactIDs()),
			len(material.Facts),
		)
	}

	for factID := range material.Facts {
		if !containsFactID(
			plan.FactIDs(),
			factID,
		) {
			t.Fatalf(
				"material contains unauthorized fact %s",
				factID,
			)
		}
	}
}

func TestPortugueseEmailGeneration(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Qual o email do João?",
		)

	if answer.Language !=
		domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			answer.Language,
		)
	}

	if !strings.Contains(
		strings.ToLower(
			answer.Text,
		),
		"joaopdias.dev@gmail.com",
	) {
		t.Fatalf(
			"expected email in answer, got %q",
			answer.Text,
		)
	}

	if len(answer.FactIDs) != 1 {
		t.Fatalf(
			"expected one fact, got %v",
			answer.FactIDs,
		)
	}

	if answer.FactIDs[0] !=
		"profile-email" {
		t.Fatalf(
			"expected profile-email, got %s",
			answer.FactIDs[0],
		)
	}
}

func TestGoCapabilityGeneration(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Ele sabe Go?",
		)

	if !strings.HasPrefix(
		answer.Text,
		"Sim.",
	) {
		t.Fatalf(
			"expected affirmative answer, got %q",
			answer.Text,
		)
	}

	if !strings.Contains(
		answer.Text,
		"Go",
	) {
		t.Fatalf(
			"expected Go in answer, got %q",
			answer.Text,
		)
	}

	if !containsFactID(
		answer.FactIDs,
		"skill-go",
	) {
		t.Fatalf(
			"expected skill-go provenance, got %v",
			answer.FactIDs,
		)
	}
}

func TestEnglishGoCapabilityGeneration(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Does João know Go?",
		)

	if answer.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			answer.Language,
		)
	}

	if !strings.HasPrefix(
		answer.Text,
		"Yes.",
	) {
		t.Fatalf(
			"expected english affirmative answer, got %q",
			answer.Text,
		)
	}

	if !strings.Contains(
		answer.Text,
		"Go",
	) {
		t.Fatalf(
			"expected Go, got %q",
			answer.Text,
		)
	}
}

func TestUnknownCapabilityAbstains(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan, material, answer :=
		generateQuestion(
			t,
			environment,
			"Ele sabe Rust?",
		)

	if plan.Status !=
		planning.PlanStatusAbstain {
		t.Fatalf(
			"expected abstention plan, got %s",
			plan.Status,
		)
	}

	if len(material.Facts) != 0 {
		t.Fatalf(
			"expected no material facts, got %v",
			material.Facts,
		)
	}

	if len(answer.FactIDs) != 0 {
		t.Fatalf(
			"expected no answer facts, got %v",
			answer.FactIDs,
		)
	}

	if !strings.Contains(
		strings.ToLower(
			answer.Text,
		),
		"evidência suficiente",
	) {
		t.Fatalf(
			"expected insufficient evidence response, got %q",
			answer.Text,
		)
	}
}

func TestEnglishUnknownCapabilityAbstains(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Does João know Rust?",
		)

	if answer.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			answer.Language,
		)
	}

	if !strings.Contains(
		strings.ToLower(
			answer.Text,
		),
		"enough evidence",
	) {
		t.Fatalf(
			"expected abstention, got %q",
			answer.Text,
		)
	}
}

func TestConcurrencyComparisonMentionsGGCompress(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	if !strings.Contains(
		answer.Text,
		"GGCompress",
	) {
		t.Fatalf(
			"expected GGCompress, got %q",
			answer.Text,
		)
	}

	if len(answer.FactIDs) < 2 {
		t.Fatalf(
			"expected multiple supporting facts, got %v",
			answer.FactIDs,
		)
	}

	if !containsFactID(
		answer.FactIDs,
		"project-ggcompress-concurrency",
	) {
		t.Fatalf(
			"expected concurrency provenance, got %v",
			answer.FactIDs,
		)
	}
}

func TestDistributedSystemsComparisonMentionsAuronix(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra sistemas distribuídos?",
		)

	if !strings.Contains(
		answer.Text,
		"Auronix",
	) {
		t.Fatalf(
			"expected Auronix, got %q",
			answer.Text,
		)
	}

	if !containsFactID(
		answer.FactIDs,
		"project-auronix-distributed",
	) {
		t.Fatalf(
			"expected distributed fact, got %v",
			answer.FactIDs,
		)
	}
}

func TestKafkaUsageMentionsXTube(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Where did João use Kafka?",
		)

	if answer.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			answer.Language,
		)
	}

	if !strings.Contains(
		answer.Text,
		"Kafka",
	) {
		t.Fatalf(
			"expected Kafka, got %q",
			answer.Text,
		)
	}

	if !strings.Contains(
		answer.Text,
		"X Tube",
	) {
		t.Fatalf(
			"expected X Tube, got %q",
			answer.Text,
		)
	}

	if !containsFactID(
		answer.FactIDs,
		"project-xtube-kafka",
	) {
		t.Fatalf(
			"expected kafka provenance, got %v",
			answer.FactIDs,
		)
	}
}

func TestGGCompressOverview(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan, _, answer :=
		generateQuestion(
			t,
			environment,
			"Me fale sobre o GGCompress",
		)

	if !strings.Contains(
		answer.Text,
		"GGCompress",
	) {
		t.Fatalf(
			"expected GGCompress, got %q",
			answer.Text,
		)
	}

	for _, factID := range answer.FactIDs {
		if !containsFactID(
			plan.FactIDs(),
			factID,
		) {
			t.Fatalf(
				"answer used unauthorized fact %s",
				factID,
			)
		}
	}
}

func TestEnglishGGCompressOverview(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Tell me about GGCompress",
		)

	if answer.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			answer.Language,
		)
	}

	if !strings.HasPrefix(
		answer.Text,
		"About GGCompress:",
	) {
		t.Fatalf(
			"expected english overview, got %q",
			answer.Text,
		)
	}
}

func TestAnswerNeverUsesFactOutsidePlan(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	questions := []string{
		"Qual projeto melhor demonstra concorrência?",
		"Ele sabe Go?",
		"Where did João use Kafka?",
		"Me fale sobre o GGCompress",
		"Qual o email do João?",
	}

	for _, questionValue := range questions {
		plan, _, answer :=
			generateQuestion(
				t,
				environment,
				questionValue,
			)

		for _, factID := range answer.FactIDs {
			if !containsFactID(
				plan.FactIDs(),
				factID,
			) {
				t.Fatalf(
					"question %q generated unauthorized fact %s",
					questionValue,
					factID,
				)
			}
		}
	}
}

func TestMissingMaterialFactFails(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		buildPlan(
			t,
			environment,
			"Qual o email do João?",
		)

	material :=
		NewMaterial()

	for _, entityID := range plan.EntityIDs() {
		entity, found :=
			environment.base.Entity(
				entityID,
			)

		if found {
			material.Entities[entityID] = entity
		}
	}

	_, err :=
		environment.generator.Generate(
			plan,
			material,
		)

	if err == nil {
		t.Fatal(
			"expected missing material error",
		)
	}
}

func TestMissingMaterialEntityFails(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		buildPlan(
			t,
			environment,
			"Me fale sobre o GGCompress",
		)

	material :=
		NewMaterial()

	for _, factID := range plan.FactIDs() {
		fact, found :=
			environment.base.Fact(
				factID,
			)

		if found {
			material.Facts[factID] = fact
		}
	}

	_, err :=
		environment.generator.Generate(
			plan,
			material,
		)

	if err == nil {
		t.Fatal(
			"expected missing entity error",
		)
	}
}

func TestAnswerFactIDsHaveNoDuplicates(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	_, _, answer :=
		generateQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, factID := range answer.FactIDs {
		if _, exists :=
			seen[factID]; exists {
			t.Fatalf(
				"duplicate provenance fact %s",
				factID,
			)
		}

		seen[factID] =
			struct{}{}
	}
}

func TestAnswerIsNotEmpty(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	questions := []string{
		"Qual projeto melhor demonstra concorrência?",
		"Ele sabe Go?",
		"Where did João use Kafka?",
		"Me fale sobre o GGCompress",
		"Qual o email do João?",
		"Ele sabe Rust?",
	}

	for _, questionValue := range questions {
		_, _, answer :=
			generateQuestion(
				t,
				environment,
				questionValue,
			)

		if answer.Empty() {
			t.Fatalf(
				"question %q produced empty answer",
				questionValue,
			)
		}
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
