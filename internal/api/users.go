package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", listUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)
}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List all users"})
}

func getUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get user by ID", "id": userID})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create a new user"})
}

func updateUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update user by ID", "id": userID})
}

func deleteUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete user by ID", "id": userID})
}
