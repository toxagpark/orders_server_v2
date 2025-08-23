package handlers

import (
	"WB_LVL_0_NEW/internal/domain/model"
	"WB_LVL_0_NEW/internal/domain/services"
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	ErrParse   = errors.New("error parse")
	ErrExecute = errors.New("error execute")
)

type OrderHandler struct {
	orderService *services.OrderService
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
	if errors.Is(err, services.ErrOrderNotFound) {
		notFound(w, r)
		return
	}

	if err := renderOrder(w, model); errors.Is(err, ErrExecute) || errors.Is(err, ErrParse) {
		http.Error(w, "Error rendering page", http.StatusInternalServerError)
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/templates/NotFound.html")
}

func renderOrder(w http.ResponseWriter, order *model.Order) error {
	tmpl, err := template.ParseFiles("static/templates/Info.html")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrParse, err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err = tmpl.Execute(w, order); err != nil {
		return fmt.Errorf("%w: %w", ErrExecute, err)
	}
	return nil
}
