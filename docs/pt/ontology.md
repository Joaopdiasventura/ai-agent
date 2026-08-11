# Ontologia

[Índice](./README.md) | [English](../en/ontology.md)

`internal/ontology` fornece o grafo semântico usado por expansão de query, indexação, retrieval, ranking e reasoning. Ele valida IDs de conceito, nomes, referências a pais, pontas de relações, pesos de relações e vínculos de entidade.

Os conceitos cobrem contato, educação, certificação, idiomas, desenvolvimento de software, linguagem de programação, framework, runtime, frontend, backend, fullstack, banco de dados, mensageria, devops, infraestrutura, infrastructure as code, orquestração, containers, cloud, AWS, Kubernetes, sistemas distribuídos, concorrência, performance e domínios relacionados.

Aliases são localizados e normalizados. `Aliases(targetLanguage)` retorna aliases ponderados com preferência de idioma. Aliases mais longos vêm antes dos curtos para que frases como `full stack development` ou `sistemas distribuidos` casem antes de palavras menores sobrepostas.

`ResolveAlias` trata aliases exatos normalizados. `Expand(id,maxDepth)` faz propagação em largura por pais e relações semânticas ponderadas. Propagação por pai usa `0.85`; propagação por relação usa o peso da relação; o melhor score por conceito é preservado.

Vínculos de entidade classificam tecnologias e idiomas. JavaScript, TypeScript, Java e Go vinculam a programming-language; Angular, React, Next.js, Spring Boot e NestJS a framework; Node.js a runtime; bancos, mensageria, cloud, containers, orquestração e infraestrutura vinculam a conceitos próprios. Português e Inglês vinculam ao conceito de idioma humano, não a programming-language.

Full Stack existe como conceito com aliases PT-BR e inglês. A expansão pode recuperar evidências relacionadas a frontend/backend, mas o reasoning ainda exige cobertura suficiente antes de conceder suporte.
