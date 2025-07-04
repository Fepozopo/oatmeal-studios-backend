package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleGetPlanogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid planogram ID", http.StatusBadRequest)
		return
	}

	planogram, err := service.GetPlanogram(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to get planogram", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(planogram); err != nil {
		http.Error(w, "Failed to encode planogram", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleListPlanograms(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	planograms, err := service.ListPlanograms(r.Context(), cfg.DbQueries)
	if err != nil {
		http.Error(w, "Failed to list planograms", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(planograms); err != nil {
		http.Error(w, "Failed to encode planograms", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleCreatePlanogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.CreatePlanogramInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	planogram, err := service.CreatePlanogram(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create planogram: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(planogram); err != nil {
		http.Error(w, "Failed to encode planogram", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleUpdatePlanogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.UpdatePlanogramInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	planogram, err := service.UpdatePlanogram(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update planogram", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(planogram); err != nil {
		http.Error(w, "Failed to encode planogram", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleDeletePlanogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid planogram ID", http.StatusBadRequest)
		return
	}

	err = service.DeletePlanogram(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to delete planogram", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *ApiConfig) HandleAssignPlanogramToLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.AssignPlanogramToLocationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	planogram, err := service.AssignPlanogramToLocation(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to assign planogram to location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(planogram); err != nil {
		http.Error(w, "Failed to encode planogram assignment", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleRemovePlanogramFromLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.RemovePlanogramFromLocationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := service.RemovePlanogramFromLocation(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to remove planogram from location", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *ApiConfig) HandleGetPlanogramsByLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	// Extract locationID from path
	prefix := "/api/planograms/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	idStr = strings.TrimSuffix(idStr, "/planograms") // Remove the suffix from the path
	locationID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid location ID", http.StatusBadRequest)
		return
	}

	planograms, err := service.GetPlanogramsByLocation(r.Context(), cfg.DbQueries, int32(locationID))
	if err != nil {
		http.Error(w, "Failed to get planograms by location", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(planograms); err != nil {
		http.Error(w, "Failed to encode planograms", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleListLocationsByPlanogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	planogramID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid planogram ID", http.StatusBadRequest)
		return
	}

	locations, err := service.ListLocationsByPlanogram(r.Context(), cfg.DbQueries, planogramID)
	if err != nil {
		http.Error(w, "Failed to list locations by planogram", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleListPocketsForPlanogram(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	// Extract planogramID from path
	prefix := "/api/planograms/"
	if !strings.HasPrefix(r.URL.Path, prefix) {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path, prefix)
	idStr = strings.TrimSuffix(idStr, "/pockets") // Remove the suffix from the path
	planogramID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid planogram ID", http.StatusBadRequest)
		return
	}

	pockets, err := service.ListPocketsForPlanogram(r.Context(), cfg.DbQueries, int32(planogramID))
	if err != nil {
		http.Error(w, "Failed to list pockets for planogram", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pockets); err != nil {
		http.Error(w, "Failed to encode pockets", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleGetPlanogramPocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid planogram pocket ID", http.StatusBadRequest)
		return
	}

	pocket, err := service.GetPlanogramPocket(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to get planogram pocket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pocket); err != nil {
		http.Error(w, "Failed to encode planogram pocket", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleCreatePlanogramPocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.CreatePlanogramPocketInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	pocket, err := service.CreatePlanogramPocket(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to create planogram pocket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(pocket); err != nil {
		http.Error(w, "Failed to encode planogram pocket", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleUpdatePlanogramPocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.UpdatePlanogramPocketInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	pocket, err := service.UpdatePlanogramPocket(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update planogram pocket", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pocket); err != nil {
		http.Error(w, "Failed to encode updated planogram pocket", http.StatusInternalServerError)
		return
	}
}

func (cfg *ApiConfig) HandleDeletePlanogramPocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid planogram pocket ID", http.StatusBadRequest)
		return
	}

	err = service.DeletePlanogramPocket(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to delete planogram pocket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cfg *ApiConfig) HandleGetPlanogramPocketByNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.GetPlanogramPocketByNumberInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	pocket, err := service.GetPlanogramPocketByNumber(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to get planogram pocket by number", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pocket); err != nil {
		http.Error(w, "Failed to encode planogram pocket", http.StatusInternalServerError)
		return
	}
}
