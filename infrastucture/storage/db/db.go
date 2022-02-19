package db

import (
	"context"
	"fmt"
	"time"

	"github.com/kamilwoloszyn/sample-gql-api/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client       *mongo.Client
	dbName       string
	CategoryRepo repository.CategoryRepo
	ProductRepo  repository.ProductRepo
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

func (d *Database) SetRepo(opts ...RepoOpts) {
	for _, opt := range opts {
		opt(d)
	}
}

type RepoOpts func(d *Database)

func SetCategoryRepo(c repository.CategoryRepo) RepoOpts {
	return func(d *Database) {
		d.CategoryRepo = c
	}
}

func SetProductRepo(p repository.ProductRepo) RepoOpts {
	return func(d *Database) {
		d.ProductRepo = p
	}
}
