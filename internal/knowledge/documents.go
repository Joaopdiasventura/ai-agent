package knowledge

import "ai-agent/internal/domain"

var staticDocuments = [...]domain.Document{
	{
		ID:       "identity-basic",
		Category: "identity",
		Content:  "João Paulo Dias Ventura é um engenheiro de software localizado em São Paulo, Brasil.",
	},
	{
		ID:       "identity-professional-summary",
		Category: "identity",
		Content:  "João Paulo é um engenheiro de software focado em sistemas distribuídos, aplicações intensivas em dados, fluxos financeiros, consistência transacional e processamento resiliente de alto volume.",
	},
	{
		ID:       "profile-focus",
		Category: "profile",
		Content:  "João Paulo trabalha com consistência transacional, processamento assíncrono, rastreabilidade operacional, arquitetura orientada a eventos, sistemas distribuídos, controle de concorrência e fluxos idempotentes.",
	},
	{
		ID:       "profile-backend-technologies",
		Category: "technology",
		Content:  "João Paulo utiliza tecnologias de backend como Node.js, TypeScript, NestJS, Fastify, Java, Spring Boot e Go.",
	},
	{
		ID:       "profile-frontend-technologies",
		Category: "technology",
		Content:  "João Paulo utiliza Angular no frontend, incluindo SSR, SSG, hydration, lazy loading e otimização de performance.",
	},
	{
		ID:       "profile-database-technologies",
		Category: "technology",
		Content:  "João Paulo trabalha com bancos de dados e armazenamento usando PostgreSQL, MongoDB, Redis e Amazon S3.",
	},
	{
		ID:       "profile-messaging-technologies",
		Category: "technology",
		Content:  "João Paulo utiliza RabbitMQ, BullMQ, WebSocket, SSE e AsyncAPI em sistemas assíncronos e aplicações em tempo real.",
	},
	{
		ID:       "profile-infrastructure-technologies",
		Category: "technology",
		Content:  "João Paulo trabalha com AWS, S3, ECS, IAM, Docker Compose, nginx e pipelines de CI/CD.",
	},
	{
		ID:       "profile-mobile-technologies",
		Category: "technology",
		Content:  "João Paulo utiliza Capacitor para aplicações móveis híbridas e Tauri para aplicações desktop.",
	},
	{
		ID:       "profile-language",
		Category: "profile",
		Content:  "João Paulo possui inglês avançado.",
	},
	{
		ID:       "profile-availability",
		Category: "profile",
		Content:  "João Paulo está aberto a oportunidades de engenharia envolvendo sistemas distribuídos, fluxos críticos de backend e responsabilidade arquitetural.",
	},
	{
		ID:       "career-first-job",
		Category: "career",
		Content:  "O primeiro emprego de João Paulo foi como estagiário em desenvolvimento de sistemas na Representa Online entre junho e agosto de 2024.",
	},
	{
		ID:       "career-current-job",
		Category: "career",
		Content:  "João Paulo trabalha atualmente como desenvolvedor pleno na uFind Tecnologia desde junho de 2025.",
	},
	{
		ID:       "career-current-activities",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo desenvolve um fluxo de faturamento de seguros que processa arquivos financeiros, reconcilia estados transacionais e mantém rastreabilidade de ponta a ponta.",
	},
	{
		ID:       "career-current-impact",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo automatizou um volume mensal de faturamento superior a um milhão de reais e reduziu atividades manuais de dias para minutos.",
	},
	{
		ID:       "career-current-architecture",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo desenvolveu pipelines de ingestão em streaming para diferentes layouts financeiros usando Strategy Pattern, validações críticas e garantias de consistência.",
	},
	{
		ID:       "career-current-aws",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo evoluiu padrões de arquitetura AWS envolvendo S3, ECS e IAM.",
	},
	{
		ID:       "career-junior-job",
		Category: "career",
		Content:  "João Paulo trabalhou como desenvolvedor júnior na Representa Online entre setembro de 2024 e maio de 2025.",
	},
	{
		ID:       "career-junior-data-pipelines",
		Category: "career",
		Content:  "Na Representa Online, João Paulo desenvolveu pipelines de dados orientados a inteligência artificial que processaram mais de 16 GB de conteúdo de televisão.",
	},
	{
		ID:       "career-junior-realtime",
		Category: "career",
		Content:  "Na Representa Online, João Paulo desenvolveu comunicação bidirecional em tempo real usando WebSocket e integração com a API da OpenAI.",
	},
	{
		ID:       "career-intern-activities",
		Category: "career",
		Content:  "Como estagiário na Representa Online, João Paulo desenvolveu funcionalidades de catálogo, descoberta de conteúdo e autenticação.",
	},
	{
		ID:       "career-intern-search",
		Category: "career",
		Content:  "Como estagiário, João Paulo implementou busca por proximidade usando a fórmula de Haversine e otimizou consultas de descoberta de catálogo.",
	},
	{
		ID:       "career-intern-authentication",
		Category: "career",
		Content:  "Como estagiário, João Paulo implementou autenticação com JWT e Google OAuth2.",
	},
	{
		ID:       "career-intern-frontend",
		Category: "career",
		Content:  "Como estagiário, João Paulo desenvolveu páginas com SSR, SSG, lazy loading, otimização de assets e controle de renderização para melhorar TTFB e LCP.",
	},
	{
		ID:       "education-fiap",
		Category: "education",
		Content:  "João Paulo cursa Inteligência Artificial na FIAP entre 2026 e 2028.",
	},
	{
		ID:       "education-etec",
		Category: "education",
		Content:  "João Paulo cursou Desenvolvimento de Sistemas na Etec Guarulhos entre 2023 e 2025.",
	},
	{
		ID:       "certificate-aws",
		Category: "certificate",
		Content:  "João Paulo possui certificação da AWS sobre modelagem de arquitetura orientada a eventos e implantação de microsserviços no Amazon EKS.",
	},
	{
		ID:       "certificate-mongodb",
		Category: "certificate",
		Content:  "João Paulo possui certificação da MongoDB sobre conhecimentos da indústria de serviços financeiros.",
	},
	{
		ID:       "contact-email",
		Category: "contact",
		Content:  "O email público de João Paulo é joaopdias.dev@gmail.com.",
	},
	{
		ID:       "contact-phone",
		Category: "contact",
		Content:  "O telefone público de João Paulo é +55 (11) 98655-3558.",
	},
	{
		ID:       "contact-linkedin",
		Category: "contact",
		Content:  "O LinkedIn de João Paulo está disponível em linkedin.com/in/joaopdias-dev.",
	},
	{
		ID:       "contact-github",
		Category: "contact",
		Content:  "O GitHub de João Paulo está disponível em github.com/Joaopdiasventura.",
	},
	{
		ID:       "services-web",
		Category: "service",
		Content:  "João Paulo pode desenvolver sites e aplicações web usando Angular, SSR, SSG, lazy loading e técnicas de otimização de performance.",
	},
	{
		ID:       "services-backend",
		Category: "service",
		Content:  "João Paulo pode desenvolver backends, APIs, sistemas transacionais, microsserviços e aplicações em tempo real.",
	},
	{
		ID:       "services-mobile",
		Category: "service",
		Content:  "João Paulo pode desenvolver aplicações móveis híbridas com Capacitor e aplicações desktop com Tauri.",
	},
	{
		ID:       "services-business-systems",
		Category: "service",
		Content:  "João Paulo pode desenvolver dashboards, sistemas administrativos, autenticação, catálogos e fluxos transacionais.",
	},
	{
		ID:       "services-data",
		Category: "service",
		Content:  "João Paulo pode desenvolver automações, integrações, pipelines de dados e fluxos orientados a inteligência artificial.",
	},
	{
		ID:       "services-commercial-limitations",
		Category: "service",
		Content:  "O portfólio público de João Paulo não informa preços, prazos, disponibilidade atual, termos contratuais ou garantias de aceitação de projetos.",
	},
	{
		ID:       "project-auronix-description",
		Category: "project",
		Content:  "Auronix é um projeto de banco digital focado em processamento transacional, transferências, liquidação assíncrona e atualização em tempo real.",
	},
	{
		ID:       "project-auronix-technologies",
		Category: "project",
		Content:  "As tecnologias utilizadas no Auronix são NestJS, Fastify, PostgreSQL, Redis, BullMQ, Angular e SSE.",
	},
	{
		ID:       "project-auronix-architecture",
		Category: "project",
		Content:  "O Auronix utiliza Redis para controle de concorrência, BullMQ para liquidação assíncrona e SSE para atualizações em tempo real.",
	},
	{
		ID:       "project-auronix-metrics",
		Category: "project",
		Content:  "O Auronix possui 248 testes automatizados, solicitações com expiração de 10 minutos e uma janela de replay SSE de 100 eventos por 24 horas.",
	},
	{
		ID:       "project-modularis-description",
		Category: "project",
		Content:  "Modularis é uma plataforma de microsserviços orientada a eventos que utiliza uma saga persistida para coordenar identidade, intenção de pagamento, webhooks assinados e ativação premium.",
	},
	{
		ID:       "project-modularis-technologies",
		Category: "project",
		Content:  "As tecnologias utilizadas no Modularis são NestJS, Spring Boot, Go, RabbitMQ, PostgreSQL, MongoDB, Docker Compose, nginx e AsyncAPI.",
	},
	{
		ID:       "project-modularis-metrics",
		Category: "project",
		Content:  "O Modularis possui seis serviços implantáveis, nove canais assíncronos e três stacks de execução.",
	},
	{
		ID:       "project-votrix-description",
		Category: "project",
		Content:  "Votrix é um runtime HTTP de alta performance desenvolvido em TypeScript com roteamento direto, parsing adiado e benchmarks reproduzíveis.",
	},
	{
		ID:       "project-votrix-technologies",
		Category: "project",
		Content:  "As tecnologias utilizadas no Votrix são TypeScript, node:http e autocannon.",
	},
	{
		ID:       "project-votrix-performance",
		Category: "project",
		Content:  "O Votrix alcançou 30124,80 requisições por segundo, superando o Fastify em 37,25 por cento e o Express em 42,08 por cento nos benchmarks medidos.",
	},
	{
		ID:       "project-ggcompress-description",
		Category: "project",
		Content:  "GGCompress é um motor de compressão e arquivamento concorrente desenvolvido em Go com processamento ordenado de chunks e extração determinística.",
	},
	{
		ID:       "project-ggcompress-technologies",
		Category: "project",
		Content:  "As tecnologias utilizadas no GGCompress são Go, goroutines, channels, gzip, SHA-256 e CLI.",
	},
	{
		ID:       "project-ggcompress-integrity",
		Category: "project",
		Content:  "O GGCompress utiliza indexação determinística e verificação de integridade com SHA-256.",
	},
	{
		ID:       "project-ggcompress-performance",
		Category: "project",
		Content:  "O GGCompress alcançou throughput de 1,23 GiB por segundo ao processar uma entrada de 9,77 GiB usando oito workers e chunks de 4 MiB.",
	},
	{
		ID:       "project-vox-description",
		Category: "project",
		Content:  "VOX é uma plataforma de votação focada em integridade, auditabilidade, estado explícito e clareza operacional sob carga concorrente.",
	},
	{
		ID:       "project-vox-technologies",
		Category: "project",
		Content:  "As tecnologias utilizadas no VOX são Tauri, NestJS e Redis.",
	},
	{
		ID:       "project-vox-metrics",
		Category: "project",
		Content:  "O VOX suporta mais de 500 usuários concorrentes, possui 29 rotas HTTP e 108 testes automatizados.",
	},
	{
		ID:       "project-etecfy-description",
		Category: "project",
		Content:  "Etecfy é uma plataforma de streaming de música construída para crescimento de catálogo, descoberta rápida e reprodução contínua.",
	},
	{
		ID:       "project-etecfy-technologies",
		Category: "project",
		Content:  "As tecnologias utilizadas no Etecfy são Angular, NestJS e Capacitor.",
	},
	{
		ID:       "project-etecfy-metrics",
		Category: "project",
		Content:  "O Etecfy recebeu 1,3 mil acessos nas primeiras seis horas, utiliza chunks de streaming de 10 segundos e possui três superfícies de entrega.",
	},
	{
		ID:       "project-comparison-financial",
		Category: "comparison",
		Content:  "Para fluxos financeiros e transacionais, Auronix é o principal exemplo público de João Paulo.",
	},
	{
		ID:       "project-comparison-distributed",
		Category: "comparison",
		Content:  "Para microsserviços distribuídos e arquitetura orientada a eventos, Modularis é o principal exemplo público de João Paulo.",
	},
	{
		ID:       "project-comparison-http-performance",
		Category: "comparison",
		Content:  "Para performance de runtime HTTP, Votrix é o principal exemplo público de João Paulo.",
	},
	{
		ID:       "project-comparison-go-performance",
		Category: "comparison",
		Content:  "Para programação de sistemas, concorrência e performance em Go, GGCompress é o principal exemplo público de João Paulo.",
	},
	{
		ID:       "project-comparison-auditability",
		Category: "comparison",
		Content:  "Para auditabilidade e integridade em sistemas de votação, VOX é o principal exemplo público de João Paulo.",
	},
	{
		ID:       "project-comparison-streaming",
		Category: "comparison",
		Content:  "Para streaming de música e entrega de produto, Etecfy é o principal exemplo público de João Paulo.",
	},
	{
		ID:       "project-comparison-best",
		Category: "comparison",
		Content:  "O portfólio de João Paulo não declara um único melhor projeto, pois cada projeto representa uma especialidade diferente.",
	},
}

var staticDocumentPointers = documentPointers()

func Documents() []*domain.Document {
	return staticDocumentPointers
}

func documentPointers() []*domain.Document {
	documents := make([]*domain.Document, 0, len(staticDocuments))

	for index := range staticDocuments {
		documents = append(documents, &staticDocuments[index])
	}

	return documents
}
