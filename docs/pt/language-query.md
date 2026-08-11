# Language e Query

[Índice](./README.md) | [English](../en/language-query.md)

`internal/language` normaliza texto, tokeniza, detecta idioma, remove stopwords e calcula similaridade. `Normalize` aplica trim, lowercase, dobra acentos comuns, mantém letras e dígitos, preserva separadores decimais entre dígitos e transforma outros separadores em espaços. `ContentTerms` remove stopwords de PT-BR e inglês preservando termos técnicos curtos como `go`.

A detecção de idioma usa scores de marcadores para português e inglês mais bônus por diacríticos portugueses. Se não há evidência ou a margem é pequena demais, o idioma é `unknown`; `AnalyzeWithLanguage` pode aplicar um fallback suportado.

As funções de similaridade são determinísticas. `EditSimilarity` é Levenshtein normalizado. `CharacterNGramSimilarity` é Jaccard sobre conjuntos de n-grams de caracteres. `FuzzySimilarity` retorna o máximo das duas.

`internal/query` compõe o analyzer. Ele processa texto, extrai entidades, extrai conceitos diretos, expande conceitos, detecta intent, detecta target e detecta escopo temporal.

A extração de entidades primeiro faz matching exato de aliases como frases. Depois faz fuzzy matching de aliases de token único com thresholds sensíveis ao tamanho e exige separação mínima de `0.08` entre o melhor e o segundo melhor candidato diferente. Isso suporta typos como `kubernts`, `dcker`, `dockr`, `jva`, `javascrit`, `typescrit`, `nodjs` e `postgre` sem mapa manual de typos.

A detecção de intent é baseada em marcadores. Marcadores de preferência e atributos pessoais sem suporte produzem propositalmente `IntentUnknown`. Marcadores de lista produzem `IntentList`. Marcadores de capability produzem `IntentCapability`. Target detection combina marcadores explícitos, defaults por intent e tipos de entidade.
