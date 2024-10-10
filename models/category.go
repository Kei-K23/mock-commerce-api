package models

type Category struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
}
