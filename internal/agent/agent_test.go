package agent

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/planning"
	"ai-agent/internal/reasoning"
)

func createService(
	t *testing.T,
) *Service {
	t.Helper()

	service, err :=
		New()

	if err != nil {
		t.Fatal(err)
	}

	return service
}

func TestServiceBuilds(
	t *testing.T,
) {
	service :=
		createService(t)

	if service == nil {
		t.Fatal(
			"expected service",
		)
	}
}

func TestEmptyQuestionFails(
	t *testing.T,
) {
	service :=
		createService(t)

	_, err :=
		service.Answer("   ")

	if !errors.Is(
		err,
		ErrEmptyQuestion,
	) {
		t.Fatalf(
			"expected ErrEmptyQuestion, got %v",
			err,
		)
	}
}

func TestEmailAnswer(
	t *testing.T,
) {
	service :=
		createService(t)

	result, err :=
		service.Answer(
			"Qual o email do João?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if !result.HasResponse {
		t.Fatal(
			"expected response",
		)
	}

	if result.Language !=
		domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			result.Language,
		)
	}

	if !strings.Contains(
		strings.ToLower(
			result.Response,
		),
		"joaopdias.dev@gmail.com",
	) {
		t.Fatalf(
			"expected email, got %q",
			result.Response,
		)
	}
}

func TestGoCapabilityAnswer(
	t *testing.T,
) {
	service :=
		createService(t)

	result, err :=
		service.Answer(
			"Ele sabe Go?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if !result.HasResponse {
		t.Fatal(
			"expected response",
		)
	}

	if !strings.Contains(
		result.Response,
		"Go",
	) {
		t.Fatalf(
			"expected Go, got %q",
			result.Response,
		)
	}

	if result.Confidence < 0 ||
		result.Confidence > 1 {
		t.Fatalf(
			"invalid confidence %f",
			result.Confidence,
		)
	}
}

func TestUnknownCapabilityHasNoResponse(
	t *testing.T,
) {
	service :=
		createService(t)

	result, err :=
		service.Answer(
			"Ele sabe Rust?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if result.HasResponse {
		t.Fatalf(
			"expected no response, got %q",
			result.Response,
		)
	}

	if result.Response != "" {
		t.Fatalf(
			"expected empty public response, got %q",
			result.Response,
		)
	}
}

func TestUnknownQueryHasNoResponse(
	t *testing.T,
) {
	service :=
		createService(t)

	result, err :=
		service.Answer(
			"xyzabc123",
		)

	if err != nil {
		t.Fatal(err)
	}

	if result.HasResponse {
		t.Fatalf(
			"expected no response, got %q",
			result.Response,
		)
	}
}

func TestEnglishKafkaAnswer(
	t *testing.T,
) {
	service :=
		createService(t)

	result, err :=
		service.Answer(
			"Where did João use Kafka?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if !result.HasResponse {
		t.Fatal(
			"expected response",
		)
	}

	if result.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			result.Language,
		)
	}

	if !strings.Contains(
		result.Response,
		"Kafka",
	) {
		t.Fatalf(
			"expected Kafka, got %q",
			result.Response,
		)
	}

	if !strings.Contains(
		result.Response,
		"X Tube",
	) {
		t.Fatalf(
			"expected X Tube, got %q",
			result.Response,
		)
	}
}

func TestDebugContainsEntirePipeline(
	t *testing.T,
) {
	service :=
		createService(t)

	debug, err :=
		service.Debug(
			"Qual projeto melhor demonstra concorrência?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if debug.Question == "" {
		t.Fatal(
			"expected question",
		)
	}

	if debug.Query.Original == "" {
		t.Fatal(
			"expected analyzed query",
		)
	}

	if len(
		debug.Retrieval.Rankings,
	) == 0 {
		t.Fatal(
			"expected retrieval rankings",
		)
	}

	if len(
		debug.Ranking.Candidates,
	) == 0 {
		t.Fatal(
			"expected ranked candidates",
		)
	}

	if debug.Reasoning.
		Conclusion.
		Type ==
		reasoning.ConclusionUnknown {
		t.Fatal(
			"expected reasoning conclusion",
		)
	}

	if debug.Plan.Type ==
		planning.PlanTypeUnknown {
		t.Fatal(
			"expected answer plan",
		)
	}

	if debug.Generation.Empty() {
		t.Fatal(
			"expected generated answer",
		)
	}

	if len(
		debug.Confidence.Signals,
	) == 0 {
		t.Fatal(
			"expected confidence signals",
		)
	}
}

