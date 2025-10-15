package api

import (
	"cloudnativedemo/gorm-demo/internals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Startserver startar API servern med GIN
func StartServer(ProductService *service.ProductService) {

	r := gin.Default()


	r.POST("/products", func(c *gin.Context) {
		
		// Struct för att binda JSON input från request body
		var input struct {
			Name  string `json:"name"`
			Price int    `json:"price"`
		}


		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ogiltig JSON"})
			return
		}

		product := ProductService.CreateProduct(input.Name, input.Price) // Skapar ny produkt via service
		c.JSON(http.StatusCreated, product)

	})


	r.GET("/products", func(c *gin.Context) {
		products := ProductService.ListProducts() // Hämtar alla produkter via service
		c.JSON(http.StatusOK, products) // Returnerar JSON array med alla produkter + 200 OK (HTTP status code
	})


	r.Run(":8082")
}