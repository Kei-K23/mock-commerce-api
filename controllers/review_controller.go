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

type ReviewController struct {
	service services.ReviewService
}

func NewReviewController(service services.ReviewService) *ReviewController {
	return &ReviewController{service: service}
}

func (p *ReviewController) CreateReview(c *gin.Context) {
	var reviewRequest dto.ReviewRequest
	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := mapper.MatchReviewRequestToReview(reviewRequest)

	createdReview, err := p.service.CreateReview(c.Request.Context(), review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdReview)
}

func (p *ReviewController) UpdateReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reviewRequest dto.ReviewRequest
	if err := c.ShouldBindJSON(&reviewRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := mapper.MatchReviewRequestToReview(reviewRequest)

	createdReview, err := p.service.UpdateReview(c.Request.Context(), id, review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdReview)
}

func (p *ReviewController) DeleteReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedId, err := p.service.DeleteReview(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedId})
}

func (p *ReviewController) GetReviewById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := p.service.GetReviewById(c.Request.Context(), id)
	if err != nil {
		if err == repository.ErrReviewNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

func (p *ReviewController) GetAllReviews(c *gin.Context) {

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productId, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	limitStr := c.Query("limit")
	sortBy := c.Query("sort")

	reviews, err := p.service.GetAllReviews(c.Request.Context(), userId, productId, limitStr, sortBy)

	if err != nil {
		if err == repository.ErrReviewNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
