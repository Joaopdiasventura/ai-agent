# Ranking

[Index](./README.md) | [Português](../pt/ranking.md) | [Mathematics](./mathematics.md)

`internal/ranking` combines Reciprocal Rank Fusion with feature scoring. The default ranker uses fusion weight $0.55$, feature weight $0.45$, and candidate pool $80$.

RRF uses $k=60$. Source weights are lexical $1.0$, entity $1.15$, concept $1.2$, and fuzzy $0.7$.

$$
\operatorname{RRF}_{\operatorname{raw}}(d)
=
\sum_{r\in R(d)}
\frac{w_r}{k+\operatorname{rank}_r(d)}
$$

$$
S_{\operatorname{fusion}}(d)
=
\frac{\operatorname{RRF}_{\operatorname{raw}}(d)}{\max_x\operatorname{RRF}_{\operatorname{raw}}(x)}
$$

Feature scoring is a weighted average of applicable signals:

$$
S_{\operatorname{feature}}(d)
=
\frac{\sum_{i=1}^{n}w_i s_i(d)}{\sum_{i=1}^{n}w_i}
$$

The final score is:

$$
S_{\operatorname{rank}}(d)
=
\operatorname{clamp}\left(0.55S_{\operatorname{fusion}}(d)+0.45S_{\operatorname{feature}}(d)\right)
$$

Entity coverage ignores person entities. Concept coverage considers query concepts with score at least $0.15$. Candidate order is deterministic with stable tie-breaking by fact ID and signal values.
