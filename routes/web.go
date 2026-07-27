package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/iskandar221201/goigniter/app/controllers"
	"github.com/iskandar221201/goigniter/app/middleware"
	"github.com/iskandar221201/goigniter/app/services"
	"github.com/iskandar221201/goigniter/config"
	"github.com/iskandar221201/goigniter/system"
)

func Register(r chi.Router, cfg *config.Config) {
	authService := services.NewAuthService(system.GetDB(), cfg)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(system.GetDB())
	userController := controllers.NewUserController(userService)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/auth/register", system.Wrap(authController.Register))
		r.Post("/auth/login", system.Wrap(authController.Login))

		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTProtected(cfg.JWT.Secret))

			r.Get("/users", system.Wrap(userController.Index))
			r.Post("/users", system.Wrap(userController.Create))
			r.Get("/users/{id}", system.Wrap(userController.Show))
			r.Put("/users/{id}", system.Wrap(userController.Update))
			r.Delete("/users/{id}", system.Wrap(userController.Delete))
		})
	})
}
