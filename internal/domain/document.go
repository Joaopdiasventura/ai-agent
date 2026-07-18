package domain

type Document struct {
	ID       string `json:"id"`
	Category string `json:"category"`
	Language string `json:"language"`
	Content  string `json:"content"`
}
