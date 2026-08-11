# Agent, CLI, and HTTP

[Index](./README.md) | [Português](../pt/interfaces.md)

`agent.Service` is the public application service. `New()` constructs and validates the full in-memory pipeline. `Answer(question)` trims the question, rejects empty input with `ErrEmptyQuestion`, executes every layer, and returns `agent.Result` with response, response flag, language, confidence score, and confidence level.

`Debug(question)` executes the same pipeline and returns `DebugResult`: question, public result, query, retrieval result, ranking result, reasoning result, plan, generated answer, and confidence result. Debug is the best interface for identifying whether a failure belongs to query, retrieval, ranking, reasoning, planning, materialization, generation, or confidence.

The public response gate requires all of these: `PlanStatusReady`, `SupportSupported`, non-empty generated answer, and confidence mode `claim`. Abstentions can still have internal generation text in debug, but `Answer` exposes an empty response.

`internal/cli` provides the terminal runner. It skips empty lines, exits on `sair`, `exit`, `quit`, or `encerrar`, prints supported answers, and prints localized fallback text for unsupported results.

HTTP is handled by `server` and `internal/handlers/ask`. `POST /ask` accepts `{ "content": "..." }`. Request bodies are capped at 2048 bytes, unknown fields are rejected, and CORS allows `https://joaopdias.dev.br` and `http://localhost:4200`. `api/index.go` is the Vercel entry point.
