package model

type Project struct {
	ID          string `json:"id"`
	UserID      string `json:"userId"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
