# Request Flow

## CLI Flow

```mermaid
sequenceDiagram
    participant User
    participant CLI as app.Run
    participant App as app.AgentResponse
    participant Search as search.Engine
    participant Answer as internal/answer

    User->>CLI: question line
    CLI->>CLI: trim input and check exit commands
    CLI->>App: AgentResponse(question)
    App->>Search: Search(question, maximumSearchResults)
    Search-->>App: SearchResult
    alt found
        App->>Answer: BuildPlan + SelectTemplateForPlan + RenderTemplate
        Answer-->>App: response text
        App-->>CLI: response, true, language
        CLI-->>User: Bot: response
    else not found
        App-->>CLI: "", false, language
        CLI-->>User: localized fallback
    end
```

The CLI reads from stdin with `bufio.Scanner`. Empty lines are ignored. The commands `sair`, `exit`, `quit`, and `encerrar` terminate the loop.

## HTTP Flow

```mermaid
sequenceDiagram
    participant Client
    participant API as api.Handler
    participant Server as server.Handler/CORS
    participant Handler as handlers.Handle
    participant App as app.AgentResponse

    Client->>API: QUERY /ask
    API->>Server: Handler()
    Server->>Handler: route QUERY /ask
    Handler->>Handler: MaxBytesReader + JSON decode
    Handler->>Handler: trim content
    Handler->>App: AgentResponse(content)
    alt answer found
        App-->>Handler: response, true, language
        Handler-->>Client: 200 JSON
    else answer not found
        App-->>Handler: "", false, language
        Handler-->>Client: 404 localized JSON
    end
```

## Search Flow

```mermaid
flowchart TD
    Input["question"] --> Limit{"limit <= 0?"}
    Limit -->|yes| NotFoundPT["Found=false, language=pt"]
    Limit -->|no| Tokens["tokenizer.Tokenize"]
    Tokens --> Empty{"tokens empty?"}
    Empty -->|yes| NotFoundLang["Found=false"]
    Empty -->|no| Language["nlp.DetectLanguage"]
    Language --> Expansion["nlp.ExpandQuery"]
    Expansion --> Entity["nlp.DetectEntity"]
    Entity --> Analysis["nlp.AnalyzeQuery"]
    Analysis --> Candidates["FilterDocumentsByIntent"]
    Candidates --> EntityTokens{"entity found?"}
    EntityTokens -->|yes| Append["append tokenized entity value"]
    EntityTokens -->|no| Vector["CalculateTFIDF"]
    Append --> Vector
    Vector --> VectorEmpty{"question vector empty?"}
    VectorEmpty -->|yes| NotFoundLang
    VectorEmpty -->|no| Scope["ShouldSearchMultipleDocuments"]
    Scope --> Ranking["FindTopDocuments"]
    Ranking --> Results{"results empty?"}
    Results -->|yes| NotFoundLang
    Results -->|no| Found["Found=true"]
```

Each not-found path preserves the detected language when detection has already happened. The limit validation path defaults to Portuguese because no token analysis is performed.
