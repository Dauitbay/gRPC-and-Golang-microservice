package main

import (
	"log"
	"net/http"
	
	handler "github.com/sikozonpc/kitchen/services/orders/handler/orders"
	"github.com/sikozonpc/kitchen/services/orders/service"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (s *HttpServer) Run() error {
	router := http.NewServeMux()
	
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)
	log.Printf("server listening at %v", s.addr)
	return http.ListenAndServe(s.addr, router)
}
