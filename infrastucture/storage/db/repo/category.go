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

type CategoryRepo struct {
	handler *mongo.Collection
}

func NewCategoryRepo(db *db.Database) *CategoryRepo {
	handler := db.GetCollectionHandler(CollectionNameCategory)
	return &CategoryRepo{
		handler: handler,
	}
}

func (c *CategoryRepo) Insert(ctx context.Context, category entity.Category) error {
	if c.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	_, err := c.handler.InsertOne(ctx, category)
	if err != nil {
		return fmt.Errorf("inserting category: %s", err)
	}
	return nil
}

func (c *CategoryRepo) DeleteSoft(ctx context.Context, category entity.Category) error {
	if c.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	productID := category.GetID()
	if productID != "" {
		return fmt.Errorf("product id is empty")
	}
	category.DeletedAt = time.Now().Unix()
	_, err := c.handler.UpdateByID(ctx, productID, category)
	if err != nil {
		fmt.Errorf("updating category: %s", err)
	}
	return nil
}

func (c *CategoryRepo) Delete(ctx context.Context, category entity.Category) error {
	if c.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	productID := category.GetID()
	if productID == "" {
		return fmt.Errorf("product id is empty")
	}
	filterQuery := bson.D{
		{
			FieldID, productID,
		},
	}
	_, err := c.handler.DeleteOne(ctx, filterQuery)
	if err != nil {
		return fmt.Errorf("could not delete a record: %s", err)
	}
	return nil

}

func (c *CategoryRepo) FindByID(ctx context.Context, id string) (entity.Category, error) {
	if c.handler == nil {
		return entity.Category{}, fmt.Errorf("handler not defined")
	}
	filterQuery := bson.D{
		{
			FieldID, id,
		},
	}
	result := c.handler.FindOne(ctx, filterQuery)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return entity.Category{}, nil
		}
		return entity.Category{}, fmt.Errorf("fetching category: %s", err)
	}
	var category entity.Category
	if err := result.Decode(&category); err != nil {
		return entity.Category{}, fmt.Errorf("decoding category: %s", err)
	}
	return category, nil
}
