package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/imanaspaul/googlemerchent/handler"
)

func SetupRouter(app *fiber.App) {
	// middleware
	api := app.Group("/api", logger.New())

	// routes
	api.Get("/", handler.Hello)
}
