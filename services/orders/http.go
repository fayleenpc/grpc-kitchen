package main

import (
	"log"
	"net/http"

	handler "github.com/fayleenpc/grpc-kitchen/services/orders/handler/orders"
	"github.com/fayleenpc/grpc-kitchen/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRoutes(router)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
