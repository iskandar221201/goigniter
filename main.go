package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/iskandar221201/goigniter/config"
	"github.com/iskandar221201/goigniter/routes"
	"github.com/iskandar221201/goigniter/system"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := system.InitDB(cfg); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: false,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return system.Error(c, code, err.Error())
		},
	})

	routes.Register(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.App.Port)))
}
