package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iskandar221201/goigniter/app/controllers"
	"github.com/iskandar221201/goigniter/app/services"
	"github.com/iskandar221201/goigniter/config"
	"github.com/iskandar221201/goigniter/system"
)

func Register(app *fiber.App, cfg *config.Config) {
	authService := services.NewAuthService(system.GetDB(), cfg)
	authController := controllers.NewAuthController(authService)

	api := app.Group("/api/v1")
	api.Post("/auth/register", authController.Register)
	api.Post("/auth/login", authController.Login)
}
