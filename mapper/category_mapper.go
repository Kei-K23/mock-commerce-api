package mapper

import (
	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/models"
)

func MatchCategoryRequestToCategory(req dto.CategoryRequest) *models.Category {
	return &models.Category{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
	}
}

func MatchCategoryToCategoryRequest(category models.Category) *dto.CategoryRequest {
	return &dto.CategoryRequest{
		ID:          category.ID,
		Title:       category.Title,
		Description: category.Description,
		Image:       category.Image,
	}
}
