# Ranking

[Índice](./README.md) | [English](../en/ranking.md)

`internal/ranking` combina Reciprocal Rank Fusion com feature scoring. O ranker padrão usa peso de fusão `0.55`, peso de features `0.45` e candidate pool `80`.

RRF usa `K=60`. Os pesos de fonte são lexical `1.0`, entity `1.15`, concept `1.2` e fuzzy `0.7`. Um candidato recebe `sourceWeight/(K+position)` de cada ranking em que aparece. Scores brutos de fusão são normalizados pelo maior score bruto.

Feature scoring é uma média ponderada dos sinais aplicáveis: importância do fato, compatibilidade de intent, compatibilidade de target, compatibilidade temporal, cobertura de entidade, cobertura de conceito e diversidade de fonte. Compatibilidade de intent favorece categorias e predicados apropriados à intenção. Contact, education e certification são rígidos; capability aceita evidência de skill, experience, project, achievement e certification com forças diferentes.

Cobertura de entidade ignora entidades de pessoa. Isso é crucial: João sozinho não deve satisfazer a cobertura semântica para um atributo pessoal específico desconhecido. Cobertura de conceito considera conceitos da query com score mínimo `0.15`.

O score final é limitado a `[0,1]`. A ordem dos candidatos é determinística, com desempates estáveis por ID de fato e valores de sinais.
