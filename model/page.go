package model

type Page struct {
	Todolist []Todo `json:"page,omitempty"`
	Total    int    `json:"total"`
	Pages    int    `json:"pages"`
}
