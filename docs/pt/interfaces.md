# Agent, CLI e HTTP

[Índice](./README.md) | [English](../en/interfaces.md)

`agent.Service` é o serviço público da aplicação. `New()` constrói e valida a pipeline inteira em memória. `Answer(question)` aplica trim, rejeita entrada vazia com `ErrEmptyQuestion`, executa todas as camadas e retorna `agent.Result` com response, flag de resposta, idioma, score de confiança e nível de confiança.

`Debug(question)` executa a mesma pipeline e retorna `DebugResult`: pergunta, resultado público, query, resultado de retrieval, resultado de ranking, resultado de reasoning, plano, resposta gerada e resultado de confidence. Debug é a melhor interface para identificar se uma falha pertence a query, retrieval, ranking, reasoning, planning, materialization, generation ou confidence.

O gate de resposta pública exige tudo isto: `PlanStatusReady`, `SupportSupported`, resposta gerada não vazia e confidence mode `claim`. Abstenções ainda podem ter texto interno de geração no debug, mas `Answer` expõe resposta vazia.

`internal/cli` fornece o runner de terminal. Ele ignora linhas vazias, sai com `sair`, `exit`, `quit` ou `encerrar`, imprime respostas suportadas e imprime fallback localizado para resultados sem suporte.

HTTP é tratado por `server` e `internal/handlers/ask`. `POST /ask` aceita `{ "content": "..." }`. Bodies são limitados a 2048 bytes, campos desconhecidos são rejeitados, e CORS permite `https://joaopdias.dev.br` e `http://localhost:4200`. `api/index.go` é a entrada da Vercel.
