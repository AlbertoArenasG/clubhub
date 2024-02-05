package middlewares

import (
	"net/http"

	"github.com/AlbertoArenasG/clubhub/internal/dtos"
	"github.com/AlbertoArenasG/clubhub/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ValidateCreateCompanyDTO(c *fiber.Ctx) error {
	var createCompanyDTO dtos.PayloadCompanyDTO
	if err := c.BodyParser(&createCompanyDTO); err != nil {
		logrus.WithError(err).Error("Failed to parse request body")
		errorResponse := entity.ErrorResponse{Error: "Invalid request body"}
		return c.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if err := createCompanyDTO.Validate(); err != nil {
		logrus.WithError(err).Error("Validation error")
		errorResponse := entity.ErrorResponse{Error: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	companyData, err := createCompanyDTO.ConvertDTOToModel()
	if err != nil {
		logrus.WithError(err).Error("Failed to convert DTO to model")
		return c.Status(http.StatusBadRequest).JSON(entity.NewErrorResponse("Invalid request body"))
	}

	c.Locals("companyData", companyData)

	return c.Next()
}

func ValidateUpdateCompanyDTO(c *fiber.Ctx) error {
	var updateCompanyDTO dtos.PayloadCompanyDTO
	if err := c.BodyParser(&updateCompanyDTO); err != nil {
		logrus.WithError(err).Error("Failed to parse request body")
		errorResponse := entity.ErrorResponse{Error: "Invalid request body"}
		return c.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	if err := updateCompanyDTO.Validate(); err != nil {
		logrus.WithError(err).Error("Validation error")
		errorResponse := entity.ErrorResponse{Error: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(errorResponse)
	}

	companyData, err := updateCompanyDTO.ConvertDTOToModel()
	if err != nil {
		logrus.WithError(err).Error("Failed to convert DTO to model")
		return c.Status(http.StatusBadRequest).JSON(entity.NewErrorResponse("Invalid request body"))
	}

	c.Locals("companyData", companyData)

	return c.Next()
}
