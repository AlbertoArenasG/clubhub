package repositories

import (
	"context"
	"time"

	"github.com/AlbertoArenasG/clubhub/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	companyCollection *mongo.Collection
}

func NewCompanyRepository(companyCollection *mongo.Collection) *CompanyRepository {
	return &CompanyRepository{companyCollection}
}

func (r *CompanyRepository) Create(company *models.Company) (*models.Company, error) {
	company.ID = primitive.NewObjectID()
	company.CreatedAt = time.Now()
	company.UpdatedAt = time.Now()

	_, err := r.companyCollection.InsertOne(context.Background(), company)
	if err != nil {
		return nil, err
	}

	return company, nil
}
