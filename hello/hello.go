package hello

import "github.com/gofiber/fiber/v2"

func get(c *fiber.Ctx) error {
	return c.SendString("hello")
}

func RegisterRoutes(router fiber.Router) {
	hello := router.Group("hello")
	hello.Get("/", get)
}
