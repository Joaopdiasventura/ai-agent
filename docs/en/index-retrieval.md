# Indexing and Retrieval

[Index](./README.md) | [Português](../pt/index-retrieval.md)

`internal/index` builds one Portuguese and one English document per fact. Documents have weighted fields: `statement`, `subject`, `object`, `concept`, `context`, `predicate`, and `category`. The index stores postings, document frequencies, vocabulary, n-gram term maps, average document lengths, average field lengths, and structural indexes by subject, entity, concept, category, and context.

`retrieval.HybridRetriever.Search` runs four retrievers independently: lexical, entity, concept, and fuzzy. Returning separate rankings preserves source evidence for ranking and confidence.

Lexical retrieval uses BM25F. Numeric terms get weight `1.35`, terms of length at most two get `1.15`, and other terms get `1.0`. Field weights are strongest for subject, concept, object, and context.

Entity retrieval uses explicit query entities and structural entity indexes. Person entities are intentionally weak (`0.25`) while projects and technologies are strong. This is one of the safeguards against a query containing João retrieving arbitrary facts about João.

Concept retrieval uses query concepts and structural concept indexes. Direct concept matches are stronger than expanded concepts. Fact importance is blended into the score.

Fuzzy retrieval is separate from entity fuzzy extraction. It considers only query terms of length at least four with no exact postings, gets candidate terms from character n-grams, applies length-sensitive thresholds, and scores with similarity, IDF, field weight, and frequency.
