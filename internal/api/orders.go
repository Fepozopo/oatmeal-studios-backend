package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RegisterOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/orders", ordersHandler)
	mux.HandleFunc("/orders/", orderHandler)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "List all orders"})
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Create a new order"})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/orders/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(map[string]string{"message": "Get order by ID", "id": id})
	case http.MethodPut:
		json.NewEncoder(w).Encode(map[string]string{"message": "Update order by ID", "id": id})
	case http.MethodDelete:
		json.NewEncoder(w).Encode(map[string]string{"message": "Delete order by ID", "id": id})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
