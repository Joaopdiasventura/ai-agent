# Decisões Técnicas E Trade-offs

## Base De Conhecimento Estática

- Decisão: armazenar documentos em código-fonte Go.
- Benefício: ausência de I/O de arquivo em runtime e nenhuma dependência de arquivo de dados externo.
- Custo: mudanças de conteúdo são mudanças de código.
- Consequência: testes conseguem validar diretamente contagem de documentos, IDs, idiomas e campos obrigatórios.

## Fluxo De Documentos Por Ponteiro

- Decisão: expor `[]*domain.Document` por `knowledge.Documents()`.
- Benefício: resultados de busca e vetores referenciam valores de documento existentes.
- Custo: callers recebem referências para dados compartilhados em memória.
- Consequência: o pipeline evita copiar structs de documento durante a recuperação.

## Recuperação Léxica

- Decisão: usar tokenização, expansão de consulta, TF-IDF, similaridade cosseno e boosts.
- Benefício: comportamento determinístico sem serviços externos.
- Custo: relevância depende de sobreposição de tokens e regras configuradas.
- Consequência: formulações desconhecidas podem falhar se não houver sobreposição de tokens, expansão, alias ou boost conectando a pergunta a um documento.

## Detecção De Idioma Baseada Em Regras

- Decisão: somar pesos de tokens para português e inglês.
- Benefício: comportamento simples que lida com tokens ambíguos como `me`.
- Custo: a detecção depende de pesos configurados manualmente.
- Consequência: empates e sinais desconhecidos caem em português.

## Núcleo Compartilhado Da Aplicação

- Decisão: CLI e HTTP usam `app.AgentResponse`.
- Benefício: a mesma lógica de recuperação e resposta atende as duas interfaces.
- Custo: comportamentos específicos de interface precisam ficar fora do núcleo.
- Consequência: testes conseguem validar o comportamento central por `internal/app` sem invocar CLI ou HTTP.

## Contrato HTTP Estrito

- Decisão: usar `DisallowUnknownFields`, limite de body de 2048 bytes, um objeto JSON por requisição e campo `content` obrigatório.
- Benefício: requisições inválidas falham antes de alcançar o agente.
- Custo: clientes precisam seguir o formato JSON exato.
- Consequência: erros de transporte ficam separados do comportamento de resposta não encontrada.

## Seleção Aleatória De Templates

- Decisão: escolher entre templates e conectores por idioma usando `math/rand/v2`.
- Benefício: respostas repetidas podem variar em redação.
- Custo: strings exatas de resposta não são estáveis.
- Consequência: testes verificam propriedades estáveis como idioma, status e presença de fatos conhecidos.

## Método HTTP `QUERY`

- Decisão: registrar `QUERY /ask`.
- Benefício: o contrato documentado corresponde à implementação atual do servidor.
- Custo: `QUERY` é menos comum que métodos como `POST`.
- Consequência: consumidores da API devem enviar requisições usando o método implementado.
