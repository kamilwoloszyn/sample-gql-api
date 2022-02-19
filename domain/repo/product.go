package repo

import (
	"context"

	"github.com/kamilwoloszyn/sample-gql-api/domain/entity"
)

type ProductRepo interface {
	NewProduct(context.Context, entity.Product) error
	BatchInsert(context.Context, []entity.Product) error
	Update(context.Context, entity.Product) error
	DeleteSoft(context.Context, entity.Product) error
	Delete(context.Context, entity.Product) error
	FindById(context.Context, string) (entity.Product, error)
}
