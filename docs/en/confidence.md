# Confidence

[Index](./README.md) | [Português](../pt/confidence.md)

`internal/confidence` computes a deterministic audit score. It is not a calibrated probability. The mode is `claim` when the plan is ready and reasoning is supported; it is `abstention` when the plan abstains or reasoning is insufficient.

Claim mode uses weighted signals: query quality, retrieval agreement, ranking separation, evidence strength, evidence directness, semantic coverage, plan grounding, and answer grounding. The weights emphasize evidence strength and semantic coverage while still checking that the generated answer is authorized by the plan.

Abstention mode uses query quality, evidence absence, plan grounding, and answer grounding. Evidence absence is strongest when reasoning is insufficient and there are no strong relevant facts.

Query quality combines language, intent, target, semantic hints, and lexical term count. Retrieval agreement compares top rankings and pairwise overlap. Separation measures the margin between top ranked candidates or top reasoning groups. Plan grounding checks planned facts against reasoning evidence. Answer grounding checks generated fact IDs against planned facts.

Scores are clamped to `[0,1]`. Levels are high at `>=0.8`, medium at `>=0.6`, and low otherwise.
