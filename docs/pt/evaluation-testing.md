# Evaluation e Testing

[Índice](./README.md) | [English](../en/evaluation-testing.md)

O repositório possui testes por pacote para todas as camadas principais. Os testes validam contratos de dados, fallback localizado, normalização, stopwords, detecção de idioma, similaridade fuzzy, aliases e expansão da ontologia, estruturas do índice, fontes de retrieval, fusão/features de ranking, suporte em reasoning, formato de planning, grounding de generation, modos de confidence, comportamento do serviço, CLI, validação HTTP e relatórios de avaliação.

`internal/evaluation` é um harness de regressão sobre `agent.Debug`. Um caso pode validar presença de resposta, idioma, intent, target, entidades, conceitos, vencedor de comparação, fatos gerados obrigatórios, fatos proibidos, status do plano, modo de confiança, limites de confiança, texto esperado na resposta, texto proibido na resposta e texto gerado no debug.

O corpus de regressão cobre respostas diretas de contato, capabilities JavaScript/TypeScript/Java/Go, typos para Docker e Kubernetes, listas de linguagens de programação, listas de frameworks, listas de tecnologias, listas de idiomas humanos, idade desconhecida, preferência desconhecida, backend, full stack, uso de tecnologia, comparações e overviews de projetos.

Execute:

```bash
go test ./...
go test -race ./...
```

O teste do corpus de regressão exige 100% de accuracy. Testes existentes não devem ser reduzidos para passar; o comportamento deve ser corrigido.
