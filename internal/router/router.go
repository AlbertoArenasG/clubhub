package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App) {
	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		logrus.Info("Health check endpoint accessed")
		return c.SendStatus(fiber.StatusOK)
	})
}
