package repo

import "github.com/kamilwoloszyn/sample-gql-api/domain/entity"

type CategoryRepo interface {
	GetCategories() []entity.Category
	InsertCategory(entity.Category) error
}
