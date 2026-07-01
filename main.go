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
	// API ของสมาชิก mind
	app.Get("/mind", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือสมาชิก mind")
	})

	// API ของสมาชิก beam
	app.Get("/beam", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือสมาชิก beam")
})
	app.Get("/nam", func(c *fiber.Ctx) error {
		return c.SendString("นี่คือสมาชิก nam")
>>>>>>> origin/dev-nam
	})

	app.Listen(":3000")
}
