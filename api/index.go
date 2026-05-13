package api

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
	apiMux   http.Handler
	initErr  error
)

func Handler(w http.ResponseWriter, r *http.Request) {
	initOnce.Do(func() {
		aiService, err := services.NewAIService(context.Background())
		if err != nil {
			initErr = err
			return
		}

		askHandler := handlers.NewAskHandler(aiService)

		mux := http.NewServeMux()
		mux.HandleFunc("POST /ask/pt", askHandler.Handle(handlers.LanguagePT))
		mux.HandleFunc("POST /ask/en", askHandler.Handle(handlers.LanguageEN))

		apiMux = cors(mux)
	})

	if initErr != nil {
		writeJSON(w, http.StatusServiceUnavailable, handlers.AskResponse{
			Response: "Service temporarily unavailable.",
		})
		return
	}

	apiMux.ServeHTTP(w, r)
}

func cors(next http.Handler) http.Handler {
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

			writeJSON(w, http.StatusForbidden, handlers.AskResponse{
				Response: "Forbidden origin.",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, response handlers.AskResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
