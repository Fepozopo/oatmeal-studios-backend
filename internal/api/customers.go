package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RegisterCustomerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/customers", customersHandler)
	mux.HandleFunc("/customers/", customerHandler)
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "List all customers"})
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Create a new customer"})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(map[string]string{"message": "Get customer by ID", "id": id})
	case http.MethodPut:
		json.NewEncoder(w).Encode(map[string]string{"message": "Update customer by ID", "id": id})
	case http.MethodDelete:
		json.NewEncoder(w).Encode(map[string]string{"message": "Delete customer by ID", "id": id})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
