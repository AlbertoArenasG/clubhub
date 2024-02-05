package handlers

import (
	"github.com/AlbertoArenasG/clubhub/internal/services"
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
