# Estrutura Do Projeto

## Arquivos Da Raiz

### `go.mod`

Define o nome do módulo `ai-agent` e a versão Go `1.26.3`. Nenhuma dependência externa de módulo é declarada.

### `vercel.json`

Define a configuração da função Vercel. Desabilita detecção de framework com `"framework": null`, deixa `buildCommand` vazio, define `maxDuration` de `api/index.go` como `10` e roteia todas as requisições para `/api/index`.

## Pontos De Entrada

### `cmd/ai-agent/main.go`

Bootstrap da CLI. Sua única responsabilidade é chamar `app.Run()`. Isso mantém a inicialização da CLI separada do workflow da aplicação.

### `api/index.go`

Bootstrap serverless. Obtém um handler HTTP de `server.Handler()` e delega a requisição. Em erro de criação do handler, retorna JSON `503`.

## Camada HTTP

### `server/server.go`

Controla criação do mux HTTP, registro de `QUERY /ask`, tratamento de CORS, funções auxiliares de JSON e cache preguiçoso do handler com mutex.

### `internal/handlers/ask/`

Controla o contrato JSON público do endpoint ask. Valida o formato do body e delega conteúdo válido para `app.AgentResponse`.

## Núcleo Da Aplicação

### `internal/app/`

Contém a interface central usada por CLI e HTTP. `agent.go` inicializa documentos e engine de busca, `config.go` armazena constantes de runtime, e `chatbot.go` implementa o loop interativo de console.

## Dados E Domínio

### `internal/domain/`

Define `Document`, a estrutura compartilhada passada entre knowledge, TF-IDF e search.

### `internal/knowledge/`

Armazena os fatos estáticos. `Documents()` retorna ponteiros para o array compilado de documentos.

## Pacotes De Processamento

### `internal/tokenizer/`

Converte texto para minúsculas, remove separadores não alfanuméricos, separa tokens e remove stopwords configuradas.

### `internal/nlp/`

Contém lógica determinística de idioma, consulta, intenção, entidade, modo de resposta e tecnologias.

### `internal/tfidf/`

Calcula mapas de frequência e vetores TF-IDF.

### `internal/search/`

Coordena recuperação, filtragem, pontuação, ordenação e checagem de relevância.

### `internal/answer/`

Cria respostas finais a partir dos resultados ranqueados.

## Testes

Testes ficam próximos ao pacote que validam:

- `internal/nlp/language_test.go`
- `internal/search/engine_test.go`
- `internal/app/agent_test.go`
- `internal/handlers/ask/ask_test.go`
- `internal/knowledge/documents_test.go`

Eles validam o comportamento atual sem depender de serviços externos.
