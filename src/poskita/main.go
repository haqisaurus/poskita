package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/haqisaurus/poskita/router"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	var errEnv error

	if len(args) == 2 {
		errEnv = godotenv.Load(args[1])
	} else {
		errEnv = godotenv.Load()
	}

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	router.LoggingRoute((app))
	router.SetupRouter(app)

	app.Listen(":3000")
}
