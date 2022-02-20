package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/kamilwoloszyn/sample-gql-api/domain/entity"
	"github.com/kamilwoloszyn/sample-gql-api/infrastucture/storage/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	handler *mongo.Collection
}

func NewProductRepo(db *db.Database) *ProductRepo {
	handler := db.GetCollectionHandler(CollectionNameProduct)
	return &ProductRepo{
		handler: handler,
	}
}

func (ph *ProductRepo) Insert(ctx context.Context, product entity.Product) error {
	if ph.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	_, err := ph.handler.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("inserting product: %s", err)
	}
	return nil
}

func (ph *ProductRepo) BatchInsert(ctx context.Context, products []entity.Product) error {
	if ph.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	records := []interface{}{}
	records = append(records, products)
	_, err := ph.handler.InsertMany(ctx, records)
	if err != nil {
		return fmt.Errorf("batch insert: %s", err)
	}
	return nil
}

func (ph *ProductRepo) Update(ctx context.Context, product entity.Product) error {
	if ph.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	productID := product.GetID()
	if productID == "" {
		return fmt.Errorf("product id doesn't exist")
	}
	_, err := ph.handler.UpdateByID(ctx, productID, product)
	if err != nil {
		return fmt.Errorf("update product: %s", err)
	}
	return nil
}
func (ph *ProductRepo) DeleteSoft(ctx context.Context, product entity.Product) error {
	if ph.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	productID := product.GetID()
	if productID == "" {
		return fmt.Errorf("product id doesn't exist")
	}
	product.DeletedAt = time.Now().Unix()
	_, err := ph.handler.UpdateByID(ctx, productID, product)
	if err != nil {
		return fmt.Errorf("soft delete product: %s", err)
	}
	return nil
}

func (ph *ProductRepo) Delete(ctx context.Context, product entity.Product) error {
	if ph.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	productID := product.GetID()
	if productID == "" {
		return fmt.Errorf("product id doesn't exist")
	}
	filterQuery := bson.D{
		{
			FieldID, productID,
		},
	}
	_, err := ph.handler.DeleteOne(ctx, filterQuery)
	if err != nil {
		return fmt.Errorf("could not delete a product: %s", err)
	}
	return nil
}

func (ph *ProductRepo) FindById(ctx context.Context, id string) (entity.Product, error) {
	if ph.handler == nil {
		return entity.Product{}, fmt.Errorf("handler not defined")
	}
	filterQuery := bson.D{
		{
			FieldID, id,
		},
	}
	result := ph.handler.FindOne(ctx, filterQuery)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Product{}, nil
		}
		return entity.Product{}, fmt.Errorf("fetching products from db")
	}
	var product entity.Product
	if err := result.Decode(&product); err != nil {
		return entity.Product{}, fmt.Errorf("decode product: %s", err)
	}
	return product, nil
}
