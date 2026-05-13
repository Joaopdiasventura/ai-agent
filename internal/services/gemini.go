package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/genai"
)

type AIService struct {
	client *genai.Client
}

func NewAIService(ctx context.Context) (*AIService, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("missing GEMINI_API_KEY or GOOGLE_API_KEY")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
		APIKey:  apiKey,
	})
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
		"gemini-2.5-flash-lite",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}

func (s *AIService) AskStream(ctx context.Context, question string, onChunk func(chunk string) error) error {
	prompt := BuildPortfolioPrompt(question)

	for result, err := range s.client.Models.GenerateContentStream(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(prompt),
		nil,
	) {
		if err != nil {
			return err
		}

		if result == nil {
			continue
		}

		text := result.Text()
		if text == "" {
			continue
		}

		if err := onChunk(text); err != nil {
			return err
		}
	}

	return nil
}
