package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliname17/confectionery/internal/models"
	"gorm.io/gorm"
)

type ReviewHandler struct {
	DB *gorm.DB
}

func (h *ReviewHandler) GetAllReviewsData() []models.Review {
	var reviews []models.Review
	h.DB.Find(&reviews)
	return reviews
}

// Для API
func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	var reviews []models.Review
	if err := h.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}
