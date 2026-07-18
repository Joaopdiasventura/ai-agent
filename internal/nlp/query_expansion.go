package nlp

var queryExpansions = map[string][]string{
	"faz": {
		"trabalha",
		"cursa",
		"estuda",
		"atualmente",
		"desenvolvedor",
	},
	"sobre": {
		"é",
		"como",
		"feito",
		"descrição",
	},
	"trabalho": {
		"trabalha",
		"atualmente",
		"desenvolvedor",
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
