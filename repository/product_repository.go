package repository

import (
	"context"
	"log"

	"github.com/Kei-K23/go-ecommerce-api/db"
	"github.com/Kei-K23/go-ecommerce-api/models"
)

type ProductRepository interface {
	GetProductById(ctx context.Context, id int) (*models.Product, error)
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (p *productRepository) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id, email, password, role FROM users WHERE email=$1 LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, id)

	var product models.Product
	// Get the product
	if err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Category, &product.Image, &product.Price); err != nil {
		log.Printf("Error when fetching product by id: %v\n", err)
		return nil, err
	}

	return &product, nil
}
