package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliname17/confectionery/internal/models"
	"gorm.io/gorm"
)

type ChefHandler struct {
	DB *gorm.DB
}

func (h *ChefHandler) GetChefData() models.Chef {
	var chef models.Chef
	h.DB.First(&chef)
	return chef
}

// Для API
func (h *ChefHandler) GetChef(c *gin.Context) {
	var chef []models.Chef
	if err := h.DB.Find(&chef).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get news"})
		return
	}
	c.JSON(http.StatusOK, chef)
}
