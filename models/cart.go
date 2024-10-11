package models

import "time"

type CartProduct struct {
	Product
	Quantity int `json:"quantity"`
}

type Cart struct {
	ID        int           `json:"id"`
	UserId    int           `json:"userId"`
	Products  []CartProduct `json:"products"`
	CreatedAt time.Time     `json:"createdAt"`
}
