package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kei-K23/mock-commerce-api/controllers"
	"github.com/Kei-K23/mock-commerce-api/db"
	"github.com/Kei-K23/mock-commerce-api/repository"
	"github.com/Kei-K23/mock-commerce-api/routes"
	"github.com/Kei-K23/mock-commerce-api/services"
)

func main() {
	// Start the db connection
	db.ConnectDB()

	// Close the db connection when before main function out of scope
	defer db.Pool.Close()

	categoryRepo := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	productRepo := repository.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	cartRepo := repository.NewCartRepository()
	cartService := services.NewCartService(cartRepo)
	cartController := controllers.NewCartController(cartService)

	reviewRepo := repository.NewReviewRepository()
	reviewService := services.NewReviewService(reviewRepo)
	reviewController := controllers.NewReviewController(reviewService)

	jwtRepo := repository.NewJWTRepository()
	jwtService := services.NewJWTService(jwtRepo)
	jwtController := controllers.NewJWTController(jwtService)

	// Get the router
	r := routes.SetupRouter(productController, categoryController, userController, cartController, jwtController, reviewController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the server
	if err := r.Run(":" + port); err != nil {
		log.Fatal("failed to run the server: ", err)
	}

	fmt.Printf("Server is running on PORT: %s\n", port)
}
