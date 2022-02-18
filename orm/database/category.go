package database

import "go.mongodb.org/mongo-driver/mongo"

type CategoryHandler struct {
	handler *mongo.Collection
}

func NewCategoryHandler(handler *mongo.Collection) *CategoryHandler {
	return &CategoryHandler{
		handler: handler,
	}
}
