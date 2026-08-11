# Domain and Knowledge

[Index](./README.md) | [Português](../pt/domain-knowledge.md)

`internal/domain` defines shared contracts: `Entity`, `Fact`, `Query`, matches, relations, fact categories, intents, query targets, periods, temporal scopes, localized text, and evidence signals.

An `Entity` has an ID, type, localized name, and aliases. Entity types are semantically important: human languages use `EntityTypeLanguage`, while technologies use `EntityTypeTechnology`. This separation prevents Portuguese and English from being treated as Java, Go, JavaScript, or TypeScript.

A `Fact` is the evidence atom. It has subject, predicate, object, category, concepts, context, localized statement, optional period, importance, and source. Validation checks IDs, object shape, current period consistency, localized statements, and importance range.

`internal/knowledge` is the source of truth. It contains João, companies, projects, roles, institutions, certifications, human languages, and technology entities. Skill facts are explicit, not generated for every technology. Current skill facts include JavaScript, TypeScript, Java, Go, Angular, React, Next.js, Spring Boot, Node.js, NestJS, PostgreSQL, MongoDB, Redis, RabbitMQ, Kafka, SQS, Docker, Terraform, Kubernetes, and AWS.

The key knowledge invariant is absence-aware behavior. If there is no fact about age, birth date, salary, height, marital status, favorite food, favorite game, or explicit preference, the agent should abstain instead of answering with unrelated João facts.
