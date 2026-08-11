# Reasoning

[Index](./README.md) | [Português](../pt/reasoning.md)

`internal/reasoning` is the evidence sufficiency layer. It receives ranked candidates and produces a conclusion with type, status, focus entity, focus concept, evidence, and optional groups.

`EvidenceBuilder` computes evidence rank, score, importance, directness, matched entities, and matched concepts. Directness combines explicit entity match and direct concept match; unrelated evidence falls back to low directness.

General direct relevance is strict. Non-person query entities must be referenced by the fact. Directly matched concepts must appear in the fact. If there is only a person entity and no supported target/category, arbitrary facts about that person are not enough.

Capability reasoning ignores person entities as capability targets. It keeps only evidence referencing an explicit non-person capability target or a directly matched concept. A supported capability requires a direct match with evidence score at least `0.35`.

Full Stack has special semantic status: JavaScript plus TypeScript alone is not enough. Support requires an explicit fullstack fact or sufficient evidence covering both frontend and backend concepts.

List reasoning groups evidence by requested entity type. Programming language, framework, runtime, database, messaging, cloud, devops, and infrastructure concepts list technology entities; human language lists use `EntityTypeLanguage`. Filtering prevents frameworks, databases, clouds, and human languages from leaking into programming-language lists.
