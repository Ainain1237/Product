package main

import (
	"log"

	"github.com/Ainain1237/Product/cache"
	"github.com/Ainain1237/Product/db"
	"github.com/Ainain1237/Product/handlers"
	"github.com/Ainain1237/Product/queue"
	"github.com/Ainain1237/Product/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize dependencies
	db.InitDB()
	cache.InitRedis()
	queue.InitQueue()

	// Create a Gin router
	router := gin.Default()

	// Define routes
	router.POST("/upload", handlers.UploadImageHandler)
	router.POST("/products", handlers.CreateProduct)
	router.GET("/products/:id", handlers.GetProductByID)
	router.GET("/products", handlers.GetProductsHandler)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Product Management!",
		})
	})

	// Start background services
	go services.StartImageProcessor(queue.Channel)

	// Start the server
	log.Fatal(router.Run(":8080"))
}
