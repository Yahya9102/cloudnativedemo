package main

import (
	"cloudnativedemo/notification-service/internals/api"
	"cloudnativedemo/notification-service/internals/repository"
	"cloudnativedemo/notification-service/internals/service"
)


func main() {
	
	repo := repository.NewRepository()
	s := service.NewNotifyService(repo)
	api.StartServer(s)

}