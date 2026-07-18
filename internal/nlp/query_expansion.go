package nlp

var queryExpansions = map[string][]string{
	"faz": {
		"desenvolvedor",
		"sistemas",
	},
	"sobre": {
		"é",
		"como",
		"feito",
		"descrição",
	},
	"joão": {
		"desenvolvedor",
		"software",
		"sistemas",
	},
	"joao": {
		"desenvolvedor",
		"software",
		"sistemas",
	},
	"who": {
		"developer",
		"software",
		"systems",
	},
	"what": {
		"developer",
		"software",
		"systems",
	},
	"projetos": {
		"projeto",
		"portfólio",
		"portfolio",
	},
	"projects": {
		"project",
		"portfolio",
	},
	"project": {
		"projects",
		"portfolio",
	},
	"portfolio": {
		"projeto",
		"projetos",
		"portfólio",
	},
	"portfólio": {
		"projeto",
		"projetos",
		"portfolio",
	},
	"contato": {
		"entrar",
		"email",
		"telefone",
		"linkedin",
	},
	"contact": {
		"email",
		"phone",
		"linkedin",
	},
	"serviços": {
		"pode",
		"desenvolver",
		"sites",
		"aplicações",
		"automações",
		"integrações",
		"sistemas",
	},
	"servicos": {
		"pode",
		"desenvolver",
		"sites",
		"aplicações",
		"automações",
		"integrações",
		"sistemas",
	},
	"serviço": {
		"pode",
		"desenvolver",
		"sites",
		"aplicações",
		"automações",
		"integrações",
		"sistemas",
	},
	"servico": {
		"pode",
		"desenvolver",
		"sites",
		"aplicações",
		"automações",
		"integrações",
		"sistemas",
	},
	"services": {
		"develop",
		"web",
		"applications",
		"automations",
		"integrations",
		"systems",
	},
	"service": {
		"develop",
		"web",
		"applications",
		"automations",
		"integrations",
		"systems",
	},
	"contratar": {
		"confiáveis",
		"negócio",
		"automações",
		"produtos",
	},
	"hire": {
		"reliable",
		"business",
		"automation",
		"products",
	},
	"hiring": {
		"reliable",
		"business",
		"automation",
		"products",
	},
	"contrataria": {
		"confiáveis",
		"negócio",
		"automações",
		"produtos",
	},
	"trabalho": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"work": {
		"currently",
		"developer",
	},
	"works": {
		"currently",
		"developer",
	},
	"trabalha": {
		"atualmente",
		"desenvolvedor",
	},
	"profissão": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"profissao": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"ocupação": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"ocupacao": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"feito": {
		"desenvolvido",
		"criado",
		"como",
		"utilizou",
	},
	"utiliza": {
		"usa",
		"utilizou",
		"fez",
		"usou",
		"feito",
		"desenvolvido",
	},
	"uses": {
		"technologies",
		"stack",
		"developed",
	},
	"use": {
		"technologies",
		"stack",
		"developed",
	},
	"cargo": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
	},
	"faculdade": {
		"cursa",
		"fiap",
		"inteligência",
		"artificial",
	},
	"estuda": {
		"cursa",
		"fiap",
	},
	"curso": {
		"cursa",
		"inteligência",
		"artificial",
	},
	"banco": {
		"auronix",
		"digital",
		"transacional",
	},
	"compressão": {
		"ggcompress",
		"gzip",
		"concorrente",
	},
	"compressao": {
		"ggcompress",
		"gzip",
		"concorrente",
	},
}

func ExpandQuery(tokens []string) []string {
	expandedTokens := make([]string, 0, len(tokens))
	seenTokens := make(map[string]struct{})

	addToken := func(token string) {
		if token == "" {
			return
		}

		if _, exists := seenTokens[token]; exists {
			return
		}

		seenTokens[token] = struct{}{}
		expandedTokens = append(expandedTokens, token)
	}

	for _, token := range tokens {
		addToken(token)
	}

	for index := 0; index < len(expandedTokens); index++ {
		token := expandedTokens[index]

		expansions, exists := queryExpansions[token]
		if !exists {
			continue
		}

		for _, expansion := range expansions {
			addToken(expansion)
		}
	}

	return expandedTokens
}
