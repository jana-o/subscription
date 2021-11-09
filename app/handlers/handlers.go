package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jana-o/subscription/app/models"
	"net/http"
)

// Routes creates and handles API routes
func Routes(router *gin.Engine) {
	router.GET("/products", GetProducts)
	router.GET("/products/:id", GetProductByID)
}

// GetProducts responds with the list of all Products as JSON
func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Products)
}

// GetProductByID responds with a single Product by ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, p := range models.Products {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}