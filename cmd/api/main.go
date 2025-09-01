package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Mini E-Commerce API v1.0.0",
	})

	// Basic health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Mini E-Commerce API is running!",
			"version": "1.0.0",
			"status":  "healthy",
		})
	})

	// Get port from environment or default to 3000
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
