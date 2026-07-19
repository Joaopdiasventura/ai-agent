package nlp

type TemporalContext string

const (
	TemporalUnspecified TemporalContext = ""
	TemporalPresent     TemporalContext = "present"
	TemporalPast        TemporalContext = "past"
	TemporalFirst       TemporalContext = "first"
)

var temporalTokenWeights = map[string]map[TemporalContext]int{
	"atual":         {TemporalPresent: 4},
	"atualmente":    {TemporalPresent: 4},
	"current":       {TemporalPresent: 4},
	"currently":     {TemporalPresent: 4},
	"hoje":          {TemporalPresent: 3},
	"now":           {TemporalPresent: 3},
	"desde":         {TemporalPresent: 3},
	"since":         {TemporalPresent: 3},
	"estuda":        {TemporalPresent: 3},
	"estudando":     {TemporalPresent: 3},
	"cursa":         {TemporalPresent: 3},
	"cursando":      {TemporalPresent: 3},
	"studies":       {TemporalPresent: 3},
	"studying":      {TemporalPresent: 3},
	"trabalha":      {TemporalPresent: 3},
	"atuando":       {TemporalPresent: 3},
	"works":         {TemporalPresent: 3},
	"working":       {TemporalPresent: 3},
	"antes":         {TemporalPast: 4},
	"anterior":      {TemporalPast: 4},
	"anteriores":    {TemporalPast: 4},
	"anteriormente": {TemporalPast: 4},
	"previous":      {TemporalPast: 4},
	"previously":    {TemporalPast: 4},
	"past":          {TemporalPast: 4},
	"formerly":      {TemporalPast: 4},
	"estudou":       {TemporalPast: 4},
	"formou":        {TemporalPast: 4},
	"concluiu":      {TemporalPast: 4},
	"completed":     {TemporalPast: 4},
	"graduated":     {TemporalPast: 4},
	"studied":       {TemporalPast: 4},
	"trabalhou":     {TemporalPast: 4},
	"atuou":         {TemporalPast: 4},
	"worked":        {TemporalPast: 4},
	"did":           {TemporalPast: 2},
	"was":           {TemporalPast: 2},
	"were":          {TemporalPast: 2},
	"primeiro":      {TemporalFirst: 5},
	"primeira":      {TemporalFirst: 5},
	"first":         {TemporalFirst: 5},
	"início":        {TemporalFirst: 3},
	"inicio":        {TemporalFirst: 3},
	"early":         {TemporalFirst: 3},
}

var temporalPriority = []TemporalContext{
	TemporalFirst,
	TemporalPast,
	TemporalPresent,
}

func DetectTemporalContext(tokens []string, intent Intent) TemporalContext {
	scores := make(map[TemporalContext]int)

	for _, token := range tokens {
		tokenScores, exists := temporalTokenWeights[token]
		if !exists {
			continue
		}

		for temporalContext, score := range tokenScores {
			scores[temporalContext] += score
		}
	}

	if intent == IntentFirstJob {
		scores[TemporalFirst] += 4
	}

	if intent == IntentCurrentJob {
		scores[TemporalPresent] += 2
	}

	bestContext := TemporalUnspecified
	bestScore := 0

	for _, temporalContext := range temporalPriority {
		score := scores[temporalContext]
		if score > bestScore {
			bestContext = temporalContext
			bestScore = score
		}
	}

	return bestContext
}
