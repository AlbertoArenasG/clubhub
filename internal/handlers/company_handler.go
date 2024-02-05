package handlers

import (
	"net/http"

	"github.com/AlbertoArenasG/clubhub/internal/entity"
	"github.com/AlbertoArenasG/clubhub/internal/models"
	"github.com/AlbertoArenasG/clubhub/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CompanyHandler struct {
	companyService *services.CompanyService
	logger         *logrus.Logger
}

func NewCompanyHandler(companyService *services.CompanyService, logger *logrus.Logger) *CompanyHandler {
	return &CompanyHandler{
		companyService: companyService,
		logger:         logger,
	}
}

func (ch *CompanyHandler) CreateCompany(c *fiber.Ctx) error {
	companyData, ok := c.Locals("companyData").(models.Company)
	if !ok {
		ch.logger.Error("Missing company data in context")
		errorResponse := entity.ErrorResponse{Error: "Invalid request body"}
		return c.Status(http.StatusInternalServerError).JSON(errorResponse)
	}

	createdCompany, err := ch.companyService.CreateCompany(&companyData)
	if err != nil {
		ch.logger.WithError(err).Error("Failed to create company")
		return c.Status(http.StatusInternalServerError).JSON(entity.NewErrorResponse("Failed to create company"))
	}

	return c.Status(http.StatusCreated).JSON(entity.NewSuccessResponse("Company created successfully", createdCompany))
}
