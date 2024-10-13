package repository

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/Kei-K23/mock-commerce-api/db"
	"github.com/Kei-K23/mock-commerce-api/models"
	"github.com/Kei-K23/mock-commerce-api/utils"
	"github.com/jackc/pgx/v5"
)

var ErrReviewNotFound = errors.New("review not found")

type ReviewRepository interface {
	CreateReview(ctx context.Context, review *models.Review) (*models.Review, error)
	UpdateReview(ctx context.Context, id int, review *models.Review) (*models.Review, error)
	DeleteReview(ctx context.Context, id int) (int, error)
	GetReviewById(ctx context.Context, id int) (*models.Review, error)
	GetAllReviews(ctx context.Context, userId, productId int, limitStr, sortBy string) ([]models.Review, error)
}

type reviewRepository struct{}

func NewReviewRepository() ReviewRepository {
	return &reviewRepository{}
}

func (p *reviewRepository) CreateReview(ctx context.Context, review *models.Review) (*models.Review, error) {

	return &models.Review{
		ID:        11,
		ProductId: review.ProductId,
		UserId:    review.UserId,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}, nil
}

func (p *reviewRepository) UpdateReview(ctx context.Context, id int, review *models.Review) (*models.Review, error) {
	return &models.Review{
		ID:        id,
		ProductId: review.ProductId,
		UserId:    review.UserId,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}, nil
}

func (p *reviewRepository) GetReviewById(ctx context.Context, id int) (*models.Review, error) {
	query := `
		SELECT 
			r.id, 
			r.product_id, 
			r.user_id, 
			r.rating,
			r.comment, 
			r.created_at
		FROM 
			reviews r
		WHERE 
			r.id = $1 
		LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, id)

	var review models.Review
	if err := row.Scan(
		&review.ID,
		&review.ProductId,
		&review.UserId,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrReviewNotFound
		}

		log.Printf("Error when fetching review by id: %v\n", err)
		return nil, err
	}

	return &review, nil
}

func (p *reviewRepository) GetAllReviews(ctx context.Context, userId, productId int, limitStr, sortBy string) ([]models.Review, error) {
	baseQuery := `
		SELECT 
			r.id, 
			r.product_id, 
			r.user_id, 
			r.rating,
			r.comment, 
			r.created_at
		FROM 
			reviews r
		`

	qb := utils.NewQueryBuilder(baseQuery)

	if userId != 0 {
		qb.AddCondition("r.user_id = $%d", userId)
	}

	if productId != 0 {
		qb.AddCondition("r.product_id = $%d", productId)
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
		log.Fatalln("Error fetching all reviews: ", err)
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review

	for rows.Next() {
		var review models.Review

		if err := rows.Scan(
			&review.ID,
			&review.ProductId,
			&review.UserId,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		); err != nil {
			if err == pgx.ErrNoRows {
				return nil, ErrReviewNotFound
			}
			log.Printf("Error when fetching reviews: %v\n", err)
			return nil, err
		}

		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with review rows: ", err)
		return nil, err
	}

	return reviews, nil
}

func (p *reviewRepository) DeleteReview(ctx context.Context, id int) (int, error) {
	return id, nil
}
