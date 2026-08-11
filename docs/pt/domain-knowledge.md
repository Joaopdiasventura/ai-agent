# Domain e Knowledge

[Índice](./README.md) | [English](../en/domain-knowledge.md)

`internal/domain` define contratos compartilhados: `Entity`, `Fact`, `Query`, matches, relações, categorias de fato, intents, query targets, períodos, escopos temporais, texto localizado e sinais de evidência.

Uma `Entity` tem ID, tipo, nome localizado e aliases. Tipos de entidade são semanticamente importantes: idiomas humanos usam `EntityTypeLanguage`, enquanto tecnologias usam `EntityTypeTechnology`. Essa separação impede que Português e Inglês sejam tratados como Java, Go, JavaScript ou TypeScript.

Um `Fact` é o átomo de evidência. Ele possui sujeito, predicado, objeto, categoria, conceitos, contexto, statement localizado, período opcional, importância e fonte. A validação verifica IDs, formato do objeto, consistência de período atual, statements localizados e faixa de importância.

`internal/knowledge` é a fonte de verdade. Ela contém João, empresas, projetos, cargos, instituições, certificações, idiomas humanos e entidades tecnológicas. Fatos de skill são explícitos, não gerados para toda tecnologia. Os fatos atuais incluem JavaScript, TypeScript, Java, Go, Angular, React, Next.js, Spring Boot, Node.js, NestJS, PostgreSQL, MongoDB, Redis, RabbitMQ, Kafka, SQS, Docker, Terraform, Kubernetes e AWS.

A invariante principal da knowledge é comportamento consciente de ausência. Se não há fato sobre idade, data de nascimento, salário, altura, estado civil, comida favorita, jogo favorito ou preferência explícita, o agente deve abster em vez de responder com fatos não relacionados de João.
