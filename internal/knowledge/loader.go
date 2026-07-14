package knowledge

import (
	"ai-agent/internal/domain"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadDocuments(path string) ([]domain.Document, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("não foi possível ler a base de conhecimento: %w", err)
	}

	var documents []domain.Document

	if err := json.Unmarshal(fileContent, &documents); err != nil {
		return nil, fmt.Errorf("não foi possível interpretar a base de conhecimento: %w", err)
	}

	if err := validateDocuments(documents); err != nil {
		return nil, err
	}

	return documents, nil
}

func validateDocuments(documents []domain.Document) error {
	if len(documents) == 0 {
		return fmt.Errorf("a base de conhecimento não contém documentos")
	}

	for index, document := range documents {
		documentID := strings.TrimSpace(document.ID)
		category := strings.TrimSpace(document.Category)
		content := strings.TrimSpace(document.Content)

		documentIDs := make(map[string]struct{})

		if documentID == "" {
			return fmt.Errorf("o documento na posição %d não possui ID", index)
		}

		if category == "" {
			return fmt.Errorf("o documento %q não possui categoria", documentID)
		}

		if content == "" {
			return fmt.Errorf("o documento %q não possui conteúdo", documentID)
		}

		if _, exists := documentIDs[documentID]; exists {
			return fmt.Errorf("o ID %q está duplicado", documentID)
		}

		documentIDs[documentID] = struct{}{}
	}

	return nil
}
