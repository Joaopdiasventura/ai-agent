package ranking

import (
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/index"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
	"ai-agent/internal/query"
	"ai-agent/internal/retrieval"
)

type testEnvironment struct {
	base      *knowledge.Knowledge
	analyzer  *query.Analyzer
	retriever *retrieval.HybridRetriever
	ranker    *Ranker
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
		New(base)

	if err != nil {
		t.Fatal(err)
	}

	return testEnvironment{
		base:      base,
		analyzer:  analyzer,
		retriever: retriever,
		ranker:    ranker,
	}
}

func rankQuestion(
	t *testing.T,
	environment testEnvironment,
	value string,
	limit int,
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

	return environment.ranker.Rank(
		retrievalResult,
		limit,
	)
}

func TestRankerBuilds(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	if environment.ranker == nil {
		t.Fatal(
			"expected ranker",
		)
	}
}

func TestConcurrencyRanksGGCompressEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual projeto demonstra melhor concorrência?",
			10,
		)

	if !containsFact(
		result,
		"project-ggcompress-concurrency",
	) {
		t.Fatalf(
			"expected ggcompress concurrency fact, got %v",
			result.Candidates,
		)
	}

	candidate, found :=
		result.Candidate(
			"project-ggcompress-concurrency",
		)

	if !found {
		t.Fatal(
			"expected candidate",
		)
	}

	if candidate.Rank > 5 {
		t.Fatalf(
			"expected ggcompress concurrency in top 5, got rank %d",
			candidate.Rank,
		)
	}
}

func TestConcurrencyPrefersProjectEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual projeto demonstra concorrência?",
			10,
		)

	if len(result.Candidates) == 0 {
		t.Fatal(
			"expected candidates",
		)
	}

	topFact, found :=
		environment.base.Fact(
			result.Candidates[0].FactID,
		)

	if !found {
		t.Fatal(
			"expected top fact",
		)
	}

	if topFact.Category !=
		domain.FactCategoryProject &&
		topFact.Category !=
			domain.FactCategoryAchievement {
		t.Fatalf(
			"expected project evidence, got category %s from %s",
			topFact.Category,
			topFact.ID,
		)
	}
}

func TestKafkaUsageRanksXTubeKafka(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Where did João use Kafka?",
			10,
		)

	candidate, found :=
		result.Candidate(
			"project-xtube-kafka",
		)

	if !found {
		t.Fatalf(
			"expected x-tube kafka fact, got %v",
			result.Candidates,
		)
	}

	if candidate.Rank > 3 {
		t.Fatalf(
			"expected x-tube kafka in top 3, got rank %d",
			candidate.Rank,
		)
	}
}

func TestEmailRanksContactFactFirst(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual o email do João?",
			10,
		)

	if len(result.Candidates) == 0 {
		t.Fatal(
			"expected candidates",
		)
	}

	if result.Candidates[0].FactID !=
		"profile-email" {
		t.Fatalf(
			"expected profile-email first, got %s",
			result.Candidates[0].FactID,
		)
	}
}

func TestEnglishEmailRanksContactFactFirst(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"What is João's email address?",
			10,
		)

	if len(result.Candidates) == 0 {
		t.Fatal(
			"expected candidates",
		)
	}

	if result.Candidates[0].FactID !=
		"profile-email" {
		t.Fatalf(
			"expected profile-email first, got %s",
			result.Candidates[0].FactID,
		)
	}
}

func TestGoCapabilityRanksSkillEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Ele sabe Go?",
			10,
		)

	candidate, found :=
		result.Candidate(
			"skill-go",
		)

	if !found {
		t.Fatalf(
			"expected skill-go, got %v",
			result.Candidates,
		)
	}

	if candidate.Rank > 5 {
		t.Fatalf(
			"expected skill-go in top 5, got %d",
			candidate.Rank,
		)
	}
}

func TestKubernetesTypoRanksAuronixEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Ele tem experiência com Kubernets?",
			10,
		)

	candidate, found :=
		result.Candidate(
			"project-auronix-kubernetes",
		)

	if !found {
		t.Fatalf(
			"expected auronix kubernetes, got %v",
			result.Candidates,
		)
	}

	if candidate.Rank > 5 {
		t.Fatalf(
			"expected auronix kubernetes in top 5, got %d",
			candidate.Rank,
		)
	}
}

func TestDistributedSystemsRanksAuronixEvidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual projeto demonstra sistemas distribuídos?",
			10,
		)

	candidate, found :=
		result.Candidate(
			"project-auronix-distributed",
		)

	if !found {
		t.Fatalf(
			"expected auronix distributed evidence, got %v",
			result.Candidates,
		)
	}

	if candidate.Rank > 5 {
		t.Fatalf(
			"expected auronix distributed in top 5, got %d",
			candidate.Rank,
		)
	}
}

