package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Joaopdiasventura/ai-agent/internal/services"
)

const maxRequestBodyBytes = 2048

type AskHandler struct {
	aiService *services.AIService
}

func NewAskHandler(aiService *services.AIService) *AskHandler {
	return &AskHandler{
		aiService: aiService,
	}
}

func (h *AskHandler) Handle(lang Language) http.HandlerFunc {
	messages := messagesForLanguage(lang)

	return func(w http.ResponseWriter, r *http.Request) {
		var body AskRequest

		r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&body)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, AskResponse{
				Response: messages.InvalidBody,
			})
			return
		}

		err = decoder.Decode(&struct{}{})
		if err != io.EOF {
			writeJSON(w, http.StatusBadRequest, AskResponse{
				Response: messages.InvalidBody,
			})
			return
		}

		content := normalizeQuestion(body.Content)

		validation := ValidateQuestion(content, lang)
		if !validation.Valid {
			writeJSON(w, http.StatusOK, AskResponse{
				Response: validation.Message,
			})
			return
		}

		response, err := h.aiService.Ask(r.Context(), content)
		if err != nil {
			writeJSON(w, http.StatusServiceUnavailable, AskResponse{
				Response: messages.ServiceError,
			})
			return
		}

		writeJSON(w, http.StatusOK, AskResponse{
			Response: response,
		})
	}
}