func TestDebugComparisonChoosesGGCompress(
	t *testing.T,
) {
	service :=
		createService(t)

	debug, err :=
		service.Debug(
			"Qual projeto melhor demonstra concorrência?",
		)

	if err != nil {
		t.Fatal(err)
	}

	group, found :=
		debug.Reasoning.TopGroup()

	if !found {
		t.Fatal(
			"expected project group",
		)
	}

	if group.EntityID !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress, got %s",
			group.EntityID,
		)
	}

	lead, found :=
		debug.Plan.Section(
			planning.SectionLead,
		)

	if !found ||
		len(lead.Items) == 0 {
		t.Fatal(
			"expected lead section",
		)
	}

	if lead.Items[0].EntityID !=
		knowledge.EntityGGCompress {
		t.Fatalf(
			"expected ggcompress plan winner, got %s",
			lead.Items[0].EntityID,
		)
	}
}

func TestDebugUnknownCapabilityShowsAbstention(
	t *testing.T,
) {
	service :=
		createService(t)

	debug, err :=
		service.Debug(
			"Ele sabe Rust?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if debug.Result.HasResponse {
		t.Fatal(
			"expected public refusal",
		)
	}

	if debug.Plan.Status !=
		planning.PlanStatusAbstain {
		t.Fatalf(
			"expected abstain plan, got %s",
			debug.Plan.Status,
		)
	}

	if debug.Reasoning.
		Conclusion.
		Status !=
		reasoning.SupportInsufficientEvidence {
		t.Fatalf(
			"expected insufficient evidence, got %s",
			debug.Reasoning.
				Conclusion.
				Status,
		)
	}

	if debug.Confidence.Mode !=
		confidence.ModeAbstention {
		t.Fatalf(
			"expected abstention confidence, got %s",
			debug.Confidence.Mode,
		)
	}

	if debug.Generation.Empty() {
		t.Fatal(
			"debug should preserve generated abstention explanation",
		)
	}

	if len(
		debug.Generation.FactIDs,
	) != 0 {
		t.Fatalf(
			"expected no factual claims, got %v",
			debug.Generation.FactIDs,
		)
	}
}

func TestDebugAndAnswerAreConsistent(
	t *testing.T,
) {
	service :=
		createService(t)

	question :=
		"Ele sabe Go?"

	result, err :=
		service.Answer(
			question,
		)

	if err != nil {
		t.Fatal(err)
	}

	debug, err :=
		service.Debug(
			question,
		)

	if err != nil {
		t.Fatal(err)
	}

	if result.Response !=
		debug.Result.Response {
		t.Fatalf(
			"response mismatch: %q != %q",
			result.Response,
			debug.Result.Response,
		)
	}

	if result.HasResponse !=
		debug.Result.HasResponse {
		t.Fatal(
			"has response mismatch",
		)
	}

	if result.Language !=
		debug.Result.Language {
		t.Fatal(
			"language mismatch",
		)
	}

	if result.Confidence !=
		debug.Result.Confidence {
		t.Fatalf(
			"confidence mismatch: %f != %f",
			result.Confidence,
			debug.Result.Confidence,
		)
	}
}

func TestServiceIsStatelessBetweenRequests(
	t *testing.T,
) {
	service :=
		createService(t)

	first, err :=
		service.Answer(
			"Me fale sobre o GGCompress",
		)

	if err != nil {
		t.Fatal(err)
	}

	if !first.HasResponse {
		t.Fatal(
			"expected first response",
		)
	}

	second, err :=
		service.Answer(
			"Me fale sobre o Auronix",
		)

	if err != nil {
		t.Fatal(err)
	}

	if !second.HasResponse {
		t.Fatal(
			"expected second response",
		)
	}

	if !strings.Contains(
		first.Response,
		"GGCompress",
	) {
		t.Fatalf(
			"expected GGCompress, got %q",
			first.Response,
		)
	}

	if !strings.Contains(
		second.Response,
		"Auronix",
	) {
		t.Fatalf(
			"expected Auronix, got %q",
			second.Response,
		)
	}

	if strings.HasPrefix(
		second.Response,
		"Sobre GGCompress",
	) {
		t.Fatalf(
			"previous request leaked into second response: %q",
			second.Response,
		)
	}
}

