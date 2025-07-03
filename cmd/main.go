package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/api"
	"github.com/Fepozopo/oatmeal-studios-backend/internal/database"
	"github.com/joho/godotenv" // godotenv for loading environment variables
	_ "github.com/lib/pq"      // PostgreSQL driver
)

func main() {
	port := ":8080" // Default port for the API

	// Open a connection to the database and environment variables
	godotenv.Load()
	tokenSecret := os.Getenv("TOKEN_SECRET")
	dbURL := os.Getenv("DB_URL")

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Failed to open a connection to the database: %v\n", err)
	}
	defer db.Close()

	// Create a new Queries instance and initialize the ApiConfig struct
	dbQueries := database.New(db)
	apiCfg := &api.ApiConfig{
		DbQueries:   dbQueries,
		TokenSecret: tokenSecret,
	}

	// Create a new HTTP ServeMux
	mux := http.NewServeMux()

	// Root route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Oatmeal Studios Backend API"}`))
	})

	// User routes
	mux.HandleFunc("GET /api/users", apiCfg.HandleListUsers)
	mux.HandleFunc("POST /api/users/register", apiCfg.HandleRegisterUser)
	mux.HandleFunc("POST /api/users/authenticate", apiCfg.HandleAuthenticateUser)
	mux.HandleFunc("GET /api/users/{id}", apiCfg.HandleGetUser)
	mux.HandleFunc("PUT /api/users/{id}/name", apiCfg.HandleUpdateUserName)
	mux.HandleFunc("PUT /api/users/{id}/password", apiCfg.HandleUpdateUserPassword)

	// Sales Rep routes
	mux.HandleFunc("GET /api/customers", apiCfg.HandleListCustomers)
	mux.HandleFunc("POST /api/customers", apiCfg.HandleCreateCustomer)
	mux.HandleFunc("GET /api/customers/{customerId}", apiCfg.HandleGetCustomer)
	mux.HandleFunc("PUT /api/customers/{customerId}", apiCfg.HandleUpdateCustomer)
	mux.HandleFunc("DELETE /api/customers/{customerId}", apiCfg.HandleDeleteCustomer)

	// Sales Rep routes
	mux.HandleFunc("GET /api/customers/{customerId}/locations", apiCfg.HandleListCustomerLocations)
	mux.HandleFunc("POST /api/customers/{customerId}/locations", apiCfg.HandleAddCustomerLocation)
	mux.HandleFunc("DELETE /api/customers/locations/{locationID}", apiCfg.HandleDeleteCustomerLocation)
	mux.HandleFunc("PUT /api/customers/locations/{locationID}", apiCfg.HandleUpdateCustomerLocation)
	mux.HandleFunc("GET /api/customers/{customerId}/locations/{locationID}", apiCfg.HandleGetCustomerLocation)

	// Sales Rep routes
	mux.HandleFunc("GET /api/sales-reps", apiCfg.HandleListSalesReps)
	mux.HandleFunc("POST /api/sales-reps", apiCfg.HandleCreateSalesRep)
	mux.HandleFunc("GET /api/sales-reps/{id}", apiCfg.HandleGetSalesRepByID)
	mux.HandleFunc("PUT /api/sales-reps/{id}", apiCfg.HandleUpdateSalesRep)
	mux.HandleFunc("DELETE /api/sales-reps/{id}", apiCfg.HandleDeleteSalesRep)

	// Product routes
	mux.HandleFunc("POST /api/products", apiCfg.HandleCreateProduct)
	mux.HandleFunc("GET /api/products/{id}", apiCfg.HandleGetProductByID)
	mux.HandleFunc("PUT /api/products/{id}", apiCfg.HandleUpdateProduct)
	mux.HandleFunc("DELETE /api/products/{id}", apiCfg.HandleDeleteProduct)
	mux.HandleFunc("GET /api/products/sku/{sku}", apiCfg.HandleGetProductBySKU)

	// Run the server on the specified port
	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, mux)
}
