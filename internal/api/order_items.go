package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleGetOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	orderID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Order ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	items, err := service.GetOrderItem(r.Context(), cfg.DbQueries, orderID)
	if err != nil {
		http.Error(w, "Failed to get order items: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (cfg *ApiConfig) HandleListOrderItemsBySKU(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	sku := r.URL.Query().Get("sku")
	if sku == "" {
		http.Error(w, "SKU is required", http.StatusBadRequest)
		return
	}

	items, err := service.ListOrderItemsBySKU(r.Context(), cfg.DbQueries, sku)
	if err != nil {
		http.Error(w, "Failed to list order items by SKU: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (cfg *ApiConfig) HandleCreateOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.CreateOrderItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	item, err := service.CreateOrderItem(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to create order item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func (cfg *ApiConfig) HandleUpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.UpdateOrderItemInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	item, err := service.UpdateOrderItem(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update order item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func (cfg *ApiConfig) HandleDeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Order Item ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteOrderItem(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to delete order item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *ApiConfig) HandleListOrderItemsByOrderID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	orderID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Order ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	items, err := service.ListOrderItemsByOrderID(r.Context(), cfg.DbQueries, orderID)
	if err != nil {
		http.Error(w, "Failed to list order items by order ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
