package services

import (
	"github.com/AlbertoArenasG/clubhub/internal/repositories"
	"github.com/sirupsen/logrus"
)

type CompanyService struct {
	companyRepository *repositories.CompanyRepository
	logger            *logrus.Logger
}

func NewCompanyService(
	companyRepository *repositories.CompanyRepository,
	logger *logrus.Logger,
) *CompanyService {
	return &CompanyService{
		companyRepository: companyRepository,
		logger:            logger,
	}
}
