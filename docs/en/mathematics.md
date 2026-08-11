# Mathematics

[Index](./README.md) | [Português](../pt/mathematics.md)

This page is the central mathematical reference for the current implementation. The formulas below preserve the Go code semantics.

## Edit Similarity

$$
S_{\operatorname{edit}}(a,b)
=
1
-
\frac{
\operatorname{lev}(a,b)
}{
\max\left(\lvert a \rvert, \lvert b \rvert\right)
}
$$

Where:

- $\operatorname{lev}(a,b)$ is the Levenshtein distance after normalization.
- $\lvert a \rvert$ and $\lvert b \rvert$ are normalized rune counts.
- Empty strings follow the code paths: both empty return $1$; only one empty returns $0$.

## Character N-Gram Similarity

$$
S_{\operatorname{ngram}}(a,b)
=
\frac{
\lvert G(a) \cap G(b) \rvert
}{
\lvert G(a) \cup G(b) \rvert
}
$$

Where $G(x)$ is the set of unique character n-grams produced for $x$ with sizes $3$ through $4$ in fuzzy similarity.

## Fuzzy Similarity

$$
S_{\operatorname{fuzzy}}(a,b)
=
\max\left(
S_{\operatorname{edit}}(a,b),
S_{\operatorname{ngram}}(a,b)
\right)
$$

## Inverse Document Frequency

$$
\operatorname{IDF}(t)
=
\ln\left(
1
+
\frac{
N - df(t) + 0.5
}{
df(t) + 0.5
}
\right)
$$

Where:

- $N$ is the number of documents for the target language.
- $df(t)$ is the document frequency of term $t$.

## BM25F Field Normalization

$$
\operatorname{norm}_{f}(d)
=
\left(1-b_f\right)
+
b_f
\frac{
\lvert d_f \rvert
}{
\operatorname{avgdl}_f
}
$$

Where:

- $f$ is an indexed field.
- $b_f$ is the field-specific normalization parameter.
- $\lvert d_f \rvert$ is the length of field $f$ in document $d$.
- $\operatorname{avgdl}_f$ is the average length for field $f$.

## BM25F Weighted Frequency

$$
\operatorname{wf}(t,d)
=
\sum_{f \in F(t,d)}
w_f
\frac{
tf_f(t,d)
}{
\operatorname{norm}_{f}(d)
}
$$

Where:

- $w_f$ is the field weight.
- $tf_f(t,d)$ is the frequency of term $t$ in field $f$ of document $d$.
- $F(t,d)$ is the set of fields where $t$ appears in $d$.

## BM25F Term Contribution

$$
S_{\operatorname{BM25F}}(t,d)
=
w_t
\cdot
\operatorname{IDF}(t)
\cdot
\frac{
\operatorname{wf}(t,d)(k_1+1)
}{
\operatorname{wf}(t,d)+k_1
}
$$

Where $w_t$ is the query term weight and $k_1=1.2$.

## Reciprocal Rank Fusion

$$
\operatorname{RRF}_{\operatorname{raw}}(d)
=
\sum_{r \in R(d)}
\frac{
w_r
}{
k + \operatorname{rank}_r(d)
}
$$

$$
S_{\operatorname{fusion}}(d)
=
\frac{
\operatorname{RRF}_{\operatorname{raw}}(d)
}{
\max_{x}\operatorname{RRF}_{\operatorname{raw}}(x)
}
$$

Where $k=60$, $r$ is a retrieval source, and $w_r$ is the source weight.

## Ranking Feature Average

$$
S_{\operatorname{feature}}(d)
=
\frac{
\sum_{i=1}^{n} w_i s_i(d)
}{
\sum_{i=1}^{n} w_i
}
$$

Only applicable ranking feature signals participate in the sum.

## Final Ranking Score

$$
S_{\operatorname{rank}}(d)
=
\operatorname{clamp}\left(
0.55S_{\operatorname{fusion}}(d)
+
0.45S_{\operatorname{feature}}(d)
\right)
$$

