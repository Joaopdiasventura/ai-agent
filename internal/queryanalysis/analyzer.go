package queryanalysis

import (
	"ai-agent/internal/domain"
	"regexp"
	"strings"
	"unicode"
)

var exactTermPattern = regexp.MustCompile(`[\w.+-]+@[\w.-]+|\+?\d[\d\s().-]{5,}\d`)

func Analyze(text string) domain.Query {
	tokens := queryTokens(text)

	return domain.Query{
		Text:           strings.TrimSpace(text),
		Tokens:         tokens,
		Language:       detectLanguage(tokens),
		Category:       detectCategory(tokens),
		Project:        detectProject(tokens),
		TemporalStatus: detectTemporalStatus(tokens),
		ExactTerms:     exactTerms(text),
	}
}

func queryTokens(text string) []string {
	text = normalize(text)
	fields := strings.FieldsFunc(text, func(value rune) bool {
		return !(unicode.IsLetter(value) || unicode.IsDigit(value) || value == '@' || value == '+' || value == '#' || value == '.')
	})

	tokens := make([]string, 0, len(fields))
	for _, field := range fields {
		field = strings.Trim(field, ".")
		if field != "" {
			tokens = append(tokens, field)
		}
	}

	return tokens
}

func detectLanguage(tokens []string) string {
	ptScore := 0
	enScore := 0

	for _, token := range tokens {
		ptScore += mapScore(token, map[string]int{
			"qual": 3, "quais": 3, "onde": 3, "que": 3, "joao": 2, "joão": 2, "ele": 3, "dele": 3,
			"estuda": 3, "estudou": 3, "trabalha": 3, "trabalhou": 3, "projeto": 2,
			"telefone": 3, "formacao": 3, "formação": 3, "experiencia": 2, "experiência": 2,
			"faz": 3, "fazer": 3, "pode": 3,
		})
		enScore += mapScore(token, map[string]int{
			"what": 3, "which": 3, "where": 3, "who": 3, "his": 3, "email": 1,
			"study": 3, "studied": 3, "work": 3, "worked": 3, "project": 2,
			"phone": 3, "education": 3, "experience": 2, "about": 2,
			"unrelated": 3, "question": 2, "does": 3, "do": 2, "can": 3,
		})
	}

	if enScore > ptScore {
		return "en"
	}

	return "pt"
}

func detectCategory(tokens []string) string {
	if containsAnyToken(tokens, "projeto", "projetos", "project", "projects") {
		return "project"
	}
	if isServiceCapabilityQuestion(tokens) {
		return "service"
	}
	if isProfileActivityQuestion(tokens) {
		return "profile"
	}

	scores := map[string]int{}

	for _, token := range tokens {
		for category, score := range categoryScores(token) {
			scores[category] += score
		}
	}

	return bestScore(scores)
}

func categoryScores(token string) map[string]int {
	switch token {
	case "faz", "does":
		return map[string]int{"profile": 4}
	case "profile", "perfil":
		return map[string]int{"profile": 4}
	case "quem", "who":
		return map[string]int{"identity": 4}
	case "email", "telefone", "phone", "contato", "contact", "linkedin", "github":
		return map[string]int{"contact": 5}
	case "estuda", "estudou", "formacao", "formação", "faculdade", "education", "study", "studied", "school", "fiap", "etec":
		return map[string]int{"education": 5}
	case "trabalha", "trabalhou", "emprego", "cargo", "work", "worked", "job", "career", "ufind", "representa":
		return map[string]int{"career": 5}
	case "tecnologia", "tecnologias", "technology", "technologies", "stack", "java", "go", "angular", "react", "spring", "kubernetes", "aws":
		return map[string]int{"technology": 4}
	case "projeto", "projetos", "project", "projects", "auronix", "xtube", "x", "tube", "ggcompress", "auditex":
		return map[string]int{"project": 4}
	case "impacto", "resultado", "result", "impact", "reduziu", "redução", "reduction", "demonstrates", "demonstra":
		return map[string]int{"impact": 3}
	case "certificacao", "certificacoes", "certificação", "certificações", "certificate", "certification", "mongodb", "edb":
		return map[string]int{"certificate": 4}
	case "servico", "servicos", "serviço", "serviços", "service", "services", "empresa", "business", "fazer", "pode", "can", "build", "deliver":
		return map[string]int{"service": 4}
	default:
		return nil
	}
}

