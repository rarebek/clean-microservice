package grpc_service_clients

import (
	"product_service/internal/pkg/config"

	"google.golang.org/grpc"
)

type ServiceClients interface {
	// SmsService()
	Close()
}

type serviceClients struct {
	services []*grpc.ClientConn
}

func New(config *config.Config) (ServiceClients, error) {
	return &serviceClients{
		services: []*grpc.ClientConn{},
	}, nil
}

func (s *serviceClients) Close() {
	// closing investment service
	for _, conn := range s.services {
		conn.Close()
	}
}
