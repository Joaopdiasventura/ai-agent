# Ontology

[Index](./README.md) | [Português](../pt/ontology.md)

`internal/ontology` provides the semantic graph used by query expansion, indexing, retrieval, ranking, and reasoning. It validates concept IDs, names, parent references, relation endpoints, relation weights, and entity bindings.

Concepts cover contact, education, certification, languages, software development, programming language, framework, runtime, frontend, backend, fullstack, database, messaging, devops, infrastructure, infrastructure as code, orchestration, containers, cloud, AWS, Kubernetes, distributed systems, concurrency, performance, and related domains.

Aliases are localized and normalized. `Aliases(targetLanguage)` returns weighted aliases with language preference. Longer aliases are sorted before shorter aliases so phrases such as `full stack development` or `sistemas distribuidos` can match before overlapping shorter words.

`ResolveAlias` handles exact normalized concept aliases. `Expand(id,maxDepth)` performs breadth-first propagation through parents and weighted semantic relations. Parent propagation uses `0.85`; relation propagation uses each relation weight; the best score per concept is retained.

Entity bindings classify technologies and languages. JavaScript, TypeScript, Java, and Go bind to programming-language; Angular, React, Next.js, Spring Boot, and NestJS bind to framework; Node.js binds to runtime; databases, messaging tools, cloud, containers, orchestration, and infrastructure tools bind to their own concepts. Portuguese and English bind to the human language concept, not to programming-language.

Full Stack exists as a concept with PT-BR and English aliases. Expansion can retrieve related frontend/backend evidence, but reasoning still requires enough coverage before support is granted.
