package server

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"

	handlers "github.com/Joaopdiasventura/ai-agent/internal/handlers/ask"
	"github.com/Joaopdiasventura/ai-agent/internal/services"
)

const allowedOrigin = "https://joaopdias.dev.br"

var (
	initOnce sync.Once
	handler  http.Handler
	initErr  error
)

func Handler() (http.Handler, error) {
	initOnce.Do(func() {
		aiService, err := services.NewAIService(context.Background())
		if err != nil {
			initErr = err
			return
		}

		handler = NewHandler(aiService)
	})

	return handler, initErr
}

func NewHandler(aiService *services.AIService) http.Handler {
	askHandler := handlers.NewAskHandler(aiService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /ask/pt", askHandler.Handle(handlers.LanguagePT))
	mux.HandleFunc("POST /ask/en", askHandler.Handle(handlers.LanguageEN))

	return CORS(mux)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		if r.Method == http.MethodOptions {
			if origin == allowedOrigin {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			WriteJSON(w, http.StatusForbidden, handlers.AskResponse{
				Response: "Forbidden origin.",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func WriteJSON(w http.ResponseWriter, statusCode int, response handlers.AskResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func WriteServiceUnavailable(w http.ResponseWriter) {
	WriteJSON(w, http.StatusServiceUnavailable, handlers.AskResponse{
		Response: "Service temporarily unavailable.",
	})
}
