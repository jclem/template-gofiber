package meta

import "github.com/gofiber/fiber/v2"

func Healthcheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("ok")
	}
}
