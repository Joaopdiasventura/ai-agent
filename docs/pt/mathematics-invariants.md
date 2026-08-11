# Matemática e Invariantes

[Índice](./README.md) | [English](../en/mathematics-invariants.md)

O IDF do BM25F é `ln(1 + (N - df + 0.5)/(df + 0.5))`. A normalização de campo é `(1-b)+b*fieldLength/avgFieldLength`. A frequência ponderada de campo é a soma de `fieldWeight * frequency / normalization`. A contribuição do termo é `termWeight * idf * (wf*(K1+1))/(wf+K1)` com `K1=1.2`.

Similaridade de edição é `1 - levenshteinDistance(a,b)/max(len(a),len(b))`. Similaridade por n-gram de caracteres é Jaccard sobre conjuntos de n-grams. Similaridade fuzzy é o máximo entre edição e n-gram.

RRF contribui `sourceWeight/(K+position)` com `K=60`. Os pesos de fonte são lexical `1.0`, entity `1.15`, concept `1.2` e fuzzy `0.7`.

O score de grupo do reasoning é `0.6*evidenceStrength + 0.2*conceptCoverage + 0.12*diversity + 0.08*quantity`. O score de seleção do planning começa com `0.62*score + 0.23*directness + 0.15*importance`.

Invariantes centrais: generation nunca busca novos fatos; materialization só carrega fatos/entidades planejados; reasoning sem suporte vira planning abstain; respostas públicas exigem suporte e confidence claim; overlap lexical apenas com pessoa não basta para atributos pessoais arbitrários; preferência exige evidência explícita de preferência; idiomas humanos e linguagens de programação são distintos; full stack exige cobertura frontend mais backend ou evidência fullstack explícita.
