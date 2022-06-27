package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func NewClient(connectURL string) (*mongo.Client, error) {
	dialTimeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	mongoOpts := options.Client().ApplyURI(connectURL)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		return nil, err
	}

	if err = checkConnection(ctx, client); err != nil {
		return nil, err
	}

	return client, nil
}

func checkConnection(ctx context.Context, client *mongo.Client) error {
	rp, err := readpref.New(readpref.PrimaryMode)
	if err != nil {
		return err
	}

	return client.Ping(ctx, rp)
}
