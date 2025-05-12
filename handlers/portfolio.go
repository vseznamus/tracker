package handlers

import (
	"net/http"
	"tracker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Создать портфель (задачу)
func CreatePortfolio(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Portfolio
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, input)
	}
}

// Получить все портфели
func ListPortfolios(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var portfolios []models.Portfolio
		if err := db.Find(&portfolios).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, portfolios)
	}
}

// Изменить портфель
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

// Удалить портфель
func DeletePortfolio(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Portfolio{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Portfolio deleted"})
	}
}
