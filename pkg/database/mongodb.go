package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB holds the MongoDB client and database
type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// Connect establishes a connection to MongoDB
func Connect(uri, dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(dbName)

	return &MongoDB{
		Client: client,
		DB:     db,
	}, nil
}

// Disconnect closes the MongoDB connection
func (m *MongoDB) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := m.Client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	return nil
}

// HealthCheck checks if the database connection is alive
func (m *MongoDB) HealthCheck(ctx context.Context) error {
	if err := m.Client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("database health check failed: %w", err)
	}
	return nil
}

// DropDatabase drops the entire database (use with caution!)
func (m *MongoDB) DropDatabase(ctx context.Context) error {
	if err := m.DB.Drop(ctx); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}
	return nil
}
