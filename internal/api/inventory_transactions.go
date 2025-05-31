package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandleGetCurrentInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	productID, err := idFromURLAsUUID(r)
	if productID == uuid.Nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	inventory, err := service.GetCurrentInventory(r.Context(), cfg.DbQueries, productID)
	if err != nil {
		http.Error(w, "Failed to get current inventory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"current_inventory": inventory})
}

func (cfg *ApiConfig) HandleGetAllCurrentInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	inventory, err := service.GetAllCurrentInventory(r.Context(), cfg.DbQueries)
	if err != nil {
		http.Error(w, "Failed to get all current inventory: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventory)
}

func (cfg *ApiConfig) HandleListInventoryTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	productID, err := idFromURLAsUUID(r)
	if productID == uuid.Nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	transactions, err := service.ListInventoryTransactions(r.Context(), cfg.DbQueries, productID)
	if err != nil {
		http.Error(w, "Failed to list inventory transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}

func (cfg *ApiConfig) HandleInsertInventoryTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.InsertInventoryTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := service.InsertInventoryTransaction(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to insert inventory transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

func (cfg *ApiConfig) HandleGetInventoryChangesByDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	productID, err := idFromURLAsUUID(r)
	if productID == uuid.Nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	rows, err := service.GetInventoryChangesByDay(r.Context(), cfg.DbQueries, productID)
	if err != nil {
		http.Error(w, "Failed to get inventory changes by day: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rows)
}
