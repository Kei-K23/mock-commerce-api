package services

import (
	"context"
	"log"

	"github.com/Kei-K23/mock-commerce-api/models"
	"github.com/Kei-K23/mock-commerce-api/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	UpdateProduct(ctx context.Context, id int, product *models.Product) (*models.Product, error)
	DeleteProduct(ctx context.Context, id int) (int, error)
	GetProductById(ctx context.Context, id int) (*models.Product, error)
	GetAllProducts(ctx context.Context, title, category, limitStr, sortBy string) ([]models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}

}

// CreateProduct implements ProductService.
func (p *productService) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	return p.repo.CreateProduct(ctx, product)
}

func (p *productService) DeleteProduct(ctx context.Context, id int) (int, error) {
	return p.repo.DeleteProduct(ctx, id)
}

// CreateProduct implements ProductService.
func (p *productService) UpdateProduct(ctx context.Context, id int, product *models.Product) (*models.Product, error) {
	return p.repo.UpdateProduct(ctx, id, product)
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
