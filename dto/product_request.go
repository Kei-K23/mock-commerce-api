package dto

type ProductRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Category    string  `json:"category"`
	Image       *string `json:"image"`
	Price       int     `json:"price"`
}
