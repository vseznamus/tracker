package handlers

import (
	"net/http"
	"tracker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Создание очереди
func CreateQueue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Queue
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&input)
		c.JSON(http.StatusOK, input)
	}
}

// Получение всех очередей
func ListQueues(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var queues []models.Queue
		db.Find(&queues)
		c.JSON(http.StatusOK, queues)
	}
}

// Обновление очереди
func UpdateQueue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var queue models.Queue
		if err := db.First(&queue, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Queue not found"})
			return
		}
		if err := c.ShouldBindJSON(&queue); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&queue)
		c.JSON(http.StatusOK, queue)
	}
}

// Удаление очереди
func DeleteQueue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&models.Queue{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "Queue deleted"})
	}
}
