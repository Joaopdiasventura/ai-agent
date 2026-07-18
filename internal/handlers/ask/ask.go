package handlers

import (
	"ai-agent/internal/app"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const maxRequestBodyBytes = 2048

func Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body AskRequest

		r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&body); err != nil {
			var maxBytesError *http.MaxBytesError
			var syntaxError *json.SyntaxError
			var typeError *json.UnmarshalTypeError

			switch {
			case errors.As(err, &maxBytesError):
				writeJSON(w, http.StatusRequestEntityTooLarge, AskResponse{
					Response: "Body da requisição excede o limite permitido.",
				})
			case errors.As(err, &syntaxError):
				writeJSON(w, http.StatusBadRequest, AskResponse{
					Response: "JSON inválido.",
				})
			case errors.As(err, &typeError):
				writeJSON(w, http.StatusBadRequest, AskResponse{
					Response: "Tipo de campo inválido.",
				})
			case errors.Is(err, io.EOF):
				writeJSON(w, http.StatusBadRequest, AskResponse{
					Response: "Body da requisição vazio.",
				})
			default:
				writeJSON(w, http.StatusBadRequest, AskResponse{
					Response: "Body da requisição inválido.",
				})
			}

			return
		}

		if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
			writeJSON(w, http.StatusBadRequest, AskResponse{
				Response: "O body deve conter apenas um objeto JSON.",
			})
			return
		}

		question := strings.TrimSpace(body.Content)

		if question == "" {
			writeJSON(w, http.StatusBadRequest, AskResponse{
				Response: "O campo content é obrigatório.",
			})
			return
		}

		response, hasResponse := app.AgentResponse(question)

		if !hasResponse {
			writeJSON(w, http.StatusNotFound, AskResponse{
				Response: "Não encontrei informações relacionadas à pergunta.",
			})
			return
		}

		writeJSON(w, http.StatusOK, AskResponse{
			Response: response,
		})
	}
}
