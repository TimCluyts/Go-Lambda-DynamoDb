package main

import (
	"goAngularTryout/server"
	"goAngularTryout/server/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVars()
}
func main() {
	// Create app
	app := fiber.New(fiber.Config{})

	// Allow all cors
	app.Use(cors.New())

	// Routing
	app.Get("/api/hello", server.HelloWorld)

	// Start app
	app.Listen(":" + os.Getenv("PORT"))
}
