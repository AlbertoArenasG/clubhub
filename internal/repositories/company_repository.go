package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/AlbertoArenasG/clubhub/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CustomErrorNotFound = errors.New("company not found")

type CompanyRepository struct {
	companyCollection *mongo.Collection
}

func NewCompanyRepository(companyCollection *mongo.Collection) *CompanyRepository {
	return &CompanyRepository{companyCollection}
}

func (r *CompanyRepository) List(limit, skip int, search string, sort int) ([]models.Company, int, error) {
	filter := bson.M{}
	if search != "" {
		filter["$or"] = []bson.M{
			{"information.name": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "i"}}},
			{"franchises.name": bson.M{"$regex": primitive.Regex{Pattern: search, Options: "i"}}},
		}
	}

	options := options.Find().
		SetSort(bson.M{"created_at": sort}).
		SetSkip(int64(skip)).
		SetLimit(int64(limit))

	cursor, err := r.companyCollection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var companies []models.Company
	for cursor.Next(context.Background()) {
		var company models.Company
		if err := cursor.Decode(&company); err != nil {
			return nil, 0, err
		}
		companies = append(companies, company)
	}

	if companies == nil {
		companies = []models.Company{}
		return companies, 0, nil
	}

	totalCount, err := r.companyCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, 0, err
	}

	return companies, int(totalCount), nil
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
