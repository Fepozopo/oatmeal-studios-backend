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
			"message": "Oatmeal Studios Backend API",
		})
	})

	// --- User Routes ---
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List all users",
		})
	})

	router.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Get user by ID",
			"id":      userID,
		})
	})

	router.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Create a new user",
		})
	})

	router.POST("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Update user by ID",
			"id":      userID,
		})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete user by ID",
			"id":      userID,
		})
	})

	// --- Customer Routes ---
	router.GET("/customers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List all customers",
		})
	})

	router.GET("/customers/:id", func(c *gin.Context) {
		customerID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Get customer by ID",
			"id":      customerID,
		})
	})

	router.POST("/customers", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Create a new customer",
		})
	})

	router.PUT("/customers/:id", func(c *gin.Context) {
		customerID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Update customer by ID",
			"id":      customerID,
		})
	})

	router.DELETE("/customers/:id", func(c *gin.Context) {
		customerID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete customer by ID",
			"id":      customerID,
		})
	})

	// --- Order Routes ---
	router.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List all orders",
		})
	})

	router.GET("/orders/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Get order by ID",
			"id":      orderID,
		})
	})

	router.POST("/orders", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "Create a new order",
		})
	})

	router.PUT("/orders/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Update order by ID",
			"id":      orderID,
		})
	})

	router.DELETE("/orders/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete order by ID",
			"id":      orderID,
		})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
