# Planning and Generation

[Index](./README.md) | [Português](../pt/planning-generation.md)

`internal/planning` converts reasoning conclusions into finite answer plans. If the conclusion is not supported, the planner emits `PlanStatusAbstain` and an abstention section. Supported conclusions become `PlanStatusReady` with sections such as lead, evidence, details, alternatives, and list.

Plan types mirror conclusion types: direct, overview, capability, experience, technology usage, comparison, and list. Plan items can represent facts, entities, support, comparison winners, alternatives, and technology usage.

`EvidenceSelector.Select` chooses compact, diverse evidence. Its base score is `0.62*score + 0.23*directness + 0.15*importance`, then it adds small bonuses for new categories/predicates and penalties for repeated ones. This reduces redundant evidence such as repeated language-level facts.

`internal/generation` verbalizes the plan. It does not retrieve, rank, or reason. `Materialize(plan, source)` loads only planned fact IDs and entity IDs, and validation ensures all planned material exists.

Generation has realizers for each plan type and supports Portuguese and English. Unknown output language falls back to English. `joinSentences` deduplicates full sentences; `joinNatural` deduplicates entity labels and joins with `e` or `and`.

For abstention, generation produces an internal explanation for debug. Public `Answer` still returns no response because public gating happens in `agent.executionHasResponse`.
