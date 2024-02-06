package middlewares

import (
	"net/http"
	"strconv"

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

func ValidateListCompaniesParams(c *fiber.Ctx) error {
	defaultLimit := 10
	maxLimit := 100
	defaultPage := 1
	defaultSort := 1

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < defaultLimit {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < defaultPage {
		page = defaultPage
	}

	search := c.Query("search")

	sort, err := strconv.Atoi(c.Query("sort"))
	if err != nil || (sort != 1 && sort != -1) {
		sort = defaultSort
	}

	c.Locals("limit", limit)
	c.Locals("page", page)
	c.Locals("search", search)
	c.Locals("sort", sort)

	return c.Next()
}
