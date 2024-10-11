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

var ErrCartNotFound = errors.New("cart not found")
var ErrCartItemNotFound = errors.New("cart item not found")

type CartRepository interface {
	CreateCart(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	UpdateCart(ctx context.Context, id int, cart *models.Cart) (*models.Cart, error)
	DeleteCart(ctx context.Context, id int) (int, error)
	GetCartById(ctx context.Context, id int) (*models.Cart, error)
	GetAllCarts(ctx context.Context, userId int, limitStr, sortBy string) ([]models.Cart, error)
}

type cartRepository struct{}

func NewCartRepository() CartRepository {
	return &cartRepository{}
}

func (p *cartRepository) CreateCart(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	return &models.Cart{
		ID:        11,
		UserId:    cart.UserId,
		Products:  cart.Products,
		CreatedAt: cart.CreatedAt,
	}, nil
}

func (p *cartRepository) UpdateCart(ctx context.Context, id int, cart *models.Cart) (*models.Cart, error) {
	return &models.Cart{
		ID:        id,
		UserId:    cart.UserId,
		Products:  cart.Products,
		CreatedAt: cart.CreatedAt,
	}, nil
}

func (p *cartRepository) GetCartById(ctx context.Context, id int) (*models.Cart, error) {
	query := `
		SELECT 
			c.id, 
			c.user_id, 
			c.created_at, 
		FROM 
			carts c
		WHERE 
			c.id = $1 
		LIMIT 1`
	row := db.Pool.QueryRow(ctx, query, id)

	var cart models.Cart
	// Get the cart
	if err := row.Scan(
		&cart.ID,
		&cart.UserId,
		&cart.CreatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrCartNotFound
		}

		log.Printf("Error when fetching cart by id: %v\n", err)
		return nil, err
	}

	cartItemsQuery := `
		SELECT 
			p.id AS id, 
			p.title AS title, 
			p.description AS description, 
			c.title AS category,
			p.image AS image, 
			p.price AS price,
			cp.quantity AS quantity
		FROM 
			cart_products cp
		JOIN 
			carts c ON cp.cart_id = c.id
		JOIN 
			products p ON cp.product_id = p.id
		JOIN 
        	categories cat ON p.category_id = cat.id`

	rows, err := db.Pool.Query(ctx, cartItemsQuery)
	if err != nil {
		log.Fatalln("Error fetching all cart items: ", err)
		return nil, err
	}

	defer rows.Close()

	var cartItems []models.CartProduct

	for rows.Next() {
		var cartItem models.CartProduct

		if err := rows.Scan(
			&cartItem.ID,
			&cartItem.Title,
			&cartItem.Description,
			&cartItem.Category,
			&cartItem.Image,
			&cartItem.Price,
			&cartItem.Quantity,
		); err != nil {
			if err == pgx.ErrNoRows {
				return nil, ErrCartItemNotFound
			}
			log.Printf("Error when fetching cart items: %v\n", err)
			return nil, err
		}

		cartItems = append(cartItems, cartItem)
	}

	// Add cart items to cart
	cart.Products = cartItems

	return &cart, nil
}

func (p *cartRepository) GetAllCarts(ctx context.Context, userId int, limitStr, sortBy string) ([]models.Cart, error) {
	baseQuery := `
		SELECT 
			c.id, 
			c.user_id, 
			c.created_at, 
		FROM 
			carts c`

	qb := utils.NewQueryBuilder(baseQuery)

	if userId != 0 {
		qb.AddCondition("c.user_id = $%d", userId)
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
		log.Fatalln("Error fetching all carts: ", err)
		return nil, err
	}
	defer rows.Close()

	var carts []models.Cart

	for rows.Next() {
		var cart models.Cart

		if err := rows.Scan(
			&cart.ID,
			&cart.UserId,
			&cart.CreatedAt,
		); err != nil {
			if err == pgx.ErrNoRows {
				return nil, ErrCartNotFound
			}
			log.Printf("Error when fetching carts: %v\n", err)
			return nil, err
		}

		carts = append(carts, cart)
	}

	if err = rows.Err(); err != nil {
		log.Println("Error with cart rows: ", err)
		return nil, err
	}

	return carts, nil
}

func (p *cartRepository) DeleteCart(ctx context.Context, id int) (int, error) {
	return id, nil
}
