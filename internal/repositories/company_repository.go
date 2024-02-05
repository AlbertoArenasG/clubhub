package repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	companyCollection *mongo.Collection
}

func NewCompanyRepository(companyCollection *mongo.Collection) *CompanyRepository {
	return &CompanyRepository{companyCollection}
}
