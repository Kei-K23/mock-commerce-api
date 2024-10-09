package models

type Product struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
}
