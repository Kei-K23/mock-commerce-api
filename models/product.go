package models

type Product struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
