package handlers

type AskRequest struct {
	Content string `json:"content"`
}

type AskResponse struct {
	Response string `json:"response"`
}
