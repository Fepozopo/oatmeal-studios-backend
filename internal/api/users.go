package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

func RegisterUserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/", userHandler)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "List all users"})
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Create a new user"})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(map[string]string{"message": "Get user by ID", "id": id})
	case http.MethodPut:
		json.NewEncoder(w).Encode(map[string]string{"message": "Update user by ID", "id": id})
	case http.MethodDelete:
		json.NewEncoder(w).Encode(map[string]string{"message": "Delete user by ID", "id": id})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
