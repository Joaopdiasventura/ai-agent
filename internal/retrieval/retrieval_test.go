package retrieval

import (
	"testing"

	"ai-agent/internal/domain"
	searchindex "ai-agent/internal/index"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
	"ai-agent/internal/query"
)

type testEnvironment struct {
	base      *knowledge.Knowledge
	analyzer  *query.Analyzer
	retriever *HybridRetriever
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
		NewHybridRetriever(
			base,
			currentIndex,
		)

	if err != nil {
		t.Fatal(err)
	}

	return testEnvironment{
		base:      base,
		analyzer:  analyzer,
		retriever: retriever,
	}
}

func TestHybridRetrieverBuilds(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Me fale sobre o GGCompress",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			10,
		)

	if len(result.Rankings) != 4 {
		t.Fatalf(
			"expected 4 rankings, got %d",
			len(result.Rankings),
		)
	}
}

func TestLexicalRetrieverFindsGGCompress(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Me fale sobre o GGCompress",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			10,
		)

	ranking, found :=
		result.Ranking(
			SourceLexical,
		)

	if !found {
		t.Fatal(
			"expected lexical ranking",
		)
	}

	if len(ranking.Candidates) == 0 {
		t.Fatal(
			"expected lexical candidates",
		)
	}

	foundGGCompress := false

	for _, candidate := range ranking.Candidates {
		fact, found :=
			environment.base.Fact(
				candidate.FactID,
			)

		if !found {
			continue
		}

		if fact.Subject ==
			knowledge.EntityGGCompress {
			foundGGCompress = true
			break
		}

		for _, contextID := range fact.Context {
			if contextID ==
				knowledge.EntityGGCompress {
				foundGGCompress = true
				break
			}
		}
	}

	if !foundGGCompress {
		t.Fatal(
			"expected ggcompress facts",
		)
	}
}

func TestEntityRetrieverFindsKafkaUsage(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Where did João use Kafka?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			10,
		)

	ranking, found :=
		result.Ranking(
			SourceEntity,
		)

	if !found {
		t.Fatal(
			"expected entity ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-xtube-kafka",
	) {
		t.Fatalf(
			"expected x-tube kafka fact, got %v",
			ranking.Candidates,
		)
	}
}

