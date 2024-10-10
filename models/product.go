package models

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Category    string  `json:"category"`
	Image       *string `json:"image"` // Pointer to string to accept NULL
	Price       int     `json:"price"`
}
