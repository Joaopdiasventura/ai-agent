# Indexação e Retrieval

[Índice](./README.md) | [English](../en/index-retrieval.md) | [Matemática](./mathematics.md)

`internal/index` constrói um documento em português e um em inglês para cada fato. Documentos têm campos ponderados: `statement`, `subject`, `object`, `concept`, `context`, `predicate` e `category`. O índice armazena postings, document frequencies, vocabulário, mapas de termo por n-gram, médias de tamanho de documento, médias por campo e índices estruturais por sujeito, entidade, conceito, categoria e contexto.

`retrieval.HybridRetriever.Search` executa quatro retrievers de forma independente: lexical, entity, concept e fuzzy. Retornar rankings separados preserva evidência de fonte para ranking e confidence.

Retrieval lexical usa BM25F. A contribuição central de termo é:

$$
S_{\operatorname{BM25F}}(t,d)
=
w_t\cdot\operatorname{IDF}(t)\cdot
\frac{\operatorname{wf}(t,d)(k_1+1)}{\operatorname{wf}(t,d)+k_1}
$$

Termos numéricos recebem peso $1.35$, termos com tamanho até dois recebem $1.15$ e os demais recebem $1.0$. Os campos mais fortes são subject, concept, object e context.

Retrieval por entidade usa entidades explícitas da query e índices estruturais de entidade. Entidades de pessoa são intencionalmente fracas com peso $0.25$, enquanto projetos e tecnologias são fortes. Essa é uma das proteções contra uma pergunta contendo João recuperar fatos arbitrários sobre João.

Retrieval por conceito usa conceitos da query e índices estruturais de conceito. Conceitos diretos são mais fortes que conceitos expandidos. A importância do fato entra no score.

Retrieval fuzzy é separado da extração fuzzy de entidades. Ele considera apenas termos da query com tamanho mínimo quatro e sem postings exatos, obtém candidatos por n-grams de caracteres, aplica thresholds sensíveis ao tamanho e pontua com similaridade, IDF, peso de campo e frequência.
