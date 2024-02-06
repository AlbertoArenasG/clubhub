package repositories

import (
	"context"

	"github.com/AlbertoArenasG/clubhub/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelDnsInfoRepository struct {
	hotelDnsInfoCollection *mongo.Collection
}

func NewHotelDnsInfoRepository(hotelDnsInfoCollection *mongo.Collection) *HotelDnsInfoRepository {
	return &HotelDnsInfoRepository{hotelDnsInfoCollection}
}

func (r *HotelDnsInfoRepository) Create(info *models.HotelDnsInfo) (*models.HotelDnsInfo, error) {
	info.ID = primitive.NewObjectID()
	_, err := r.hotelDnsInfoCollection.InsertOne(context.Background(), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
