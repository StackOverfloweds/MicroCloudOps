package main

import (
	"log"

	"github.com/StackOverfloweds/db-gateway-services/src/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRoutes(app)

	log.Println("DB gateway Running at localhost:8080")
	log.Fatal(app.Listen(":8080"))

}
