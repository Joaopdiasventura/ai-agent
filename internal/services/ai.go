package services

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const groqURL = "https://api.groq.com/openai/v1/chat/completions"
const groqModel = "openai/gpt-oss-120b"

type AIService struct {
	apiKey string
	client *http.Client
}

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqRequest struct {
	Model       string        `json:"model"`
	Messages    []groqMessage `json:"messages"`
	Temperature float64       `json:"temperature"`
	Stream      bool          `json:"stream"`
}

type groqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

type groqStreamResponse struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error,omitempty"`
}

func NewAIService(ctx context.Context) (*AIService, error) {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing GROQ_API_KEY")
	}

	return &AIService{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}, nil
}

func (s *AIService) Ask(ctx context.Context, question string) (string, error) {
	prompt := BuildPortfolioPrompt(question)

	payload := groqRequest{
		Model: groqModel,
		Messages: []groqMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.4,
		Stream:      false,
	}

	body, err := s.doRequest(ctx, payload)
	if err != nil {
		return "", err
	}
	defer body.Close()

	var response groqResponse
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", errors.New(response.Error.Message)
	}

	if len(response.Choices) == 0 {
		return "", errors.New("empty response from Groq")
	}

	return response.Choices[0].Message.Content, nil
}

func (s *AIService) AskStream(ctx context.Context, question string, onChunk func(chunk string) error) error {
	prompt := BuildPortfolioPrompt(question)

	payload := groqRequest{
		Model: groqModel,
		Messages: []groqMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
		Temperature: 0.4,
		Stream:      true,
	}

	body, err := s.doRequest(ctx, payload)
	if err != nil {
		return err
	}
	defer body.Close()

	scanner := bufio.NewScanner(body)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, "data:") {
			continue
		}

		data := strings.TrimSpace(strings.TrimPrefix(line, "data:"))

		if data == "[DONE]" {
			return nil
		}

		var response groqStreamResponse
		if err := json.Unmarshal([]byte(data), &response); err != nil {
			return err
		}

		if response.Error != nil {
			return errors.New(response.Error.Message)
		}

		if len(response.Choices) == 0 {
			continue
		}

		chunk := response.Choices[0].Delta.Content
		if chunk == "" {
			continue
		}

		if err := onChunk(chunk); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (s *AIService) doRequest(ctx context.Context, payload groqRequest) (io.ReadCloser, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, groqURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+s.apiKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := s.client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		defer response.Body.Close()

		var groqError groqResponse
		if err := json.NewDecoder(response.Body).Decode(&groqError); err == nil && groqError.Error != nil {
			return nil, errors.New(groqError.Error.Message)
		}

		return nil, fmt.Errorf("groq request failed with status %d", response.StatusCode)
	}

	return response.Body, nil
}
