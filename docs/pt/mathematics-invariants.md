# Matemática e Invariantes

[Índice](./README.md) | [English](../en/mathematics-invariants.md) | [Referência Matemática Completa](./mathematics.md)

Esta página resume as fórmulas e invariantes mais importantes. A referência matemática completa está em [Matemática](./mathematics.md).

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

Onde $k_1=1.2$ e $w_t$ é o peso do termo da query.

## Similaridade

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

## Ranking, Reasoning e Planning

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

## Confidence e Avaliação

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

## Invariantes

Generation nunca busca novos fatos; materialization só carrega fatos/entidades planejados; reasoning sem suporte vira planning abstain; respostas públicas exigem suporte e confidence claim; overlap lexical apenas com pessoa não basta para atributos pessoais arbitrários; preferência exige evidência explícita de preferência; idiomas humanos e linguagens de programação são distintos; full stack exige cobertura frontend mais backend ou evidência fullstack explícita.
