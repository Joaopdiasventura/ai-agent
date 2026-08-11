# Language and Query

[Index](./README.md) | [Português](../pt/language-query.md)

`internal/language` normalizes text, tokenizes, detects language, removes stopwords, and computes similarity. `Normalize` trims, lowercases, folds common accents, keeps letters and digits, preserves decimal separators between digits, and turns other separators into spaces. `ContentTerms` removes PT-BR and English stopwords while preserving technical short words such as `go`.

Language detection uses marker scores for Portuguese and English plus a Portuguese diacritic bonus. If there is no evidence or the margin is too small, the language is `unknown`; `AnalyzeWithLanguage` can apply a supported fallback.

Similarity functions are deterministic. `EditSimilarity` is normalized Levenshtein similarity. `CharacterNGramSimilarity` is Jaccard over character n-gram sets. `FuzzySimilarity` returns the maximum of both.

`internal/query` composes the analyzer. It processes text, extracts entities, extracts direct concepts, expands concepts, detects intent, detects target, and detects temporal scope.

Entity extraction first performs exact alias phrase matching. It then performs fuzzy single-token alias matching with length-sensitive thresholds and requires at least `0.08` separation between the best and second-best different candidates. This supports typos such as `kubernts`, `dcker`, `dockr`, `jva`, `javascrit`, `typescrit`, `nodjs`, and `postgre` without a manual typo map.

Intent detection is marker-based. Preference markers and unsupported personal attributes intentionally produce `IntentUnknown`. List markers produce `IntentList`. Capability markers produce `IntentCapability`. Target detection combines explicit markers, intent defaults, and entity types.
