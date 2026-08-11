# Planning e Generation

[Índice](./README.md) | [English](../en/planning-generation.md)

`internal/planning` converte conclusões de reasoning em planos finitos de resposta. Se a conclusão não é suportada, o planner emite `PlanStatusAbstain` e uma seção de abstenção. Conclusões suportadas viram `PlanStatusReady` com seções como lead, evidence, details, alternatives e list.

Os tipos de plano espelham os tipos de conclusão: direct, overview, capability, experience, technology usage, comparison e list. Itens de plano podem representar fatos, entidades, suporte, vencedores de comparação, alternativas e uso de tecnologia.

`EvidenceSelector.Select` escolhe evidência compacta e diversa. O score base é `0.62*score + 0.23*directness + 0.15*importance`, depois adiciona pequenos bônus para novas categorias/predicados e penalidades para repetidos. Isso reduz evidência redundante, como fatos repetidos de nível de idioma.

`internal/generation` verbaliza o plano. Ele não recupera, ranqueia nem raciocina. `Materialize(plan, source)` carrega apenas os IDs de fatos e entidades planejados, e a validação garante que todo material planejado existe.

Generation tem realizers para cada tipo de plano e suporta português e inglês. Idioma unknown cai para inglês. `joinSentences` deduplica frases completas; `joinNatural` deduplica labels de entidade e junta com `e` ou `and`.

Para abstenção, generation produz uma explicação interna para debug. O `Answer` público ainda retorna sem resposta porque o gate público acontece em `agent.executionHasResponse`.
