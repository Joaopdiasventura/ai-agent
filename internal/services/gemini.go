package services

import (
	"context"
	"time"

	"google.golang.org/genai"
)

type AIService struct {
	client *genai.Client
}

func NewAIService(ctx context.Context) (*AIService, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &AIService{
		client: client,
	}, nil
}

func (s *AIService) Ask(ctx context.Context, question string) (string, error) {
	prompt := BuildPortfolioPrompt(question)

	result, err := s.client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
