package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleAddCustomerLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.AddCustomerLocationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	location, err := service.AddCustomerLocation(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to add customer location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func (cfg *ApiConfig) HandleDeleteCustomerLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	locationID := r.URL.Query().Get("id")
	if locationID == "" {
		http.Error(w, "Location ID is required", http.StatusBadRequest)
		return
	}

	var locationIDInt int32
	_, err := fmt.Sscanf(locationID, "%d", &locationIDInt)
	if err != nil {
		http.Error(w, "Invalid Location ID format", http.StatusBadRequest)
		return
	}

	if err := service.DeleteCustomerLocation(r.Context(), cfg.DbQueries, locationIDInt); err != nil {
		http.Error(w, "Failed to delete customer location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *ApiConfig) HandleUpdateCustomerLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.UpdateCustomerLocationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	location, err := service.UpdateCustomerLocation(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update customer location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func (cfg *ApiConfig) HandleGetCustomerLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	locationID := r.URL.Query().Get("id")
	if locationID == "" {
		http.Error(w, "Location ID is required", http.StatusBadRequest)
		return
	}

	var locationIDInt int32
	_, err := fmt.Sscanf(locationID, "%d", &locationIDInt)
	if err != nil {
		http.Error(w, "Invalid Location ID format", http.StatusBadRequest)
		return
	}

	location, err := service.GetCustomerLocationByID(r.Context(), cfg.DbQueries, locationIDInt)
	if err != nil {
		http.Error(w, "Failed to get customer location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(location)
}

func (cfg *ApiConfig) HandleListCustomerLocations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var userIDInt int32
	_, err := fmt.Sscanf(userID, "%d", &userIDInt)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	locations, err := service.ListCustomerLocations(r.Context(), cfg.DbQueries, userIDInt)
	if err != nil {
		http.Error(w, "Failed to list customer locations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}
