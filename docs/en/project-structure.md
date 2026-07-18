# Project Structure

## Root Files

### `go.mod`

Defines module name `ai-agent` and Go version `1.26.3`. No external module dependencies are declared.

### `vercel.json`

Defines the Vercel function configuration. It disables framework detection with `"framework": null`, leaves `buildCommand` empty, sets `api/index.go` `maxDuration` to `10`, and routes all requests to `/api/index`.

## Entrypoints

### `cmd/ai-agent/main.go`

CLI bootstrap. Its only responsibility is to call `app.Run()`. This keeps CLI startup separate from the application workflow.

### `api/index.go`

Serverless bootstrap. It obtains an HTTP handler from `server.Handler()` and delegates the request. On handler creation error, it returns a JSON `503`.

## HTTP Layer

### `server/server.go`

Owns HTTP mux creation, `QUERY /ask` registration, CORS handling, JSON helper functions, and lazy handler caching with a mutex.

### `internal/handlers/ask/`

Owns the public JSON contract for the ask endpoint. It validates body shape and delegates valid content to `app.AgentResponse`.

## Application Core

### `internal/app/`

Contains the core interface used by both CLI and HTTP. `agent.go` initializes documents and the search engine, `config.go` stores runtime constants, and `chatbot.go` implements the interactive console loop.

## Data And Domain

### `internal/domain/`

Defines `Document`, the shared structure passed across knowledge, TF-IDF, and search.

### `internal/knowledge/`

Stores the static facts. `Documents()` returns pointers to the compiled document array.

## Processing Packages

### `internal/tokenizer/`

Lowercases text, removes non-alphanumeric separators, splits tokens, and removes configured stopwords.

### `internal/nlp/`

Contains deterministic language, query, intent, entity, answer mode, and technology logic.

### `internal/tfidf/`

Calculates frequency maps and TF-IDF vectors.

### `internal/search/`

Coordinates retrieval, filtering, scoring, sorting, and relevance checks.

### `internal/answer/`

Builds final responses from ranked results.

## Tests

Tests are placed next to the package they validate:

- `internal/nlp/language_test.go`
- `internal/search/engine_test.go`
- `internal/app/agent_test.go`
- `internal/handlers/ask/ask_test.go`
- `internal/knowledge/documents_test.go`

They validate current behavior without depending on external services.
