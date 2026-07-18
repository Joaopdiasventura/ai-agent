package knowledge

import (
	"ai-agent/internal/domain"
	"encoding/json"
	"fmt"
	"os"
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
