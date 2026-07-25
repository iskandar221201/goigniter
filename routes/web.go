package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	api := app.Group("/api/v1")
	_ = api
}
