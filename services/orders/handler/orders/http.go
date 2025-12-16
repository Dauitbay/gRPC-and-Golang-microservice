package orders

import (
	"net/http"
	
	"github.com/sikozonpc/kitchen/services/common/genproto/orders"
	"github.com/sikozonpc/kitchen/services/common/util"
	"github.com/sikozonpc/kitchen/services/orders/types"
)

type HttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(ordersService types.OrderService) *HttpHandler {
	handler := &HttpHandler{
		ordersService: ordersService,
	}
	return handler
}

func (h *HttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *HttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	order := &orders.Order{
		OrderId:    42,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
	}
	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	res := &orders.CreateOrderResponse{
		Status: "Order created successfully",
	}
	err = util.WriteJSON(w, http.StatusOK, res)
	if err != nil {
		return
	}
}
