package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/paincake00/order-service/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	application := app.New(app.LoadConfig())

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
