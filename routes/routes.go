package routes

import (
	"github.com/Kei-K23/go-ecommerce-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(productController *controllers.ProductController, categoryController *controllers.CategoryController) *gin.Engine {
	r := gin.Default()

	r.GET("/categories", categoryController.GetAllProducts)
	r.GET("/categories/:id", categoryController.GetCategoryById)

	// Product routes
	r.GET("/products", productController.GetAllProducts)
	r.GET("/products/:id", productController.GetProductById)
	return r
}
