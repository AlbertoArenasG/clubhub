package db

import (
	"context"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logrus.WithError(err).Error("Failed to create MongoDB client")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to MongoDB")
		return nil, err
	}

	logrus.Info("Connected to MongoDB")

	// Accessing the database to check and create if not exists
	databaseName := os.Getenv("MONGO_DB")
	db := client.Database(databaseName)

	return db, nil
}
