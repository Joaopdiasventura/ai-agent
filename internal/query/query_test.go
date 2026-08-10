package query

import (
	"testing"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
)

func createAnalyzer(
	t *testing.T,
) *Analyzer {
	t.Helper()

	base, err := knowledge.New()

	if err != nil {
		t.Fatal(err)
	}

	currentOntology, err := ontology.New()

	if err != nil {
		t.Fatal(err)
	}

	analyzer, err := NewAnalyzer(
		base,
		currentOntology,
	)

	if err != nil {
		t.Fatal(err)
	}

	return analyzer
}

func TestPortugueseComparisonQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Qual projeto melhor demonstra experiência com concorrência?",
	)

	if result.Language !=
		domain.LanguagePortuguese {
		t.Fatalf(
			"expected portuguese, got %s",
			result.Language,
		)
	}

	if result.Intent !=
		domain.IntentComparison {
		t.Fatalf(
			"expected comparison, got %s",
			result.Intent,
		)
	}

	if result.Target !=
		domain.QueryTargetProject {
		t.Fatalf(
			"expected project target, got %s",
			result.Target,
		)
	}

	if !result.HasConcept(
		ontology.ConceptConcurrency,
	) {
		t.Fatal(
			"expected concurrency concept",
		)
	}
}

func TestEnglishKafkaUsageQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Where did João use Kafka?",
	)

	if result.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			result.Language,
		)
	}

	if result.Intent !=
		domain.IntentTechnologyUsage {
		t.Fatalf(
			"expected technology usage, got %s",
			result.Intent,
		)
	}

	if !result.HasEntity(
		knowledge.EntityJoao,
	) {
		t.Fatal(
			"expected joao entity",
		)
	}

	if !result.HasEntity(
		knowledge.EntityKafka,
	) {
		t.Fatal(
			"expected kafka entity",
		)
	}

	if !result.HasConcept(
		ontology.ConceptMessaging,
	) {
		t.Fatal(
			"expected messaging concept",
		)
	}

	if !result.HasConcept(
		ontology.ConceptEventDriven,
	) {
		t.Fatal(
			"expected event-driven concept",
		)
	}
}

func TestPortugueseGGCompressOverview(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Me fale sobre o GGCompress",
	)

	if result.Intent !=
		domain.IntentOverview {
		t.Fatalf(
			"expected overview, got %s",
			result.Intent,
		)
	}

	if result.Target !=
		domain.QueryTargetProject {
		t.Fatalf(
			"expected project target, got %s",
			result.Target,
		)
	}

	if !result.HasEntity(
		knowledge.EntityGGCompress,
	) {
		t.Fatal(
			"expected ggcompress entity",
		)
	}
}

func TestEnglishGGCompressOverview(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Tell me about GGCompress",
	)

	if result.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english, got %s",
			result.Language,
		)
	}

	if result.Intent !=
		domain.IntentOverview {
		t.Fatalf(
			"expected overview, got %s",
			result.Intent,
		)
	}

	if !result.HasEntity(
		knowledge.EntityGGCompress,
	) {
		t.Fatal(
			"expected ggcompress entity",
		)
	}
}

func TestPortugueseContactQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Qual o email do João?",
	)

	if result.Intent !=
		domain.IntentContact {
		t.Fatalf(
			"expected contact, got %s",
			result.Intent,
		)
	}

	if result.Target !=
		domain.QueryTargetContact {
		t.Fatalf(
			"expected contact target, got %s",
			result.Target,
		)
	}

	if !result.HasConcept(
		ontology.ConceptEmail,
	) {
		t.Fatal(
			"expected email concept",
		)
	}
}

func TestPortugueseGoCapability(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Ele sabe Go?",
	)

	if result.Intent !=
		domain.IntentCapability {
		t.Fatalf(
			"expected capability, got %s",
			result.Intent,
		)
	}

	if result.Target !=
		domain.QueryTargetSkill {
		t.Fatalf(
			"expected skill target, got %s",
			result.Target,
		)
	}

	if !result.HasEntity(
		knowledge.EntityJoao,
	) {
		t.Fatal(
			"expected joao entity",
		)
	}

	if !result.HasEntity(
		knowledge.EntityGo,
	) {
		t.Fatal(
			"expected go entity",
		)
	}
}

func TestEnglishGoCapability(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Does João know Go?",
	)

	if result.Intent !=
		domain.IntentCapability {
		t.Fatalf(
			"expected capability, got %s",
			result.Intent,
		)
	}

	if !result.HasEntity(
		knowledge.EntityGo,
	) {
		t.Fatal(
			"expected go entity",
		)
	}
}

func TestEnglishGoVerbIsNotProgrammingLanguage(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Where did he go to school?",
	)

	if result.HasEntity(
		knowledge.EntityGo,
	) {
		t.Fatal(
			"did not expect Go programming language",
		)
	}

	if result.Intent !=
		domain.IntentEducation {
		t.Fatalf(
			"expected education, got %s",
			result.Intent,
		)
	}
}

