package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/mapper"
	"github.com/Kei-K23/mock-commerce-api/repository"
	"github.com/Kei-K23/mock-commerce-api/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var productRequest dto.ProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := mapper.MatchProductRequestToProduct(productRequest)

	createdProduct, err := p.service.CreateProduct(c.Request.Context(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productRequest dto.ProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := mapper.MatchProductRequestToProduct(productRequest)

	createdProduct, err := p.service.UpdateProduct(c.Request.Context(), id, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdProduct)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedId, err := p.service.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedId})
}

func (p *ProductController) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := p.service.GetProductById(c.Request.Context(), id)
	if err != nil {
		if err == repository.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) GetAllProducts(c *gin.Context) {

	limitStr := c.Query("limit")
	title := c.Query("title")
	category := c.Query("category")
	sortBy := c.Query("sort")

	products, err := p.service.GetAllProducts(c.Request.Context(), title, category, limitStr, sortBy)

	if err != nil {
		if err == repository.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p *ProductController) GetAllProductsByCategory(c *gin.Context) {

	category := c.Param("category")
	limitStr := c.Query("limit")
	title := c.Query("title")
	sortBy := c.Query("sort")

	products, err := p.service.GetAllProducts(c.Request.Context(), title, category, limitStr, sortBy)

	if err != nil {
		if err == repository.ErrProductNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
