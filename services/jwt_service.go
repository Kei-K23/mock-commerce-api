package services

import (
	"context"

	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/repository"
)

type JWTService interface {
	CreateJWT(ctx context.Context, userReq *dto.LoginRequest) (string, error)
}

type jwtService struct {
	repo repository.JWTRepository
}

func NewJWTService(repo repository.JWTRepository) JWTService {
	return &jwtService{repo: repo}

}

// CreateJWT implements JWTService.
func (p *jwtService) CreateJWT(ctx context.Context, userReq *dto.LoginRequest) (string, error) {
	return p.repo.CreateJWT(ctx, userReq)
}
