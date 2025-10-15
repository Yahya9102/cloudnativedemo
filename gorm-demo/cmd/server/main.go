package main

import (
	"cloudnativedemo/gorm-demo/internals/api"
	"cloudnativedemo/gorm-demo/internals/repository"
	"cloudnativedemo/gorm-demo/internals/service"
)

/*
	// Bibliotek vi använder
	gin : go get github.com/gin-gonic/gin
	gorm : go get gorm.io/gorm
	mysql driver : go get gorm.io/driver/mysql


*/


func main() {
	repo := repository.NewProductRepository()
	service := service.NewProductService(repo)
	api.StartServer(service)
}