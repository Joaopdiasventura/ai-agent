package domain

type Document struct {
	ID             string   `json:"id"`
	Language       string   `json:"language"`
	Category       string   `json:"category"`
	Subject        string   `json:"subject"`
	Project        string   `json:"project,omitempty"`
	TemporalStatus string   `json:"temporal_status"`
	Keywords       []string `json:"keywords"`
	Content        string   `json:"content"`
}

const (
	TemporalPast     = "past"
	TemporalCurrent  = "current"
	TemporalFuture   = "future"
	TemporalTimeless = "timeless"
)