func TestRepeatedCallsAreDeterministic(
	t *testing.T,
) {
	service :=
		createService(t)

	question :=
		"Qual projeto melhor demonstra concorrência?"

	first, err :=
		service.Answer(
			question,
		)

	if err != nil {
		t.Fatal(err)
	}

	for index := 0; index < 10; index++ {
		current, err :=
			service.Answer(
				question,
			)

		if err != nil {
			t.Fatal(err)
		}

		if current.Response !=
			first.Response {
			t.Fatalf(
				"non deterministic response at iteration %d",
				index,
			)
		}

		if current.HasResponse !=
			first.HasResponse {
			t.Fatalf(
				"non deterministic support at iteration %d",
				index,
			)
		}

		if current.Confidence !=
			first.Confidence {
			t.Fatalf(
				"non deterministic confidence at iteration %d",
				index,
			)
		}
	}
}

func TestConcurrentCalls(
	t *testing.T,
) {
	service :=
		createService(t)

	questions := []string{
		"Ele sabe Go?",
		"Qual projeto melhor demonstra concorrência?",
		"Where did João use Kafka?",
		"Me fale sobre o GGCompress",
		"Qual o email do João?",
	}

	var waitGroup sync.WaitGroup

	errChannel :=
		make(
			chan error,
			len(questions),
		)

	for _, question := range questions {
		waitGroup.Add(1)

		go func(
			value string,
		) {
			defer waitGroup.Done()

			result, err :=
				service.Answer(value)

			if err != nil {
				errChannel <- err
				return
			}

			if !result.HasResponse {
				errChannel <- errors.New(
					"expected supported response",
				)
			}
		}(question)
	}

	waitGroup.Wait()

	close(errChannel)

	for err := range errChannel {
		t.Fatal(err)
	}
}

func TestResultConfidenceIsNormalized(
	t *testing.T,
) {
	service :=
		createService(t)

	questions := []string{
		"Ele sabe Go?",
		"Ele sabe Rust?",
		"Where did João use Kafka?",
		"Qual projeto melhor demonstra concorrência?",
		"Me fale sobre o GGCompress",
		"xyzabc123",
	}

	for _, question := range questions {
		result, err :=
			service.Answer(
				question,
			)

		if err != nil {
			t.Fatal(err)
		}

		if result.Confidence < 0 ||
			result.Confidence > 1 {
			t.Fatalf(
				"question %q produced invalid confidence %f",
				question,
				result.Confidence,
			)
		}
	}
}

func TestSupportedAnswerUsesClaimConfidence(
	t *testing.T,
) {
	service :=
		createService(t)

	debug, err :=
		service.Debug(
			"Ele sabe Go?",
		)

	if err != nil {
		t.Fatal(err)
	}

	if !debug.Result.HasResponse {
		t.Fatal(
			"expected supported answer",
		)
	}

	if debug.Confidence.Mode !=
		confidence.ModeClaim {
		t.Fatalf(
			"expected claim confidence, got %s",
			debug.Confidence.Mode,
		)
	}
}

func TestUnknownAgeAbstains(
	t *testing.T,
) {
	service := createService(t)

	debug, err := service.Debug("quantos anos tem João?")

	if err != nil {
		t.Fatal(err)
	}

	if debug.Result.HasResponse {
		t.Fatalf("expected no public response, got %q", debug.Result.Response)
	}

	if debug.Reasoning.Conclusion.Status != reasoning.SupportInsufficientEvidence {
		t.Fatalf("expected insufficient evidence, got %s", debug.Reasoning.Conclusion.Status)
	}

	if debug.Plan.Status != planning.PlanStatusAbstain {
		t.Fatalf("expected abstain, got %s", debug.Plan.Status)
	}

	if len(debug.Generation.FactIDs) != 0 {
		t.Fatalf("expected no generated facts, got %v", debug.Generation.FactIDs)
	}
}

func TestUnknownLanguagePreferenceAbstains(
	t *testing.T,
) {
	service := createService(t)

	result, err := service.Answer("João tem preferência de linguagem?")

	if err != nil {
		t.Fatal(err)
	}

	if result.HasResponse {
		t.Fatalf("expected no response, got %q", result.Response)
	}
}

