package services

import (
	"context"
	"log"

	"github.com/Kei-K23/mock-commerce-api/models"
	"github.com/Kei-K23/mock-commerce-api/repository"
)

type ReviewService interface {
	CreateReview(ctx context.Context, review *models.Review) (*models.Review, error)
	UpdateReview(ctx context.Context, id int, review *models.Review) (*models.Review, error)
	DeleteReview(ctx context.Context, id int) (int, error)
	GetReviewById(ctx context.Context, id int) (*models.Review, error)
	GetAllReviews(ctx context.Context, userId, productId int, limitStr, sortBy string) ([]models.Review, error)
}

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) ReviewService {
	return &reviewService{repo: repo}

}

// CreateReview implements ReviewService.
func (p *reviewService) CreateReview(ctx context.Context, review *models.Review) (*models.Review, error) {
	return p.repo.CreateReview(ctx, review)
}

func (p *reviewService) DeleteReview(ctx context.Context, id int) (int, error) {
	return p.repo.DeleteReview(ctx, id)
}

// CreateReview implements ReviewService.
func (p *reviewService) UpdateReview(ctx context.Context, id int, review *models.Review) (*models.Review, error) {
	return p.repo.UpdateReview(ctx, id, review)
}

func (p *reviewService) GetReviewById(ctx context.Context, id int) (*models.Review, error) {
	review, err := p.repo.GetReviewById(ctx, id)
	if err != nil {
		log.Println("Error in GetReviewById:", err)
		return nil, err
	}

	return review, nil
}

func (p *reviewService) GetAllReviews(ctx context.Context, userId, productId int, limitStr, sortBy string) ([]models.Review, error) {
	reviews, err := p.repo.GetAllReviews(ctx, userId, productId, limitStr, sortBy)
	if err != nil {
		log.Println("Error in GetAllReviews:", err)
		return nil, err
	}

	return reviews, nil
}
