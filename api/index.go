package api

import (
	"ai-agent/server"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiHandler, err := server.Handler()
	if err != nil {
		server.WriteServiceUnavailable(w)
		return
	}

	apiHandler.ServeHTTP(w, r)
}
