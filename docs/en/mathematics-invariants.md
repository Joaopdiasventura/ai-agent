# Mathematics and Invariants

[Index](./README.md) | [Português](../pt/mathematics-invariants.md) | [Full Mathematics Reference](./mathematics.md)

This page summarizes the most important formulas and invariants. The full mathematical reference is in [Mathematics](./mathematics.md).

## BM25F

$$
\operatorname{IDF}(t)
=
\ln\left(
1+
\frac{N-df(t)+0.5}{df(t)+0.5}
\right)
$$

$$
\operatorname{norm}_{f}(d)
=
(1-b_f)+b_f\frac{\lvert d_f\rvert}{\operatorname{avgdl}_f}
$$

$$
S_{\operatorname{BM25F}}(t,d)
=
w_t\cdot\operatorname{IDF}(t)\cdot
\frac{\operatorname{wf}(t,d)(k_1+1)}{\operatorname{wf}(t,d)+k_1}
$$

Where $k_1=1.2$ and $w_t$ is the query term weight.

## Similarity

$$
S_{\operatorname{edit}}(a,b)
=
1-
\frac{\operatorname{lev}(a,b)}{\max(\lvert a\rvert,\lvert b\rvert)}
$$

$$
S_{\operatorname{ngram}}(a,b)
=
\frac{\lvert G(a)\cap G(b)\rvert}{\lvert G(a)\cup G(b)\rvert}
$$

$$
S_{\operatorname{fuzzy}}(a,b)
=
\max\left(S_{\operatorname{edit}}(a,b),S_{\operatorname{ngram}}(a,b)\right)
$$

## Ranking, Reasoning, and Planning

$$
\operatorname{RRF}_{\operatorname{raw}}(d)
=
\sum_{r\in R(d)}\frac{w_r}{k+\operatorname{rank}_r(d)}
$$

$$
S_{\operatorname{rank}}(d)
=
\operatorname{clamp}\left(0.55S_{\operatorname{fusion}}(d)+0.45S_{\operatorname{feature}}(d)\right)
$$

$$
S_{\operatorname{group}}
=
\operatorname{clamp}\left(
0.60S_{\operatorname{groupEvidence}}
+0.20S_{\operatorname{conceptCoverage}}
+0.12S_{\operatorname{diversity}}
+0.08S_{\operatorname{quantity}}
\right)
$$

$$
S_{\operatorname{select}}
=
0.62S+0.23D+0.15I+B_{\operatorname{category}}+B_{\operatorname{predicate}}-P_{\operatorname{category}}-P_{\operatorname{predicate}}
$$

## Confidence and Evaluation

$$
C
=
\operatorname{clamp}\left(
\frac{\sum_{i=1}^{n}w_i s_i}{\sum_{i=1}^{n}w_i}
\right)
$$

$$
\operatorname{Accuracy}
=
\frac{\operatorname{passed}}{\operatorname{total}}
$$

## Invariants

Generation never fetches new facts; materialization only loads planned facts/entities; unsupported reasoning becomes abstain planning; public answers require support and claim confidence; person-only lexical overlap is not enough for arbitrary personal attributes; preference requires explicit preference evidence; human languages and programming languages are distinct; full stack requires frontend plus backend coverage or explicit fullstack evidence.
