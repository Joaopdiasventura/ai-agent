package main

import (
	"context"
	"log"
	"net/http"
	"os"

	handlers "github.com/Joaopdiasventura/ai-agent/internal/handlers/ask"
	"github.com/Joaopdiasventura/ai-agent/internal/services"
)

const allowedOrigin = "https://joaopdias.dev.br"

func main() {
	ctx := context.Background()

	aiService, err := services.NewAIService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	askHandler := handlers.NewAskHandler(aiService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /ask/pt", askHandler.Handle(handlers.LanguagePT))
	mux.HandleFunc("POST /ask/en", askHandler.Handle(handlers.LanguageEN))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("server running on http://localhost:"+port)

	err = http.ListenAndServe(":"+port, cors(mux))

	if err != nil {
		log.Fatal(err)
	}
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

			http.Error(w, "forbidden origin", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