func isServiceCapabilityQuestion(tokens []string) bool {
	return containsAnyToken(tokens, "pode", "can", "could") &&
		containsAnyToken(tokens, "fazer", "faz", "do", "build", "deliver", "oferecer", "oferece")
}

func isProfileActivityQuestion(tokens []string) bool {
	return containsAnyToken(tokens, "que", "what") &&
		containsAnyToken(tokens, "faz", "does", "do")
}

func detectTemporalStatus(tokens []string) string {
	scores := map[string]int{}

	for _, token := range tokens {
		switch token {
		case "atual", "atualmente", "current", "currently", "trabalha", "works":
			scores[domain.TemporalCurrent] += 4
		case "previsto", "prevista", "planned", "future", "vai":
			scores[domain.TemporalFuture] += 4
		case "estudou", "trabalhou", "antes", "anterior", "anteriormente", "primeiro", "worked", "studied", "previous", "before", "first":
			scores[domain.TemporalPast] += 4
		case "estuda", "cursa", "study":
			scores[domain.TemporalFuture] += 2
		}
	}

	best := bestScore(scores)
	if best == "" {
		return domain.TemporalTimeless
	}

	return best
}

func detectProject(tokens []string) string {
	joined := strings.Join(tokens, " ")

	switch {
	case strings.Contains(joined, "auronix"):
		return "auronix"
	case strings.Contains(joined, "x tube"), strings.Contains(joined, "xtube"):
		return "x-tube"
	case strings.Contains(joined, "ggcompress"):
		return "ggcompress"
	case strings.Contains(joined, "auditex"):
		return "auditex"
	case containsAnyToken(tokens, "lideranca", "liderança", "leadership"):
		return "x-tube"
	case containsAnyToken(tokens, "complexo", "complexos", "complex", "dificil", "difícil", "dificeis", "difíceis", "desafio", "challenge"):
		return "x-tube"
	case containsAnyToken(tokens, "concorrencia", "concorrência", "desempenho", "performance", "throughput", "benchmark", "go", "golang"):
		return "ggcompress"
	case containsAnyToken(tokens, "auditabilidade", "auditability", "criptografia", "cryptography", "historica", "histórica", "historical"):
		return "auditex"
	case containsAnyToken(tokens, "destacaria", "destacar", "recrutador", "recruiter", "highlight"):
		return "auronix"
	case containsAnyToken(tokens, "financeiro", "financeiros", "financial", "banco", "banking", "capacidade", "tecnica", "técnica"):
		return "auronix"
	default:
		return ""
	}
}

func exactTerms(text string) []string {
	matches := exactTermPattern.FindAllString(text, -1)
	terms := make([]string, 0, len(matches))
	seen := make(map[string]struct{}, len(matches))

	for _, match := range matches {
		match = strings.TrimSpace(match)
		if match == "" {
			continue
		}

		key := strings.ToLower(match)
		if _, exists := seen[key]; exists {
			continue
		}

		seen[key] = struct{}{}
		terms = append(terms, match)
	}

	return terms
}

func bestScore(scores map[string]int) string {
	bestKey := ""
	bestValue := 0

	for key, value := range scores {
		if value > bestValue || value == bestValue && (bestKey == "" || key < bestKey) {
			bestKey = key
			bestValue = value
		}
	}

	return bestKey
}

func mapScore(token string, scores map[string]int) int {
	return scores[token]
}

func containsAnyToken(tokens []string, terms ...string) bool {
	for _, token := range tokens {
		for _, term := range terms {
			if token == term {
				return true
			}
		}
	}

	return false
}

func normalize(text string) string {
	replacer := strings.NewReplacer(
		"á", "a", "à", "a", "â", "a", "ã", "a", "ä", "a",
		"é", "e", "ê", "e", "è", "e", "ë", "e",
		"í", "i", "î", "i", "ì", "i", "ï", "i",
		"ó", "o", "ô", "o", "õ", "o", "ò", "o", "ö", "o",
		"ú", "u", "û", "u", "ù", "u", "ü", "u",
		"ç", "c",
	)

	return replacer.Replace(strings.ToLower(text))
}
