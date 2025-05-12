package main

import (
	"log"
	"tracker/handlers"
	"tracker/models"

	"github.com/gin-gonic/gin"
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
