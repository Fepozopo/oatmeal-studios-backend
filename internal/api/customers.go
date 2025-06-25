package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.CreateCustomerInput
	// Decode the JSON request body into the input struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := service.CreateCustomer(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to create customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)

}

func (cfg *ApiConfig) HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	// Extract customerId from path value
	customerIdStr := r.PathValue("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil || customerId <= 0 {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	customer, err := service.GetCustomerByID(r.Context(), cfg.DbQueries, int32(customerId))
	if err != nil {
		http.Error(w, "Failed to get customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (cfg *ApiConfig) HandleUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.UpdateCustomerInput
	// Decode the JSON request body into the input struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := service.UpdateCustomer(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (cfg *ApiConfig) HandleListCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	customers, err := service.ListCustomers(r.Context(), cfg.DbQueries)
	if err != nil {
		http.Error(w, "Failed to list customers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (cfg *ApiConfig) HandleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	// Extract customerId from path value
	customerIdStr := r.PathValue("customerId")
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil || customerId <= 0 {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}

	err = service.DeleteCustomer(r.Context(), cfg.DbQueries, int32(customerId))
	if err != nil {
		http.Error(w, "Failed to delete customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
