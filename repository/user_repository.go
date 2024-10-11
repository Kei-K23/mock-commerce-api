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

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, id int, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (int, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context, title, category, limitStr, sortBy string) ([]models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (p *userRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return &models.User{
		ID:            11,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Password:      user.Password,
		Phone:         user.Phone,
		City:          user.City,
		Street:        user.Street,
		AddressNumber: user.AddressNumber,
		ZipCode:       user.ZipCode,
		Lat:           user.Lat,
		Long:          user.Long,
	}, nil
}

func (p *userRepository) UpdateUser(ctx context.Context, id int, user *models.User) (*models.User, error) {
	// Simulate the update data. This process will not actually update data in to database
	return &models.User{
		ID:            id,
		Username:      user.Username,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Password:      user.Password,
		Phone:         user.Phone,
		City:          user.City,
		Street:        user.Street,
		AddressNumber: user.AddressNumber,
		ZipCode:       user.ZipCode,
		Lat:           user.Lat,
		Long:          user.Long,
	}, nil
}

func (p *userRepository) GetUserById(ctx context.Context, id int) (*models.User, error) {
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
			u.id = $1 
		LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, id)

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
			return nil, ErrUserNotFound
		}

		log.Printf("Error when fetching user by id: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (p *userRepository) GetAllUsers(ctx context.Context, username, city, limitStr, sortBy string) ([]models.User, error) {
	baseQuery := `
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
		FROM users u`

	qb := utils.NewQueryBuilder(baseQuery)

	if username != "" {
		qb.AddCondition("u.username ILIKE $%d", "%"+username+"%")
	}

	if city != "" {
		qb.AddCondition("u.city ILIKE $%d", "%"+city+"%")
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
		log.Fatalln("Error fetching all users: ", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
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
				return nil, ErrUserNotFound
			}
			log.Printf("Error when fetching users: %v\n", err)
			return nil, err
		}

		products = append(products, user)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with user rows: ", err)
		return nil, err
	}

	return products, nil
}

func (p *userRepository) DeleteUser(ctx context.Context, id int) (int, error) {
	return id, nil
}
