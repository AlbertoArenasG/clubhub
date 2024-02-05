package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/AlbertoArenasG/clubhub/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var CustomErrorNotFound = errors.New("company not found")

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

func (r *CompanyRepository) FindByID(id string) (*models.Company, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var company models.Company
	err = r.companyCollection.FindOne(context.Background(), primitive.M{"_id": objectID}).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, CustomErrorNotFound
		}
		return nil, err
	}

	return &company, nil
}

func (r *CompanyRepository) Update(id string, company *models.Company) (*models.Company, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := primitive.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"owner":       company.Owner,
			"information": company.Information,
			"franchises":  company.Franchises,
			"updated_at":  time.Now(),
		},
	}

	result, err := r.companyCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, CustomErrorNotFound
	}

	company.ID = objectID
	return company, nil
}
