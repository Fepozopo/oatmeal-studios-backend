package api

import (
	"encoding/json"
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

	locationID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Location ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := service.DeleteCustomerLocation(r.Context(), cfg.DbQueries, locationID); err != nil {
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

	locationID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Location ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	location, err := service.GetCustomerLocationByID(r.Context(), cfg.DbQueries, locationID)
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

	userID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid User ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	locations, err := service.ListCustomerLocations(r.Context(), cfg.DbQueries, userID)
	if err != nil {
		http.Error(w, "Failed to list customer locations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
}
