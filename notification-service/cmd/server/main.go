package main

import (
	"cloudnativedemo/notification-service/internals/api"
	"cloudnativedemo/notification-service/internals/client"
	"cloudnativedemo/notification-service/internals/service"
)

func main() {
	
	// Skapa en client som pekar mot user-service
	userClient := client.NewUserClient("http://localhost:8080")

	// Skapa vår notiferingstjänst
	notifyService := service.NewNotifyService(userClient)

	// Starta vår server
	api.StartServer(notifyService)

}