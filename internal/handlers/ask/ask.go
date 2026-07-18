package handlers

import (
	"ai-agent/internal/app"
	"encoding/json"
	"net/http"
)

const maxRequestBodyBytes = 2048

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body AskRequest

		r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(r.Body)

		if err != nil {
			writeJSON(w, http.StatusBadRequest, AskResponse{
				Response: "Body da requisição inválido.",
			})
			return
		}

		question := body.Content

		response, hasResponse := app.AgentResponse(question)

		if !hasResponse {
			writeJSON(w, http.StatusNotFound, AskResponse{
				Response: "Não encontrei informações relacionadas à pergunta.",
			})
			return
		}

		writeJSON(w, http.StatusOK, AskResponse{Response: response})
	}
}
