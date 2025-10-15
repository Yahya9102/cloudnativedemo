package main // startpunkt

import (
	"cloudnativedemo/gorm-demo/internals/api"
	"cloudnativedemo/gorm-demo/internals/repository"
	"cloudnativedemo/gorm-demo/internals/service"
)

func main() {
	// Initiera repository (ansluter till MySQL)
	repo := repository.NewProductRepository()

	// Initiera service
	service := service.NewProductService(repo)

	// Starta Gin-server
	api.StartServer(service)
}
