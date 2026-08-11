package evaluation

import (
	"ai-agent/internal/confidence"
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/ontology"
	"ai-agent/internal/planning"
)

func boolValue(
	value bool,
) *bool {
	return &value
}

func floatValue(
	value float64,
) *float64 {
	return &value
}

func RegressionCases() []Case {
	return []Case{
		{
			ID:       "pt-email",
			Question: "Qual o email do João?",
			Category: CategoryDirect,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentContact,
				Facts: []domain.FactID{
					"profile-email",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"joaopdias.dev@gmail.com",
				},
			},
		},
		{
			ID:       "en-email",
			Question: "What is João's email address?",
			Category: CategoryDirect,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentContact,
				Facts: []domain.FactID{
					"profile-email",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"joaopdias.dev@gmail.com",
				},
			},
		},
		{
			ID:       "pt-go-capability",
			Question: "Ele sabe Go?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityGo,
				},
				Facts: []domain.FactID{
					"skill-go",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				MinConfidence:  floatValue(0.4),
				ResponseContains: []string{
					"Go",
				},
			},
		},
		{
			ID:       "en-go-capability",
			Question: "Does João know Go?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityGo,
				},
				Facts: []domain.FactID{
					"skill-go",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Go",
				},
			},
		},
		{
			ID:       "pt-rust-abstention",
			Question: "Ele sabe Rust?",
			Category: CategoryAbstention,
			Expectation: Expectation{
				HasResponse:    boolValue(false),
				Language:       domain.LanguagePortuguese,
				Intent:         domain.IntentCapability,
				PlanStatus:     planning.PlanStatusAbstain,
				ConfidenceMode: confidence.ModeAbstention,
				MinConfidence:  floatValue(0.6),
			},
		},
		{
			ID:       "en-rust-abstention",
			Question: "Does João know Rust?",
			Category: CategoryAbstention,
			Expectation: Expectation{
				HasResponse:    boolValue(false),
				Language:       domain.LanguageEnglish,
				Intent:         domain.IntentCapability,
				PlanStatus:     planning.PlanStatusAbstain,
				ConfidenceMode: confidence.ModeAbstention,
			},
		},
		{
			ID:       "pt-ggcompress-overview",
			Question: "Me fale sobre o GGCompress",
			Category: CategoryOverview,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentOverview,
				Entities: []domain.EntityID{
					knowledge.EntityGGCompress,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"GGCompress",
				},
			},
		},
		{
			ID:       "en-ggcompress-overview",
			Question: "Tell me about GGCompress",
			Category: CategoryOverview,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentOverview,
				Entities: []domain.EntityID{
					knowledge.EntityGGCompress,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"GGCompress",
				},
			},
		},
		{
			ID:       "en-kafka-usage",
			Question: "Where did João use Kafka?",
			Category: CategoryTechnologyUsage,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentTechnologyUsage,
				Entities: []domain.EntityID{
					knowledge.EntityKafka,
				},
				Facts: []domain.FactID{
					"project-xtube-kafka",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Kafka",
					"X Tube",
				},
			},
		},
		{
			ID:       "pt-concurrency-comparison",
			Question: "Qual projeto melhor demonstra concorrência?",
			Category: CategoryComparison,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentComparison,
				Concepts: []domain.ConceptID{
					ontology.ConceptConcurrency,
				},
				Winner: knowledge.EntityGGCompress,
				Facts: []domain.FactID{
					"project-ggcompress-concurrency",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"GGCompress",
				},
			},
		},
		{
			ID:       "pt-distributed-comparison",
			Question: "Qual projeto melhor demonstra sistemas distribuídos?",
			Category: CategoryComparison,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentComparison,
				Concepts: []domain.ConceptID{
					ontology.ConceptDistributedSystems,
				},
				Winner: knowledge.EntityAuronix,
				Facts: []domain.FactID{
					"project-auronix-distributed",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Auronix",
				},
			},
		},
		{
			ID:       "pt-kubernetes-typo",
			Question: "Ele tem experiência com Kubernets?",
			Category: CategoryTypo,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Entities: []domain.EntityID{
					knowledge.EntityKubernetes,
				},
				Facts: []domain.FactID{
					"project-auronix-kubernetes",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
			},
		},
		{
			ID:       "pt-javascript-capability",
			Question: "Ele sabe JavaScript?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityJavaScript,
				},
				Facts: []domain.FactID{
					"skill-javascript",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"JavaScript",
				},
			},
		},
		{
			ID:       "pt-typescript-capability",
			Question: "Ele sabe TypeScript?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityTypeScript,
				},
				Facts: []domain.FactID{
					"skill-typescript",
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"TypeScript",
				},
			},
		},
		{
			ID:       "pt-java-typo",
			Question: "ele sabe jva?",
			Category: CategoryTypo,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityJava,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Java",
				},
			},
		},
		{
			ID:       "pt-docker-typo",
			Question: "ele sabe dcker?",
			Category: CategoryTypo,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityDocker,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Docker",
				},
			},
		},
		{
			ID:       "pt-programming-languages-list",
			Question: "Quais linguagens ele sabe?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentList,
				Concepts: []domain.ConceptID{
					ontology.ConceptProgrammingLanguage,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"JavaScript",
					"TypeScript",
					"Java",
					"Go",
				},
				ResponseNotContains: []string{
					"Português",
					"Inglês",
				},
			},
		},
		{
			ID:       "pt-frameworks-list",
			Question: "Quais frameworks ele sabe?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentList,
				Concepts: []domain.ConceptID{
					ontology.ConceptFramework,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Angular",
					"React",
					"Next.js",
					"Spring Boot",
					"NestJS",
				},
			},
		},
		{
			ID:       "pt-technologies-list",
			Question: "Quais tecnologias ele sabe?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse:    boolValue(true),
				Language:       domain.LanguagePortuguese,
				Intent:         domain.IntentList,
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Java",
					"Go",
					"Docker",
				},
			},
		},
		{
			ID:       "pt-human-languages-list",
			Question: "quais idiomas ele fala?",
			Category: CategoryDirect,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentList,
				Concepts: []domain.ConceptID{
					ontology.ConceptLanguage,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Português",
					"Inglês",
				},
				ResponseNotContains: []string{
					"JavaScript",
				},
			},
		},
		{
			ID:       "pt-unknown-age",
			Question: "quantos anos tem João?",
			Category: CategoryAbstention,
			Expectation: Expectation{
				HasResponse:    boolValue(false),
				Language:       domain.LanguagePortuguese,
				Intent:         domain.IntentUnknown,
				PlanStatus:     planning.PlanStatusAbstain,
				ConfidenceMode: confidence.ModeAbstention,
			},
		},
		{
			ID:       "pt-unknown-preference",
			Question: "João tem preferência de linguagem?",
			Category: CategoryAbstention,
			Expectation: Expectation{
				HasResponse:    boolValue(false),
				Language:       domain.LanguagePortuguese,
				Intent:         domain.IntentUnknown,
				PlanStatus:     planning.PlanStatusAbstain,
				ConfidenceMode: confidence.ModeAbstention,
			},
		},
		{
			ID:       "pt-backend-capability",
			Question: "ele sabe backend?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Concepts: []domain.ConceptID{
					ontology.ConceptBackend,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
			},
		},
		{
			ID:       "pt-fullstack-capability",
			Question: "ele sabe fulstack?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguagePortuguese,
				Intent:      domain.IntentCapability,
				Concepts: []domain.ConceptID{
					ontology.ConceptFullStack,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
			},
		},
		{
			ID:       "en-programming-languages-list",
			Question: "What programming languages does João know?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentList,
				Concepts: []domain.ConceptID{
					ontology.ConceptProgrammingLanguage,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"JavaScript",
					"TypeScript",
					"Java",
					"Go",
				},
			},
		},
		{
			ID:       "en-frameworks-list",
			Question: "Which frameworks does João know?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentList,
				Concepts: []domain.ConceptID{
					ontology.ConceptFramework,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Angular",
					"React",
					"Next.js",
				},
			},
		},
		{
			ID:       "en-docker-capability",
			Question: "Does João know Docker?",
			Category: CategoryCapability,
			Expectation: Expectation{
				HasResponse: boolValue(true),
				Language:    domain.LanguageEnglish,
				Intent:      domain.IntentCapability,
				Entities: []domain.EntityID{
					knowledge.EntityDocker,
				},
				PlanStatus:     planning.PlanStatusReady,
				ConfidenceMode: confidence.ModeClaim,
				ResponseContains: []string{
					"Docker",
				},
			},
		},
		{
			ID:       "en-unknown-age",
			Question: "How old is João?",
			Category: CategoryAbstention,
			Expectation: Expectation{
				HasResponse:    boolValue(false),
				Language:       domain.LanguageEnglish,
				Intent:         domain.IntentUnknown,
				PlanStatus:     planning.PlanStatusAbstain,
				ConfidenceMode: confidence.ModeAbstention,
			},
		},
		{
			ID:       "unknown-gibberish",
			Question: "xyzabc123",
			Category: CategoryUnknown,
			Expectation: Expectation{
				HasResponse:    boolValue(false),
				Intent:         domain.IntentUnknown,
				PlanStatus:     planning.PlanStatusAbstain,
				ConfidenceMode: confidence.ModeAbstention,
			},
		},
	}
}

