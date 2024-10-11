package services

import (
	"context"
	"log"

	"github.com/Kei-K23/go-ecommerce-api/models"
	"github.com/Kei-K23/go-ecommerce-api/repository"
)

type CartService interface {
	CreateCart(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	UpdateCart(ctx context.Context, id int, cart *models.Cart) (*models.Cart, error)
	DeleteCart(ctx context.Context, id int) (int, error)
	GetCartById(ctx context.Context, id int) (*models.Cart, error)
	GetAllCarts(ctx context.Context, userId int, limitStr, sortBy string) ([]models.Cart, error)
}

type cartService struct {
	repo repository.CartRepository
}

func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{repo: repo}

}

// CreateCart implements CartService.
func (p *cartService) CreateCart(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	return p.repo.CreateCart(ctx, cart)
}

func (p *cartService) DeleteCart(ctx context.Context, id int) (int, error) {
	return p.repo.DeleteCart(ctx, id)
}

// CreateCart implements CartService.
func (p *cartService) UpdateCart(ctx context.Context, id int, cart *models.Cart) (*models.Cart, error) {
	return p.repo.UpdateCart(ctx, id, cart)
}

func (p *cartService) GetCartById(ctx context.Context, id int) (*models.Cart, error) {
	cart, err := p.repo.GetCartById(ctx, id)
	if err != nil {
		log.Println("Error in GetCartById:", err)
		return nil, err
	}

	return cart, nil
}

func (p *cartService) GetAllCarts(ctx context.Context, userId int, limitStr, sortBy string) ([]models.Cart, error) {
	carts, err := p.repo.GetAllCarts(ctx, userId, limitStr, sortBy)
	if err != nil {
		log.Println("Error in GetAllCarts:", err)
		return nil, err
	}

	return carts, nil
}
