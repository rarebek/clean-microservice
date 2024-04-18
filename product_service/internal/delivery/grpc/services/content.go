package services

import (
	"context"
	"go.uber.org/zap"
	pb "product_service/genproto/product"
	"product_service/internal/delivery/grpc"
	"product_service/internal/entity"
	"product_service/internal/usecase"
)

type contentRPC struct {
	logger         *zap.Logger
	productUseCase usecase.Product
	pb.UnimplementedProductServiceServer
	//brokerProducer event.BrokerProducer

}

func NewRPC(logger *zap.Logger, productUseCase usecase.Product) *contentRPC {
	//brokerProducer event.BrokerProducer) pb.ProductServiceServer {
	return &contentRPC{
		logger:         logger,
		productUseCase: productUseCase,
		//brokerProducer: brokerProducer,
	}
}

func (s contentRPC) AddProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	addedProduct, err := s.productUseCase.AddProduct(ctx, &entity.Product{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		CategoryID:  in.CategoryID,
		UnitPrice:   float64(in.UnitPrice),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	})
	if err != nil {
		s.logger.Error("articleCategoriesUsecase.Get", zap.Error(err))
		return &pb.Product{}, grpc.Error(ctx, err)
	}

	return &pb.Product{
		ID:          addedProduct.ID,
		Name:        addedProduct.Name,
		Description: addedProduct.Description,
		CategoryID:  addedProduct.CategoryID,
		UnitPrice:   float32(addedProduct.UnitPrice),
		CreatedAt:   addedProduct.CreatedAt,
		UpdatedAt:   addedProduct.UpdatedAt,
	}, nil
}

func (s contentRPC) GetProduct(ctx context.Context, in *pb.IdRequest) (*pb.Product, error) {
	gotProduct, err := s.productUseCase.GetProduct(ctx, in.ID)
	if err != nil {
		s.logger.Error("articleCategoriesUsecase.Get", zap.Error(err))
		return &pb.Product{}, grpc.Error(ctx, err)
	}

	return &pb.Product{
		ID:          gotProduct.ID,
		Name:        gotProduct.Name,
		Description: gotProduct.Description,
		CategoryID:  gotProduct.CategoryID,
		UnitPrice:   float32(gotProduct.UnitPrice),
		CreatedAt:   gotProduct.CreatedAt,
		UpdatedAt:   gotProduct.UpdatedAt,
	}, nil
}

func (s contentRPC) UpdateProduct(ctx context.Context, in *pb.Product) (*pb.Product, error) {
	addedProduct, err := s.productUseCase.UpdateProduct(ctx, &entity.Product{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		CategoryID:  in.CategoryID,
		UnitPrice:   float64(in.UnitPrice),
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	})
	if err != nil {
		s.logger.Error("articleCategoriesUsecase.Get", zap.Error(err))
		return &pb.Product{}, grpc.Error(ctx, err)
	}

	return &pb.Product{
		ID:          addedProduct.ID,
		Name:        addedProduct.Name,
		Description: addedProduct.Description,
		CategoryID:  addedProduct.CategoryID,
		UnitPrice:   float32(addedProduct.UnitPrice),
		CreatedAt:   addedProduct.CreatedAt,
		UpdatedAt:   addedProduct.UpdatedAt,
	}, nil
}

func (s contentRPC) DeleteProduct(ctx context.Context, in *pb.IdRequest) (*pb.Product, error) {
	gotProduct, err := s.productUseCase.DeleteProduct(ctx, in.ID)
	if err != nil {
		s.logger.Error("articleCategoriesUsecase.Get", zap.Error(err))
		return &pb.Product{}, grpc.Error(ctx, err)
	}

	return &pb.Product{
		ID:          gotProduct.ID,
		Name:        gotProduct.Name,
		Description: gotProduct.Description,
		CategoryID:  gotProduct.CategoryID,
		UnitPrice:   float32(gotProduct.UnitPrice),
		CreatedAt:   gotProduct.CreatedAt,
		UpdatedAt:   gotProduct.UpdatedAt,
	}, nil
}
