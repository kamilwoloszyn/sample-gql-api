package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/kamilwoloszyn/sample-gql-api/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryHandler struct {
	handler *mongo.Collection
}

func NewCategoryHandler(handler *mongo.Collection) *CategoryHandler {
	return &CategoryHandler{
		handler: handler,
	}
}

func (c *CategoryHandler) InsertCategory(ctx context.Context, category entity.Category) error {
	if c.handler == nil {
		return fmt.Errorf("handler not defined")
	}
	_, err := c.handler.InsertOne(ctx, category)
	if err != nil {
		return fmt.Errorf("inserting category: %s", err)
	}
	return nil
}

func (c *CategoryHandler) DeleteSoft(ctx context.Context, category entity.Category) error {
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

func (c *CategoryHandler) Delete(ctx context.Context, category entity.Category) error {
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

func (c *CategoryHandler) FindByID(ctx context.Context, id string) (entity.Category, error) {
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
