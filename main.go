package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/imanaspaul/googlemerchent/database"
	"github.com/imanaspaul/googlemerchent/router"
)

func main() {
	app := fiber.New()

	// midlldeware
	app.Use(cors.New())

	// database connection
	if err := database.ConnectDB(); err != nil {
		log.Fatal("Database Connection Error: ", err)
	}

	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
