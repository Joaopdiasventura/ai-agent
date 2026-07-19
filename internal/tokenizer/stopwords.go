package tokenizer

var stopWords = map[string]bool{
	"a": true, "o": true, "as": true, "os": true,
	"de": true, "da": true, "do": true, "das": true, "dos": true,
	"e": true, "é": true, "em": true, "um": true, "uma": true,
	"para": true, "com": true, "que": true, "por": true, "sem": true,
	"qual": true, "quais": true, "quem": true, "quando": true,
	"onde": true, "como": true, "porque": true,
	"foi": true, "era": true, "ser": true, "são": true,
	"ele": true, "ela": true, "dele": true, "dela": true,
	"seu": true, "sua": true, "seus": true, "suas": true,
	"the": true, "an": true, "to": true,
	"is": true, "are": true, "was": true, "were": true,
	"does": true, "did": true, "where": true, "what": true,
	"which": true, "who": true, "how": true, "his": true,
	"him": true, "he": true,
}
