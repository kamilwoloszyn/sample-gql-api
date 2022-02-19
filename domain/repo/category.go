package repo

import (
	"context"

	"github.com/kamilwoloszyn/sample-gql-api/domain/entity"
)

type CategoryRepo interface {
	InsertCategory(context.Context, entity.Category) error
	DeleteSoft(context.Context, entity.Category) error
	Delete(context.Context, entity.Category) error
	FindByID(context.Context, string) (entity.Category, error)
}
