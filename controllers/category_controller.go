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

type CategoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (p *CategoryController) CreateCategory(c *gin.Context) {
	var categoryRequest dto.CategoryRequest
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := mapper.MatchCategoryRequestToCategory(categoryRequest)

	createdCategory, err := p.service.CreateCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func (p *CategoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var categoryRequest dto.CategoryRequest
	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := mapper.MatchCategoryRequestToCategory(categoryRequest)

	updatedCategory, err := p.service.UpdateCategory(c.Request.Context(), id, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (p *CategoryController) GetCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := p.service.GetCategoryById(c.Request.Context(), id)
	if err != nil {
		if err == repository.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *CategoryController) GetAllProducts(c *gin.Context) {

	limitStr := c.Query("limit")
	title := c.Query("title")
	sortBy := c.Query("sort")

	products, err := p.service.GetAllCategories(c.Request.Context(), title, limitStr, sortBy)

	if err != nil {
		if err == repository.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (p *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedId, err := p.service.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedId})
}
