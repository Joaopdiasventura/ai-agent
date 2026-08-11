# Ranking

[Index](./README.md) | [Português](../pt/ranking.md)

`internal/ranking` combines Reciprocal Rank Fusion with feature scoring. The default ranker uses fusion weight `0.55`, feature weight `0.45`, and candidate pool `80`.

RRF uses `K=60`. Source weights are lexical `1.0`, entity `1.15`, concept `1.2`, and fuzzy `0.7`. A candidate receives `sourceWeight/(K+position)` from each ranking where it appears. Raw fusion scores are normalized by the maximum raw score.

Feature scoring is a weighted average of applicable signals: fact importance, intent compatibility, target compatibility, temporal compatibility, entity coverage, concept coverage, and source diversity. Intent compatibility favors categories and predicates appropriate to the intent. Contact, education, and certification are strict; capability accepts skill, experience, project, achievement, and certification evidence with different strengths.

Entity coverage ignores person entities. This is crucial: João alone should not satisfy the semantic coverage for a specific unknown personal attribute. Concept coverage considers query concepts with score at least `0.15`.

The final score is clamped to `[0,1]`. Candidate order is deterministic with stable tie-breaking by fact ID and signal values.
