# Planning and Generation

[Index](./README.md) | [Português](../pt/planning-generation.md) | [Mathematics](./mathematics.md)

`internal/planning` converts reasoning conclusions into finite answer plans. If the conclusion is not supported, the planner emits `PlanStatusAbstain` and an abstention section. Supported conclusions become `PlanStatusReady` with sections such as lead, evidence, details, alternatives, and list.

Plan types mirror conclusion types: direct, overview, capability, experience, technology usage, comparison, and list. Plan items can represent facts, entities, support, comparison winners, alternatives, and technology usage.

`EvidenceSelector.Select` chooses compact, diverse evidence. Its base scoring term is:

$$
S_{\operatorname{baseSelect}}
=
0.62S
+
0.23D
+
0.15I
$$

The complete selection score adds the diversity adjustments implemented in `selectionScore`:

$$
S_{\operatorname{select}}
=
S_{\operatorname{baseSelect}}
+
B_{\operatorname{category}}
+
B_{\operatorname{predicate}}
-
P_{\operatorname{category}}
-
P_{\operatorname{predicate}}
$$

Where $B_{\operatorname{category}}=0.08$ for a new category, $B_{\operatorname{predicate}}=0.08$ for a new predicate, $P_{\operatorname{category}}=0.04$ once that category has already been selected at least twice, and $P_{\operatorname{predicate}}=0.06$ once that predicate has already been selected at least twice.

`internal/generation` verbalizes the plan. It does not retrieve, rank, or reason. `Materialize(plan, source)` loads only planned fact IDs and entity IDs, and validation ensures all planned material exists.

Generation has realizers for each plan type and supports Portuguese and English. Unknown output language falls back to English. `joinSentences` deduplicates full sentences; `joinNatural` deduplicates entity labels and joins with `e` or `and`.

For abstention, generation produces an internal explanation for debug. Public `Answer` still returns no response because public gating happens in `agent.executionHasResponse`.
