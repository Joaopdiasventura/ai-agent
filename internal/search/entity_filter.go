package search

import (
	"ai-agent/internal/domain"
	"ai-agent/internal/nlp"
	"strings"
)

func FilterDocumentsByEntity(documents []domain.Document, entity nlp.Entity) []domain.Document {
	filteredDocuments := make([]domain.Document, 0)

	for _, document := range documents {
		if matchesEntity(document, entity) {
			filteredDocuments = append(filteredDocuments, document)
		}
	}

	return filteredDocuments
}

func matchesEntity(document domain.Document, entity nlp.Entity) bool {
	documentID := strings.ToLower(document.ID)
	documentContent := strings.ToLower(document.Content)
	entityValue := strings.ToLower(entity.Value)

	switch entity.Type {
	case nlp.EntityProject:
		projectID := "project-" + normalizeEntityID(entityValue)

		return documentID == projectID ||
			strings.Contains(documentContent, entityValue)

	case nlp.EntityCompany:
		return strings.Contains(documentContent, entityValue)

	case nlp.EntityInstitution:
		return strings.Contains(documentContent, entityValue)

	case nlp.EntityTechnology:
		return strings.Contains(documentContent, entityValue)

	default:
		return false
	}
}

func normalizeEntityID(value string) string {
	value = strings.ToLower(value)
	value = strings.ReplaceAll(value, " ", "-")
	return value
}
