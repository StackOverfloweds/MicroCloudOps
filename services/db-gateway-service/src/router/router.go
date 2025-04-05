package router

import (
	"github.com/StackOverfloweds/db-gateway-services/src/db/mysql"
	"github.com/StackOverfloweds/db-gateway-services/src/db/postgres"
	rediss "github.com/StackOverfloweds/db-gateway-services/src/db/redis"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("DB Gateway Service is Healty")
	})

	app.Get("/mysql", func(c *fiber.Ctx) error {
		dbConn, err := mysql.ConnectMySQL()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to MySQL: " + err.Error())
		}
		defer dbConn.Close()
		return c.SendString("Successfully connected to MySQL!")
	})

	app.Get("/postgres", func(c *fiber.Ctx) error {
		dbConn, err := postgres.ConnectPostgres()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to PostgreSQL: " + err.Error())
		}
		defer dbConn.Close()
		return c.SendString("Successfully connected to PostgreSQL!")
	})

	app.Get("/redis", func(c *fiber.Ctx) error {
		client, err := rediss.ConnectRedis()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to connect to Redis: " + err.Error())
		}
		defer client.Close()
		return c.SendString("Successfully connected to Redis")
	})
}
