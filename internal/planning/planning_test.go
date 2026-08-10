package planning

import (
	"slices"
	"testing"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
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
	planner   *Planner
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
		searchindex.New(
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
		New(base)

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
	}
}

func planQuestion(
	t *testing.T,
	environment testEnvironment,
	value string,
) Plan {
	t.Helper()

	currentQuery :=
		environment.analyzer.Analyze(
			value,
		)

	retrievalResult :=
		environment.retriever.Search(
			currentQuery,
			80,
		)

	ranked :=
		environment.ranker.Rank(
			retrievalResult,
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

func TestPlannerBuilds(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	if environment.planner == nil {
		t.Fatal(
			"expected planner",
		)
	}
}

func TestConcurrencyComparisonPlansGGCompressWinner(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	if plan.Status !=
		PlanStatusReady {
		t.Fatalf(
			"expected ready, got %s",
			plan.Status,
		)
	}

	if plan.Type !=
		PlanTypeComparison {
		t.Fatalf(
			"expected comparison, got %s",
			plan.Type,
		)
	}

	lead, found :=
		plan.Section(
			SectionLead,
		)

	if !found ||
		len(lead.Items) == 0 {
		t.Fatal(
			"expected lead section",
		)
	}

	winner :=
		lead.Items[0]

	if winner.Kind !=
		ItemComparisonWinner {
		t.Fatalf(
			"expected comparison winner, got %s",
			winner.Kind,
		)
	}

	if winner.EntityID !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress, got %s",
			winner.EntityID,
		)
	}

	if len(
		winner.EvidenceIDs,
	) < 2 {
		t.Fatalf(
			"expected multiple winner evidence, got %v",
			winner.EvidenceIDs,
		)
	}
}

func TestComparisonContainsWinnerEvidenceSection(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	section, found :=
		plan.Section(
			SectionEvidence,
		)

	if !found {
		t.Fatal(
			"expected evidence section",
		)
	}

	if len(section.Items) == 0 {
		t.Fatal(
			"expected evidence items",
		)
	}

	if !sectionHasFact(
		section,
		"project-ggcompress-concurrency",
	) {
		t.Fatalf(
			"expected concurrency fact, got %v",
			section.Items,
		)
	}
}

func TestDistributedSystemsPlansAuronixWinner(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra sistemas distribuídos?",
		)

	lead, found :=
		plan.Section(
			SectionLead,
		)

	if !found ||
		len(lead.Items) == 0 {
		t.Fatal(
			"expected lead",
		)
	}

	if lead.Items[0].EntityID !=
		knowledge.EntityAuronix {
		t.Fatalf(
			"expected auronix, got %s",
			lead.Items[0].EntityID,
		)
	}
}

func TestGoCapabilityCreatesSupportedLead(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Ele sabe Go?",
		)

	if plan.Type !=
		PlanTypeCapability {
		t.Fatalf(
			"expected capability, got %s",
			plan.Type,
		)
	}

	if plan.Status !=
		PlanStatusReady {
		t.Fatalf(
			"expected ready, got %s",
			plan.Status,
		)
	}

	lead, found :=
		plan.Section(
			SectionLead,
		)

	if !found ||
		len(lead.Items) != 1 {
		t.Fatal(
			"expected support lead",
		)
	}

	item :=
		lead.Items[0]

	if item.Kind !=
		ItemSupport {
		t.Fatalf(
			"expected support item, got %s",
			item.Kind,
		)
	}

	if item.Support !=
		reasoning.SupportSupported {
		t.Fatalf(
			"expected supported, got %s",
			item.Support,
		)
	}

	if item.EntityID !=
		knowledge.EntityGo {
		t.Fatalf(
			"expected Go focus, got %s",
			item.EntityID,
		)
	}
}

func TestGoCapabilityIncludesEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Ele sabe Go?",
		)

	if !planContainsFact(
		plan,
		"skill-go",
	) {
		t.Fatalf(
			"expected skill-go, got %v",
			plan.FactIDs(),
		)
	}
}

func TestUnknownCapabilityCreatesAbstention(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Ele sabe Rust?",
		)

	if plan.Status !=
		PlanStatusAbstain {
		t.Fatalf(
			"expected abstain, got %s",
			plan.Status,
		)
	}

	section, found :=
		plan.Section(
			SectionAbstention,
		)

	if !found {
		t.Fatal(
			"expected abstention section",
		)
	}

	if len(section.Items) != 1 {
		t.Fatal(
			"expected one abstention item",
		)
	}

	if section.Items[0].Support !=
		reasoning.SupportInsufficientEvidence {
		t.Fatalf(
			"expected insufficient evidence, got %s",
			section.Items[0].Support,
		)
	}
}

func TestUnknownCapabilityDoesNotPlanFacts(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Ele sabe Rust?",
		)

	if len(plan.FactIDs()) != 0 {
		t.Fatalf(
			"expected no fact claims, got %v",
			plan.FactIDs(),
		)
	}
}

func TestKafkaUsagePlansXTube(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Where did João use Kafka?",
		)

	if plan.Type !=
		PlanTypeTechnologyUsage {
		t.Fatalf(
			"expected technology usage, got %s",
			plan.Type,
		)
	}

	if plan.FocusEntity !=
		knowledge.EntityKafka {
		t.Fatalf(
			"expected kafka focus, got %s",
			plan.FocusEntity,
		)
	}

	section, found :=
		plan.Section(
			SectionList,
		)

	if !found {
		t.Fatal(
			"expected list section",
		)
	}

	foundXTube := false

	for _, item := range section.Items {
		if item.EntityID ==
			knowledge.EntityXTube {
			foundXTube = true

			if item.Kind !=
				ItemTechnologyUsage {
				t.Fatalf(
					"expected technology usage item, got %s",
					item.Kind,
				)
			}

			if !containsFactID(
				item.EvidenceIDs,
				"project-xtube-kafka",
			) {
				t.Fatalf(
					"expected kafka evidence, got %v",
					item.EvidenceIDs,
				)
			}
		}
	}

	if !foundXTube {
		t.Fatal(
			"expected x-tube usage",
		)
	}
}

