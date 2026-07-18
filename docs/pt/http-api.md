# API HTTP

## Rota

```text
QUERY /ask
Content-Type: application/json
```

A rota é registrada em `server.NewHandler()` com:

```go
mux.HandleFunc("QUERY /ask", handlers.Handle())
```

`vercel.json` roteia todos os caminhos recebidos para `/api/index`, que delega para `server.Handler()`.

## Body Da Requisição

```json
{
  "content": "Me fale sobre o auronix"
}
```

Regras:

- o body é limitado a 2048 bytes;
- o body deve ser um único objeto JSON;
- campos desconhecidos são rejeitados;
- `content` deve ser uma string não vazia após trim.

## Resposta De Sucesso

```json
{
  "response": "De forma simples: Auronix é um projeto de banco digital criado para simular transferências, acompanhar movimentações e mostrar atualizações em tempo real."
}
```

Status: `200`.

O texto exato pode variar porque templates de resposta são selecionados aleatoriamente.

## Resposta De Não Encontrado

```json
{
  "response": "I don't have that specific information, but I can talk about João's experience, projects, technologies, services, or contact details."
}
```

Status: `404`.

O idioma do fallback é baseado no idioma detectado na consulta.

## Respostas De Erro

| Status | Condição | Resposta |
| --- | --- | --- |
| `400` | Sintaxe JSON inválida | `JSON inválido.` |
| `400` | Tipo inválido de campo | `Tipo de campo inválido.` |
| `400` | Body vazio | `Body da requisição vazio.` |
| `400` | Body genericamente inválido, incluindo campos desconhecidos | `Body da requisição inválido.` |
| `400` | Múltiplos objetos JSON | `O body deve conter apenas um objeto JSON.` |
| `400` | `content` vazio | `O campo content é obrigatório.` |
| `413` | Body acima de 2048 bytes | `Body da requisição excede o limite permitido.` |
| `403` | Origem de preflight CORS não permitida | `Forbidden origin.` |
| `503` | Erro de criação de handler no ponto de entrada da API | `Service temporarily unavailable.` |

## CORS

Origens permitidas:

- `https://joaopdias.dev.br`
- `http://localhost:4200`

Métodos permitidos:

- `QUERY`
- `OPTIONS`

Headers permitidos:

- `Content-Type`
