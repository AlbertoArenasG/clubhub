package main

import (
	"context"

	"github.com/AlbertoArenasG/clubhub/internal/db"
	"github.com/AlbertoArenasG/clubhub/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	})

	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	dbClient, err := db.ConnectDB()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to MongoDB")
	}

	defer dbClient.Client().Disconnect(context.Background())

	// Create Fiber app
	app := fiber.New()

	// Define routes
	router.SetupRoutes(app)

	// Start server
	port := ":3000"
	logrus.Infof("Server is running on port %s", port)
	err = app.Listen(port)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
}
