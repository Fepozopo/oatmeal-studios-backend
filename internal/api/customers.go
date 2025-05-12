package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine) {
	r.GET("/customers", listCustomers)
	r.GET("/customers/:id", getCustomer)
	r.POST("/customers", createCustomer)
	r.PUT("/customers/:id", updateCustomer)
	r.DELETE("/customers/:id", deleteCustomer)
}

func listCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List all customers"})
}

func getCustomer(c *gin.Context) {
	customerID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get customer by ID", "id": customerID})
}

func createCustomer(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Create a new customer"})
}

func updateCustomer(c *gin.Context) {
	customerID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Update customer by ID", "id": customerID})
}

func deleteCustomer(c *gin.Context) {
	customerID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Delete customer by ID", "id": customerID})
}
