package services

import (
	"context"
	"log"

	"github.com/Kei-K23/go-ecommerce-api/models"
	"github.com/Kei-K23/go-ecommerce-api/repository"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error)
	UpdateCategory(ctx context.Context, id int, category *models.Category) (*models.Category, error)
	GetCategoryById(ctx context.Context, id int) (*models.Category, error)
	GetAllCategories(ctx context.Context, title, limitStr, sortBy string) ([]models.Category, error)
	DeleteCategory(ctx context.Context, id int) (int, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// CreateProduct implements ProductService.
func (p *categoryService) CreateCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return p.repo.CreateCategory(ctx, category)
}

// CreateProduct implements ProductService.
func (p *categoryService) UpdateCategory(ctx context.Context, id int, category *models.Category) (*models.Category, error) {
	return p.repo.UpdateCategory(ctx, id, category)
}

func (p *categoryService) GetCategoryById(ctx context.Context, id int) (*models.Category, error) {
	category, err := p.repo.GetCategoryById(ctx, id)
	if err != nil {
		log.Println("Error in GetCategoryById: ", err)
		return nil, err
	}

	return category, nil
}

func (p *categoryService) GetAllCategories(ctx context.Context, title, limitStr, sortBy string) ([]models.Category, error) {
	categories, err := p.repo.GetAllCategories(ctx, title, limitStr, sortBy)
	if err != nil {
		log.Println("Error in GetAllCategories: ", err)
		return nil, err
	}

	return categories, nil
}

func (p *categoryService) DeleteCategory(ctx context.Context, id int) (int, error) {
	return p.repo.DeleteCategory(ctx, id)
}
