# Testing

## Command

```powershell
go test ./...
```

## Current Coverage Areas

### Language Detection

`internal/nlp/language_test.go` validates weighted language detection, including the case where Portuguese `Me fale sobre o auronix` must not be classified as English because of `me`.

### Search Pipeline

`internal/search/engine_test.go` validates that Portuguese and English questions return documents in the detected language, that Auronix can be found in Portuguese, that English technology questions return multiple results, and that invalid inputs do not produce answers.

### App Contract

`internal/app/agent_test.go` validates `AgentResponse` for Portuguese and English Auronix questions and localized fallback behavior.

### HTTP Contract

`internal/handlers/ask/ask_test.go` validates successful JSON requests, empty content, invalid JSON, unknown fields, oversized bodies, and not-found responses.

### Knowledge Base Invariants

`internal/knowledge/documents_test.go` validates:

- 140 documents;
- 70 Portuguese documents;
- 70 English documents;
- unique IDs;
- required fields;
- `-pt` and `-en` suffixes.

## Testing Boundaries

The tests do not assert exact full response text because template and connector selection is random. They also avoid external services because the application does not require network, database, or model calls for the documented flows.
