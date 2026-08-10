package reasoning

import (
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/index"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
	"ai-agent/internal/query"
	"ai-agent/internal/ranking"
	"ai-agent/internal/retrieval"
)

type testEnvironment struct {
	base      *knowledge.Knowledge
	analyzer  *query.Analyzer
	retriever *retrieval.HybridRetriever
	ranker    *ranking.Ranker
	reasoner  *Reasoner
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
		ranking.New(
			base,
		)

	if err != nil {
		t.Fatal(err)
	}

	reasoner, err :=
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
	}
}

func reasonQuestion(
	t *testing.T,
	environment testEnvironment,
	value string,
) Result {
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

	return environment.reasoner.Reason(
		currentQuery,
		ranked,
	)
}

func TestReasonerBuilds(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	if environment.reasoner == nil {
		t.Fatal(
			"expected reasoner",
		)
	}
}

func TestConcurrencyComparisonChoosesGGCompress(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	if result.Conclusion.Type !=
		ConclusionComparison {
		t.Fatalf(
			"expected comparison, got %s",
			result.Conclusion.Type,
		)
	}

	if result.Conclusion.Status !=
		SupportSupported {
		t.Fatalf(
			"expected supported, got %s",
			result.Conclusion.Status,
		)
	}

	group, found :=
		result.TopGroup()

	if !found {
		t.Fatal(
			"expected project groups",
		)
	}

	if group.EntityID !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress first, got %s",
			group.EntityID,
		)
	}
}

func TestConcurrencyComparisonHasMultipleGGCompressEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	group, found :=
		findGroup(
			result.Conclusion.Groups,
			knowledge.EntityGGCompress,
		)

	if !found {
		t.Fatal(
			"expected ggcompress group",
		)
	}

	if len(group.Evidence) < 2 {
		t.Fatalf(
			"expected multiple ggcompress evidence, got %d",
			len(group.Evidence),
		)
	}

	if !groupHasFact(
		group,
		"project-ggcompress-concurrency",
	) {
		t.Fatal(
			"expected concurrency fact",
		)
	}

	if !groupHasFact(
		group,
		"project-ggcompress-goroutines",
	) {
		t.Fatal(
			"expected goroutines fact",
		)
	}
}

func TestDistributedSystemsComparisonChoosesAuronix(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra sistemas distribuídos?",
		)

	group, found :=
		result.TopGroup()

	if !found {
		t.Fatal(
			"expected project group",
		)
	}

	if group.EntityID !=
		knowledge.EntityAuronix {
		t.Fatalf(
			"expected auronix first, got %s",
			group.EntityID,
		)
	}

	if !groupHasFact(
		group,
		"project-auronix-distributed",
	) {
		t.Fatal(
			"expected auronix distributed evidence",
		)
	}
}

func TestGoCapabilityIsSupported(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Ele sabe Go?",
		)

	if result.Conclusion.Type !=
		ConclusionCapability {
		t.Fatalf(
			"expected capability, got %s",
			result.Conclusion.Type,
		)
	}

	if result.Conclusion.Status !=
		SupportSupported {
		t.Fatalf(
			"expected supported, got %s",
			result.Conclusion.Status,
		)
	}

	if !evidenceContainsFact(
		result.Conclusion.Evidence,
		"skill-go",
	) {
		t.Fatalf(
			"expected skill-go, got %v",
			result.Conclusion.Evidence,
		)
	}
}

func TestConcurrencyCapabilityIsSupported(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Ele tem experiência com concorrência?",
		)

	if result.Conclusion.Status !=
		SupportSupported {
		t.Fatalf(
			"expected supported, got %s",
			result.Conclusion.Status,
		)
	}

	if result.Conclusion.FocusConcept !=
		ontology.ConceptConcurrency {
		t.Fatalf(
			"expected concurrency focus, got %s",
			result.Conclusion.FocusConcept,
		)
	}

	if len(
		result.Conclusion.Evidence,
	) == 0 {
		t.Fatal(
			"expected capability evidence",
		)
	}
}

func TestUnknownCapabilityDoesNotBecomeFalseClaim(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Ele sabe Rust?",
		)

	if result.Conclusion.Type !=
		ConclusionCapability {
		t.Fatalf(
			"expected capability, got %s",
			result.Conclusion.Type,
		)
	}

	if result.Conclusion.Status !=
		SupportInsufficientEvidence {
		t.Fatalf(
			"expected insufficient evidence, got %s",
			result.Conclusion.Status,
		)
	}
}

func TestKafkaTechnologyUsage(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Where did João use Kafka?",
		)

	if result.Conclusion.Type !=
		ConclusionTechnologyUsage {
		t.Fatalf(
			"expected technology usage, got %s",
			result.Conclusion.Type,
		)
	}

	if result.Conclusion.Status !=
		SupportSupported {
		t.Fatalf(
			"expected supported, got %s",
			result.Conclusion.Status,
		)
	}

	if result.Conclusion.FocusEntity !=
		knowledge.EntityKafka {
		t.Fatalf(
			"expected kafka focus, got %s",
			result.Conclusion.FocusEntity,
		)
	}

	if !evidenceContainsFact(
		result.Conclusion.Evidence,
		"project-xtube-kafka",
	) {
		t.Fatalf(
			"expected x-tube kafka fact, got %v",
			result.Conclusion.Evidence,
		)
	}
}

func TestKafkaUsageGroupsXTube(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Where did João use Kafka?",
		)

	group, found :=
		findGroup(
			result.Conclusion.Groups,
			knowledge.EntityXTube,
		)

	if !found {
		t.Fatal(
			"expected x-tube group",
		)
	}

	if !groupHasFact(
		group,
		"project-xtube-kafka",
	) {
		t.Fatal(
			"expected kafka fact in x-tube group",
		)
	}
}

func TestGGCompressOverviewFiltersEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Me fale sobre o GGCompress",
		)

	if result.Conclusion.Type !=
		ConclusionOverview {
		t.Fatalf(
			"expected overview, got %s",
			result.Conclusion.Type,
		)
	}

	if result.Conclusion.FocusEntity !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress focus, got %s",
			result.Conclusion.FocusEntity,
		)
	}

	for _, currentEvidence := range result.Conclusion.Evidence {
		fact, found :=
			environment.base.Fact(
				currentEvidence.FactID,
			)

		if !found {
			t.Fatal(
				"expected fact",
			)
		}

		if !factReferencesEntity(
			fact,
			knowledge.EntityGGCompress,
		) {
			t.Fatalf(
				"unexpected non-ggcompress evidence %s",
				fact.ID,
			)
		}
	}
}

func TestEmailDirectConclusion(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual o email do João?",
		)

	if result.Conclusion.Status !=
		SupportSupported {
		t.Fatalf(
			"expected supported, got %s",
			result.Conclusion.Status,
		)
	}

	if len(
		result.Conclusion.Evidence,
	) == 0 {
		t.Fatal(
			"expected evidence",
		)
	}

	if result.Conclusion.Evidence[0].FactID !=
		"profile-email" {
		t.Fatalf(
			"expected profile-email first, got %s",
			result.Conclusion.Evidence[0].FactID,
		)
	}
}

func TestCurrentRoleEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual é o cargo atual do João?",
		)

	if !evidenceContainsFact(
		result.Conclusion.Evidence,
		"experience-current-role",
	) {
		t.Fatalf(
			"expected current role evidence, got %v",
			result.Conclusion.Evidence,
		)
	}
}

func TestProjectGroupScoresNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto demonstra performance?",
		)

	for _, group := range result.Conclusion.Groups {
		if group.Score < 0 ||
			group.Score > 1 {
			t.Fatalf(
				"invalid group score %f",
				group.Score,
			)
		}

		if group.EvidenceStrength < 0 ||
			group.EvidenceStrength > 1 {
			t.Fatalf(
				"invalid evidence strength %f",
				group.EvidenceStrength,
			)
		}

		if group.ConceptCoverage < 0 ||
			group.ConceptCoverage > 1 {
			t.Fatalf(
				"invalid coverage %f",
				group.ConceptCoverage,
			)
		}

		if group.Diversity < 0 ||
			group.Diversity > 1 {
			t.Fatalf(
				"invalid diversity %f",
				group.Diversity,
			)
		}

		if group.Quantity < 0 ||
			group.Quantity > 1 {
			t.Fatalf(
				"invalid quantity %f",
				group.Quantity,
			)
		}
	}
}

func TestGroupsAreSorted(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto demonstra performance?",
		)

	for index := 1; index <
		len(result.Conclusion.Groups); index++ {
		previous :=
			result.Conclusion.Groups[index-1]

		current :=
			result.Conclusion.Groups[index]

		if current.Score >
			previous.Score {
			t.Fatal(
				"groups are not sorted",
			)
		}

		if current.Rank != index+1 {
			t.Fatalf(
				"expected rank %d, got %d",
				index+1,
				current.Rank,
			)
		}
	}
}

func TestEvidenceHasNoDuplicates(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, currentEvidence := range result.Conclusion.Evidence {
		if _, exists :=
			seen[currentEvidence.FactID]; exists {
			t.Fatalf(
				"duplicate evidence %s",
				currentEvidence.FactID,
			)
		}

		seen[currentEvidence.FactID] =
			struct{}{}
	}
}

func TestUnknownQueryHasInsufficientEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"xyzabc123",
		)

	if result.Conclusion.Status !=
		SupportInsufficientEvidence {
		t.Fatalf(
			"expected insufficient evidence, got %s",
			result.Conclusion.Status,
		)
	}
}

func TestComparisonDoesNotInventProjects(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Qual projeto demonstra concorrência?",
		)

	for _, group := range result.Conclusion.Groups {
		entity, found :=
			environment.base.Entity(
				group.EntityID,
			)

		if !found {
			t.Fatalf(
				"unknown group entity %s",
				group.EntityID,
			)
		}

		if entity.Type !=
			domain.EntityTypeProject {
			t.Fatalf(
				"expected project, got %s",
				entity.Type,
			)
		}
	}
}

func findGroup(
	groups []EntityGroup,
	entityID domain.EntityID,
) (EntityGroup, bool) {
	for _, group := range groups {
		if group.EntityID ==
			entityID {
			return group, true
		}
	}

	return EntityGroup{}, false
}

func groupHasFact(
	group EntityGroup,
	factID domain.FactID,
) bool {
	return evidenceContainsFact(
		group.Evidence,
		factID,
	)
}

func evidenceContainsFact(
	evidence []Evidence,
	factID domain.FactID,
) bool {
	for _, currentEvidence := range evidence {
		if currentEvidence.FactID ==
			factID {
			return true
		}
	}

	return false
}

func TestUnknownCapabilityDoesNotUsePersonAsEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		reasonQuestion(
			t,
			environment,
			"Ele sabe Rust?",
		)

	if result.Conclusion.Status !=
		SupportInsufficientEvidence {
		t.Fatalf(
			"expected insufficient evidence, got %s",
			result.Conclusion.Status,
		)
	}

	if result.Conclusion.FocusEntity != "" {
		t.Fatalf(
			"expected no capability focus entity, got %s",
			result.Conclusion.FocusEntity,
		)
	}

	if len(
		result.Conclusion.Evidence,
	) != 0 {
		t.Fatalf(
			"expected no capability evidence, got %v",
			result.Conclusion.Evidence,
		)
	}
}
