package main

import (
	"cloudnativedemo/notification-service/internals/api"
	"cloudnativedemo/notification-service/internals/repository"
	"cloudnativedemo/notification-service/internals/service"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	

		
	if err:= godotenv.Load(); err != nil {
		log.Fatal("Kunde inte läsa .env filen",err)
	}

	dbUser := os.Getenv("DB_USER");
	dbPass := os.Getenv("DB_PASS");
	dbName := os.Getenv("DB_NAME");
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println(dbUser,dbPass,dbName,dbHost,dbPort)


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",dbUser,dbPass,dbHost,dbPort,dbName)




	repo := repository.NewRepository(dsn)
	s := service.NewNotifyService(repo)
	api.StartServer(s)

}