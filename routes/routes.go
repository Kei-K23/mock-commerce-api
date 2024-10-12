package routes

import (
	"net/http"

	"github.com/Kei-K23/go-ecommerce-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(productController *controllers.ProductController, categoryController *controllers.CategoryController, userController *controllers.UserController, cartController *controllers.CartController, jwtController *controllers.JWTController) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Version 1
	v1 := r.Group("/api/v1")
	{
		// Category routes
		v1.POST("/category", categoryController.CreateCategory)
		v1.PATCH("/categories/:id", categoryController.UpdateCategory)
		v1.PUT("/categories/:id", categoryController.UpdateCategory)
		v1.DELETE("/categories/:id", categoryController.DeleteCategory)
		v1.GET("/categories", categoryController.GetAllProducts)
		v1.GET("/categories/:id", categoryController.GetCategoryById)

		// Product routes
		v1.POST("/products", productController.CreateProduct)
		v1.PATCH("/products/:id", productController.UpdateProduct)
		v1.PUT("/products/:id", productController.UpdateProduct)
		v1.DELETE("/products/:id", productController.DeleteProduct)
		v1.GET("/products", productController.GetAllProducts)
		v1.GET("/products/:id", productController.GetProductById)
		v1.GET("/products/category/:category", productController.GetAllProductsByCategory)

		// User routes
		v1.POST("/users", userController.CreateUser)
		v1.PATCH("/users/:id", userController.UpdateUser)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)
		v1.GET("/users", userController.GetAllUsers)
		v1.GET("/users/:id", userController.GetUserById)

		// Cart routes
		v1.POST("/carts", cartController.CreateCart)
		v1.PATCH("/carts/:id", cartController.UpdateCart)
		v1.PUT("/carts/:id", cartController.UpdateCart)
		v1.DELETE("/carts/:id", cartController.DeleteCart)
		v1.GET("/carts", cartController.GetAllCarts)
		v1.GET("/carts/:id", cartController.GetCartById)
		v1.GET("/carts/user/:userId", cartController.GetAllCartsByUserId)

		// Auth routes
		v1.POST("/auth/login", jwtController.CreateJWT)
	}

	// Documentation HTML site
	{
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.templ", gin.H{
				"title": "Hello from Templ + Go + Gin",
			})
		})
	}

	return r
}
