package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliname17/confectionery/internal/models"
	"gorm.io/gorm"
)

type DessertHandler struct {
	DB *gorm.DB
}

func (h *DessertHandler) GetAllDessertsData() []models.Dessert {
	var desserts []models.Dessert
	h.DB.Find(&desserts)
	return desserts
}

// Для API
func (h *DessertHandler) GetAllDesserts(c *gin.Context) {
	var desserts []models.Dessert
	if err := h.DB.Find(&desserts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get desserts"})
		return
	}
	c.JSON(http.StatusOK, desserts)
}
