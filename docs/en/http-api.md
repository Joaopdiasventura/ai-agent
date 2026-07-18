# HTTP API

## Route

```text
QUERY /ask
Content-Type: application/json
```

The route is registered in `server.NewHandler()` with:

```go
mux.HandleFunc("QUERY /ask", handlers.Handle())
```

`vercel.json` routes all incoming paths to `/api/index`, which delegates to `server.Handler()`.

## Request Body

```json
{
  "content": "Me fale sobre o auronix"
}
```

Rules:

- body size is limited to 2048 bytes;
- body must be a single JSON object;
- unknown fields are rejected;
- `content` must be a non-empty string after trimming.

## Successful Response

```json
{
  "response": "De forma simples: Auronix é um projeto de banco digital criado para simular transferências, acompanhar movimentações e mostrar atualizações em tempo real."
}
```

Status: `200`.

The exact wording can vary because answer templates are selected randomly.

## Not Found Response

```json
{
  "response": "I don't have that specific information, but I can talk about João's experience, projects, technologies, services, or contact details."
}
```

Status: `404`.

The fallback language is based on detected query language.

## Error Responses

| Status | Condition | Response |
| --- | --- | --- |
| `400` | Invalid JSON syntax | `JSON inválido.` |
| `400` | Invalid field type | `Tipo de campo inválido.` |
| `400` | Empty body | `Body da requisição vazio.` |
| `400` | Generic invalid body, including unknown fields | `Body da requisição inválido.` |
| `400` | Multiple JSON objects | `O body deve conter apenas um objeto JSON.` |
| `400` | Empty `content` | `O campo content é obrigatório.` |
| `413` | Body above 2048 bytes | `Body da requisição excede o limite permitido.` |
| `403` | Disallowed CORS preflight origin | `Forbidden origin.` |
| `503` | Handler creation error at API entrypoint | `Service temporarily unavailable.` |

## CORS

Allowed origins:

- `https://joaopdias.dev.br`
- `http://localhost:4200`

Allowed methods:

- `QUERY`
- `OPTIONS`

Allowed headers:

- `Content-Type`
