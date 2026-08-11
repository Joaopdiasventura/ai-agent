# Reasoning

[Índice](./README.md) | [English](../en/reasoning.md) | [Matemática](./mathematics.md)

`internal/reasoning` é a camada de suficiência de evidência. Ela recebe candidatos ranqueados e produz uma conclusão com tipo, status, entidade foco, conceito foco, evidências e grupos opcionais.

`EvidenceBuilder` calcula rank, score, importância, directness, entidades casadas e conceitos casados. Directness segue os casos do código:

$$
D(e)
=
\begin{cases}
\operatorname{clamp}\left(0.55S_{\operatorname{entity}} + 0.45S_{\operatorname{concept}}\right), & S_{\operatorname{entity}}>0 \land S_{\operatorname{concept}}>0 \\
\operatorname{clamp}\left(0.85S_{\operatorname{entity}}\right), & S_{\operatorname{entity}}>0 \\
\operatorname{clamp}\left(0.90S_{\operatorname{concept}}\right), & S_{\operatorname{concept}}>0 \\
0.25, & \text{otherwise}
\end{cases}
$$

A relevância direta geral é rígida. Entidades da query que não sejam pessoa precisam aparecer no fato. Conceitos casados diretamente precisam aparecer no fato. Se há apenas entidade pessoa e nenhum target/categoria suportado, fatos arbitrários sobre essa pessoa não bastam.

Reasoning de capability ignora entidades pessoa como alvo de capability. Ele mantém apenas evidências que referenciam um alvo explícito não-pessoa ou um conceito diretamente casado. Uma capability suportada exige match direto com score de evidência mínimo $0.35$.

Full Stack tem status semântico especial: JavaScript mais TypeScript sozinhos não bastam. Suporte exige fato fullstack explícito ou evidência suficiente cobrindo conceitos frontend e backend.

Reasoning de lista agrupa evidência pelo tipo de entidade pedido. Conceitos de programming language, framework, runtime, database, messaging, cloud, devops e infrastructure listam tecnologias; listas de idiomas humanos usam `EntityTypeLanguage`. A filtragem impede que frameworks, bancos, clouds e idiomas humanos vazem para listas de linguagens de programação.

