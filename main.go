// Entry point aplikasi. Load config, inisialisasi database, jalanin migration, setup Fiber, register routes, start server.
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)

	m, err := migrate.New("file://database/migrations", dsn)
	if err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migration: %v", err)
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

	routes.Register(app, cfg)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.App.Port)))
}
