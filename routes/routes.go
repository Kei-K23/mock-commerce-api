package routes

import (
	"github.com/Kei-K23/go-ecommerce-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(productController *controllers.ProductController, categoryController *controllers.CategoryController, userController *controllers.UserController, cartController *controllers.CartController) *gin.Engine {
	r := gin.Default()

	// Category routes
	r.POST("/category", categoryController.CreateCategory)
	r.PATCH("/categories/:id", categoryController.UpdateCategory)
	r.PUT("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)
	r.GET("/categories", categoryController.GetAllProducts)
	r.GET("/categories/:id", categoryController.GetCategoryById)

	// Product routes
	r.POST("/products", productController.CreateProduct)
	r.PATCH("/products/:id", productController.UpdateProduct)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.GET("/products", productController.GetAllProducts)
	r.GET("/products/:id", productController.GetProductById)

	// User routes
	r.POST("/users", userController.CreateUser)
	r.PATCH("/users/:id", userController.UpdateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUserById)

	// Cart routes
	r.POST("/carts", cartController.CreateCart)
	r.PATCH("/carts/:id", cartController.UpdateCart)
	r.PUT("/carts/:id", cartController.UpdateCart)
	r.DELETE("/carts/:id", cartController.DeleteCart)
	r.GET("/carts", cartController.GetAllCarts)
	r.GET("/carts/:id", cartController.GetCartById)

	return r
}
