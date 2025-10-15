package main

import (
	"cloudnativedemo/notification-service/internals/api"
	"cloudnativedemo/notification-service/internals/client"
	"cloudnativedemo/notification-service/internals/service"
)

/*
	TODO från förra lektionen:
	1- Filtrera vuxna användare
	2- Visa notifieringshistorik
	3- Rensa loggar
	4- Skicka notis till en specifik användare

*/


func main() {
	
	// Skapa en client som pekar mot user-service
	userClient := client.NewUserClient("http://localhost:8080")

	// Skapa vår notiferingstjänst
	notifyService := service.NewNotifyService(userClient)

	// Starta vår server
	api.StartServer(notifyService)

}