package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	dbName string
}

func InitializeDatabase(
	ctx context.Context,
	timeout time.Duration,
	host, port, dbName string,
) (*Database, error) {
	clientOpts := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s/", host, port),
	)
	ctx, cancelFunc := context.WithTimeout(ctx, timeout)
	defer cancelFunc()
	client, err := mongo.Connect(
		ctx,
		clientOpts,
	)
	if err != nil {
		return nil, fmt.Errorf("connecting to db: %s", err)
	}
	return &Database{
		client: client,
		dbName: dbName,
	}, nil
}

func (d *Database) TestConnection(ctx context.Context, timeout time.Duration) error {
	ctx, cancelFunc := context.WithTimeout(ctx, timeout)
	defer cancelFunc()
	err := d.client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetCollectionHandler(collectonName string) *mongo.Collection {
	return d.client.Database(d.dbName).Collection(collectonName)
}

func (d *Database) CloseConnection(ctx context.Context) error {
	return d.client.Disconnect(ctx)
}
