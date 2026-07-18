package search

import "ai-agent/internal/domain"

func documentIDMatches(document *domain.Document, documentID string) bool {
	return document.ID == documentID ||
		document.ID == documentID+"-pt" ||
		document.ID == documentID+"-en"
}
