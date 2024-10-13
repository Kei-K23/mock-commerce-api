package mapper

import (
	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/models"
)

func MatchReviewRequestToReview(req dto.ReviewRequest) *models.Review {
	return &models.Review{
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Rating:    req.Rating,
		Comment:   req.Comment,
		CreatedAt: req.CreatedAt,
	}
}

func MatchReviewToReviewRequest(review models.Review) *dto.ReviewRequest {
	return &dto.ReviewRequest{
		ID:        review.ID,
		ProductId: review.ProductId,
		UserId:    review.UserId,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}
}
