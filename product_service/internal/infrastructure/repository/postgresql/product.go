package postgresql

import (
	"context"
	"fmt"
	"product_service/internal/entity"
	"product_service/internal/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
)

const (
	productTableName      = "product"
	productServiceName    = "productService"
	productSpanRepoPrefix = "productRepo"
)

type productRepo struct {
	tableName string
	db        *postgres.PostgresDB
}

func NewProductRepo(db *postgres.PostgresDB) *productRepo {
	return &productRepo{
		tableName: productTableName,
		db:        db,
	}
}

//func (p *productRepo) productSelectQueryPrefix() squirrel.SelectBuilder {
//	return p.db.Sq.Builder.
//		Select(
//			"guid",
//			"chapter_id",
//			"title",
//			"lang_id",
//			"text",
//			"icon",
//			"media",
//			"status",
//			"created_at",
//			"updated_at",
//		).From(p.tableName)
//}

func (p productRepo) AddProduct(ctx context.Context, t entity.Product) (entity.Product, error) {
	sql, args, err := p.db.Sq.Builder.
		Insert("products").
		Columns("id", "name", "description", "category_id", "unit_price").
		Values(t.ID, t.Name, t.Description, t.CategoryID, t.UnitPrice).
		ToSql()
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - AddProduct - r.Builder: %w", err)
	}

	sql += "RETURNING id, name, description, category_id, unit_price"

	var addedProduct entity.Product
	fmt.Println(t)
	err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(&addedProduct.ID, &addedProduct.Name, &addedProduct.Description, &addedProduct.CategoryID, &addedProduct.UnitPrice)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - AddProduct - r.Pool.QueryRow: %w", err)
	}

	return addedProduct, nil
}

func (p productRepo) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	sql, args, err := p.db.Sq.Builder.
		Select("id", "name", "description", "category_id", "unit_price", "created_at", "updated_at").
		From("products").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - GetProduct - r.Builder: %w", err)
	}

	var product entity.Product
	var createdAtTime time.Time
	var updatedAtTime time.Time
	if err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(&product.ID, &product.Name, &product.Description, &product.CategoryID, &product.UnitPrice, &createdAtTime, &updatedAtTime); err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - GetProduct - r.Pool.QueryRow: %w", err)
	}

	product.CreatedAt = createdAtTime.String()
	product.UpdatedAt = updatedAtTime.String()

	return product, nil
}

func (p productRepo) UpdateProduct(ctx context.Context, t entity.Product) (entity.Product, error) {
	sql, args, err := p.db.Sq.Builder.
		Update("products").
		Set("name", t.Name).
		Set("description", t.Description).
		Set("category_id", t.CategoryID).
		Set("unit_price", t.UnitPrice).
		Where(squirrel.Eq{"id": t.ID}).
		ToSql()
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - UpdateProduct - r.Builder: %w", err)
	}

	sql += " RETURNING id, name, description, category_id, unit_price, created_at, updated_at"

	var updatedProduct entity.Product
	var createdAtTime time.Time
	var updatedAtTime time.Time

	err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&updatedProduct.ID,
		&updatedProduct.Name,
		&updatedProduct.Description,
		&updatedProduct.CategoryID,
		&updatedProduct.UnitPrice,
		&createdAtTime,
		&updatedAtTime,
	)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - UpdateProduct - r.Pool.QueryRow: %w", err)
	}

	updatedProduct.CreatedAt = createdAtTime.String()
	updatedProduct.UpdatedAt = updatedAtTime.String()
	return updatedProduct, nil
}

func (p productRepo) DeleteProduct(ctx context.Context, id string) (entity.Product, error) {
	sql, args, err := p.db.Sq.Builder.
		Delete("products").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - DeleteProduct - r.Builder: %w", err)
	}

	var deletedProduct entity.Product
	var createdAtTime time.Time
	var updatedAtTime time.Time
	sql += " RETURNING id, name, description, category_id, unit_price, created_at, updated_at"
	if err = p.db.Pool.QueryRow(ctx, sql, args...).Scan(&deletedProduct.ID, &deletedProduct.Name, &deletedProduct.Description, &deletedProduct.CategoryID, &deletedProduct.UnitPrice, &createdAtTime, &updatedAtTime); err != nil {
		return entity.Product{}, fmt.Errorf("ProductRepo - DeleteProduct - r.Pool.QueryRow: %w", err)
	}

	deletedProduct.CreatedAt = createdAtTime.String()
	deletedProduct.UpdatedAt = updatedAtTime.String()
	return deletedProduct, nil
}
