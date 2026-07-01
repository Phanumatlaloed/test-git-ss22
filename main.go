package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// API ของหัวหน้า
	app.Get("/head", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือหัวหน้า")
	})
	app.Get("/name", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือสมาชิก")
	})

	app.Listen(":3000")
}
