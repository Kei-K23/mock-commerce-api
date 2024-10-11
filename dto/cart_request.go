package dto

import (
	"time"

	"github.com/Kei-K23/go-ecommerce-api/models"
)

type CartRequest struct {
	ID        int                  `json:"id"`
	UserId    int                  `json:"userId"`
	Products  []models.CartProduct `json:"products"`
	CreatedAt time.Time            `json:"createdAt"`
}
