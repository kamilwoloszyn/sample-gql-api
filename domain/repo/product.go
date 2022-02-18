package repo

import "github.com/kamilwoloszyn/sample-gql-api/domain/entity"

type ProductRepo interface {
	GetProductByID(id string) (entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
	InsertProduct([]entity.Product) error
}
