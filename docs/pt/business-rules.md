# Regras De Negócio

## A Pergunta Deve Gerar Tokens Pesquisáveis

- Condição: a pergunta fica vazia após tokenização ou não contém token utilizável.
- Processamento: `search.Engine.Search` retorna `Found=false`.
- Resultado: CLI ou HTTP seguem o caminho de fallback.
- Componente: `internal/search` e `internal/tokenizer`.

## O Body HTTP Deve Ser Um Único Objeto JSON Válido

- Condição: a API recebe JSON inválido, body vazio, mais de um objeto JSON, campo desconhecido ou tipo inválido de campo.
- Processamento: `internal/handlers/ask` rejeita a requisição antes de chamar `app.AgentResponse`.
- Resultado: o handler retorna `400` com campo JSON `response`.
- Componente: `internal/handlers/ask`.

## O Campo Content É Obrigatório

- Condição: `content` está ausente, vazio ou contém apenas espaços após trim.
- Processamento: o handler rejeita a requisição.
- Resultado: `400` com `O campo content é obrigatório.`
- Componente: `internal/handlers/ask`.

## O Tamanho Do Body É Limitado

- Condição: o body excede `maxRequestBodyBytes`, atualmente `2048`.
- Processamento: `http.MaxBytesReader` faz o decode falhar com `http.MaxBytesError`.
- Resultado: `413` com resposta JSON de erro.
- Componente: `internal/handlers/ask`.

## O Idioma Controla A Seleção De Documentos

- Condição: uma pergunta é detectada como português ou inglês.
- Processamento: `FilterDocumentsByIntent` ignora documentos cujo `Document.Language` não corresponde a `QueryAnalysis.Language`.
- Resultado: perguntas em português retornam documentos `pt`; perguntas em inglês retornam documentos `en`.
- Componente: `internal/nlp` e `internal/search`.

## O Threshold De Similaridade Controla A Possibilidade De Resposta

- Condição: resultados ranqueados têm similaridade menor que `minimumSimilarity`, atualmente `0.1`.
- Processamento: `FilterRelevantResults` remove esses resultados.
- Resultado: se nenhum resultado permanecer, a aplicação retorna fallback localizado.
- Componente: `internal/search`.

## Resultados Específicos De Entidade Devem Mencionar A Entidade

- Condição: uma entidade é detectada na pergunta.
- Processamento: `FilterRelevantResults` exige que `result.Document.Content` contenha `analysis.Entity.Value`.
- Resultado: resultados que não mencionam o valor da entidade resolvida são excluídos.
- Componente: `internal/search`.

## Intenções De Projeto E Tecnologia Podem Retornar Múltiplos Documentos

- Condição: a intenção é `IntentTechnologies` ou `IntentProject`.
- Processamento: `ShouldSearchMultipleDocuments` permite usar o limite configurado de resultados.
- Resultado: até `maximumSearchResults`, atualmente `5`, podem ser usados na resposta.
- Componente: `internal/search` e `internal/app`.

## Fallbacks Localizados São Retornados Em Vez De Respostas Fabricadas

- Condição: a busca retorna `Found=false`.
- Processamento: `app.NotFoundMessage` seleciona fallback em português ou inglês.
- Resultado: o sistema não gera resposta fora dos fatos recuperados.
- Componente: `internal/app`.
