# Matemática

[Índice](./README.md) | [English](../en/mathematics.md)

Esta página é a referência matemática central da implementação atual. As fórmulas abaixo preservam a semântica do código Go.

## Similaridade de Edição

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

Onde:

- $\operatorname{lev}(a,b)$ é a distância de Levenshtein após normalização.
- $\lvert a \rvert$ e $\lvert b \rvert$ são contagens de runes normalizadas.
- Strings vazias seguem os caminhos do código: ambas vazias retornam $1$; apenas uma vazia retorna $0$.

## Similaridade por N-Grams de Caracteres

$$
S_{\operatorname{ngram}}(a,b)
=
\frac{
\lvert G(a) \cap G(b) \rvert
}{
\lvert G(a) \cup G(b) \rvert
}
$$

Onde $G(x)$ é o conjunto de n-grams de caracteres únicos produzido para $x$ com tamanhos $3$ a $4$ em fuzzy similarity.

## Similaridade Fuzzy

$$
S_{\operatorname{fuzzy}}(a,b)
=
\max\left(
S_{\operatorname{edit}}(a,b),
S_{\operatorname{ngram}}(a,b)
\right)
$$

## Frequência Documental Inversa

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

Onde:

- $N$ é o número de documentos do idioma alvo.
- $df(t)$ é a frequência documental do termo $t$.

## Normalização de Campo do BM25F

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

Onde:

- $f$ é um campo indexado.
- $b_f$ é o parâmetro de normalização específico do campo.
- $\lvert d_f \rvert$ é o tamanho do campo $f$ no documento $d$.
- $\operatorname{avgdl}_f$ é o tamanho médio do campo $f$.

## Frequência Ponderada do BM25F

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

Onde:

- $w_f$ é o peso do campo.
- $tf_f(t,d)$ é a frequência do termo $t$ no campo $f$ do documento $d$.
- $F(t,d)$ é o conjunto de campos em que $t$ aparece em $d$.

## Contribuição de Termo do BM25F

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

Onde $w_t$ é o peso do termo da query e $k_1=1.2$.

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

Onde $k=60$, $r$ é uma fonte de retrieval e $w_r$ é o peso da fonte.

## Média de Features do Ranking

$$
S_{\operatorname{feature}}(d)
=
\frac{
\sum_{i=1}^{n} w_i s_i(d)
}{
\sum_{i=1}^{n} w_i
}
$$

Somente sinais de feature aplicáveis participam da soma.

## Score Final de Ranking

$$
S_{\operatorname{rank}}(d)
=
\operatorname{clamp}\left(
0.55S_{\operatorname{fusion}}(d)
+
0.45S_{\operatorname{feature}}(d)
\right)
$$

## Cobertura de Entidade e Conceito

$$
S_{\operatorname{coverage}}
=
\frac{
\sum_{m \in M} w_m \mathbf{1}_{\operatorname{matched}}(m)
}{
\sum_{m \in M} w_m
}
$$

Na cobertura de entidades do ranking, entidades pessoa são ignoradas. Na cobertura de conceitos, conceitos com score abaixo de $0.15$ são ignorados.

## Directness da Evidência

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

## Força de Evidência do Grupo

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

## Diversidade do Grupo

$$
S_{\operatorname{diversity}}
=
\operatorname{clamp}\left(
0.45\min\left(\frac{\lvert C \rvert}{3},1\right)
+
0.55\min\left(\frac{\lvert P \rvert}{4},1\right)
\right)
$$

## Quantidade do Grupo

$$
S_{\operatorname{quantity}}(n)
=
\operatorname{clamp}\left(
1-e^{-n/3}
\right)
$$

## Score do Grupo de Reasoning

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

## Score de Seleção do Planning

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

Os bônus são $0.08$ para uma nova categoria e $0.08$ para um novo predicado; as penalidades são $0.04$ para categoria já selecionada pelo menos duas vezes e $0.06$ para predicado já selecionado pelo menos duas vezes.

## Concordância de Retrieval

$$
S_{\operatorname{agreement}}
=
\operatorname{clamp}\left(
0.65S_{\operatorname{topConsensus}}
+
0.35S_{\operatorname{pairwiseOverlap}}
\right)
$$

O overlap par a par usa similaridade de Jaccard sobre conjuntos de fatos top-$5$:

$$
J(A,B)
=
\frac{
\lvert A \cap B \rvert
}{
\lvert A \cup B \rvert
}
$$

## Separação

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

Onde $s_1$ e $s_2$ são os dois maiores scores de candidatos ou grupos.

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

Somente sinais de confidence aplicáveis e com peso positivo participam.

## Accuracy da Avaliação

$$
\operatorname{Accuracy}
=
\frac{
\operatorname{passed}
}{
\operatorname{total}
}
$$

Se $\operatorname{total}=0$, a implementação retorna $0$.

