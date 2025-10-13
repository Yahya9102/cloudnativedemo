package api

import (
	"cloudnativedemo/notification-service/internals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Startar vå GIN server
func StartServer(NotifyService *service.NotifyService){


	router := gin.Default()

	// Get anrop mot notify/users	
	router.GET("/notify/users", func(c *gin.Context) {
		
		users, err := NotifyService.FetchUsers()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	router.Run(":8081") // Starta server på port 8081
}