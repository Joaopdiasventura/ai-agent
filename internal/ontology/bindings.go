package ontology

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
)

type EntityConceptBinding struct {
	EntityID domain.EntityID
	Concepts []WeightedConcept
}

func allEntityBindings() []EntityConceptBinding {
	return []EntityConceptBinding{
		binding(
			knowledge.EntityAngular,
			weighted(ConceptFramework, 1),
			weighted(ConceptFrontend, 1),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityReact,
			weighted(ConceptFramework, 1),
			weighted(ConceptFrontend, 1),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityNextJS,
			weighted(ConceptFramework, 1),
			weighted(ConceptFrontend, 0.9),
			weighted(ConceptBackend, 0.45),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityJava,
			weighted(ConceptProgrammingLanguage, 1),
			weighted(ConceptBackend, 0.85),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntitySpringBoot,
			weighted(ConceptFramework, 1),
			weighted(ConceptBackend, 1),
			weighted(ConceptSoftwareDevelopment, 0.85),
		),
		binding(
			knowledge.EntityGo,
			weighted(ConceptProgrammingLanguage, 1),
			weighted(ConceptBackend, 0.85),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityNodeJS,
			weighted(ConceptRuntime, 1),
			weighted(ConceptBackend, 0.9),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityNestJS,
			weighted(ConceptFramework, 1),
			weighted(ConceptBackend, 1),
			weighted(ConceptSoftwareDevelopment, 0.85),
		),
		binding(
			knowledge.EntityJavaScript,
			weighted(ConceptProgrammingLanguage, 1),
			weighted(ConceptFrontend, 0.75),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityTypeScript,
			weighted(ConceptProgrammingLanguage, 1),
			weighted(ConceptFrontend, 0.75),
			weighted(ConceptSoftwareDevelopment, 0.8),
		),
		binding(
			knowledge.EntityPostgreSQL,
			weighted(ConceptPostgreSQL, 1),
			weighted(ConceptDatabase, 1),
		),
		binding(
			knowledge.EntityMongoDB,
			weighted(ConceptMongoDB, 1),
			weighted(ConceptDatabase, 1),
		),
		binding(
			knowledge.EntityRedis,
			weighted(ConceptDatabase, 0.65),
			weighted(ConceptDistributedSystems, 0.45),
		),
		binding(
			knowledge.EntityRabbitMQ,
			weighted(ConceptMessaging, 1),
			weighted(ConceptAsyncProcessing, 0.75),
			weighted(ConceptEventDriven, 0.7),
		),
		binding(
			knowledge.EntityKafka,
			weighted(ConceptMessaging, 1),
			weighted(ConceptEventDriven, 1),
			weighted(ConceptAsyncProcessing, 0.65),
		),
		binding(
			knowledge.EntitySQS,
			weighted(ConceptMessaging, 1),
			weighted(ConceptAsyncProcessing, 0.9),
			weighted(ConceptCloud, 0.7),
			weighted(ConceptAWS, 0.7),
		),
		binding(
			knowledge.EntityDocker,
			weighted(ConceptContainerization, 1),
			weighted(ConceptContainers, 1),
			weighted(ConceptDevOps, 0.7),
		),
		binding(
			knowledge.EntityTerraform,
			weighted(ConceptInfrastructureAsCode, 1),
			weighted(ConceptInfrastructure, 0.85),
			weighted(ConceptDevOps, 0.7),
		),
		binding(
			knowledge.EntityKubernetes,
			weighted(ConceptKubernetes, 1),
			weighted(ConceptOrchestration, 1),
			weighted(ConceptContainerization, 0.65),
			weighted(ConceptDistributedSystems, 0.55),
		),
		binding(
			knowledge.EntityAWS,
			weighted(ConceptAWS, 1),
			weighted(ConceptCloud, 1),
			weighted(ConceptInfrastructure, 0.7),
		),
		binding(
			knowledge.EntityECS,
			weighted(ConceptAWS, 0.8),
			weighted(ConceptCloud, 0.8),
			weighted(ConceptContainerization, 0.9),
		),
		binding(
			knowledge.EntityEKS,
			weighted(ConceptAWS, 0.9),
			weighted(ConceptCloud, 0.9),
			weighted(ConceptKubernetes, 1),
			weighted(ConceptOrchestration, 1),
		),
		binding(
			knowledge.EntityS3,
			weighted(ConceptAWS, 0.85),
			weighted(ConceptCloud, 0.85),
			weighted(ConceptStorage, 1),
		),
		binding(
			knowledge.EntityIAM,
			weighted(ConceptAWS, 0.85),
			weighted(ConceptCloud, 0.8),
			weighted(ConceptSecurity, 1),
		),
		binding(
			knowledge.EntityCloudflare,
			weighted(ConceptInfrastructure, 0.8),
			weighted(ConceptDNS, 1),
			weighted(ConceptProxy, 0.9),
			weighted(ConceptRouting, 0.9),
		),
		binding(
			knowledge.EntityFFmpeg,
			weighted(ConceptMediaProcessing, 1),
			weighted(ConceptStreaming, 0.55),
		),
		binding(
			knowledge.EntityPrometheus,
			weighted(ConceptObservability, 1),
			weighted(ConceptReliability, 0.55),
		),
		binding(
			knowledge.EntityTauri,
			weighted(ConceptCrossPlatform, 1),
			weighted(ConceptFrontend, 0.55),
		),
		binding(
			knowledge.EntityJWT,
			weighted(ConceptAuthentication, 1),
			weighted(ConceptSecurity, 0.9),
		),
		binding(
			knowledge.EntityOAuth2,
			weighted(ConceptAuthentication, 1),
			weighted(ConceptSecurity, 0.9),
		),
		binding(
			knowledge.EntityWorkerThread,
			weighted(ConceptConcurrency, 1),
			weighted(ConceptAsyncProcessing, 0.75),
			weighted(ConceptFaultIsolation, 0.55),
		),
		binding(
			knowledge.EntityWebWorker,
			weighted(ConceptConcurrency, 0.9),
			weighted(ConceptFrontend, 0.75),
			weighted(ConceptPerformance, 0.6),
		),
		binding(
			knowledge.EntitySHA256,
			weighted(ConceptChecksum, 1),
			weighted(ConceptDataIntegrity, 1),
			weighted(ConceptReliability, 0.7),
		),
		binding(
			knowledge.EntityVectorStores,
			weighted(ConceptSemanticSearch, 0.85),
			weighted(ConceptRetrieval, 0.9),
			weighted(ConceptArtificialIntelligence, 0.8),
		),
		binding(
			knowledge.EntityFileSearch,
			weighted(ConceptSemanticSearch, 0.95),
			weighted(ConceptRetrieval, 1),
			weighted(ConceptArtificialIntelligence, 0.8),
		),
		binding(
			knowledge.EntityPortuguese,
			weighted(ConceptLanguage, 1),
		),
		binding(
			knowledge.EntityEnglish,
			weighted(ConceptLanguage, 1),
		),
	}
}

func binding(
	entityID domain.EntityID,
	concepts ...WeightedConcept,
) EntityConceptBinding {
	return EntityConceptBinding{
		EntityID: entityID,
		Concepts: concepts,
	}
}

func weighted(
	id domain.ConceptID,
	weight float64,
) WeightedConcept {
	return WeightedConcept{
		ID:     id,
		Weight: weight,
	}
}
