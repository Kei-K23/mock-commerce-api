package mapper

import (
	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/models"
)

func MatchCartRequestToCart(req dto.CartRequest) *models.Cart {
	return &models.Cart{
		UserId:    req.UserId,
		Products:  req.Products,
		CreatedAt: req.CreatedAt,
	}
}

func MatchCartToCartRequest(category models.Cart) *dto.CartRequest {
	return &dto.CartRequest{
		ID:        category.ID,
		UserId:    category.UserId,
		Products:  category.Products,
		CreatedAt: category.CreatedAt,
	}
}
