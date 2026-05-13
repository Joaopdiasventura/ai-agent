package handlers

import (
	"encoding/json"
	"io"
	"log"
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

		flush := func() {}
		if flusher, ok := w.(http.Flusher); ok {
			flush = flusher.Flush
		} else {
			log.Printf("response writer does not support http.Flusher; SSE response may be buffered")
		}

		writeSSEHeaders(w)
		w.WriteHeader(http.StatusOK)
		flush()

		err = h.aiService.AskStream(r.Context(), content, func(chunk string) error {
			err := writeSSE(w, "", sseChunk{
				Chunk: chunk,
			})
			if err != nil {
				return err
			}

			flush()
			return nil
		})
		if err != nil {
			log.Printf("failed to stream Gemini response: %v", err)
			_ = writeSSE(w, "error", sseError{
				Message: "service unavailable",
			})
			flush()
			return
		}

		_ = writeSSE(w, "done", struct{}{})
		flush()
	}
}
