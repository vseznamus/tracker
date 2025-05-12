package handlers

import (
	"net/http"
	"tracker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Создание портфеля (или задачи)
func CreatePortfolio(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Portfolio
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&input)
		c.JSON(http.StatusOK, input)
	}
}

// Получение всех портфелей
func ListPortfolios(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var portfolios []models.Portfolio
		db.Find(&portfolios)
		c.JSON(http.StatusOK, portfolios)
	}
}

// Обновление портфеля
func UpdatePortfolio(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var portfolio models.Portfolio
		if err := db.First(&portfolio, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
			return
		}
		if err := c.ShouldBindJSON(&portfolio); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&portfolio)
		c.JSON(http.StatusOK, portfolio)
	}
}

// Удаление портфеля
func DeletePortfolio(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&models.Portfolio{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "Portfolio deleted"})
	}
}
