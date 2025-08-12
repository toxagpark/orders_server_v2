package http

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/services"
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	orderService *services.OrderService // Инжектим зависимость
}

func NewOrderHandler(s *services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: s}
}

func (h *OrderHandler) InterID(w http.ResponseWriter, r *http.Request) {
	path := "static/templates/InterID.html"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, path)
}

func (h *OrderHandler) CheckID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	model, err := h.orderService.HandleOrderGet(ctx, orderID)
	if err != nil {
		notFound(w, r)
		return
	}

	if err := renderOrder(w, model); err != nil {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/NotFound.html")
}

func renderOrder(w http.ResponseWriter, order *model.Order) error {
	tmpl, err := template.ParseFiles("static/templates/Info.html")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.Execute(w, order)
}
