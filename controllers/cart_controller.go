package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kei-K23/go-ecommerce-api/dto"
	"github.com/Kei-K23/go-ecommerce-api/mapper"
	"github.com/Kei-K23/go-ecommerce-api/repository"
	"github.com/Kei-K23/go-ecommerce-api/services"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	service services.CartService
}

func NewCartController(service services.CartService) *CartController {
	return &CartController{service: service}
}

func (p *CartController) CreateCart(c *gin.Context) {
	var cartRequest dto.CartRequest
	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart := mapper.MatchCartRequestToCart(cartRequest)

	createdCart, err := p.service.CreateCart(c.Request.Context(), cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdCart)
}

func (p *CartController) UpdateCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cartRequest dto.CartRequest
	if err := c.ShouldBindJSON(&cartRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart := mapper.MatchCartRequestToCart(cartRequest)

	createdCart, err := p.service.UpdateCart(c.Request.Context(), id, cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdCart)
}

func (p *CartController) DeleteCart(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedId, err := p.service.DeleteCart(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedId})
}

func (p *CartController) GetCartById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart, err := p.service.GetCartById(c.Request.Context(), id)
	if err != nil {
		if err == repository.ErrCartNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (p *CartController) GetAllCarts(c *gin.Context) {

	limitStr := c.Query("limit")
	sortBy := c.Query("sort")
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carts, err := p.service.GetAllCarts(c.Request.Context(), userId, limitStr, sortBy)

	if err != nil {
		if err == repository.ErrCartNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, carts)
}
