package main

import (
	"log"
	"tracker/handlers"
	"tracker/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=99#dfgDD56678 dbname=trackerdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.User{}, &models.Queue{}, &models.Portfolio{})

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/users_page", UsersPage(db))
	r.POST("/users_page", UsersPage(db))

	// Пользователи
	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db))
	r.GET("/users", handlers.ListUsers(db))
	r.PATCH("/users/:id", handlers.AuthMiddleware(), handlers.UpdateUser(db))
	r.DELETE("/users/:id", handlers.AuthMiddleware(), handlers.DeleteUser(db))

	// Очереди
	r.POST("/queues", handlers.AuthMiddleware(), handlers.CreateQueue(db))
	r.GET("/queues", handlers.ListQueues(db))
	r.PATCH("/queues/:id", handlers.AuthMiddleware(), handlers.UpdateQueue(db))
	r.DELETE("/queues/:id", handlers.AuthMiddleware(), handlers.DeleteQueue(db))

	// Портфели
	r.POST("/portfolios", handlers.AuthMiddleware(), handlers.CreatePortfolio(db))
	r.GET("/portfolios", handlers.ListPortfolios(db))
	r.PATCH("/portfolios/:id", handlers.AuthMiddleware(), handlers.UpdatePortfolio(db))
	r.DELETE("/portfolios/:id", handlers.AuthMiddleware(), handlers.DeletePortfolio(db))

	r.Run(":8080")
}

func UsersPage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		db.Find(&users)

		if c.Request.Method == "POST" {
			login := c.PostForm("login")
			password := c.PostForm("password")
			fullname := c.PostForm("fullname")
			role := c.PostForm("role")
			birthdate := c.PostForm("birthdate")
			phonenumber := c.PostForm("phonenumber")

			hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			user := models.User{
				Login:          login,
				FullName:       fullname,
				Role:           role,
				BirthDate:      birthdate,
				PhoneNumber:    phonenumber,
				ActivityStatus: "active",
				PasswordHash:   string(hash),
			}
			if err := db.Create(&user).Error; err != nil {
				db.Find(&users)
				c.HTML(200, "users.html", gin.H{
					"Users": users,
					"Error": "Ошибка при создании пользователя: " + err.Error(),
				})
				return
			}
			db.Find(&users)
			c.HTML(200, "users.html", gin.H{
				"Users":   users,
				"Message": "Пользователь успешно добавлен!",
			})
			return
		}

		c.HTML(200, "users.html", gin.H{
			"Users": users,
		})
	}
}
