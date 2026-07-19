# Testes

## Comando

```powershell
go test ./...
```

## Áreas Cobertas Atualmente

### Detecção De Idioma

`internal/nlp/language_test.go` valida a detecção ponderada de idioma, incluindo o caso em que `Me fale sobre o auronix` em português não deve ser classificado como inglês por causa de `me`.

### Pipeline De Busca

`internal/search/engine_test.go` valida que perguntas em português e inglês retornam documentos no idioma detectado, que Auronix pode ser encontrado em português, que perguntas de tecnologia em inglês retornam múltiplos resultados e que entradas inválidas não produzem respostas.

### Contrato Da Aplicação

`internal/app/agent_test.go` valida `AgentResponse` para perguntas sobre Auronix em português e inglês e o comportamento de fallback localizado.

### Contrato HTTP

`internal/handlers/ask/ask_test.go` valida requisições JSON bem-sucedidas, `content` vazio, JSON inválido, campos desconhecidos, bodies grandes demais e respostas de não encontrado.

### Invariantes Da Base De Conhecimento

`internal/knowledge/documents_test.go` valida:

- 140 documentos;
- 70 documentos em português;
- 70 documentos em inglês;
- IDs únicos;
- campos obrigatórios;
- sufixos `-pt` e `-en`.

## Limites Dos Testes

Os testes não verificam o texto completo exato da resposta porque a seleção de templates e conectores é aleatória. Eles também evitam serviços externos porque a aplicação não exige rede, banco de dados ou chamadas a modelos para os fluxos documentados.
