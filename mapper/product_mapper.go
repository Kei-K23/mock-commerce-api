package mapper

import (
	"github.com/Kei-K23/go-ecommerce-api/dto"
	"github.com/Kei-K23/go-ecommerce-api/models"
)

// MatchProductRequestToProduct maps a ProductRequest DTO to a Product model
func MatchProductRequestToProduct(req dto.ProductRequest) *models.Product {
	return &models.Product{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Image:       req.Image,
		Price:       req.Price,
	}
}

// MatchProductToProductRequest maps a Product model back to a ProductRequest DTO (optional)
func MatchProductToProductRequest(product models.Product) *dto.ProductRequest {
	return &dto.ProductRequest{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Category:    product.Category,
		Image:       product.Image,
		Price:       product.Price,
	}
}
