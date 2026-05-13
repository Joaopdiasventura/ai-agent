package api

import (
	"net/http"

	"github.com/Joaopdiasventura/ai-agent/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiHandler, err := server.Handler()
	if err != nil {
		server.WriteServiceUnavailable(w)
		return
	}

	apiHandler.ServeHTTP(w, r)
}
