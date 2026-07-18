package knowledge

import "ai-agent/internal/domain"

var staticDocuments = [...]domain.Document{
	{
		ID:       "identity-basic",
		Category: "identity",
		Content:  "João Paulo Dias Ventura é um desenvolvedor de software de São Paulo que cria sistemas digitais para empresas.",
	},
	{
		ID:       "identity-professional-summary",
		Category: "identity",
		Content:  "João Paulo cria sistemas digitais confiáveis, principalmente para empresas que precisam automatizar processos importantes e lidar com muitos dados.",
	},
	{
		ID:       "profile-focus",
		Category: "profile",
		Content:  "João Paulo gosta de transformar processos complexos em sistemas organizados, seguros e fáceis de acompanhar.",
	},
	{
		ID:       "profile-backend-technologies",
		Category: "technology",
		Content:  "João Paulo trabalha na parte dos sistemas que fica por trás das telas, usando Node.js, TypeScript, NestJS, Fastify, Java, Spring Boot e Go.",
	},
	{
		ID:       "profile-frontend-technologies",
		Category: "technology",
		Content:  "João Paulo também desenvolve interfaces web com Angular, buscando páginas rápidas, organizadas e agradáveis de usar.",
	},
	{
		ID:       "profile-database-technologies",
		Category: "technology",
		Content:  "João Paulo trabalha com armazenamento de informações usando PostgreSQL, MongoDB, Redis e Amazon S3.",
	},
	{
		ID:       "profile-messaging-technologies",
		Category: "technology",
		Content:  "João Paulo cria sistemas que trocam informações em tempo real, usando RabbitMQ, BullMQ, WebSocket, SSE e AsyncAPI.",
	},
	{
		ID:       "profile-infrastructure-technologies",
		Category: "technology",
		Content:  "João Paulo sabe colocar aplicações no ar e organizar ambientes usando AWS, S3, ECS, IAM, Docker Compose, nginx e pipelines de CI/CD.",
	},
	{
		ID:       "profile-mobile-technologies",
		Category: "technology",
		Content:  "João Paulo consegue criar aplicações para celular com Capacitor e aplicações para computador com Tauri.",
	},
	{
		ID:       "profile-language",
		Category: "profile",
		Content:  "João Paulo possui inglês avançado.",
	},
	{
		ID:       "profile-availability",
		Category: "profile",
		Content:  "João Paulo está aberto a oportunidades em que possa criar sistemas importantes para o negócio, especialmente automações, produtos digitais e plataformas internas.",
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
		Content:  "Na uFind Tecnologia, João Paulo desenvolve um sistema que automatiza o faturamento de seguros e ajuda a acompanhar cada etapa do processo.",
	},
	{
		ID:       "career-current-impact",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo ajudou a automatizar mais de um milhão de reais por mês em faturamento, reduzindo tarefas manuais de dias para minutos.",
	},
	{
		ID:       "career-current-architecture",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo criou formas de receber e validar diferentes arquivos financeiros sem depender de trabalho manual repetitivo.",
	},
	{
		ID:       "career-current-aws",
		Category: "career",
		Content:  "Na uFind Tecnologia, João Paulo ajudou a organizar aplicações e arquivos na AWS usando S3, ECS e IAM.",
	},
	{
		ID:       "career-junior-job",
		Category: "career",
		Content:  "João Paulo trabalhou como desenvolvedor júnior na Representa Online entre setembro de 2024 e maio de 2025.",
	},
	{
		ID:       "career-junior-data-pipelines",
		Category: "career",
		Content:  "Na Representa Online, João Paulo criou automações que prepararam mais de 16 GB de conteúdo de televisão para uso com inteligência artificial.",
	},
	{
		ID:       "career-junior-realtime",
		Category: "career",
		Content:  "Na Representa Online, João Paulo criou recursos em tempo real e integrou sistemas com a API da OpenAI.",
	},
	{
		ID:       "career-intern-activities",
		Category: "career",
		Content:  "Como estagiário na Representa Online, João Paulo ajudou a criar recursos de catálogo, busca de conteúdo e login de usuários.",
	},
	{
		ID:       "career-intern-search",
		Category: "career",
		Content:  "Como estagiário, João Paulo criou uma busca por proximidade para ajudar usuários a encontrar conteúdos ou opções próximas.",
	},
	{
		ID:       "career-intern-authentication",
		Category: "career",
		Content:  "Como estagiário, João Paulo trabalhou em recursos de login, incluindo acesso com conta Google.",
	},
	{
		ID:       "career-intern-frontend",
		Category: "career",
		Content:  "Como estagiário, João Paulo melhorou páginas web para que carregassem mais rápido e oferecessem uma experiência melhor ao usuário.",
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
		Content:  "João Paulo possui certificação da AWS relacionada à criação e publicação de sistemas em nuvem.",
	},
	{
		ID:       "certificate-mongodb",
		Category: "certificate",
		Content:  "João Paulo possui certificação da MongoDB com foco em tecnologia para serviços financeiros.",
	},
	{
		ID:       "contact-email",
		Category: "contact",
		Content:  "Para entrar em contato com João Paulo por email, use joaopdias.dev@gmail.com.",
	},
	{
		ID:       "contact-phone",
		Category: "contact",
		Content:  "Para entrar em contato com João Paulo por telefone, use +55 (11) 98655-3558.",
	},
	{
		ID:       "contact-linkedin",
		Category: "contact",
		Content:  "Para conhecer o perfil profissional de João Paulo, acesse linkedin.com/in/joaopdias-dev.",
	},
	{
		ID:       "contact-github",
		Category: "contact",
		Content:  "Para ver códigos e projetos de João Paulo, acesse github.com/Joaopdiasventura.",
	},
	{
		ID:       "services-web",
		Category: "service",
		Content:  "João Paulo pode desenvolver sites e aplicações web rápidas, organizadas e fáceis de usar.",
	},
	{
		ID:       "services-backend",
		Category: "service",
		Content:  "João Paulo pode desenvolver a parte interna de sistemas, APIs, automações, integrações e aplicações em tempo real.",
	},
	{
		ID:       "services-mobile",
		Category: "service",
		Content:  "João Paulo pode desenvolver aplicativos para celular e também aplicações para computador.",
	},
	{
		ID:       "services-business-systems",
		Category: "service",
		Content:  "João Paulo pode desenvolver painéis, sistemas administrativos, áreas de login, catálogos e fluxos de pagamento ou operação.",
	},
	{
		ID:       "services-data",
		Category: "service",
		Content:  "João Paulo pode criar automações, integrações entre sistemas e fluxos que organizam dados para uso com inteligência artificial.",
	},
	{
		ID:       "services-commercial-limitations",
		Category: "service",
		Content:  "O portfólio público de João Paulo não informa preços, prazos, disponibilidade atual, termos contratuais ou garantias de aceitação de projetos.",
	},
	{
		ID:       "project-auronix-description",
		Category: "project",
		Content:  "Auronix é um projeto de banco digital criado para simular transferências, acompanhar movimentações e mostrar atualizações em tempo real.",
	},
	{
		ID:       "project-auronix-technologies",
		Category: "project",
		Content:  "No Auronix, João usou NestJS, Fastify, PostgreSQL, Redis, BullMQ, Angular e SSE para criar uma experiência parecida com a de um banco digital.",
	},
	{
		ID:       "project-auronix-architecture",
		Category: "project",
		Content:  "O Auronix foi pensado para evitar conflitos em operações financeiras e mostrar ao usuário o andamento das transações em tempo real.",
	},
	{
		ID:       "project-auronix-metrics",
		Category: "project",
		Content:  "O Auronix possui 248 testes automatizados e regras para expirar solicitações, reforçando a confiabilidade do projeto.",
	},
	{
		ID:       "project-modularis-description",
		Category: "project",
		Content:  "Modularis é uma plataforma que divide um sistema grande em partes menores para organizar login, pagamentos e ativação de planos premium.",
	},
	{
		ID:       "project-modularis-technologies",
		Category: "project",
		Content:  "No Modularis, João usou NestJS, Spring Boot, Go, RabbitMQ, PostgreSQL, MongoDB, Docker Compose, nginx e AsyncAPI.",
	},
	{
		ID:       "project-modularis-metrics",
		Category: "project",
		Content:  "O Modularis mostra a capacidade de João de organizar sistemas grandes em vários serviços que trabalham juntos.",
	},
	{
		ID:       "project-votrix-description",
		Category: "project",
		Content:  "Votrix é um projeto em TypeScript criado para estudar como servidores web podem responder mais rápido.",
	},
	{
		ID:       "project-votrix-technologies",
		Category: "project",
		Content:  "No Votrix, João usou TypeScript, node:http e autocannon para medir e melhorar desempenho.",
	},
	{
		ID:       "project-votrix-performance",
		Category: "project",
		Content:  "O Votrix alcançou mais de 30 mil requisições por segundo em testes, mostrando foco em velocidade e eficiência.",
	},
	{
		ID:       "project-ggcompress-description",
		Category: "project",
		Content:  "GGCompress é uma ferramenta em Go criada para compactar e organizar arquivos grandes de forma rápida e confiável.",
	},
	{
		ID:       "project-ggcompress-technologies",
		Category: "project",
		Content:  "No GGCompress, João usou Go, goroutines, channels, gzip, SHA-256 e CLI para lidar com arquivos grandes pelo terminal.",
	},
	{
		ID:       "project-ggcompress-integrity",
		Category: "project",
		Content:  "O GGCompress verifica se os arquivos continuam íntegros depois do processamento, usando SHA-256.",
	},
	{
		ID:       "project-ggcompress-performance",
		Category: "project",
		Content:  "O GGCompress conseguiu processar quase 10 GiB de dados com alta velocidade em testes controlados.",
	},
	{
		ID:       "project-vox-description",
		Category: "project",
		Content:  "VOX é uma plataforma de votação criada para manter clareza, segurança e confiança mesmo com muitos usuários ao mesmo tempo.",
	},
	{
		ID:       "project-vox-technologies",
		Category: "project",
		Content:  "No VOX, João usou Tauri, NestJS e Redis para criar uma experiência de votação rápida e organizada.",
	},
	{
		ID:       "project-vox-metrics",
		Category: "project",
		Content:  "O VOX foi testado para suportar mais de 500 usuários ao mesmo tempo e possui 108 testes automatizados.",
	},
	{
		ID:       "project-etecfy-description",
		Category: "project",
		Content:  "Etecfy é uma plataforma de música criada para permitir descoberta de faixas, crescimento de catálogo e reprodução contínua.",
	},
	{
		ID:       "project-etecfy-technologies",
		Category: "project",
		Content:  "No Etecfy, João usou Angular, NestJS e Capacitor para entregar a experiência em web e celular.",
	},
	{
		ID:       "project-etecfy-metrics",
		Category: "project",
		Content:  "O Etecfy recebeu 1,3 mil acessos nas primeiras seis horas, mostrando boa recepção inicial do projeto.",
	},
	{
		ID:       "project-comparison-financial",
		Category: "comparison",
		Content:  "Para mostrar experiência com sistemas financeiros, o melhor exemplo público de João Paulo é o Auronix.",
	},
	{
		ID:       "project-comparison-distributed",
		Category: "comparison",
		Content:  "Para mostrar organização de sistemas grandes e divididos em partes, o melhor exemplo público de João Paulo é o Modularis.",
	},
	{
		ID:       "project-comparison-http-performance",
		Category: "comparison",
		Content:  "Para mostrar foco em velocidade de servidores web, o melhor exemplo público de João Paulo é o Votrix.",
	},
	{
		ID:       "project-comparison-go-performance",
		Category: "comparison",
		Content:  "Para mostrar habilidade com Go e processamento de arquivos grandes, o melhor exemplo público de João Paulo é o GGCompress.",
	},
	{
		ID:       "project-comparison-auditability",
		Category: "comparison",
		Content:  "Para mostrar cuidado com confiança e segurança em votações, o melhor exemplo público de João Paulo é o VOX.",
	},
	{
		ID:       "project-comparison-streaming",
		Category: "comparison",
		Content:  "Para mostrar experiência com produto digital e música por streaming, o melhor exemplo público de João Paulo é o Etecfy.",
	},
	{
		ID:       "project-comparison-best",
		Category: "comparison",
		Content:  "João Paulo já fez projetos como Auronix, Modularis, Votrix, GGCompress, VOX e Etecfy; cada um mostra uma habilidade diferente.",
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
