package main

import (
	"mini-ecommerce/config"
	"mini-ecommerce/internal/interfaces/http/routes"
	"mini-ecommerce/pkg/logger"
	"mini-ecommerce/pkg/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	validation.InitValidator()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err, "[ErrMain-1]Failed to load config")
	}
	app := fiber.New()

	app.Use(recover.New())

	db, err := config.Connect(cfg)
	if err != nil {
		logger.Fatal(err, "[ErrMain-2]Failed to connect to database")
	}
	defer db.Close()
	routes.SetupRoutes(app, db.DB)

	addr := cfg.Server.Host + ":" + cfg.Server.Port
	logger.Info("Starting server on: " + addr)

	if err := app.Listen(addr); err != nil {
		logger.Fatal(err, "[ErrMain-3]Failed to start server")
	}
}
