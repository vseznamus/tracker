package handlers

import (
	"net/http"
	"tracker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Создать очередь
func CreateQueue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Queue
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

// Получить все очереди
func ListQueues(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var queues []models.Queue
		if err := db.Find(&queues).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, queues)
	}
}

// Изменить очередь
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

// Удалить очередь
func DeleteQueue(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := db.Delete(&models.Queue{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Queue deleted"})
	}
}
