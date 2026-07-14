package domain

type Document struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Content  string `json:"content"`
}
