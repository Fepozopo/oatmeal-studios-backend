package main

import (
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/api"
)

func main() {
	mux := http.NewServeMux()

	// Root route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Oatmeal Studios Backend API"}`))
	})

	// Register resource routes
	api.RegisterUserRoutes(mux)
	api.RegisterCustomerRoutes(mux)
	api.RegisterOrderRoutes(mux)

	// Run the server on port 8080
	http.ListenAndServe(":8080", mux)
}
