package api // API-lagret

import (
	"cloudnativedemo/gorm-demo/internals/service" // importerar service-lagret
	"fmt"
	"net/http" // statuskoder

	"github.com/gin-gonic/gin" // webbramverk
)

// StartServer startar Gin-servern
func StartServer(productService *service.ProductService) {
	r := gin.Default() // standardserver med logger & recovery

	// CREATE – POST /products
	r.POST("/products", func(c *gin.Context) {
		var input struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ogiltig JSON"})
			return
		}
		product := productService.CreateProduct(input.Name, input.Price)
		c.JSON(http.StatusCreated, product)
	})

	// READ – GET /products
	r.GET("/products", func(c *gin.Context) {
		products := productService.ListProducts()
		c.JSON(http.StatusOK, products)
	})

	// UPDATE – PUT /products/:id
	r.PUT("/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		var id uint
		fmt.Sscanf(idStr, "%d", &id)

		var input struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ogiltig JSON"})
			return
		}
		product, ok := productService.UpdateProduct(id, input.Name, input.Price)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produkten hittades inte"})
			return
		}
		c.JSON(http.StatusOK, product)
	})

	// DELETE – DELETE /products/:id
	r.DELETE("/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		var id uint
		fmt.Sscanf(idStr, "%d", &id)
		if !productService.DeleteProduct(id) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produkten hittades inte"})
			return
		}
		c.Status(http.StatusNoContent)
	})

	// Starta servern på port 8082
	r.Run(":8082")
}
