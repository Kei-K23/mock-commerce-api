package repository

import (
	"context"
	"log"

	"github.com/Kei-K23/go-ecommerce-api/db"
	"github.com/Kei-K23/go-ecommerce-api/dto"
	"github.com/Kei-K23/go-ecommerce-api/models"
	"github.com/Kei-K23/go-ecommerce-api/utils"
	"github.com/jackc/pgx/v5"
)

type JWTRepository interface {
	CreateJWT(ctx context.Context, userReq *dto.LoginRequest) (string, error)
}

type jwtRepository struct{}

func NewJWTRepository() JWTRepository {
	return &jwtRepository{}
}

func (p *jwtRepository) CreateJWT(ctx context.Context, userReq *dto.LoginRequest) (string, error) {
	query := `
		SELECT 
			u.id, 
			u.username, 
			u.firstname, 
			u.lastname,
			u.email, 
			u.password, 
			u.phone, 
			u.city, 
			u.street, 
			u.address_number, 
			u.zip_code, 
			u.lat, 
			u.long 
		FROM users u 
		WHERE 
			u.username = $1 
		AND
			u.password = $2
		LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, userReq.Username, userReq.Password)

	var user models.User
	// Get the user
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.City,
		&user.Street,
		&user.AddressNumber,
		&user.ZipCode,
		&user.Lat,
		&user.Long,
	); err != nil {
		if err == pgx.ErrNoRows {
			return "", ErrUserNotFound
		}

		log.Printf("Error when fetching user by username and password: %v\n", err)
		return "", err
	}

	jwtToken, err := utils.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
