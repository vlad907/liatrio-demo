package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

const candidateName = "Vladimir Avdeev"

func buildApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		payload := fiber.Map{
			"message":   "My name is " + candidateName,
			"timestamp": time.Now().UTC().Unix(),
		}

		return c.JSON(payload)
	})

	return app
}

func main() {
	app := buildApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
