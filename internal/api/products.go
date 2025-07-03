package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
)

func (cfg *ApiConfig) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	product, err := service.CreateProduct(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (cfg *ApiConfig) HandleGetProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsUUID(r)
	if err != nil {
		http.Error(w, "Invalid UUID: "+err.Error(), http.StatusBadRequest)
		return
	}

	product, err := service.GetProductByID(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (cfg *ApiConfig) HandleGetProductBySKU(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	// Extract SKU from path
	// Example: /api/products/sku/ABC123
	prefix := "/api/products/sku/"
	if len(r.URL.Path) <= len(prefix) {
		http.Error(w, "Missing SKU", http.StatusBadRequest)
		return
	}
	sku := r.URL.Path[len(prefix):]
	if sku == "" {
		http.Error(w, "Missing SKU", http.StatusBadRequest)
		return
	}

	product, err := service.GetProductBySKU(r.Context(), cfg.DbQueries, sku)
	if err != nil {
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (cfg *ApiConfig) HandleListProductsByType(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	typ := r.URL.Query().Get("type")
	if typ == "" {
		http.Error(w, "Missing type", http.StatusBadRequest)
		return
	}

	products, err := service.ListProductsByType(r.Context(), cfg.DbQueries, typ)
	if err != nil {
		http.Error(w, "Failed to list products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (cfg *ApiConfig) HandleListProductsByCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	category := r.URL.Query().Get("category")
	if category == "" {
		http.Error(w, "Missing category", http.StatusBadRequest)
		return
	}

	products, err := service.ListProductsByCategory(r.Context(), cfg.DbQueries, category)
	if err != nil {
		http.Error(w, "Failed to list products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (cfg *ApiConfig) HandleListProductsByArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	artist := r.URL.Query().Get("artist")
	if artist == "" {
		http.Error(w, "Missing artist", http.StatusBadRequest)
		return
	}

	products, err := service.ListProductsByArtist(r.Context(), cfg.DbQueries, artist)
	if err != nil {
		http.Error(w, "Failed to list products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (cfg *ApiConfig) HandleListProductsByStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	status := r.URL.Query().Get("status")
	if status == "" {
		http.Error(w, "Missing status", http.StatusBadRequest)
		return
	}

	products, err := service.ListProductsByStatus(r.Context(), cfg.DbQueries, status)
	if err != nil {
		http.Error(w, "Failed to list products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func (cfg *ApiConfig) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPatch {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var input service.UpdateProductInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	product, err := service.UpdateProduct(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (cfg *ApiConfig) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	id, err := idFromURLAsUUID(r)
	if err != nil {
		http.Error(w, "Invalid UUID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteProduct(r.Context(), cfg.DbQueries, id)
	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
