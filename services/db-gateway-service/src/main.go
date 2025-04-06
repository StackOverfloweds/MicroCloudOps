package main

import (
	"log"

	"github.com/StackOverfloweds/db-gateway-services/src/config"
	"github.com/StackOverfloweds/db-gateway-services/src/db/mysql"
	"github.com/StackOverfloweds/db-gateway-services/src/db/postgres"
	rediss "github.com/StackOverfloweds/db-gateway-services/src/db/redis"
	"github.com/StackOverfloweds/db-gateway-services/src/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	mysqlDB, err := mysql.ConnectMySQL()
	if err != nil {
		log.Println("Failed to connect to MySQL:", err)
	} else {
		log.Println("Connected to MySQL")
		defer mysqlDB.Close()
	}

	postgresDB, err := postgres.ConnectPostgres()
	if err != nil {
		log.Println("Failed to connect to PostgreSQL:", err)
	} else {
		log.Println("Connected to PostgreSQL")
		defer postgresDB.Close()
	}

	redisClient, err := rediss.ConnectRedis()
	if err != nil {
		log.Println("Failed to connect to Redis:", err)
	} else {
		log.Println("Connected to Redis")
		defer redisClient.Close()
	}

	app := fiber.New()

	router.SetupRoutes(app)

	log.Println("DB gateway Running at localhost:8080")
	log.Fatal(app.Listen(":8080"))

}
