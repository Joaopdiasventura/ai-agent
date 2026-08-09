package ontology

import "ai-agent/internal/domain"

type RelationType string

const (
	RelationRelatedTo  RelationType = "related_to"
	RelationSupports   RelationType = "supports"
	RelationEnables    RelationType = "enables"
	RelationAssociated RelationType = "associated_with"
)

type ConceptRelation struct {
	From          domain.ConceptID
	To            domain.ConceptID
	Type          RelationType
	Weight        float64
	Bidirectional bool
}

func allRelations() []ConceptRelation {
	return []ConceptRelation{
		{
			From:          ConceptMessaging,
			To:            ConceptAsyncProcessing,
			Type:          RelationSupports,
			Weight:        0.75,
			Bidirectional: true,
		},
		{
			From:          ConceptMessaging,
			To:            ConceptEventDriven,
			Type:          RelationSupports,
			Weight:        0.8,
			Bidirectional: true,
		},
		{
			From:          ConceptEventDriven,
			To:            ConceptDistributedSystems,
			Type:          RelationAssociated,
			Weight:        0.65,
			Bidirectional: true,
		},
		{
			From:          ConceptOrchestration,
			To:            ConceptDistributedSystems,
			Type:          RelationSupports,
			Weight:        0.7,
			Bidirectional: true,
		},
		{
			From:          ConceptCloud,
			To:            ConceptDistributedSystems,
			Type:          RelationAssociated,
			Weight:        0.45,
			Bidirectional: true,
		},
		{
			From:          ConceptStateless,
			To:            ConceptDistributedSystems,
			Type:          RelationSupports,
			Weight:        0.65,
			Bidirectional: true,
		},
		{
			From:          ConceptScalability,
			To:            ConceptDistributedSystems,
			Type:          RelationAssociated,
			Weight:        0.55,
			Bidirectional: true,
		},
		{
			From:          ConceptConcurrency,
			To:            ConceptPerformance,
			Type:          RelationSupports,
			Weight:        0.55,
			Bidirectional: true,
		},
		{
			From:          ConceptConcurrency,
			To:            ConceptAsyncProcessing,
			Type:          RelationAssociated,
			Weight:        0.5,
			Bidirectional: true,
		},
		{
			From:          ConceptDataIntegrity,
			To:            ConceptReliability,
			Type:          RelationSupports,
			Weight:        0.8,
			Bidirectional: true,
		},
		{
			From:          ConceptObservability,
			To:            ConceptReliability,
			Type:          RelationSupports,
			Weight:        0.65,
			Bidirectional: true,
		},
		{
			From:          ConceptFaultIsolation,
			To:            ConceptReliability,
			Type:          RelationSupports,
			Weight:        0.8,
			Bidirectional: true,
		},
		{
			From:          ConceptTransactionalConsistency,
			To:            ConceptFinancialSystems,
			Type:          RelationAssociated,
			Weight:        0.6,
			Bidirectional: true,
		},
		{
			From:          ConceptOptimisticLocking,
			To:            ConceptTransactionalConsistency,
			Type:          RelationSupports,
			Weight:        0.9,
			Bidirectional: false,
		},
		{
			From:          ConceptAtomicity,
			To:            ConceptTransactionalConsistency,
			Type:          RelationSupports,
			Weight:        0.85,
			Bidirectional: false,
		},
		{
			From:          ConceptAuditability,
			To:            ConceptFinancialSystems,
			Type:          RelationAssociated,
			Weight:        0.5,
			Bidirectional: true,
		},
		{
			From:          ConceptAuthentication,
			To:            ConceptSecurity,
			Type:          RelationSupports,
			Weight:        0.8,
			Bidirectional: true,
		},
		{
			From:          ConceptChecksum,
			To:            ConceptDataIntegrity,
			Type:          RelationSupports,
			Weight:        0.9,
			Bidirectional: false,
		},
		{
			From:          ConceptSafeExtraction,
			To:            ConceptDataIntegrity,
			Type:          RelationSupports,
			Weight:        0.75,
			Bidirectional: true,
		},
		{
			From:          ConceptInfrastructureAsCode,
			To:            ConceptDevOps,
			Type:          RelationAssociated,
			Weight:        0.75,
			Bidirectional: true,
		},
		{
			From:          ConceptContainerization,
			To:            ConceptDevOps,
			Type:          RelationAssociated,
			Weight:        0.75,
			Bidirectional: true,
		},
		{
			From:          ConceptDeployment,
			To:            ConceptDevOps,
			Type:          RelationAssociated,
			Weight:        0.75,
			Bidirectional: true,
		},
		{
			From:          ConceptLeadership,
			To:            ConceptTeamwork,
			Type:          RelationAssociated,
			Weight:        0.55,
			Bidirectional: true,
		},
		{
			From:          ConceptSemanticSearch,
			To:            ConceptArtificialIntelligence,
			Type:          RelationAssociated,
			Weight:        0.45,
			Bidirectional: true,
		},
		{
			From:          ConceptRetrieval,
			To:            ConceptArtificialIntelligence,
			Type:          RelationAssociated,
			Weight:        0.35,
			Bidirectional: true,
		},
		{
			From:          ConceptServerless,
			To:            ConceptStateless,
			Type:          RelationAssociated,
			Weight:        0.55,
			Bidirectional: true,
		},
		{
			From:          ConceptStreaming,
			To:            ConceptMediaProcessing,
			Type:          RelationAssociated,
			Weight:        0.75,
			Bidirectional: true,
		},
	}
}
