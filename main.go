package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"product-api/database"
	"product-api/models"
)

func main() {
	// Connect to database
	database.Connect()

	// Auto migrate
	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()

	// Root route
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello there, Welcome to My Store!")
	})

	// Products route
	r.GET("/products", func(c *gin.Context) {
		var products []models.Product
		database.DB.Find(&products)
		c.JSON(http.StatusOK, products)
	})

	r.Run(":8080")
}
