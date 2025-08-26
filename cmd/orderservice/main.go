package main

import (
	"github.com/joho/godotenv"
	"github.com/paincake00/order-service/internal/app"
	"github.com/paincake00/order-service/internal/logs"
)

func main() {
	logger := logs.NewLogger()
	defer func() {
		_ = logger.Sync()
	}()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	config := app.LoadConfig()

	application := app.New(config, logger)

	err = application.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
