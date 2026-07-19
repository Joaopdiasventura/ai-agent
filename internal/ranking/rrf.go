package ranking

import (
	"ai-agent/internal/domain"
	"sort"
)

const DefaultRRFK = 60

func Fuse(vectorResults []domain.SearchResult, lexicalResults []domain.SearchResult, limit int) []domain.SearchResult {
	return FuseWithK(vectorResults, lexicalResults, limit, DefaultRRFK)
}

func FuseWithK(vectorResults []domain.SearchResult, lexicalResults []domain.SearchResult, limit int, k float64) []domain.SearchResult {
	if limit <= 0 {
		return []domain.SearchResult{}
	}

	if k <= 0 {
		k = DefaultRRFK
	}

	combined := make(map[string]domain.SearchResult)

	addRanking := func(results []domain.SearchResult, source string) {
		for index, result := range results {
			if result.Document == nil || result.Document.ID == "" {
				continue
			}

			rank := index + 1
			existing, exists := combined[result.Document.ID]
			if !exists {
				existing = result
				existing.Score = 0
				existing.Sources = []string{}
			}

			existing.Score += 1 / (k + float64(rank))
			existing.Sources = addSource(existing.Sources, source)

			if source == "vector" {
				existing.VectorRank = rank
			}

			if source == "lexical" {
				existing.LexicalRank = rank
			}

			combined[result.Document.ID] = existing
		}
	}

	addRanking(vectorResults, "vector")
	addRanking(lexicalResults, "lexical")

	fused := make([]domain.SearchResult, 0, len(combined))
	for _, result := range combined {
		fused = append(fused, result)
	}

	sort.Slice(fused, func(firstIndex int, secondIndex int) bool {
		if fused[firstIndex].Score == fused[secondIndex].Score {
			return fused[firstIndex].Document.ID < fused[secondIndex].Document.ID
		}

		return fused[firstIndex].Score > fused[secondIndex].Score
	})

	if limit > len(fused) {
		limit = len(fused)
	}

	fused = fused[:limit]
	for index := range fused {
		fused[index].FusedRank = index + 1
	}

	return fused
}

func addSource(sources []string, source string) []string {
	for _, existing := range sources {
		if existing == source {
			return sources
		}
	}

	return append(sources, source)
}
