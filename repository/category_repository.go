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

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepository interface {
	GetCategoryById(ctx context.Context, id int) (*models.Category, error)
	GetAllCategories(ctx context.Context, title, limitStr, sortBy string) ([]models.Category, error)
}

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (p *categoryRepository) GetCategoryById(ctx context.Context, id int) (*models.Category, error) {
	query := `SELECT id, title, description, image FROM categories WHERE id=$1 LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, id)

	var category models.Category
	// Get the category
	if err := row.Scan(
		&category.ID,
		&category.Title,
		&category.Description,
		&category.Image,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrCategoryNotFound
		}

		log.Printf("Error when fetching category by id: %v\n", err)
		return nil, err
	}

	return &category, nil
}

func (p *categoryRepository) GetAllCategories(ctx context.Context, title, limitStr, sortBy string) ([]models.Category, error) {

	// Base query
	baseQuery := "SELECT id, title, description, image FROM categories"

	qb := utils.NewQueryBuilder(baseQuery)

	if title != "" {
		qb.AddCondition("title ILIKE $%d", "%"+title+"%")
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
		log.Fatalln("Error fetching all categories: ", err)
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var category models.Category
		// Get the category
		if err := rows.Scan(
			&category.ID,
			&category.Title,
			&category.Description,
			&category.Image,
		); err != nil {
			if err == pgx.ErrNoRows {
				return nil, ErrCategoryNotFound
			}
			log.Printf("Error when fetching categories: %v\n", err)
			return nil, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with category rows: ", err)
		return nil, err
	}

	return categories, nil
}
