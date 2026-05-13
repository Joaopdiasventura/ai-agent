package data

const ProfileContext = `
Identity:
- Name: João Paulo Dias Ventura.
- Location: São Paulo, Brazil.
- Portfolio description: portfolio focused on distributed systems, financial workflows, and resilient data-intensive architecture.
- Professional summary: software engineer focused on distributed systems and data-intensive applications, building financial and operational flows with transactional consistency, event-driven communication, and resilient high-volume processing.

Profile:
- Hands-on full-stack delivery across backend, frontend, cloud, and messaging layers.
- Focus: transactional consistency, asynchronous processing, operational traceability, event-driven architecture, distributed systems, concurrency control, idempotent flows, financial systems, billing workflows, streaming ingestion, data pipelines, distributed processing, Angular SSR, hydration, lazy loading, and frontend performance optimization.
- Technical stack listed in the portfolio: Node.js, Java, Go, Angular, PostgreSQL, MongoDB, RabbitMQ, Redis, AWS, CI/CD, S3, ECS, IAM, NestJS, Spring Boot, Fastify, BullMQ, SSE, Docker Compose, nginx, AsyncAPI, TypeScript, node:http, autocannon, Tauri, Capacitor, gzip, SHA-256, goroutines, channels, CLI.
- Language: Advanced English.
- Availability: open to engineering roles involving distributed systems, critical backend flows, and architectural responsibility.

Education and certificates:
- Fatec Osasco - Multiplatform Software Development (2026-2028).
- Etec Guarulhos - Systems Development (2023-2025).
- AWS - Event-Driven Architecture Modeling and Deploying Microservices on Amazon EKS.
- MongoDB - Financial Services Industry Knowledge.

Public contact:
- Email: joaopdias.dev@gmail.com.
- Phone: +55 (11) 98655-3558.
- LinkedIn: linkedin.com/in/joaopdias-dev.
- GitHub: github.com/Joaopdiasventura.

Experience:
- uFind Tecnologia, Mid-level Developer / Desenvolvedor Pleno, Jun 2025 - Present. João Paulo leads implementation of an insurance billing workflow that ingests heterogeneous financial files, reconciles transactional state, and automates monthly processing above R$1M with end-to-end traceability. Highlights: stream-based ingestion pipelines for multiple financial layouts with Strategy Pattern, critical validations, and consistency guarantees across the billing flow; reduced manual processing from days to minutes; evolved AWS architecture standards around S3, ECS, and IAM with high autonomy.
- Representa Online, Junior Developer / Desenvolvedor Júnior, Sep 2024 - May 2025. João Paulo designed AI-oriented data pipelines for more than 16 GB of television media data, structuring ingestion, transformation, and vector-store storage for model consumption. Highlights: event-driven Node.js pipelines with flow control, fault isolation, and distributed processing; bidirectional real-time communication over WebSocket integrated with the OpenAI API.
- Representa Online, Systems Development Intern / Estagiário em Desenvolvimento de Sistemas, Jun 2024 - Aug 2024. João Paulo built catalog and authentication features focused on scalable discovery, SSR/SSG delivery, and frontend performance for content-heavy experiences. Highlights: proximity search with the Haversine formula; optimized queries for catalog discovery; JWT and Google OAuth2 authentication; SSR/SSG pages with lazy loading, asset optimization, and render control to improve TTFB and LCP.

Impact metrics:
- +16GB television media data processed.
- R$1M+ monthly billed volume automated.
- 1.23 GiB/s compression throughput measured.

Projects:
- Auronix: digital banking and transactional settlement. Teaser: digital banking core with asynchronous settlement, Redis-backed concurrency control, and live transfer feedback over SSE. Role: full-stack architecture, transactional backend design, real-time workflow UX. Timeline: independent engineering project. Stack: NestJS, Fastify, PostgreSQL, Redis, BullMQ, Angular, SSE. Metrics: 248 automated tests; 10 min request expiry; 100 / 24h SSE replay window. Links: https://auronix.joaopdias.dev.br, https://github.com/joaopdiasventura/Auronix-server, https://github.com/joaopdiasventura/Auronix-client.
- Modularis: event-driven microservices and distributed onboarding. Teaser: polyglot microservice platform where a persisted onboarding saga coordinates identity, payment intent, signed webhooks, and premium activation through RabbitMQ contracts. Role: distributed backend architecture, saga orchestration, async contracts, operational resilience. Timeline: independent engineering project. Stack: NestJS, Spring Boot, Go, RabbitMQ, PostgreSQL, MongoDB, Docker Compose, nginx, AsyncAPI. Metrics: 6 deployable services; 9 async channels; 3 runtime stacks. Link: https://github.com/joaopdiasventura/Modularis.
- Votrix: high-performance Node.js HTTP runtime. Teaser: TypeScript HTTP runtime built for direct routing, deferred parsing, and inspectable benchmark gains. Role: runtime architecture, performance engineering, benchmark methodology. Timeline: independent engineering project. Stack: TypeScript, node:http, autocannon. Metrics: 30.1K local RPS; 37.25% lead vs Fastify; 42.08% lead vs Express. Benchmark artifact dated 2026-04-03T03:18:21.722Z shows Votrix leading all four measured scenarios and peaking at 30,124.80 RPS on GET /health. Link: https://github.com/joaopdiasventura/Votrix.
- GGCompress: concurrent archive engine and deterministic extraction. Teaser: Go archive engine built around ordered chunk pipelines, deterministic indexing, integrity verification, and measured multi-gigabyte throughput. Role: systems design, archive format engineering, concurrency pipeline, and benchmark methodology. Timeline: independent engineering project. Stack: Go, goroutines, channels, gzip, SHA-256, CLI. Metrics: 1.23 GiB/s observed throughput; 9.77 GiB benchmark input; 8 workers / 4 MiB chunk size. Measured CLI run dated April 23, 2026 finished in 7.952s while preserving deterministic archive order. Link: https://github.com/Joaopdiasventura/ggc.
- VOX: auditability and voting integrity. Teaser: voting platform focused on explicit state, auditability, and operator clarity under concurrent load. Role: product architecture, backend design, audit-oriented modeling. Timeline: independent engineering project. Stack: Tauri, NestJS, Redis. Metrics: +500 concurrent users; 29 HTTP routes; 108 automated tests. Link: https://vox.joaopdias.dev.br.
- Etecfy: music streaming. Teaser: streaming platform built for catalog growth, fast discovery, and smooth playback. Role: architecture, product structure, media-oriented frontend. Timeline: independent engineering project. Stack: Angular, NestJS, Capacitor. Metrics: 1.3K launch accesses in 6h; 10s chunk size; 3 delivery surfaces. Links: https://etecfy.joaopdias.dev.br, https://github.com/joaopdiasventura/etecfy-server, https://github.com/joaopdiasventura/etecfy-client.
`
