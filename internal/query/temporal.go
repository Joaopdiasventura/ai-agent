package query

import "ai-agent/internal/domain"

type TemporalDetector struct{}

func NewTemporalDetector() *TemporalDetector {
	return &TemporalDetector{}
}

func (d *TemporalDetector) Detect(
	normalized string,
) domain.TemporalScope {
	if hasAnyPhrase(
		normalized,
		currentTemporalMarkers,
	) {
		return domain.TemporalScopeCurrent
	}

	if hasAnyPhrase(
		normalized,
		pastTemporalMarkers,
	) {
		return domain.TemporalScopePast
	}

	return domain.TemporalScopeAny
}

var currentTemporalMarkers = []string{
	"atual",
	"atualmente",
	"hoje",
	"agora",
	"cargo atual",
	"empresa atual",
	"trabalho atual",
	"current",
	"currently",
	"today",
	"now",
	"present role",
	"current role",
	"current company",
	"current job",
}

var pastTemporalMarkers = []string{
	"anterior",
	"anteriores",
	"antes",
	"passado",
	"passados",
	"antigo",
	"antigos",
	"emprego anterior",
	"experiencia anterior",
	"previous",
	"previously",
	"past",
	"former",
	"formerly",
	"before",
	"previous job",
	"previous role",
	"past experience",
}
