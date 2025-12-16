package main

import (
	"log"
	"net"
	
	"github.com/sikozonpc/kitchen/services/orders/handler/orders"
	"github.com/sikozonpc/kitchen/services/orders/service"
	
	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{addr: addr}
}

func (s *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	
	// register gRPC services
	orderService := service.NewOrderService()
	orders.NewGrpcOrdersService(grpcServer, orderService)
	
	log.Printf("server listening at %v", lis.Addr())
	
	return grpcServer.Serve(lis)
}
