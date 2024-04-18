package repository

import (
	"context"
	"product_service/internal/entity"
)

type Product interface {
	AddProduct(context.Context, entity.Product) (entity.Product, error)
	GetProduct(context.Context, string) (entity.Product, error)
	UpdateProduct(context.Context, entity.Product) (entity.Product, error)
	DeleteProduct(context.Context, string) (entity.Product, error)
}
