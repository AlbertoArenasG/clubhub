package bootstrap

import (
	"github.com/AlbertoArenasG/clubhub/internal/handlers"
	"github.com/AlbertoArenasG/clubhub/internal/repositories"
	"github.com/AlbertoArenasG/clubhub/internal/services"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiContainer struct {
	companyHandler *handlers.CompanyHandler
}

func NewApiContainer(
	companyCollection *mongo.Collection,
	logger *logrus.Logger,
) *ApiContainer {
	companyRepository := repositories.NewCompanyRepository(companyCollection)
	companyService := services.NewCompanyService(companyRepository, logger)

	companyHandler := handlers.NewCompanyHandler(companyService, logger)

	return &ApiContainer{
		companyHandler: companyHandler,
	}
}

func (container *ApiContainer) GetCompanyHandler() *handlers.CompanyHandler {
	return container.companyHandler
}
