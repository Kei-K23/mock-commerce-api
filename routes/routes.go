package routes

import (
	"github.com/Kei-K23/go-ecommerce-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(productController *controllers.ProductController) *gin.Engine {
	r := gin.Default()

	r.GET("/products/:id", productController.GetProductById)

	return r
}
