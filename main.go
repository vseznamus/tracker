package main

import (
	"log"
	"tracker/models"

	"tracker/handlers"

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
	err = db.AutoMigrate(&models.User{}, &models.Queue{}, &models.Portfolio{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database connected and migrated successfully!")

	r := gin.Default()
	r.POST("/register", handlers.Register(db))

	r.POST("/queues", handlers.CreateQueue(db))
	r.GET("/queues", handlers.ListQueues(db))
	r.PATCH("/queues/:id", handlers.UpdateQueue(db))
	r.DELETE("/queues/:id", handlers.DeleteQueue(db))

	// Портфели
	r.POST("/portfolios", handlers.CreatePortfolio(db))
	r.GET("/portfolios", handlers.ListPortfolios(db))
	r.PATCH("/portfolios/:id", handlers.UpdatePortfolio(db))
	r.DELETE("/portfolios/:id", handlers.DeletePortfolio(db))

	r.Run(":8080")
}
