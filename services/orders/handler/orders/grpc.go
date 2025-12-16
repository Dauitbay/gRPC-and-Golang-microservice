package orders

import (
	"context"
	
	"github.com/sikozonpc/kitchen/services/common/genproto/orders"
	"github.com/sikozonpc/kitchen/services/orders/types"
	
	"google.golang.org/grpc"
)

type GrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &GrpcHandler{
		ordersService: ordersService,
	}
	// register gRPC service
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *GrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	o := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrdersResponse{
		Orders: o,
	}
	return res, nil
}

func (h *GrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := orders.Order{
		OrderId:    42,
		CustomerId: 2,
		ProductId:  1,
		Quantity:   10,
	}
	err := h.ordersService.CreateOrder(ctx, &order)
	if err != nil {
		return nil, err
	}
	
	res := &orders.CreateOrderResponse{
		Status: "Order created successfully",
	}
	return res, nil
}
