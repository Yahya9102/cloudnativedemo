package main

import (
	"cloudnativedemo/s3-uploader/internals/api"
	"cloudnativedemo/s3-uploader/internals/s3client"
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("S3_BUCKET")



	// Skapa S3-klient
	ctx := context.Background()
	


	s3c, err := s3client.New(ctx,accessKey,secretKey,region,bucket)
	if err != nil {
		log.Fatalf("Kunde inte skapa s3-klient: %v", err)
	}

	api.StartServer(s3c)


}