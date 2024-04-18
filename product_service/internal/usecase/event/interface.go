package event

import (
	"context"
	"product_service/internal/entity"
)

type ConsumerConfig interface {
	GetBrokers() []string
	GetTopic() string
	GetGroupID() string
	GetHandler() func(ctx context.Context, key, value []byte) error
}

type BrokerConsumer interface {
	Run()
	RegisterConsumer(config ConsumerConfig)
	Close()
}

// Product -.
type Product interface {
	AddProduct(context.Context, entity.Product) (entity.Product, error)
	GetProduct(context.Context, string) (entity.Product, error)
	UpdateProduct(context.Context, entity.Product) (entity.Product, error)
	DeleteProduct(context.Context, string) (entity.Product, error)
}

//
//type BrokerProducer interface {
//	ProduceContent(ctx context.Context, key string, value *entity.ArticleCategories) error
//	Close()
//}
