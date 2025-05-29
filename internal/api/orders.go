package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Order ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	order, err := service.GetOrder(r.Context(), cfg.DbQueries, orderID)
	if err != nil {
		http.Error(w, "Failed to get order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (cfg *ApiConfig) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.CreateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	order, err := service.CreateOrder(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to create order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (cfg *ApiConfig) HandleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.UpdateOrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	order, err := service.UpdateOrder(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (cfg *ApiConfig) HandleListOrdersOpen(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	orders, err := service.ListOrdersOpen(r.Context(), cfg.DbQueries)
	if err != nil {
		http.Error(w, "Failed to list orders: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (cfg *ApiConfig) HandleListOrdersByCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	customerID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Customer ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := service.ListOrdersByCustomer(r.Context(), cfg.DbQueries, customerID)
	if err != nil {
		http.Error(w, "Failed to list orders: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (cfg *ApiConfig) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Order ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteOrder(r.Context(), cfg.DbQueries, orderID)
	if err != nil {
		http.Error(w, "Failed to delete order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
