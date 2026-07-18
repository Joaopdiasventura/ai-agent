package server

import (
	handlers "ai-agent/internal/handlers/ask"
	"encoding/json"
	"net/http"
	"sync"
)

var allowedOrigins = []string{
	"https://joaopdias.dev.br",
	"http://localhost:4200",
}

var (
	handlerMu sync.Mutex
	handler   http.Handler
)

func Handler() (http.Handler, error) {
	handlerMu.Lock()
	defer handlerMu.Unlock()

	if handler != nil {
		return handler, nil
	}

	handler = NewHandler()
	return handler, nil
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("QUERY /ask", handlers.Handle())
	return CORS(mux)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "QUERY, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		if r.Method == http.MethodOptions {
			if isAllowedOrigin(origin) {
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

func isAllowedOrigin(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}

	return false
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
