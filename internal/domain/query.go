package domain

type Query struct {
	Text           string
	Tokens         []string
	Language       string
	Category       string
	Project        string
	TemporalStatus string
	ExactTerms     []string
	Vector         []float32
}

type SearchResult struct {
	Document       *Document
	Score          float64
	VectorRank     int
	LexicalRank    int
	FusedRank      int
	Sources        []string
	MetadataScore  float64
	FinalScore     float64
	PenaltyReasons []string
}

type Evidence struct {
	DocumentID     string
	Language       string
	Category       string
	Project        string
	TemporalStatus string
	Content        string
	Score          float64
	Sources        []string
}
