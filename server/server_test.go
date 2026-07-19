package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCORSAllowsConfiguredOriginAndPreflight(t *testing.T) {
	request := httptest.NewRequest(http.MethodOptions, "/ask", nil)
	request.Header.Set("Origin", "https://joaopdias.dev.br")
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, request)

	if response.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusNoContent)
	}

	if response.Header().Get("Access-Control-Allow-Origin") != "https://joaopdias.dev.br" {
		t.Fatalf("allow origin = %q", response.Header().Get("Access-Control-Allow-Origin"))
	}
}

func TestCORSRejectsUnknownOrigin(t *testing.T) {
	request := httptest.NewRequest("QUERY", "/ask", strings.NewReader(`{"content":"Qual é o email dele?"}`))
	request.Header.Set("Origin", "https://example.com")
	response := httptest.NewRecorder()

	NewHandler().ServeHTTP(response, request)

	if response.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusForbidden)
	}
}

func TestServerHandlesMultipleCallsInSameInstance(t *testing.T) {
	handler := NewHandler()

	for range 2 {
		request := httptest.NewRequest("QUERY", "/ask", strings.NewReader(`{"content":"Qual é o email dele?"}`))
		request.Header.Set("Origin", "http://localhost:4200")
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d; body: %s", response.Code, http.StatusOK, response.Body.String())
		}
	}
}
