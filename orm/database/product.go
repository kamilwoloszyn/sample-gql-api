package database

import (
	"context"
	"fmt"
	"time"

	"github.com/kamilwoloszyn/sample-gql-api/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	FieldID = "id"
)

type ProductHandler struct {
	handler *mongo.Collection
}

func NewProductHandler(handler *mongo.Collection) *ProductHandler {
	return &ProductHandler{
		handler: handler,
	}
}

func (ph *ProductHandler) Insert(ctx context.Context, product entity.Product) error {
	if ph.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	_, err := ph.handler.InsertOne(ctx, product)
	if err != nil {
		return fmt.Errorf("inserting product: %s", err)
	}
	return nil
}

func (ph *ProductHandler) BatchInsert(ctx context.Context, products []entity.Product) error {
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

func (ph *ProductHandler) Update(ctx context.Context, product entity.Product) error {
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
func (ph *ProductHandler) DeleteSoft(ctx context.Context, product entity.Product) error {
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

func (ph *ProductHandler) Delete(ctx context.Context, product entity.Product) error {
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

func (ph *ProductHandler) FindById(ctx context.Context, id string) (entity.Product, error) {
	if ph.handler == nil {
		return entity.Product{}, fmt.Errorf("handler not defined")
	}
	filterQuery := bson.D{
		{
			FieldID, id,
		},
	}
	result := ph.handler.FindOne(ctx, filterQuery)
	var product entity.Product
	if err := result.Decode(&product); err != nil {
		return entity.Product{}, fmt.Errorf("decode product: %s", err)
	}
	return product, nil
}