# Ranking

[Índice](./README.md) | [English](../en/ranking.md) | [Matemática](./mathematics.md)

`internal/ranking` combina Reciprocal Rank Fusion com feature scoring. O ranker padrão usa peso de fusão $0.55$, peso de features $0.45$ e candidate pool $80$.

RRF usa $k=60$. Os pesos de fonte são lexical $1.0$, entity $1.15$, concept $1.2$ e fuzzy $0.7$.

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

Feature scoring é uma média ponderada dos sinais aplicáveis:

$$
S_{\operatorname{feature}}(d)
=
\frac{\sum_{i=1}^{n}w_i s_i(d)}{\sum_{i=1}^{n}w_i}
$$

O score final é:

$$
S_{\operatorname{rank}}(d)
=
\operatorname{clamp}\left(0.55S_{\operatorname{fusion}}(d)+0.45S_{\operatorname{feature}}(d)\right)
$$

Cobertura de entidade ignora entidades de pessoa. Cobertura de conceito considera conceitos da query com score mínimo $0.15$. A ordem dos candidatos é determinística, com desempates estáveis por ID de fato e valores de sinais.
