# Evaluation and Testing

[Index](./README.md) | [Português](../pt/evaluation-testing.md)

The repository has package tests for every major layer. Tests validate data validation, localized fallback, normalization, stopwords, language detection, fuzzy similarity, ontology aliasing and expansion, index structures, retrieval sources, ranking fusion/features, reasoning support, planning shape, generation grounding, confidence modes, service behavior, CLI behavior, HTTP request validation, and evaluation reporting.

`internal/evaluation` is a regression harness over `agent.Debug`. A case can assert response presence, language, intent, target, entities, concepts, comparison winner, required generated facts, forbidden facts, plan status, confidence mode, confidence bounds, expected response text, forbidden response text, and generated debug text.

The regression corpus covers direct contact answers, JavaScript/TypeScript/Java/Go capabilities, typos for Docker and Kubernetes, programming language lists, framework lists, technology lists, human language lists, unknown age, unknown preference, backend, full stack, technology usage, comparisons, and project overviews.

Run:

```bash
go test ./...
go test -race ./...
```

The regression corpus test requires 100% accuracy. Existing tests should not be reduced to make changes pass; behavior should be corrected instead.
