package ranking

import (
	"errors"

	"ai-agent/internal/domain"
	"ai-agent/internal/knowledge"
	"ai-agent/internal/retrieval"
)

type Config struct {
	FusionWeight  float64
	FeatureWeight float64
	CandidatePool int
}

func DefaultConfig() Config {
	return Config{
		FusionWeight:  0.55,
		FeatureWeight: 0.45,
		CandidatePool: 80,
	}
}

type Ranker struct {
	base     *knowledge.Knowledge
	fusion   *RRFFusion
	features *FeatureScorer
	config   Config
}

func New(
	base *knowledge.Knowledge,
) (*Ranker, error) {
	return NewWithConfig(
		base,
		DefaultConfig(),
		DefaultRRFConfig(),
	)
}

func NewWithConfig(
	base *knowledge.Knowledge,
	config Config,
	rrfConfig RRFConfig,
) (*Ranker, error) {
	if base == nil {
		return nil, errors.New(
			"knowledge base is required",
		)
	}

	if config.FusionWeight < 0 ||
		config.FeatureWeight < 0 {
		return nil, errors.New(
			"ranking weights cannot be negative",
		)
	}

	totalWeight :=
		config.FusionWeight +
			config.FeatureWeight

	if totalWeight <= 0 {
		return nil, errors.New(
			"ranking weights cannot both be zero",
		)
	}

	config.FusionWeight /=
		totalWeight

	config.FeatureWeight /=
		totalWeight

	if config.CandidatePool <= 0 {
		config.CandidatePool = 80
	}

	return &Ranker{
		base: base,
		fusion: NewRRFFusion(
			rrfConfig,
		),
		features: NewFeatureScorer(
			base,
		),
		config: config,
	}, nil
}

func (r *Ranker) Rank(
	retrievalResult retrieval.Result,
	limit int,
) Result {
	if limit <= 0 {
		limit = 20
	}

	poolSize :=
		r.config.CandidatePool

	if minimum :=
		limit * 4; minimum > poolSize {
		poolSize = minimum
	}

	fused :=
		r.fusion.Fuse(
			retrievalResult,
			poolSize,
		)

	if len(fused) == 0 {
		return Result{
			Query: retrievalResult.Query,
		}
	}

	values := make(
		[]Candidate,
		0,
		len(fused),
	)

	for _, candidate := range fused {
		fact, found :=
			r.base.Fact(
				candidate.FactID,
			)

		if !found {
			continue
		}

		featureScore, signals :=
			r.features.Score(
				retrievalResult.Query,
				fact,
				candidate,
			)

		candidate.FeatureScore =
			featureScore

		candidate.Score =
			clamp(
				r.config.FusionWeight*
					candidate.FusionScore +
					r.config.FeatureWeight*
						candidate.FeatureScore,
			)

		candidate.Signals =
			append(
				candidate.Signals,
				signals...,
			)

		values = append(
			values,
			candidate,
		)
	}

	sortFinalCandidates(
		values,
	)

	if len(values) > limit {
		values =
			values[:limit]
	}

	for index := range values {
		values[index].Rank =
			index + 1
	}

	return Result{
		Query:      retrievalResult.Query,
		Candidates: values,
	}
}

func (r *Ranker) RankQuery(
	currentQuery domain.Query,
	retriever *retrieval.HybridRetriever,
	limit int,
) Result {
	if retriever == nil {
		return Result{
			Query: currentQuery,
		}
	}

	retrievalResult :=
		retriever.Search(
			currentQuery,
			r.config.CandidatePool,
		)

	return r.Rank(
		retrievalResult,
		limit,
	)
}
