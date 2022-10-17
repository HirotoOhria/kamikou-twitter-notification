package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/joho/godotenv"

	_ "hiroto.ohira/kamikou-twitter-notification/functoin"
)

const port = "8080"

func init() {
	if err := godotenv.Load("local/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	fmt.Println(os.Getenv("FUNCTION_TARGET"))
}

func main() {
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("functionframework.Start: %v\n", err)
	}
}
