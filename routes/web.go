// Package routes mendefinisikan semua endpoint API di satu tempat, mirip routing CI4.
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iskandar221201/goigniter/app/controllers"
	"github.com/iskandar221201/goigniter/app/middleware"
	"github.com/iskandar221201/goigniter/app/services"
	"github.com/iskandar221201/goigniter/config"
	"github.com/iskandar221201/goigniter/system"
)

func Register(app *fiber.App, cfg *config.Config) {
	authService := services.NewAuthService(system.GetDB(), cfg)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(system.GetDB())
	userController := controllers.NewUserController(userService)

	api := app.Group("/api/v1")

	api.Post("/auth/register", authController.Register)
	api.Post("/auth/login", authController.Login)

	protected := api.Group("/", middleware.JWTProtected(cfg.JWT.Secret))
	protected.Get("/users", userController.Index)
	protected.Post("/users", userController.Create)
	protected.Get("/users/:id", userController.Show)
	protected.Put("/users/:id", userController.Update)
	protected.Delete("/users/:id", userController.Delete)
}
