package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleValidQuestionReturnsResponse(t *testing.T) {
	response := performAskRequest(t, `{"content":"Me fale sobre o auronix"}`)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", response.Code, http.StatusOK, response.Body.String())
	}

	body := decodeAskResponse(t, response)
	if !strings.Contains(body.Response, "Auronix") {
		t.Fatalf("response %q does not mention Auronix", body.Response)
	}
}

func TestHandleRejectsInvalidRequests(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantText   string
	}{
		{
			name:       "empty content",
			body:       `{"content":"   "}`,
			wantStatus: http.StatusBadRequest,
			wantText:   "O campo content é obrigatório.",
		},
		{
			name:       "invalid json",
			body:       `{"content":}`,
			wantStatus: http.StatusBadRequest,
			wantText:   "JSON inválido.",
		},
		{
			name:       "unknown field",
			body:       `{"content":"oi","extra":true}`,
			wantStatus: http.StatusBadRequest,
			wantText:   "Body da requisição inválido.",
		},
		{
			name:       "body too large",
			body:       `{"content":"` + strings.Repeat("a", maxRequestBodyBytes) + `"}`,
			wantStatus: http.StatusRequestEntityTooLarge,
			wantText:   "Body da requisição excede o limite permitido.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := performAskRequest(t, test.body)

			if response.Code != test.wantStatus {
				t.Fatalf("status = %d, want %d; body: %s", response.Code, test.wantStatus, response.Body.String())
			}

			body := decodeAskResponse(t, response)
			if body.Response != test.wantText {
				t.Fatalf("response = %q, want %q", body.Response, test.wantText)
			}
		})
	}
}

func TestHandleUnknownQuestionReturnsLocalizedNotFound(t *testing.T) {
	response := performAskRequest(t, `{"content":"Unrelated question"}`)

	if response.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d; body: %s", response.Code, http.StatusNotFound, response.Body.String())
	}

	body := decodeAskResponse(t, response)
	if !strings.Contains(body.Response, "I don't have that specific information") {
		t.Fatalf("response = %q, want english not found message", body.Response)
	}
}

func performAskRequest(t *testing.T, body string) *httptest.ResponseRecorder {
	t.Helper()

	request := httptest.NewRequest(http.MethodPost, "/api/ask", strings.NewReader(body))
	response := httptest.NewRecorder()

	Handle().ServeHTTP(response, request)

	return response
}

func decodeAskResponse(t *testing.T, response *httptest.ResponseRecorder) AskResponse {
	t.Helper()

	var body AskResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("failed to decode response %q: %v", response.Body.String(), err)
	}

	return body
}
