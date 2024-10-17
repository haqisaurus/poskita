package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/haqisaurus/poskita/router"
	"github.com/joho/godotenv"
	"github.com/gofiber/template/html/v2"
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
	engine := html.New("./", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		
	})
	app.Static("/", "./")
	router.LoggingRoute((app))
	router.SetupRouter(app)

	app.Listen(":3001")
}
