package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleCreateInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.CreateInvoiceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := service.CreateInvoice(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to create invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

func (cfg *ApiConfig) HandleGetInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Invoice ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := service.GetInvoice(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to get invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

func (cfg *ApiConfig) HandleGetInvoicesByOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	orderID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Order ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := service.GetInvoicesByOrder(r.Context(), cfg.DbQueries, orderID)
	if err != nil {
		http.Error(w, "Failed to get invoice by order: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

func (cfg *ApiConfig) HandleListInvoicesByCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	customerID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Customer ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	invoices, err := service.ListInvoicesByCustomer(r.Context(), cfg.DbQueries, customerID)
	if err != nil {
		http.Error(w, "Failed to list invoices by customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func (cfg *ApiConfig) HandleListInvoicesByCustomerLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	customerLocationID, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Customer Location ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	invoices, err := service.ListInvoicesByCustomerLocation(r.Context(), cfg.DbQueries, customerLocationID)
	if err != nil {
		http.Error(w, "Failed to list invoices by customer location: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func (cfg *ApiConfig) HandleUpdateInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.UpdateInvoiceInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	invoice, err := service.UpdateInvoice(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

func (cfg *ApiConfig) HandleDeleteInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := idFromURLAsInt32(r)
	if err != nil {
		http.Error(w, "Invalid Invoice ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteInvoice(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to delete invoice: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
