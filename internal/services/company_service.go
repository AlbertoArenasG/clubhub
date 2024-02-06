package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/AlbertoArenasG/clubhub/external"
	"github.com/AlbertoArenasG/clubhub/internal/models"
	"github.com/AlbertoArenasG/clubhub/internal/repositories"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CompanyService struct {
	companyRepository         *repositories.CompanyRepository
	hotelServerInfoRepository *repositories.HotelServerInfoRepository
	hotelDnsInfoRepository    *repositories.HotelDnsInfoRepository
	logger                    *logrus.Logger
}

func NewCompanyService(
	companyRepository *repositories.CompanyRepository,
	hotelServerInfoRepository *repositories.HotelServerInfoRepository,
	hotelDnsInfoRepository *repositories.HotelDnsInfoRepository,
	logger *logrus.Logger,
) *CompanyService {
	return &CompanyService{
		companyRepository:         companyRepository,
		hotelServerInfoRepository: hotelServerInfoRepository,
		hotelDnsInfoRepository:    hotelDnsInfoRepository,
		logger:                    logger,
	}
}

func (s *CompanyService) ListCompanies(limit, page int, search string, sort int) ([]models.Company, int, int, error) {
	offset := (page - 1) * limit

	companies, totalCompanies, err := s.companyRepository.List(limit, offset, search, sort)
	if err != nil {
		s.logger.WithError(err).Error("Failed to list companies")
		return nil, 0, 0, err
	}

	totalPages := (totalCompanies + limit - 1) / limit

	return companies, totalCompanies, totalPages, nil
}

func (s *CompanyService) CreateCompany(company *models.Company) (*models.Company, error) {
	createdCompany, err := s.companyRepository.Create(company)
	if err != nil {
		s.logger.WithError(err).Error("Failed to create company")
		return nil, err
	}

	go s.GetHotelsInfo(createdCompany.ID, createdCompany.Franchises)

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

	go s.GetHotelsInfo(updatedCompany.ID, updatedCompany.Franchises)

	s.logger.Infof("Company updated successfully: %v", updatedCompany)
	return updatedCompany, nil
}

func (s *CompanyService) GetHotelsInfo(companyID primitive.ObjectID, franchises []models.Franchise) {
	errCh := make(chan error, len(franchises))
	doneCh := make(chan struct{})

	go s.GetWhoisInfoForFranchises(companyID, franchises, errCh)

	go s.GetSSLLabsInfoForFranchises(companyID, franchises, errCh)

	go func() {
		for i := 0; i < 2; i++ {
			select {
			case <-doneCh:
				continue
			case err := <-errCh:
				if err != nil {
					s.logger.WithError(err).Error("Failed to get hotel info")
				}
			}
		}
	}()

	time.Sleep(time.Second * 10)

	close(doneCh)
}

func (s *CompanyService) GetWhoisInfoForFranchises(companyID primitive.ObjectID, franchises []models.Franchise, errCh chan<- error) {
	whoisClient := external.NewWhoisClient()

	for _, franchise := range franchises {
		go func(franchise models.Franchise) {
			const maxRetries = 3
			for i := 0; i < maxRetries; i++ {
				cleanedURL := strings.TrimPrefix(strings.TrimPrefix(franchise.URL, "http://"), "https://")
				cleanedURL = strings.TrimPrefix(cleanedURL, "www.")

				domainInfo, err := whoisClient.GetWhoisInfo(cleanedURL)
				if err != nil {
					s.logger.WithError(err).Errorf("Failed to get whois information for URL %s, retrying", cleanedURL)
					time.Sleep(time.Second * 10)
					continue
				}
				whoisData, err := external.ParseWhoisData(domainInfo)
				if err != nil {
					s.logger.WithError(err).Errorf("Failed to parse whois information for URL %s", cleanedURL)
					errCh <- fmt.Errorf("failed to parse whois information for URL %s: %v", cleanedURL, err)
					return
				}

				dnsInfo := models.MapWhoisDataToDnsInfo(companyID, cleanedURL, &whoisData)
				_, err = s.hotelDnsInfoRepository.Create(dnsInfo)
				if err != nil {
					s.logger.WithError(err).Errorf("Failed to save DNS info to repository for URL %s", cleanedURL)
					errCh <- fmt.Errorf("failed to save DNS info to repository for URL %s: %v", cleanedURL, err)
					return
				}
				time.Sleep(time.Second * 10)
				return
			}
			errCh <- fmt.Errorf("maximum number of retries exceeded for URL %s", franchise.URL)
		}(franchise)
	}
}

func (s *CompanyService) GetSSLLabsInfoForFranchises(companyID primitive.ObjectID, franchises []models.Franchise, errCh chan<- error) {
	for _, franchise := range franchises {
		go func(franchise models.Franchise) {
			const maxRetries = 3
			for i := 0; i < maxRetries; i++ {
				cleanedURL := strings.TrimPrefix(strings.TrimPrefix(franchise.URL, "http://"), "https://")
				sslClient := external.NewSSLLabsClient()
				result, err := sslClient.GetSSLInfo(cleanedURL)
				if err != nil {
					s.logger.WithError(err).Errorf("Failed to fetch SSL Labs info for URL %s, retrying", franchise.URL)
					time.Sleep(time.Second * 15)
					continue
				}

				serverInfo := models.MapSSLLabsResultToServerInfo(companyID, result)
				_, err = s.hotelServerInfoRepository.Create(serverInfo)
				if err != nil {
					s.logger.WithError(err).Errorf("Failed to save server info to repository for URL %s", franchise.URL)
					errCh <- fmt.Errorf("failed to save server info to repository for URL %s: %v", franchise.URL, err)
					return
				}
				time.Sleep(time.Second * 10)
				return
			}
			errCh <- fmt.Errorf("maximum number of retries exceeded for URL %s", franchise.URL)
		}(franchise)
	}
}
