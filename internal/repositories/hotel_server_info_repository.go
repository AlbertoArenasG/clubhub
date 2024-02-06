package repositories

import (
	"context"
	"time"

	"github.com/AlbertoArenasG/clubhub/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelServerInfoRepository struct {
	hotelServerInfoCollection *mongo.Collection
}

func NewHotelServerInfoRepository(hotelServerInfoCollection *mongo.Collection) *HotelServerInfoRepository {
	return &HotelServerInfoRepository{hotelServerInfoCollection}
}

func (r *HotelServerInfoRepository) Create(info *models.HotelServerInfo) (*models.HotelServerInfo, error) {
	info.ID = primitive.NewObjectID()
	info.CreatedAt = time.Now()
	info.UpdatedAt = time.Now()
	_, err := r.hotelServerInfoCollection.InsertOne(context.Background(), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
