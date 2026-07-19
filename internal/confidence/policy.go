package confidence

import (
	"ai-agent/internal/domain"
)

type Level string

const (
	High   Level = "high"
	Medium Level = "medium"
	Low    Level = "low"
)

type Assessment struct {
	Level  Level
	Score  float64
	Reason string
}

type Policy struct {
	HighThreshold   float64
	MediumThreshold float64
}

func DefaultPolicy() Policy {
	return Policy{
		HighThreshold:   75,
		MediumThreshold: 0,
	}
}

func (policy Policy) Assess(query domain.Query, results []domain.SearchResult, evidence []domain.Evidence) Assessment {
	if len(results) == 0 || len(evidence) == 0 {
		return Assessment{Level: Low, Reason: "no_evidence"}
	}

	if query.Category == "" && query.Project == "" && len(query.ExactTerms) == 0 {
		return Assessment{Level: Low, Reason: "unspecified_query"}
	}

	first := results[0]
	if first.Document != nil && query.Language != "" && first.Document.Language != query.Language {
		return Assessment{Level: Low, Reason: "language_mismatch"}
	}

	score := first.FinalScore

	if score == 0 {
		score = first.Score
	}

	if hasSource(first.Sources, "vector") && hasSource(first.Sources, "lexical") {
		score += 12
	}

	if query.Language != "" && first.Document != nil && first.Document.Language == query.Language {
		score += 10
	}

	if query.Category != "" && first.Document != nil && first.Document.Category == query.Category {
		score += 8
	}

	if query.Project != "" && first.Document != nil && first.Document.Project == query.Project {
		score += 8
	}

	if len(evidence) > 1 {
		score += 4
	}

	if len(results) > 1 {
		gap := first.FinalScore - results[1].FinalScore
		if first.FinalScore == 0 {
			gap = first.Score - results[1].Score
		}

		if gap >= 15 {
			score += 8
		} else if gap < 3 {
			score -= 8
		}
	}

	if len(first.PenaltyReasons) > 0 {
		score -= float64(len(first.PenaltyReasons) * 15)
	}

	if query.Category == "" && len(query.ExactTerms) == 0 && !hasSource(first.Sources, "lexical") {
		score -= 45
	}

	switch {
	case score >= policy.HighThreshold:
		return Assessment{Level: High, Score: score, Reason: "high_confidence"}
	case score >= policy.MediumThreshold:
		return Assessment{Level: Medium, Score: score, Reason: "medium_confidence"}
	default:
		return Assessment{Level: Low, Score: score, Reason: "low_score"}
	}
}

func hasSource(sources []string, source string) bool {
	for _, existing := range sources {
		if existing == source {
			return true
		}
	}

	return false
}
