# Fluxo De Requisição

## Fluxo Da CLI

```mermaid
sequenceDiagram
    participant User as Usuário
    participant CLI as app.Run
    participant App as app.AgentResponse
    participant Search as search.Engine
    participant Answer as internal/answer

    User->>CLI: linha da pergunta
    CLI->>CLI: trim da entrada e checagem de comandos de saída
    CLI->>App: AgentResponse(question)
    App->>Search: Search(question, maximumSearchResults)
    Search-->>App: SearchResult
    alt encontrado
        App->>Answer: BuildPlan + SelectTemplateForPlan + RenderTemplate
        Answer-->>App: texto da resposta
        App-->>CLI: response, true, language
        CLI-->>User: Bot: resposta
    else não encontrado
        App-->>CLI: "", false, language
        CLI-->>User: fallback localizado
    end
```

A CLI lê stdin com `bufio.Scanner`. Linhas vazias são ignoradas. Os comandos `sair`, `exit`, `quit` e `encerrar` encerram o loop.

## Fluxo HTTP

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as api.Handler
    participant Server as server.Handler/CORS
    participant Handler as handlers.Handle
    participant App as app.AgentResponse

    Client->>API: QUERY /ask
    API->>Server: Handler()
    Server->>Handler: rota QUERY /ask
    Handler->>Handler: MaxBytesReader + decode JSON
    Handler->>Handler: trim de content
    Handler->>App: AgentResponse(content)
    alt resposta encontrada
        App-->>Handler: response, true, language
        Handler-->>Client: 200 JSON
    else resposta não encontrada
        App-->>Handler: "", false, language
        Handler-->>Client: 404 JSON localizado
    end
```

## Fluxo De Busca

```mermaid
flowchart TD
    Input["question"] --> Limit{"limit <= 0?"}
    Limit -->|sim| NotFoundPT["Found=false, language=pt"]
    Limit -->|não| Tokens["tokenizer.Tokenize"]
    Tokens --> Empty{"tokens vazios?"}
    Empty -->|sim| NotFoundLang["Found=false"]
    Empty -->|não| Language["nlp.DetectLanguage"]
    Language --> Expansion["nlp.ExpandQuery"]
    Expansion --> Entity["nlp.DetectEntity"]
    Entity --> Analysis["nlp.AnalyzeQuery"]
    Analysis --> Candidates["FilterDocumentsByIntent"]
    Candidates --> EntityTokens{"entidade encontrada?"}
    EntityTokens -->|sim| Append["anexa valor da entidade tokenizado"]
    EntityTokens -->|não| Vector["CalculateTFIDF"]
    Append --> Vector
    Vector --> VectorEmpty{"vetor da pergunta vazio?"}
    VectorEmpty -->|sim| NotFoundLang
    VectorEmpty -->|não| Scope["ShouldSearchMultipleDocuments"]
    Scope --> Ranking["FindTopDocuments"]
    Ranking --> Results{"resultados vazios?"}
    Results -->|sim| NotFoundLang
    Results -->|não| Found["Found=true"]
```

Cada caminho de não encontrado preserva o idioma detectado quando a detecção já aconteceu. O caminho de validação de limite usa português como padrão porque nenhuma análise de tokens é executada.