func DefaultCases() []Case {
	result := append(
		[]Case{},
		RegressionCases()...,
	)

	result = append(
		result,
		[]Case{
			{
				ID:       "pt-current-role",
				Question: "Qual é o cargo atual do João?",
				Category: CategoryExperience,
				Expectation: Expectation{
					HasResponse: boolValue(true),
					Language:    domain.LanguagePortuguese,
					Facts: []domain.FactID{
						"experience-current-role",
					},
					PlanStatus:     planning.PlanStatusReady,
					ConfidenceMode: confidence.ModeClaim,
				},
			},
			{
				ID:       "pt-concurrency-capability",
				Question: "Ele tem experiência com concorrência?",
				Category: CategoryCapability,
				Expectation: Expectation{
					HasResponse: boolValue(true),
					Language:    domain.LanguagePortuguese,
					Intent:      domain.IntentCapability,
					Concepts: []domain.ConceptID{
						ontology.ConceptConcurrency,
					},
					PlanStatus:     planning.PlanStatusReady,
					ConfidenceMode: confidence.ModeClaim,
				},
			},
			{
				ID:       "en-concurrency-comparison",
				Question: "Which project best demonstrates concurrency?",
				Category: CategoryComparison,
				Expectation: Expectation{
					HasResponse: boolValue(true),
					Language:    domain.LanguageEnglish,
					Intent:      domain.IntentComparison,
					Concepts: []domain.ConceptID{
						ontology.ConceptConcurrency,
					},
					Winner:         knowledge.EntityGGCompress,
					PlanStatus:     planning.PlanStatusReady,
					ConfidenceMode: confidence.ModeClaim,
					ResponseContains: []string{
						"GGCompress",
					},
				},
			},
			{
				ID:       "pt-auronix-overview",
				Question: "Me fale sobre o Auronix",
				Category: CategoryOverview,
				Expectation: Expectation{
					HasResponse: boolValue(true),
					Language:    domain.LanguagePortuguese,
					Intent:      domain.IntentOverview,
					Entities: []domain.EntityID{
						knowledge.EntityAuronix,
					},
					PlanStatus:     planning.PlanStatusReady,
					ConfidenceMode: confidence.ModeClaim,
					ResponseContains: []string{
						"Auronix",
					},
				},
			},
			{
				ID:       "pt-xtube-overview",
				Question: "Me fale sobre o X Tube",
				Category: CategoryOverview,
				Expectation: Expectation{
					HasResponse: boolValue(true),
					Language:    domain.LanguagePortuguese,
					Intent:      domain.IntentOverview,
					Entities: []domain.EntityID{
						knowledge.EntityXTube,
					},
					PlanStatus:     planning.PlanStatusReady,
					ConfidenceMode: confidence.ModeClaim,
					ResponseContains: []string{
						"X Tube",
					},
				},
			},
			{
				ID:       "pt-vox-overview",
				Question: "Me fale sobre o Vox",
				Category: CategoryOverview,
				Expectation: Expectation{
					HasResponse: boolValue(true),
					Language:    domain.LanguagePortuguese,
					Intent:      domain.IntentOverview,
					Entities: []domain.EntityID{
						knowledge.EntityVox,
					},
					PlanStatus:     planning.PlanStatusReady,
					ConfidenceMode: confidence.ModeClaim,
					ResponseContains: []string{
						"Vox",
					},
				},
			},
			{
				ID:       "en-kafka-unknown-negative",
				Question: "Did João use Kafka in GGCompress?",
				Category: CategoryTechnologyUsage,
				Expectation: Expectation{
					Language: domain.LanguageEnglish,
					Entities: []domain.EntityID{
						knowledge.EntityKafka,
						knowledge.EntityGGCompress,
					},
					ForbiddenFacts: []domain.FactID{
						"project-ggcompress-kafka",
					},
				},
			},
		}...,
	)

	return result
}
