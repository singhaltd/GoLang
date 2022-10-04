package main

import (
	"log"

	"api.shsoft.com/m/v2/database"
	"api.shsoft.com/m/v2/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	route.SetupRoutes(app)
	log.Fatal(app.Listen(":3300"))
}