## Entity and Concept Coverage

$$
S_{\operatorname{coverage}}
=
\frac{
\sum_{m \in M} w_m \mathbf{1}_{\operatorname{matched}}(m)
}{
\sum_{m \in M} w_m
}
$$

For ranking entity coverage, person entities are skipped. For concept coverage, concepts with score below $0.15$ are skipped.

## Evidence Directness

$$
D(e)
=
\begin{cases}
\operatorname{clamp}\left(0.55S_{\operatorname{entity}} + 0.45S_{\operatorname{concept}}\right), & S_{\operatorname{entity}}>0 \land S_{\operatorname{concept}}>0 \\
\operatorname{clamp}\left(0.85S_{\operatorname{entity}}\right), & S_{\operatorname{entity}}>0 \\
\operatorname{clamp}\left(0.90S_{\operatorname{concept}}\right), & S_{\operatorname{concept}}>0 \\
0.25, & \text{otherwise}
\end{cases}
$$

## Reasoning Group Evidence Strength

$$
S_{\operatorname{groupEvidence}}
=
\operatorname{clamp}\left(
\frac{
\sum_{i=1}^{n} w_i
\operatorname{clamp}\left(
0.7S_i+0.2D_i+0.1I_i
\right)
}{
\sum_{i=1}^{n} w_i
}
\right)
$$

## Reasoning Group Diversity

$$
S_{\operatorname{diversity}}
=
\operatorname{clamp}\left(
0.45\min\left(\frac{\lvert C \rvert}{3},1\right)
+
0.55\min\left(\frac{\lvert P \rvert}{4},1\right)
\right)
$$

## Reasoning Group Quantity

$$
S_{\operatorname{quantity}}(n)
=
\operatorname{clamp}\left(
1-e^{-n/3}
\right)
$$

## Reasoning Group Score

$$
S_{\operatorname{group}}
=
\operatorname{clamp}\left(
0.60S_{\operatorname{groupEvidence}}
+
0.20S_{\operatorname{conceptCoverage}}
+
0.12S_{\operatorname{diversity}}
+
0.08S_{\operatorname{quantity}}
\right)
$$

## Planning Selection Score

$$
S_{\operatorname{select}}
=
0.62S
+
0.23D
+
0.15I
+
B_{\operatorname{category}}
+
B_{\operatorname{predicate}}
-
P_{\operatorname{category}}
-
P_{\operatorname{predicate}}
$$

Where the bonuses are $0.08$ for a new category and $0.08$ for a new predicate; the penalties are $0.04$ for a category already selected at least twice and $0.06$ for a predicate already selected at least twice.

## Retrieval Agreement

$$
S_{\operatorname{agreement}}
=
\operatorname{clamp}\left(
0.65S_{\operatorname{topConsensus}}
+
0.35S_{\operatorname{pairwiseOverlap}}
\right)
$$

Pairwise overlap uses Jaccard similarity over top-$5$ fact sets:

$$
J(A,B)
=
\frac{
\lvert A \cap B \rvert
}{
\lvert A \cup B \rvert
}
$$

## Separation

$$
S_{\operatorname{separation}}
=
\operatorname{clamp}\left(
\frac{
\frac{s_1-s_2}{\max(s_1,10^{-6})}
}{
0.25
}
\right)
$$

Where $s_1$ and $s_2$ are the top two candidate or group scores.

## Confidence

$$
C
=
\operatorname{clamp}\left(
\frac{
\sum_{i=1}^{n} w_i s_i
}{
\sum_{i=1}^{n} w_i
}
\right)
$$

Only applicable confidence signals with positive weights participate.

## Evaluation Accuracy

$$
\operatorname{Accuracy}
=
\frac{
\operatorname{passed}
}{
\operatorname{total}
}
$$

If $\operatorname{total}=0$, the implementation returns $0$.
