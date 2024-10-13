package services

import (
	"context"
	"log"

	"github.com/Kei-K23/mock-commerce-api/models"
	"github.com/Kei-K23/mock-commerce-api/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, product *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, id int, product *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context, username, city, limitStr, sortBy string) ([]models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}

}

// CreateUser implements UserService.
func (p *userService) CreateUser(ctx context.Context, product *models.User) (*models.User, error) {
	return p.repo.CreateUser(ctx, product)
}

func (p *userService) DeleteUser(ctx context.Context, id int) (int, error) {
	return p.repo.DeleteUser(ctx, id)
}

// CreateUser implements UserService.
func (p *userService) UpdateUser(ctx context.Context, id int, product *models.User) (*models.User, error) {
	return p.repo.UpdateUser(ctx, id, product)
}

func (p *userService) GetUserById(ctx context.Context, id int) (*models.User, error) {
	product, err := p.repo.GetUserById(ctx, id)
	if err != nil {
		log.Println("Error in GetUserById:", err)
		return nil, err
	}

	return product, nil
}

func (p *userService) GetAllUsers(ctx context.Context, title, category, limitStr, sortBy string) ([]models.User, error) {
	products, err := p.repo.GetAllUsers(ctx, title, category, limitStr, sortBy)
	if err != nil {
		log.Println("Error in GetAllUsers:", err)
		return nil, err
	}

	return products, nil
}