func TestProgrammingLanguageList(
	t *testing.T,
) {
	service := createService(t)

	debug, err := service.Debug("Quais linguagens ele sabe?")

	if err != nil {
		t.Fatal(err)
	}

	if !debug.Result.HasResponse {
		t.Fatal("expected response")
	}

	if debug.Query.Intent != domain.IntentList {
		t.Fatalf("expected list intent, got %s", debug.Query.Intent)
	}

	for _, expected := range []string{"JavaScript", "TypeScript", "Java", "Go"} {
		if !strings.Contains(debug.Result.Response, expected) {
			t.Fatalf("expected %s in %q", expected, debug.Result.Response)
		}
	}

	for _, forbidden := range []string{"Português", "Inglês"} {
		if strings.Contains(debug.Result.Response, forbidden) {
			t.Fatalf("did not expect %s in %q", forbidden, debug.Result.Response)
		}
	}
}

func TestFrameworkList(
	t *testing.T,
) {
	service := createService(t)

	result, err := service.Answer("Quais frameworks ele sabe?")

	if err != nil {
		t.Fatal(err)
	}

	if !result.HasResponse {
		t.Fatal("expected response")
	}

	for _, expected := range []string{"Angular", "React", "Next.js", "Spring Boot", "NestJS"} {
		if !strings.Contains(result.Response, expected) {
			t.Fatalf("expected %s in %q", expected, result.Response)
		}
	}

	for _, forbidden := range []string{"PostgreSQL", "MongoDB", "JavaScript", "TypeScript"} {
		if strings.Contains(result.Response, forbidden) {
			t.Fatalf("did not expect %s in %q", forbidden, result.Response)
		}
	}
}

func TestTechnologyList(
	t *testing.T,
) {
	service := createService(t)

	result, err := service.Answer("quais tecnologias ele sabe?")

	if err != nil {
		t.Fatal(err)
	}

	if !result.HasResponse {
		t.Fatal("expected response")
	}

	for _, expected := range []string{"Java", "Go", "Docker"} {
		if !strings.Contains(result.Response, expected) {
			t.Fatalf("expected %s in %q", expected, result.Response)
		}
	}
}

func TestFuzzyTechnologyCapabilities(
	t *testing.T,
) {
	service := createService(t)

	cases := map[string]string{
		"ele sabe kubernts?":     "Kubernetes",
		"ele sabe dcker?":        "Docker",
		"ele sabe dockr?":        "Docker",
		"ele sabe jva?":          "Java",
		"ele sabe javascrit?":    "JavaScript",
		"ele sabe typescrit?":    "TypeScript",
		"ele sabe nodjs?":        "Node.js",
		"ele sabe postgre?":      "PostgreSQL",
		"Does João know Docker?": "Docker",
	}

	for question, expected := range cases {
		result, err := service.Answer(question)

		if err != nil {
			t.Fatal(err)
		}

		if !result.HasResponse {
			t.Fatalf("question %q expected response", question)
		}

		if !strings.Contains(result.Response, expected) {
			t.Fatalf("question %q expected %s in %q", question, expected, result.Response)
		}
	}
}

func TestAbstractCapabilities(
	t *testing.T,
) {
	service := createService(t)

	for _, question := range []string{"ele sabe backend?", "ele sabe fulstack?"} {
		result, err := service.Answer(question)

		if err != nil {
			t.Fatal(err)
		}

		if !result.HasResponse {
			t.Fatalf("question %q expected response", question)
		}
	}
}

func TestHumanLanguageList(
	t *testing.T,
) {
	service := createService(t)

	result, err := service.Answer("quais idiomas ele fala?")

	if err != nil {
		t.Fatal(err)
	}

	if !result.HasResponse {
		t.Fatal("expected response")
	}

	for _, expected := range []string{"Português", "Inglês"} {
		if !strings.Contains(result.Response, expected) {
			t.Fatalf("expected %s in %q", expected, result.Response)
		}
	}

	for _, forbidden := range []string{"Java", "Go", "JavaScript"} {
		if strings.Contains(result.Response, forbidden) {
			t.Fatalf("did not expect %s in %q", forbidden, result.Response)
		}
	}
}

func TestEnglishListsAndUnknownAge(
	t *testing.T,
) {
	service := createService(t)

	for _, question := range []string{
		"What programming languages does João know?",
		"Which frameworks does João know?",
	} {
		result, err := service.Answer(question)

		if err != nil {
			t.Fatal(err)
		}

		if !result.HasResponse {
			t.Fatalf("question %q expected response", question)
		}
	}

	result, err := service.Answer("How old is João?")

	if err != nil {
		t.Fatal(err)
	}

	if result.HasResponse {
		t.Fatalf("expected no response, got %q", result.Response)
	}
}
