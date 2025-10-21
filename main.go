package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

const candidateName = "Vladimir Avdeev"

func buildApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		body := fmt.Sprintf(
			`{"message":"My name is %s","timestamp":%d}`,
			candidateName,
			time.Now().UTC().UnixMilli(),
		)

		return c.Type("json").SendString(body)
	})

	return app
}

func main() {
	app := buildApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	log.Printf("Starting server on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
