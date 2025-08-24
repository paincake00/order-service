package http

import (
	"net/http"
)

type OrderHandler struct {
	// передавать сюда OrderService
}

func (o *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {

}
