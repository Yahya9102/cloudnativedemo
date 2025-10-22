package api

import (
	"cloudnativedemo/notification-service/internals/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Startar vå GIN server
func StartServer(NotifyService *service.NotifyService){


	router := gin.Default()


	router.POST("/notifications", func (c *gin.Context) {

		var input struct {
			UserID uint `json:"user_id"`
			Message string `json:"message"`

		}


		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": "Ogiltig json"})
		}

		notis, err := NotifyService.CreateNotification(input.UserID, input.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Kunde inte skapa notis"})
			return
		}

		c.JSON(http.StatusCreated, notis)

	})


	router.GET("/notifications", func(c *gin.Context) {
		
		notifications, err := NotifyService.ListNotifications()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Kunde inte hämta notiserna"})			
		}

		c.JSON(http.StatusOK,notifications)
	})



	router.GET("/notifications/combined", func(c *gin.Context) {
		
		result, err := NotifyService.GetCombinedNotifications()

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Ok(fallback används om user-service är nere)",
			"data": result,
		})
	})

	router.Run(":8081") // Starta server på port 8081
}