func TestKubernetesTypo(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Ele tem experiência com Kubernets?",
	)

	if !result.HasEntity(
		knowledge.EntityKubernetes,
	) {
		t.Fatalf(
			"expected kubernetes entity, got %v",
			result.Entities,
		)
	}

	if !result.HasConcept(
		ontology.ConceptOrchestration,
	) {
		t.Fatal(
			"expected orchestration concept",
		)
	}
}

func TestPostgreSQLAlias(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Onde ele usou Postgres?",
	)

	if !result.HasEntity(
		knowledge.EntityPostgreSQL,
	) {
		t.Fatal(
			"expected postgresql entity",
		)
	}
}

func TestGolangAlias(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Quais projetos usam Golang?",
	)

	if !result.HasEntity(
		knowledge.EntityGo,
	) {
		t.Fatal(
			"expected Go entity",
		)
	}

	if result.Target !=
		domain.QueryTargetProject {
		t.Fatalf(
			"expected project target, got %s",
			result.Target,
		)
	}
}

func TestEnglishParallelism(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Which project best demonstrates parallelism?",
	)

	if !result.HasConcept(
		ontology.ConceptConcurrency,
	) {
		t.Fatal(
			"expected concurrency concept",
		)
	}
}

func TestPortugueseConcurrencyTypo(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Qual projeto demonstra concorencia?",
	)

	if !result.HasConcept(
		ontology.ConceptConcurrency,
	) {
		t.Fatalf(
			"expected concurrency concept, got %v",
			result.Concepts,
		)
	}
}

func TestCurrentTemporalScope(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Qual é o cargo atual do João?",
	)

	if result.TemporalScope !=
		domain.TemporalScopeCurrent {
		t.Fatalf(
			"expected current temporal scope, got %s",
			result.TemporalScope,
		)
	}
}

func TestPastTemporalScope(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Qual foi o emprego anterior do João?",
	)

	if result.TemporalScope !=
		domain.TemporalScopePast {
		t.Fatalf(
			"expected past temporal scope, got %s",
			result.TemporalScope,
		)
	}
}

func TestUnknownLanguageWithFallback(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result :=
		analyzer.AnalyzeWithLanguage(
			"GGCompress",
			domain.LanguageEnglish,
		)

	if result.Language !=
		domain.LanguageEnglish {
		t.Fatalf(
			"expected english fallback, got %s",
			result.Language,
		)
	}

	if result.Intent !=
		domain.IntentOverview {
		t.Fatalf(
			"expected overview, got %s",
			result.Intent,
		)
	}
}

func TestUnknownQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"xyzabc123",
	)

	if result.Intent !=
		domain.IntentUnknown {
		t.Fatalf(
			"expected unknown intent, got %s",
			result.Intent,
		)
	}

	if len(result.Entities) != 0 {
		t.Fatalf(
			"expected no entities, got %v",
			result.Entities,
		)
	}
}

func TestKafkaAddsSemanticConcepts(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Ele conhece Kafka?",
	)

	if !result.HasConcept(
		ontology.ConceptMessaging,
	) {
		t.Fatal(
			"expected messaging",
		)
	}

	if !result.HasConcept(
		ontology.ConceptEventDriven,
	) {
		t.Fatal(
			"expected event-driven",
		)
	}
}

func TestAuronixEntity(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Como funciona o Auronix?",
	)

	if !result.HasEntity(
		knowledge.EntityAuronix,
	) {
		t.Fatal(
			"expected auronix",
		)
	}

	if result.Target !=
		domain.QueryTargetProject {
		t.Fatalf(
			"expected project target, got %s",
			result.Target,
		)
	}
}

func TestEducationQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"O que ele estuda na FIAP?",
	)

	if result.Intent !=
		domain.IntentEducation {
		t.Fatalf(
			"expected education, got %s",
			result.Intent,
		)
	}

	if !result.HasEntity(
		knowledge.EntityFIAP,
	) {
		t.Fatal(
			"expected FIAP entity",
		)
	}
}

func TestCertificationQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Quais certificações ele possui?",
	)

	if result.Intent !=
		domain.IntentCertification {
		t.Fatalf(
			"expected certification, got %s",
			result.Intent,
		)
	}

	if result.Target !=
		domain.QueryTargetCertification {
		t.Fatalf(
			"expected certification target, got %s",
			result.Target,
		)
	}
}

func TestProfessionalExperienceQuery(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Me fale sobre a experiência profissional do João",
	)

	if result.Intent !=
		domain.IntentExperience {
		t.Fatalf(
			"expected experience, got %s",
			result.Intent,
		)
	}

	if result.Target !=
		domain.QueryTargetExperience {
		t.Fatalf(
			"expected experience target, got %s",
			result.Target,
		)
	}
}

func TestFuzzyMatchingDoesNotInventTechnology(
	t *testing.T,
) {
	analyzer := createAnalyzer(t)

	result := analyzer.Analyze(
		"Ele tem experiência com futebol?",
	)

	technologies := []domain.EntityID{
		knowledge.EntityKubernetes,
		knowledge.EntityPostgreSQL,
		knowledge.EntityMongoDB,
		knowledge.EntityRabbitMQ,
		knowledge.EntityKafka,
		knowledge.EntityGo,
	}

	for _, technology := range technologies {
		if result.HasEntity(technology) {
			t.Fatalf(
				"did not expect technology %s",
				technology,
			)
		}
	}
}
