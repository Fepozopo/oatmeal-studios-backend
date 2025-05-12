package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/api"
)

func main() {
	router := gin.Default()

	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Oatmeal Studios Backend API"})
	})

	// Register resource routes
	api.RegisterUserRoutes(router)
	api.RegisterCustomerRoutes(router)
	api.RegisterOrderRoutes(router)

	// Run the server on port 8080
	router.Run(":8080")
}
