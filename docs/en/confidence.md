# Confidence

[Index](./README.md) | [Português](../pt/confidence.md) | [Mathematics](./mathematics.md)

`internal/confidence` computes a deterministic audit score. It is not a calibrated probability. The mode is `claim` when the plan is ready and reasoning is supported; it is `abstention` when the plan abstains or reasoning is insufficient.

All confidence scores use the same weighted-average shape over applicable signals:

$$
C
=
\operatorname{clamp}\left(
\frac{\sum_{i=1}^{n}w_i s_i}{\sum_{i=1}^{n}w_i}
\right)
$$

Claim mode uses query quality, retrieval agreement, ranking separation, evidence strength, evidence directness, semantic coverage, plan grounding, and answer grounding. The weights emphasize evidence strength and semantic coverage while still checking that the generated answer is authorized by the plan.

Retrieval agreement is:

$$
S_{\operatorname{agreement}}
=
\operatorname{clamp}\left(
0.65S_{\operatorname{topConsensus}}
+
0.35S_{\operatorname{pairwiseOverlap}}
\right)
$$

Separation is:

$$
S_{\operatorname{separation}}
=
\operatorname{clamp}\left(
\frac{\frac{s_1-s_2}{\max(s_1,10^{-6})}}{0.25}
\right)
$$

Abstention mode uses query quality, evidence absence, plan grounding, and answer grounding. Evidence absence is strongest when reasoning is insufficient and there are no strong relevant facts.

Scores are clamped to $[0,1]$. Levels are high at $C \ge 0.8$, medium at $C \ge 0.6$, and low otherwise.
