package main

import (
	"runtime/debug"

	"github.com/chirpx/backend/hello"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Commit = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}

	return ""
}()

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString(Commit)
	})

	api := app.Group("api")
	v1 := api.Group("v1")

	hello.RegisterRoutes(v1)

	app.Listen(":3000")
}
