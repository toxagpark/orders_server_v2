package http

import (
	"WB_LVL_0_NEW/internal/domain/services"

	"github.com/gorilla/mux"
)

type Router struct {
	Router       *mux.Router
	orderHandler *OrderHandler
}

func NewRouter(orderService *services.OrderService) *Router {
	r := mux.NewRouter()
	orderHandler := NewOrderHandler(orderService)
	r.HandleFunc("/order", orderHandler.InterID)
	r.HandleFunc("/order/{id}", orderHandler.CheckID)
	return &Router{
		Router:       r,
		orderHandler: orderHandler,
	}
}
