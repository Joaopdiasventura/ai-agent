package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type sseChunk struct {
	Chunk string `json:"chunk"`
}

type sseError struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, statusCode int, response AskResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func writeSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
}

func writeSSE(w http.ResponseWriter, event string, data any) error {
	payload, err := json.Marshal(data)
	if err != nil {
		payload = []byte("{}")
	}

	if event != "" {
		if _, err := fmt.Fprintf(w, "event: %s\n", event); err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "data: %s\n\n", payload)
	return err
}
