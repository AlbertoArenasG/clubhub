package services

import (
	"github.com/AlbertoArenasG/clubhub/internal/models"
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

func (s *CompanyService) CreateCompany(company *models.Company) (*models.Company, error) {
	createdCompany, err := s.companyRepository.Create(company)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create company")
		return nil, err
	}

	s.logger.Infof("Company created successfully: %v", createdCompany)
	return createdCompany, nil
}

func (s *CompanyService) UpdateCompany(id string, companyData *models.Company) (*models.Company, error) {
	company, err := s.companyRepository.FindByID(id)
	if err != nil {
		s.logger.WithError(err).Error("Failed to find company")
		return nil, err
	}

	company.Owner = companyData.Owner
	company.Information = companyData.Information
	company.Franchises = companyData.Franchises

	updatedCompany, err := s.companyRepository.Update(id, company)
	if err != nil {
		s.logger.WithError(err).Error("Failed to update company")
		return nil, err
	}

	s.logger.Infof("Company updated successfully: %v", updatedCompany)
	return updatedCompany, nil
}
