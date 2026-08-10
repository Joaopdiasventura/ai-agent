package agent

import (
	"strings"

	"ai-agent/internal/confidence"
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

type Service struct {
	base       *knowledge.Knowledge
	analyzer   *query.Analyzer
	retriever  *retrieval.HybridRetriever
	ranker     *ranking.Ranker
	reasoner   *reasoning.Reasoner
	planner    *planning.Planner
	generator  *generation.Generator
	confidence *confidence.Calculator
}

func New() (*Service, error) {
	base, err :=
		knowledge.New()

	if err != nil {
		return nil, err
	}

	currentOntology, err :=
		ontology.New()

	if err != nil {
		return nil, err
	}

	currentIndex, err :=
		index.New(
			base,
			currentOntology,
		)

	if err != nil {
		return nil, err
	}

	analyzer, err :=
		query.NewAnalyzer(
			base,
			currentOntology,
		)

	if err != nil {
		return nil, err
	}

	retriever, err :=
		retrieval.NewHybridRetriever(
			base,
			currentIndex,
		)

	if err != nil {
		return nil, err
	}

	ranker, err :=
		ranking.New(base)

	if err != nil {
		return nil, err
	}

	reasoner, err :=
		reasoning.New(base)

	if err != nil {
		return nil, err
	}

	planner, err :=
		planning.New(base)

	if err != nil {
		return nil, err
	}

	return &Service{
		base:       base,
		analyzer:   analyzer,
		retriever:  retriever,
		ranker:     ranker,
		reasoner:   reasoner,
		planner:    planner,
		generator:  generation.New(),
		confidence: confidence.New(),
	}, nil
}

func (s *Service) Answer(
	question string,
) (Result, error) {
	currentExecution, err :=
		s.execute(question)

	if err != nil {
		return Result{}, err
	}

	hasResponse :=
		executionHasResponse(
			currentExecution,
		)

	response := ""

	if hasResponse {
		response =
			currentExecution.answer.Text
	}

	return Result{
		Response:        response,
		HasResponse:     hasResponse,
		Language:        currentExecution.answer.Language,
		Confidence:      currentExecution.confidence.Score,
		ConfidenceLevel: currentExecution.confidence.Level,
	}, nil
}

func (s *Service) execute(
	question string,
) (execution, error) {
	question =
		strings.TrimSpace(
			question,
		)

	if question == "" {
		return execution{},
			ErrEmptyQuestion
	}

	currentQuery :=
		s.analyzer.Analyze(
			question,
		)

	retrieved :=
		s.retriever.Search(
			currentQuery,
			80,
		)

	ranked :=
		s.ranker.Rank(
			retrieved,
			40,
		)

	reasoned :=
		s.reasoner.Reason(
			currentQuery,
			ranked,
		)

	plan :=
		s.planner.Plan(
			reasoned,
		)

	material, err :=
		generation.Materialize(
			plan,
			s.base,
		)

	if err != nil {
		return execution{}, err
	}

	answer, err :=
		s.generator.Generate(
			plan,
			material,
		)

	if err != nil {
		return execution{}, err
	}

	confidenceInput :=
		confidence.Input{
			Query:     currentQuery,
			Retrieval: retrieved,
			Ranking:   ranked,
			Reasoning: reasoned,
			Plan:      plan,
			Answer:    answer,
		}

	if err :=
		s.confidence.Validate(
			confidenceInput,
		); err != nil {
		return execution{}, err
	}

	confidenceResult :=
		s.confidence.Assess(
			confidenceInput,
		)

	return execution{
		question:   question,
		query:      currentQuery,
		retrieval:  retrieved,
		ranking:    ranked,
		reasoning:  reasoned,
		plan:       plan,
		answer:     answer,
		confidence: confidenceResult,
	}, nil
}

func executionHasResponse(
	current execution,
) bool {
	if current.plan.Status !=
		planning.PlanStatusReady {
		return false
	}

	if current.reasoning.
		Conclusion.
		Status !=
		reasoning.SupportSupported {
		return false
	}

	if current.answer.Empty() {
		return false
	}

	if current.confidence.Mode !=
		confidence.ModeClaim {
		return false
	}

	return true
}
