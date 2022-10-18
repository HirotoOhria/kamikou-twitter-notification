package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"

	_ "hiroto.ohira/kamikou-twitter-notification"
)

const port = "8080"

func init() {
	if err := godotenv.Load("local/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}
}

func main() {
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("functionframework.Start: %v\n", err)
	}
}