func TestRRFUsesRankInsteadOfRawRetrieverScore(
	t *testing.T,
) {
	fusion :=
		NewRRFFusion(
			DefaultRRFConfig(),
		)

	retrievalResult :=
		retrieval.Result{
			Query: domain.Query{
				Language: domain.LanguagePortuguese,
			},
			Rankings: []retrieval.Ranking{
				{
					Source: retrieval.SourceLexical,
					Candidates: []retrieval.Candidate{
						{
							FactID: "fact-a",
							Score:  1000,
						},
						{
							FactID: "fact-b",
							Score:  999,
						},
					},
				},
				{
					Source: retrieval.SourceConcept,
					Candidates: []retrieval.Candidate{
						{
							FactID: "fact-b",
							Score:  0.9,
						},
						{
							FactID: "fact-a",
							Score:  0.1,
						},
					},
				},
			},
		}

	values :=
		fusion.Fuse(
			retrievalResult,
			10,
		)

	if len(values) != 2 {
		t.Fatalf(
			"expected 2 candidates, got %d",
			len(values),
		)
	}

	if values[0].FactID !=
		"fact-b" {
		t.Fatalf(
			"expected fact-b first because concept source has greater weight, got %s",
			values[0].FactID,
		)
	}
}

func TestMultiSourceCandidatePreservesSources(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Me fale sobre o GGCompress",
			20,
		)

	foundMultiSource := false

	for _, candidate := range result.Candidates {
		if len(candidate.Sources) > 1 {
			foundMultiSource = true
			break
		}
	}

	if !foundMultiSource {
		t.Fatal(
			"expected at least one multi-source candidate",
		)
	}
}

func TestScoresAreNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual projeto demonstra performance?",
			30,
		)

	for _, candidate := range result.Candidates {
		if candidate.Score < 0 ||
			candidate.Score > 1 {
			t.Fatalf(
				"invalid score %f for %s",
				candidate.Score,
				candidate.FactID,
			)
		}

		if candidate.FusionScore < 0 ||
			candidate.FusionScore > 1 {
			t.Fatalf(
				"invalid fusion score %f",
				candidate.FusionScore,
			)
		}

		if candidate.FeatureScore < 0 ||
			candidate.FeatureScore > 1 {
			t.Fatalf(
				"invalid feature score %f",
				candidate.FeatureScore,
			)
		}
	}
}

func TestRankingIsSorted(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual projeto demonstra performance e concorrência?",
			30,
		)

	for index := 1; index < len(result.Candidates); index++ {
		previous :=
			result.Candidates[index-1]

		current :=
			result.Candidates[index]

		if current.Score >
			previous.Score {
			t.Fatalf(
				"ranking is not sorted at %d",
				index,
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

func TestRankingContainsNoDuplicates(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Me fale sobre o Auronix",
			50,
		)

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, candidate := range result.Candidates {
		if _, exists :=
			seen[candidate.FactID]; exists {
			t.Fatalf(
				"duplicated fact %s",
				candidate.FactID,
			)
		}

		seen[candidate.FactID] =
			struct{}{}
	}
}

func TestUnknownQueryProducesEmptyRanking(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"xyzabc123",
			20,
		)

	if len(result.Candidates) != 0 {
		t.Fatalf(
			"expected empty ranking, got %v",
			result.Candidates,
		)
	}
}

func TestSignalsAreRecorded(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"Qual projeto demonstra concorrência?",
			10,
		)

	if len(result.Candidates) == 0 {
		t.Fatal(
			"expected candidates",
		)
	}

	candidate :=
		result.Candidates[0]

	if _, found :=
		candidate.Signal(
			SignalFusion,
		); !found {
		t.Fatal(
			"expected fusion signal",
		)
	}

	if _, found :=
		candidate.Signal(
			SignalImportance,
		); !found {
		t.Fatal(
			"expected importance signal",
		)
	}

	if _, found :=
		candidate.Signal(
			SignalIntent,
		); !found {
		t.Fatal(
			"expected intent signal",
		)
	}
}

func TestRankingLimit(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		rankQuestion(
			t,
			environment,
			"experiência",
			3,
		)

	if len(result.Candidates) > 3 {
		t.Fatalf(
			"expected maximum 3 candidates, got %d",
			len(result.Candidates),
		)
	}
}

func containsFact(
	result Result,
	expected domain.FactID,
) bool {
	for _, candidate := range result.Candidates {
		if candidate.FactID ==
			expected {
			return true
		}
	}

	return false
}
