package router

import (
	"github.com/AlbertoArenasG/clubhub/bootstrap"
	"github.com/AlbertoArenasG/clubhub/internal/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, apiContainer *bootstrap.ApiContainer) {
	companyHandler := apiContainer.GetCompanyHandler()

	// Health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		logrus.Info("Health check endpoint accessed")
		return c.SendStatus(fiber.StatusOK)
	})

	// Company routes
	companyRouter := app.Group("/companies")
	companyRouter.Get("/", middlewares.ValidateListCompaniesParams, companyHandler.ListCompanies)
	companyRouter.Post("/", middlewares.ValidateCreateCompanyDTO, companyHandler.CreateCompany)
	companyRouter.Put("/:id", middlewares.ValidateUpdateCompanyDTO, companyHandler.UpdateCompany)
}