func TestConceptRetrieverFindsConcurrency(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto demonstra concorrência?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	ranking, found :=
		result.Ranking(
			SourceConcept,
		)

	if !found {
		t.Fatal(
			"expected concept ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-ggcompress-concurrency",
	) {
		t.Fatalf(
			"expected ggcompress concurrency fact, got %v",
			ranking.Candidates,
		)
	}
}

func TestConceptRetrieverFindsGoroutinesThroughConcurrency(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto demonstra concorrência?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	ranking, found :=
		result.Ranking(
			SourceConcept,
		)

	if !found {
		t.Fatal(
			"expected concept ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-ggcompress-goroutines",
	) {
		t.Fatalf(
			"expected ggcompress goroutines fact, got %v",
			ranking.Candidates,
		)
	}
}

func TestFuzzyRetrieverFindsKubernetesTypo(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Ele tem experiência com Kubernets?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	ranking, found :=
		result.Ranking(
			SourceFuzzy,
		)

	if !found {
		t.Fatal(
			"expected fuzzy ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-auronix-kubernetes",
	) {
		t.Fatalf(
			"expected auronix kubernetes fact, got %v",
			ranking.Candidates,
		)
	}
}

func TestFuzzyRetrieverFindsConcurrencyTypo(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto demonstra concorencia?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	ranking, found :=
		result.Ranking(
			SourceFuzzy,
		)

	if !found {
		t.Fatal(
			"expected fuzzy ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-ggcompress-concurrency",
	) {
		t.Fatalf(
			"expected ggcompress concurrency fact, got %v",
			ranking.Candidates,
		)
	}
}

func TestEnglishRetrieval(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Which project demonstrates concurrency?",
		)

	if currentQuery.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			currentQuery.Language,
		)
	}

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	if result.Empty() {
		t.Fatal(
			"expected retrieval results",
		)
	}

	ranking, found :=
		result.Ranking(
			SourceConcept,
		)

	if !found {
		t.Fatal(
			"expected concept ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-ggcompress-concurrency",
	) {
		t.Fatal(
			"expected ggcompress concurrency",
		)
	}
}

func TestBM25FindsThroughput(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto atingiu 1.23 GB/s?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	ranking, found :=
		result.Ranking(
			SourceLexical,
		)

	if !found {
		t.Fatal(
			"expected lexical ranking",
		)
	}

	if !rankingContainsFact(
		ranking,
		"project-ggcompress-throughput",
	) {
		t.Fatalf(
			"expected throughput fact, got %v",
			ranking.Candidates,
		)
	}
}

func TestGoCapabilityRetrieval(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Ele sabe Go?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	entityRanking, found :=
		result.Ranking(
			SourceEntity,
		)

	if !found {
		t.Fatal(
			"expected entity ranking",
		)
	}

	if !rankingContainsFact(
		entityRanking,
		"skill-go",
	) {
		t.Fatalf(
			"expected go skill fact, got %v",
			entityRanking.Candidates,
		)
	}
}

func TestAuronixDistributedSystems(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto demonstra sistemas distribuídos?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	conceptRanking, found :=
		result.Ranking(
			SourceConcept,
		)

	if !found {
		t.Fatal(
			"expected concept ranking",
		)
	}

	if !rankingContainsFact(
		conceptRanking,
		"project-auronix-distributed",
	) {
		t.Fatalf(
			"expected auronix distributed fact, got %v",
			conceptRanking.Candidates,
		)
	}
}

func TestUnknownQueryReturnsNoUsefulResults(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"xyzabc123",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	if !result.Empty() {
		t.Fatalf(
			"expected empty result, got %v",
			result.Rankings,
		)
	}
}

func TestRankingHasNoDuplicateFacts(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto demonstra performance e concorrência?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			50,
		)

	for _, ranking := range result.Rankings {
		seen := make(
			map[domain.FactID]struct{},
		)

		for _, candidate := range ranking.Candidates {
			if _, exists :=
				seen[candidate.FactID]; exists {
				t.Fatalf(
					"duplicate fact %s in %s ranking",
					candidate.FactID,
					ranking.Source,
				)
			}

			seen[candidate.FactID] =
				struct{}{}
		}
	}
}

func TestRankingsAreSorted(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Me fale sobre o GGCompress",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			50,
		)

	for _, ranking := range result.Rankings {
		for current := 1; current <
			len(ranking.Candidates); current++ {
			previous :=
				ranking.Candidates[current-1]

			value :=
				ranking.Candidates[current]

			if value.Score >
				previous.Score {
				t.Fatalf(
					"ranking %s is not sorted",
					ranking.Source,
				)
			}
		}
	}
}

func TestResultFactIDsAreUnique(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"Qual projeto demonstra concorrência?",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			20,
		)

	factIDs :=
		result.FactIDs()

	seen := make(
		map[domain.FactID]struct{},
	)

	for _, factID := range factIDs {
		if _, exists :=
			seen[factID]; exists {
			t.Fatalf(
				"duplicate fact id %s",
				factID,
			)
		}

		seen[factID] =
			struct{}{}
	}
}

func TestDefaultLimit(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	currentQuery :=
		environment.analyzer.Analyze(
			"experiência",
		)

	result :=
		environment.retriever.Search(
			currentQuery,
			0,
		)

	for _, ranking := range result.Rankings {
		if len(ranking.Candidates) >
			DefaultLimit {
			t.Fatalf(
				"expected maximum %d candidates, got %d",
				DefaultLimit,
				len(ranking.Candidates),
			)
		}
	}
}

func rankingContainsFact(
	ranking Ranking,
	expected domain.FactID,
) bool {
	for _, candidate := range ranking.Candidates {
		if candidate.FactID ==
			expected {
			return true
		}
	}

	return false
}
