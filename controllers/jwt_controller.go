package controllers

import (
	"net/http"

	"github.com/Kei-K23/mock-commerce-api/dto"
	"github.com/Kei-K23/mock-commerce-api/services"
	"github.com/gin-gonic/gin"
)

type JWTController struct {
	service services.JWTService
}

func NewJWTController(service services.JWTService) *JWTController {
	return &JWTController{service: service}
}

func (p *JWTController) CreateJWT(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdJWT, err := p.service.CreateJWT(c.Request.Context(), &loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": createdJWT})
}
