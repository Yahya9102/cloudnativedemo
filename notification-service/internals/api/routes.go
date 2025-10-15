package api

import (
	"cloudnativedemo/notification-service/internals/service"
	"fmt"
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
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})




	router.POST("/notify/send", func(c *gin.Context) {
		var input struct {
			Message string `json:"message"` // Vi förväntar oss ett fält som heter message i JSON body
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ogiltig json"})
		}



		logs := NotifyService.SendNotification(input.Message)

		c.JSON(http.StatusOK, gin.H{
			"status": "notiser skickade",
			"logs": logs,
		
		})
	})


	router.GET("/notify/adults", func(c *gin.Context) {
		adults, err := NotifyService.FetchAdultUsers()
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, adults)
	})


	router.GET("/notify/logs", func(c *gin.Context) {
		logs := NotifyService.GetLogs()
		c.JSON(http.StatusOK, gin.H{"log:": logs})
	})


	router.DELETE("/notify/clear", func(c *gin.Context) {
		NotifyService.ClearLogs()
		c.JSON(http.StatusOK, gin.H{"Status": "Loggar rensade"})
	})




	router.POST("/notify/user/:id", func(c *gin.Context) {

		idString := c.Param("id")
		var id int
		fmt.Sscanf(idString, "%d", &id)


		var input struct {
			Message string `json:"message"` // Vi förväntar oss ett fält som heter message i JSON body
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ogiltig JSON"})
			return
		}

		logMsg , err := NotifyService.SendNotificationToUser(id, input.Message)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "notis skickad",
			"log": logMsg,
		})

	})

	router.Run(":8081") // Starta server på port 8081
}