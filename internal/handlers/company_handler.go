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

func (ch *CompanyHandler) ListCompanies(c *fiber.Ctx) error {
	limit := c.Locals("limit").(int)
	page := c.Locals("page").(int)
	sort := c.Locals("sort").(int)
	search := c.Locals("search").(string)

	companies, totalCompanies, totalPages, err := ch.companyService.ListCompanies(limit, page, search, sort)
	if err != nil {
		ch.logger.WithError(err).Error("Failed to list companies")
		return c.Status(http.StatusInternalServerError).JSON(entity.NewErrorResponse("Failed to list companies"))
	}

	response := entity.NewSuccessListResponse("Companies listed successfully", companies, map[string]interface{}{
		"current_page":  page,
		"page_size":     limit,
		"total_records": totalCompanies,
		"total_pages":   totalPages,
	})

	return c.Status(http.StatusOK).JSON(response)
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

func (ch *CompanyHandler) UpdateCompany(c *fiber.Ctx) error {
	id := c.Params("id")

	companyData, ok := c.Locals("companyData").(models.Company)
	if !ok {
		ch.logger.Error("Missing update company data in context")
		errorResponse := entity.ErrorResponse{Error: "Invalid request body"}
		return c.Status(http.StatusInternalServerError).JSON(errorResponse)
	}

	updatedCompany, err := ch.companyService.UpdateCompany(id, &companyData)
	if err != nil {
		ch.logger.WithError(err).Error("Failed to update company")
		if err.Error() == "company not found" {
			return c.Status(http.StatusNotFound).JSON(entity.NewErrorResponse(err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(entity.NewErrorResponse("Failed to update company"))
	}

	return c.Status(http.StatusOK).JSON(entity.NewSuccessResponse("Company updated successfully", updatedCompany))
}
