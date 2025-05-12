package handlers

import (
	"net/http"
	"tracker/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Login       string
			Password    string
			FullName    string
			Role        string
			BirthDate   string
			PhoneNumber string
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		user := models.User{
			Login:          input.Login,
			FullName:       input.FullName,
			Role:           input.Role,
			BirthDate:      input.BirthDate,
			PhoneNumber:    input.PhoneNumber,
			ActivityStatus: "active",
			PasswordHash:   string(hash),
		}
		db.Create(&user)
		c.JSON(http.StatusOK, gin.H{"message": "User registered"})
	}
}
