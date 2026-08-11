# Reasoning

[Índice](./README.md) | [English](../en/reasoning.md)

`internal/reasoning` é a camada de suficiência de evidência. Ela recebe candidatos ranqueados e produz uma conclusão com tipo, status, entidade foco, conceito foco, evidências e grupos opcionais.

`EvidenceBuilder` calcula rank, score, importância, directness, entidades casadas e conceitos casados. Directness combina match de entidade explícita e match de conceito direto; evidência não relacionada cai para baixa directness.

A relevância direta geral é rígida. Entidades da query que não sejam pessoa precisam aparecer no fato. Conceitos casados diretamente precisam aparecer no fato. Se há apenas entidade pessoa e nenhum target/categoria suportado, fatos arbitrários sobre essa pessoa não bastam.

Reasoning de capability ignora entidades pessoa como alvo de capability. Ele mantém apenas evidências que referenciam um alvo explícito não-pessoa ou um conceito diretamente casado. Uma capability suportada exige match direto com score de evidência mínimo `0.35`.

Full Stack tem status semântico especial: JavaScript mais TypeScript sozinhos não bastam. Suporte exige fato fullstack explícito ou evidência suficiente cobrindo conceitos frontend e backend.

Reasoning de lista agrupa evidência pelo tipo de entidade pedido. Conceitos de programming language, framework, runtime, database, messaging, cloud, devops e infrastructure listam tecnologias; listas de idiomas humanos usam `EntityTypeLanguage`. A filtragem impede que frameworks, bancos, clouds e idiomas humanos vazem para listas de linguagens de programação.
