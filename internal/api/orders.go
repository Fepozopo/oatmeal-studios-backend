package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {
	r.GET("/orders", listOrders)
	r.GET("/orders/:id", getOrder)
	r.POST("/orders", createOrder)
	r.PUT("/orders/:id", updateOrder)
	r.DELETE("/orders/:id", deleteOrder)
}

func listOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List all orders"})
}

func getOrder(c *gin.Context) {
	orderID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get order by ID", "id": orderID})
}

func createOrder(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create a new order"})
}

func updateOrder(c *gin.Context) {
	orderID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update order by ID", "id": orderID})
}

func deleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete order by ID", "id": orderID})
}
