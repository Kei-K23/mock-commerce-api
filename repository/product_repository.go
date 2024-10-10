package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/Kei-K23/go-ecommerce-api/db"
	"github.com/Kei-K23/go-ecommerce-api/models"
	"github.com/Kei-K23/go-ecommerce-api/utils"
	"github.com/jackc/pgx/v5"
)

var ErrProductNotFound = errors.New("product not found")

type ProductRepository interface {
	GetProductById(ctx context.Context, id int) (*models.Product, error)
	GetAllProducts(ctx context.Context, title, category, limitStr, sortBy string) ([]models.Product, error)
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (p *productRepository) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id, title, description, category, image, price FROM products WHERE id=$1 LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, id)

	var product models.Product
	// Get the product
	if err := row.Scan(
		&product.ID,
		&product.Title,
		&product.Description,
		&product.Category,
		&product.Image, // This can now be NULL
		&product.Price,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrProductNotFound
		}

		log.Printf("Error when fetching product by id: %v\n", err)
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) GetAllProducts(ctx context.Context, title, category, limitStr, sortBy string) ([]models.Product, error) {

	// Base query
	baseQuery := "SELECT id, title, description, category, image, price FROM products"

	qb := utils.NewQueryBuilder(baseQuery)

	if title != "" {
		qb.AddCondition("title ILIKE $%d", "%"+title+"%")
	}

	if category != "" {
		qb.AddCondition("category = $%d", category)
	}

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			log.Fatalln("Error when parsing limit string to int: ", err)
			return nil, err
		}

		qb.SetLimit(limit)
	}

	if sortBy != "" {
		qb.SetSortBy(sortBy)
	}

	query, params := qb.Build()

	rows, err := db.Pool.Query(ctx, query, params...)
	if err != nil {
		log.Fatalln("Error fetching all products: ", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		// Get the product
		if err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.Category,
			&product.Image,
			&product.Price,
		); err != nil {
			if err == pgx.ErrNoRows {
				return nil, ErrProductNotFound
			}
			log.Printf("Error when fetching products: %v\n", err)
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with product rows: ", err)
		return nil, err
	}

	return products, nil
}
