package server

import "github.com/gofiber/fiber/v2"

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).SendString("Hello World!")
}
