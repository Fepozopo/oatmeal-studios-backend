package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Define a handler function for the root path ("/")
	router.GET("/", func(c *gin.Context) {
		// Respond with JSON
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
