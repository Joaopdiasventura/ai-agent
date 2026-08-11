# Planning e Generation

[Índice](./README.md) | [English](../en/planning-generation.md) | [Matemática](./mathematics.md)

`internal/planning` converte conclusões de reasoning em planos finitos de resposta. Se a conclusão não é suportada, o planner emite `PlanStatusAbstain` e uma seção de abstenção. Conclusões suportadas viram `PlanStatusReady` com seções como lead, evidence, details, alternatives e list.

Os tipos de plano espelham os tipos de conclusão: direct, overview, capability, experience, technology usage, comparison e list. Itens de plano podem representar fatos, entidades, suporte, vencedores de comparação, alternativas e uso de tecnologia.

`EvidenceSelector.Select` escolhe evidência compacta e diversa. O termo base de pontuação é:

$$
S_{\operatorname{baseSelect}}
=
0.62S
+
0.23D
+
0.15I
$$

O score completo de seleção adiciona os ajustes de diversidade implementados em `selectionScore`:

$$
S_{\operatorname{select}}
=
S_{\operatorname{baseSelect}}
+
B_{\operatorname{category}}
+
B_{\operatorname{predicate}}
-
P_{\operatorname{category}}
-
P_{\operatorname{predicate}}
$$

Onde $B_{\operatorname{category}}=0.08$ para uma nova categoria, $B_{\operatorname{predicate}}=0.08$ para um novo predicado, $P_{\operatorname{category}}=0.04$ quando essa categoria já foi selecionada pelo menos duas vezes, e $P_{\operatorname{predicate}}=0.06$ quando esse predicado já foi selecionado pelo menos duas vezes.

`internal/generation` verbaliza o plano. Ele não recupera, ranqueia nem raciocina. `Materialize(plan, source)` carrega apenas os IDs de fatos e entidades planejados, e a validação garante que todo material planejado existe.

Generation tem realizers para cada tipo de plano e suporta português e inglês. Idioma unknown cai para inglês. `joinSentences` deduplica frases completas; `joinNatural` deduplica labels de entidade e junta com `e` ou `and`.

Para abstenção, generation produz uma explicação interna para debug. O `Answer` público ainda retorna sem resposta porque o gate público acontece em `agent.executionHasResponse`.
