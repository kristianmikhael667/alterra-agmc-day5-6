package main

import (
	"log"
	"main/config"
	"main/middleware"
	"main/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	e := routes.New()
	// implementasi middleware logger
	middleware.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
