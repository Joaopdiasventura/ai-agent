package confidence

import (
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/generation"
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
	base       *knowledge.Knowledge
	analyzer   *query.Analyzer
	retriever  *retrieval.HybridRetriever
	ranker     *ranking.Ranker
	reasoner   *reasoning.Reasoner
	planner    *planning.Planner
	generator  *generation.Generator
	calculator *Calculator
}

type pipelineResult struct {
	query      domain.Query
	retrieval  retrieval.Result
	ranking    ranking.Result
	reasoning  reasoning.Result
	plan       planning.Plan
	answer     generation.Answer
	confidence Result
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
		base:       base,
		analyzer:   analyzer,
		retriever:  retriever,
		ranker:     ranker,
		reasoner:   reasoner,
		planner:    planner,
		generator:  generation.New(),
		calculator: New(),
	}
}

func runPipeline(
	t *testing.T,
	environment testEnvironment,
	value string,
) pipelineResult {
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

	plan :=
		environment.planner.Plan(
			reasoned,
		)

	material, err :=
		generation.Materialize(
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

	input := Input{
		Query:     currentQuery,
		Retrieval: retrieved,
		Ranking:   ranked,
		Reasoning: reasoned,
		Plan:      plan,
		Answer:    answer,
	}

	result :=
		environment.calculator.Assess(
			input,
		)

	return pipelineResult{
		query:      currentQuery,
		retrieval:  retrieved,
		ranking:    ranked,
		reasoning:  reasoned,
		plan:       plan,
		answer:     answer,
		confidence: result,
	}
}

func TestCalculatorBuilds(
	t *testing.T,
) {
	calculator :=
		New()

	if calculator == nil {
		t.Fatal(
			"expected calculator",
		)
	}
}

func TestDefaultConfigIsValid(
	t *testing.T,
) {
	if err :=
		DefaultConfig().
			Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestGoCapabilityUsesClaimMode(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Go?",
		)

	if result.confidence.Mode !=
		ModeClaim {
		t.Fatalf(
			"expected claim mode, got %s",
			result.confidence.Mode,
		)
	}

	if result.confidence.Score <
		0.55 {
		t.Fatalf(
			"expected meaningful confidence, got %f",
			result.confidence.Score,
		)
	}
}

func TestRustCapabilityUsesAbstentionMode(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Rust?",
		)

	if result.confidence.Mode !=
		ModeAbstention {
		t.Fatalf(
			"expected abstention mode, got %s",
			result.confidence.Mode,
		)
	}

	if result.confidence.Score <
		0.7 {
		t.Fatalf(
			"expected strong abstention confidence, got %f",
			result.confidence.Score,
		)
	}
}

func TestRustEvidenceAbsenceIsHigh(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Rust?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalEvidenceAbsence,
		)

	if !found {
		t.Fatal(
			"expected evidence absence signal",
		)
	}

	if signal.Score < 0.8 {
		t.Fatalf(
			"expected high evidence absence, got %f",
			signal.Score,
		)
	}
}

func TestGoAnswerIsGrounded(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Go?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalAnswerGrounding,
		)

	if !found {
		t.Fatal(
			"expected answer grounding signal",
		)
	}

	if signal.Score != 1 {
		t.Fatalf(
			"expected perfect answer grounding, got %f",
			signal.Score,
		)
	}
}

func TestComparisonIsClaimMode(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	if result.confidence.Mode !=
		ModeClaim {
		t.Fatalf(
			"expected claim mode, got %s",
			result.confidence.Mode,
		)
	}

	if result.confidence.Score <= 0 {
		t.Fatalf(
			"expected positive confidence, got %f",
			result.confidence.Score,
		)
	}
}

func TestComparisonHasSeparationSignal(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalSeparation,
		)

	if !found {
		t.Fatal(
			"expected separation signal",
		)
	}

	if !signal.Applicable {
		t.Fatal(
			"expected separation to be applicable",
		)
	}

	if signal.Score < 0 ||
		signal.Score > 1 {
		t.Fatalf(
			"invalid separation %f",
			signal.Score,
		)
	}
}

func TestRetrievalAgreementIsNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalRetrievalAgreement,
		)

	if !found {
		t.Fatal(
			"expected retrieval agreement",
		)
	}

	if signal.Applicable &&
		(signal.Score < 0 ||
			signal.Score > 1) {
		t.Fatalf(
			"invalid agreement %f",
			signal.Score,
		)
	}
}

func TestEvidenceStrengthIsNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Go?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalEvidenceStrength,
		)

	if !found {
		t.Fatal(
			"expected evidence strength",
		)
	}

	if !signal.Applicable {
		t.Fatal(
			"expected evidence strength to be applicable",
		)
	}

	if signal.Score < 0 ||
		signal.Score > 1 {
		t.Fatalf(
			"invalid evidence strength %f",
			signal.Score,
		)
	}
}

func TestDirectnessIsNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Where did João use Kafka?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalEvidenceDirectness,
		)

	if !found {
		t.Fatal(
			"expected directness signal",
		)
	}

	if !signal.Applicable {
		t.Fatal(
			"expected directness to be applicable",
		)
	}

	if signal.Score < 0 ||
		signal.Score > 1 {
		t.Fatalf(
			"invalid directness %f",
			signal.Score,
		)
	}
}

func TestAllConfidenceScoresAreNormalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	questions := []string{
		"Ele sabe Go?",
		"Ele sabe Rust?",
		"Where did João use Kafka?",
		"Qual projeto melhor demonstra concorrência?",
		"Me fale sobre o GGCompress",
		"Qual o email do João?",
		"xyzabc123",
	}

	for _, questionValue := range questions {
		result :=
			runPipeline(
				t,
				environment,
				questionValue,
			)

		if result.confidence.Score < 0 ||
			result.confidence.Score > 1 {
			t.Fatalf(
				"question %q produced invalid confidence %f",
				questionValue,
				result.confidence.Score,
			)
		}

		for _, signal := range result.confidence.Signals {
			if signal.Score < 0 ||
				signal.Score > 1 {
				t.Fatalf(
					"question %q produced invalid signal %s=%f",
					questionValue,
					signal.Name,
					signal.Score,
				)
			}
		}
	}
}

func TestConfidenceLevelMatchesScore(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Go?",
		)

	expected :=
		confidenceLevel(
			result.confidence.Score,
		)

	if result.confidence.Level !=
		expected {
		t.Fatalf(
			"expected %s, got %s",
			expected,
			result.confidence.Level,
		)
	}
}

func TestTamperedAnswerLowersGrounding(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Go?",
		)

	original :=
		result.confidence

	tamperedAnswer :=
		result.answer

	tamperedAnswer.FactIDs =
		append(
			tamperedAnswer.FactIDs,
			domain.FactID(
				"profile-email",
			),
		)

	tampered :=
		environment.calculator.Assess(
			Input{
				Query:     result.query,
				Retrieval: result.retrieval,
				Ranking:   result.ranking,
				Reasoning: result.reasoning,
				Plan:      result.plan,
				Answer:    tamperedAnswer,
			},
		)

	grounding, found :=
		tampered.Signal(
			SignalAnswerGrounding,
		)

	if !found {
		t.Fatal(
			"expected answer grounding",
		)
	}

	if grounding.Score >= 1 {
		t.Fatalf(
			"expected grounding penalty, got %f",
			grounding.Score,
		)
	}

	if tampered.Score >=
		original.Score {
		t.Fatalf(
			"expected tampered confidence %f below original %f",
			tampered.Score,
			original.Score,
		)
	}
}

func TestAbstentionAnswerWithFactIsPenalized(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Rust?",
		)

	tamperedAnswer :=
		result.answer

	tamperedAnswer.FactIDs =
		[]domain.FactID{
			"profile-email",
		}

	tampered :=
		environment.calculator.Assess(
			Input{
				Query:     result.query,
				Retrieval: result.retrieval,
				Ranking:   result.ranking,
				Reasoning: result.reasoning,
				Plan:      result.plan,
				Answer:    tamperedAnswer,
			},
		)

	signal, found :=
		tampered.Signal(
			SignalAnswerGrounding,
		)

	if !found {
		t.Fatal(
			"expected answer grounding",
		)
	}

	if signal.Score != 0 {
		t.Fatalf(
			"expected zero grounding for factual abstention, got %f",
			signal.Score,
		)
	}
}

func TestPlanGroundingIsPerfectForValidPipeline(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Qual projeto melhor demonstra concorrência?",
		)

	signal, found :=
		result.confidence.Signal(
			SignalPlanGrounding,
		)

	if !found {
		t.Fatal(
			"expected plan grounding",
		)
	}

	if signal.Score != 1 {
		t.Fatalf(
			"expected perfect plan grounding, got %f",
			signal.Score,
		)
	}
}

func TestEnglishQueryProducesConfidence(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Where did João use Kafka?",
		)

	if result.query.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			result.query.Language,
		)
	}

	if result.confidence.Score <= 0 {
		t.Fatalf(
			"expected positive confidence, got %f",
			result.confidence.Score,
		)
	}
}

func TestUnknownQueryAbstains(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"xyzabc123",
		)

	if result.confidence.Mode !=
		ModeAbstention {
		t.Fatalf(
			"expected abstention, got %s",
			result.confidence.Mode,
		)
	}

	if result.reasoning.Conclusion.Status !=
		reasoning.SupportInsufficientEvidence {
		t.Fatalf(
			"expected insufficient evidence, got %s",
			result.reasoning.Conclusion.Status,
		)
	}
}

func TestSignalsContainExpectedClaimSignals(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Go?",
		)

	expected := []SignalName{
		SignalQueryQuality,
		SignalRetrievalAgreement,
		SignalSeparation,
		SignalEvidenceStrength,
		SignalEvidenceDirectness,
		SignalSemanticCoverage,
		SignalPlanGrounding,
		SignalAnswerGrounding,
	}

	for _, name := range expected {
		if _, found :=
			result.confidence.Signal(
				name,
			); !found {
			t.Fatalf(
				"expected signal %s",
				name,
			)
		}
	}
}

func TestSignalsContainExpectedAbstentionSignals(
	t *testing.T,
) {
	environment :=
		createEnvironment(t)

	result :=
		runPipeline(
			t,
			environment,
			"Ele sabe Rust?",
		)

	expected := []SignalName{
		SignalQueryQuality,
		SignalEvidenceAbsence,
		SignalPlanGrounding,
		SignalAnswerGrounding,
	}

	for _, name := range expected {
		if _, found :=
			result.confidence.Signal(
				name,
			); !found {
			t.Fatalf(
				"expected signal %s",
				name,
			)
		}
	}
}
