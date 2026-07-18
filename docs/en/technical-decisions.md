# Technical Decisions And Trade-offs

## Static Knowledge Base

- Decision: store documents in Go source code.
- Benefit: no runtime file I/O and no dependency on an external data file.
- Cost: content changes are code changes.
- Consequence: tests can validate document count, IDs, languages, and required fields directly.

## Pointer-Based Document Flow

- Decision: expose `[]*domain.Document` from `knowledge.Documents()`.
- Benefit: search results and vectors refer to existing document values.
- Cost: callers receive references to shared in-memory data.
- Consequence: the pipeline avoids copying document structs during retrieval.

## Lexical Retrieval

- Decision: use tokenization, query expansion, TF-IDF, cosine similarity, and boosts.
- Benefit: deterministic behavior without external services.
- Cost: relevance depends on token overlap and configured rules.
- Consequence: unknown wording can fail if no token overlap, expansion, alias, or boost connects it to a document.

## Rule-Based Language Detection

- Decision: sum token weights for Portuguese and English.
- Benefit: simple behavior that handles ambiguous tokens such as `me`.
- Cost: detection depends on manually configured weights.
- Consequence: ties and unknown signals default to Portuguese.

## Shared App Core

- Decision: CLI and HTTP both use `app.AgentResponse`.
- Benefit: the same retrieval and answer logic serves both interfaces.
- Cost: interface-specific behavior must stay outside the core.
- Consequence: tests can validate core behavior through `internal/app` without invoking CLI or HTTP.

## Strict HTTP Contract

- Decision: use `DisallowUnknownFields`, a 2048-byte body limit, one JSON object per request, and a required `content` field.
- Benefit: invalid requests fail before reaching the agent.
- Cost: clients must match the exact JSON shape.
- Consequence: transport errors are separated from answer-not-found behavior.

## Random Template Selection

- Decision: choose among language-specific templates and connectors using `math/rand/v2`.
- Benefit: repeated answers can vary in wording.
- Cost: exact response strings are not stable.
- Consequence: tests assert stable properties such as language, status, and presence of known facts.

## HTTP Method `QUERY`

- Decision: register `QUERY /ask`.
- Benefit: the documented contract matches the current server implementation.
- Cost: `QUERY` is less common than methods such as `POST`.
- Consequence: API consumers must send requests using the implemented method.
