package services

import (
	"context"
	"log"

	"github.com/Kei-K23/go-ecommerce-api/models"
	"github.com/Kei-K23/go-ecommerce-api/repository"
)

type ProductService interface {
	GetProductById(ctx context.Context, id int) (*models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

// GetProductById implements ProductService.
func (p *productService) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	product, err := p.repo.GetProductById(ctx, id)
	if err != nil {
		log.Println("Error in GetProductById:", err)
		return nil, err
	}

	return product, nil
}