func TestEmailPlansOnlyPrimaryFact(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual o email do João?",
		)

	if plan.Type !=
		PlanTypeDirect {
		t.Fatalf(
			"expected direct, got %s",
			plan.Type,
		)
	}

	facts :=
		plan.FactIDs()

	if len(facts) != 1 {
		t.Fatalf(
			"expected exactly one fact, got %v",
			facts,
		)
	}

	if facts[0] !=
		"profile-email" {
		t.Fatalf(
			"expected profile-email, got %s",
			facts[0],
		)
	}
}

func TestGGCompressOverviewFocus(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Me fale sobre o GGCompress",
		)

	if plan.Type !=
		PlanTypeOverview {
		t.Fatalf(
			"expected overview, got %s",
			plan.Type,
		)
	}

	if plan.FocusEntity !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress, got %s",
			plan.FocusEntity,
		)
	}

	lead, found :=
		plan.Section(
			SectionLead,
		)

	if !found ||
		len(lead.Items) == 0 {
		t.Fatal(
			"expected lead",
		)
	}

	if lead.Items[0].EntityID !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress lead, got %s",
			lead.Items[0].EntityID,
		)
	}
}

func TestGGCompressOverviewOnlyUsesRelatedFacts(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Me fale sobre o GGCompress",
		)

	for _, factID := range plan.FactIDs() {
		fact, found :=
			environment.base.Fact(
				factID,
			)

		if !found {
			t.Fatalf(
				"unknown fact %s",
				factID,
			)
		}

		if !factReferencesEntity(
			fact,
			knowledge.EntityGGCompress,
		) {
			t.Fatalf(
				"unexpected fact %s",
				factID,
			)
		}
	}
}

func TestEnglishLanguageIsPreserved(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Tell me about GGCompress",
		)

	if plan.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			plan.Language,
		)
	}
}

func TestPortugueseLanguageIsPreserved(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Me fale sobre o GGCompress",
		)

	if plan.Language !=
		domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			plan.Language,
		)
	}
}

func TestPlanDoesNotDuplicateFacts(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	facts :=
		plan.FactIDs()

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, factID := range facts {
		if _, exists :=
			seen[factID]; exists {
			t.Fatalf(
				"duplicate fact %s",
				factID,
			)
		}

		seen[factID] =
			struct{}{}
	}
}

func TestPlanFactIDsExistInKnowledge(
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
		plan :=
			planQuestion(
				t,
				environment,
				questionValue,
			)

		for _, factID := range plan.FactIDs() {
			if _, found :=
				environment.base.Fact(
					factID,
				); !found {
				t.Fatalf(
					"question %q planned unknown fact %s",
					questionValue,
					factID,
				)
			}
		}
	}
}

func TestComparisonAlternativesDoNotReplaceWinner(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	lead, found :=
		plan.Section(
			SectionLead,
		)

	if !found ||
		len(lead.Items) == 0 {
		t.Fatal(
			"expected winner",
		)
	}

	winner :=
		lead.Items[0].EntityID

	alternatives, found :=
		plan.Section(
			SectionAlternatives,
		)

	if !found {
		return
	}

	for _, item := range alternatives.Items {
		if item.EntityID == winner {
			t.Fatalf(
				"winner %s cannot appear as alternative",
				winner,
			)
		}
	}
}

func TestComparisonWinnerScoreNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	lead, found :=
		plan.Section(
			SectionLead,
		)

	if !found ||
		len(lead.Items) == 0 {
		t.Fatal(
			"expected lead",
		)
	}

	score :=
		lead.Items[0].Score

	if score < 0 ||
		score > 1 {
		t.Fatalf(
			"expected normalized score, got %f",
			score,
		)
	}
}

func TestUnknownQueryAbstains(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	plan :=
		planQuestion(
			t,
			environment,
			"xyzabc123",
		)

	if plan.Status !=
		PlanStatusAbstain {
		t.Fatalf(
			"expected abstain, got %s",
			plan.Status,
		)
	}

	if len(plan.FactIDs()) != 0 {
		t.Fatalf(
			"expected no claims, got %v",
			plan.FactIDs(),
		)
	}
}

func sectionHasFact(
	section Section,
	expected domain.FactID,
) bool {
	for _, item := range section.Items {
		if item.FactID == expected {
			return true
		}

		if containsFactID(
			item.EvidenceIDs,
			expected,
		) {
			return true
		}
	}

	return false
}

func planContainsFact(
	plan Plan,
	expected domain.FactID,
) bool {
	return containsFactID(
		plan.FactIDs(),
		expected,
	)
}

func containsFactID(
	values []domain.FactID,
	expected domain.FactID,
) bool {
	return slices.Contains(values, expected)
}

func factReferencesEntity(
	fact domain.Fact,
	entityID domain.EntityID,
) bool {
	if fact.Subject == entityID {
		return true
	}

	if fact.Object.Kind ==
		domain.FactObjectEntity &&
		fact.Object.EntityID ==
			entityID {
		return true
	}

	for _, contextID := range fact.Context {
		if contextID == entityID {
			return true
		}
	}

	return false
}
