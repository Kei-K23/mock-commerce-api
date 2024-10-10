package services

import (
	"context"
	"log"

	"github.com/Kei-K23/go-ecommerce-api/models"
	"github.com/Kei-K23/go-ecommerce-api/repository"
)

type ProductService interface {
	GetProductById(ctx context.Context, id int) (*models.Product, error)
	GetAllProducts(ctx context.Context, title, category, limitStr, sortBy string) ([]models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (p *productService) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	product, err := p.repo.GetProductById(ctx, id)
	if err != nil {
		log.Println("Error in GetProductById:", err)
		return nil, err
	}

	return product, nil
}

func (p *productService) GetAllProducts(ctx context.Context, title, category, limitStr, sortBy string) ([]models.Product, error) {
	products, err := p.repo.GetAllProducts(ctx, title, category, limitStr, sortBy)
	if err != nil {
		log.Println("Error in GetAllProducts:", err)
		return nil, err
	}

	return products, nil
}
