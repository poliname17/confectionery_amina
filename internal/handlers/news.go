package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/poliname17/confectionery/internal/models"
	"gorm.io/gorm"
)

type NewsHandler struct {
	DB *gorm.DB
}

func (h *NewsHandler) GetAllNewsData() []models.News {
	var news []models.News
	h.DB.Find(&news)
	return news
}

// Для API
func (h *NewsHandler) GetAllNews(c *gin.Context) {
	var news []models.News
	if err := h.DB.Find(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get news"})
		return
	}
	c.JSON(http.StatusOK, news)
}
