package usecase

import (
	"context"
	"product_service/internal/entity"
	"product_service/internal/infrastructure/repository"
	"time"
)

const (
	serviceNameArticle = "contentService"
	spanNameArticle    = "contentUsecase"
)

// Product -.
type Product interface {
	AddProduct(context.Context, *entity.Product) (*entity.Product, error)
	GetProduct(context.Context, string) (*entity.Product, error)
	UpdateProduct(context.Context, *entity.Product) (*entity.Product, error)
	DeleteProduct(context.Context, string) (*entity.Product, error)
}

type productService struct {
	BaseUseCase
	repo       repository.Product
	ctxTimeout time.Duration
}

func NewProductService(ctxTimeout time.Duration, repo repository.Product) productService {
	return productService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u productService) AddProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	//ctx, span := otlp.Start(ctx, serviceNameArticle, spanNameArticle+"Create")
	//defer span.End()

	u.beforeRequest(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	rproduct, err := u.repo.AddProduct(ctx, *product)
	if err != nil {
		return nil, err
	}

	return &rproduct, nil
}
func (u productService) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	//ctx, span := otlp.Start(ctx, serviceNameArticle, spanNameArticle+"Get")
	//defer span.End()

	gproduct, err := u.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return &gproduct, nil
}

func (u productService) UpdateProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()
	//
	//ctx, span := otlp.Start(ctx, serviceNameArticle, spanNameArticle+"Update")
	//defer span.End()

	//u.beforeRequest(nil, nil, &product.UpdatedAt)
	uproduct, err := u.repo.UpdateProduct(ctx, *product)
	if err != nil {
		return nil, err
	}
	return &uproduct, nil
}

func (u productService) DeleteProduct(ctx context.Context, id string) (*entity.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	//ctx, span := otlp.Start(ctx, serviceNameArticle, spanNameArticle+"Delete")
	//defer span.End()

	dproduct, err := u.repo.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dproduct, nil
}
