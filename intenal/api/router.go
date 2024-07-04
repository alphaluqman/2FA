package api

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, authHandler *AuthHandler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Post("/sendcode", authHandler.SendCode)
	app.Post("/verifycode", authHandler.VerifyCode)
}
