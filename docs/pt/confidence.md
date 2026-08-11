# Confidence

[Índice](./README.md) | [English](../en/confidence.md) | [Matemática](./mathematics.md)

`internal/confidence` calcula um score determinístico de auditoria. Não é uma probabilidade calibrada. O modo é `claim` quando o plano está ready e o reasoning é supported; é `abstention` quando o plano abstém ou o reasoning é insufficient.

Todos os scores de confidence usam a mesma forma de média ponderada sobre sinais aplicáveis:

$$
C
=
\operatorname{clamp}\left(
\frac{\sum_{i=1}^{n}w_i s_i}{\sum_{i=1}^{n}w_i}
\right)
$$

O modo claim usa qualidade da query, concordância de retrieval, separação de ranking, força da evidência, directness da evidência, cobertura semântica, grounding do plano e grounding da resposta. Os pesos enfatizam força da evidência e cobertura semântica, ainda verificando que a resposta gerada foi autorizada pelo plano.

A concordância de retrieval é:

$$
S_{\operatorname{agreement}}
=
\operatorname{clamp}\left(
0.65S_{\operatorname{topConsensus}}
+
0.35S_{\operatorname{pairwiseOverlap}}
\right)
$$

A separação é:

$$
S_{\operatorname{separation}}
=
\operatorname{clamp}\left(
\frac{\frac{s_1-s_2}{\max(s_1,10^{-6})}}{0.25}
\right)
$$

O modo abstention usa qualidade da query, ausência de evidência, grounding do plano e grounding da resposta. Ausência de evidência é mais forte quando o reasoning é insufficient e não há fatos relevantes fortes.

Scores são limitados a $[0,1]$. Níveis são high quando $C \ge 0.8$, medium quando $C \ge 0.6$ e low abaixo disso.
