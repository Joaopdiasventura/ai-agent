# Mathematics and Invariants

[Index](./README.md) | [Português](../pt/mathematics-invariants.md)

BM25F IDF is `ln(1 + (N - df + 0.5)/(df + 0.5))`. Field normalization is `(1-b)+b*fieldLength/avgFieldLength`. Weighted field frequency is the sum of `fieldWeight * frequency / normalization`. Term contribution is `termWeight * idf * (wf*(K1+1))/(wf+K1)` with `K1=1.2`.

Edit similarity is `1 - levenshteinDistance(a,b)/max(len(a),len(b))`. Character n-gram similarity is Jaccard over n-gram sets. Fuzzy similarity is the maximum of edit and n-gram similarity.

RRF contributes `sourceWeight/(K+position)` with `K=60`. Source weights are lexical `1.0`, entity `1.15`, concept `1.2`, and fuzzy `0.7`.

Reasoning group score is `0.6*evidenceStrength + 0.2*conceptCoverage + 0.12*diversity + 0.08*quantity`. Planning selection score starts with `0.62*score + 0.23*directness + 0.15*importance`.

Core invariants: generation never fetches new facts; materialization only loads planned facts/entities; unsupported reasoning becomes abstain planning; public answers require support and claim confidence; person-only lexical overlap is not enough for arbitrary personal attributes; preference requires explicit preference evidence; human languages and programming languages are distinct; full stack requires frontend plus backend coverage or explicit fullstack evidence.
