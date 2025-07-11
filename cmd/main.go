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
		DB:          db,
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

	// Customer routes
	mux.HandleFunc("GET /api/customers", apiCfg.HandleListCustomers)
	mux.HandleFunc("POST /api/customers", apiCfg.HandleCreateCustomer)
	mux.HandleFunc("GET /api/customers/{customerId}", apiCfg.HandleGetCustomer)
	mux.HandleFunc("PUT /api/customers/{customerId}", apiCfg.HandleUpdateCustomer)
	mux.HandleFunc("DELETE /api/customers/{customerId}", apiCfg.HandleDeleteCustomer)

	// Customer Location routes
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

	// Order routes
	mux.HandleFunc("POST /api/orders", apiCfg.HandleCreateOrder)
	mux.HandleFunc("GET /api/orders/{id}", apiCfg.HandleGetOrder)
	mux.HandleFunc("PUT /api/orders/{id}", apiCfg.HandleUpdateOrder)
	mux.HandleFunc("DELETE /api/orders/{id}", apiCfg.HandleDeleteOrder)

	// Planogram routes
	mux.HandleFunc("POST /api/planograms", apiCfg.HandleCreatePlanogram)
	mux.HandleFunc("GET /api/planograms/{id}", apiCfg.HandleGetPlanogram)
	mux.HandleFunc("PUT /api/planograms/{id}", apiCfg.HandleUpdatePlanogram)
	mux.HandleFunc("DELETE /api/planograms/{id}", apiCfg.HandleDeletePlanogram)
	mux.HandleFunc("POST /api/planograms/{id}/assign", apiCfg.HandleAssignPlanogramToLocation)
	mux.HandleFunc("DELETE /api/planograms/{id}/remove", apiCfg.HandleRemovePlanogramFromLocation)
	mux.HandleFunc("GET /api/planograms", apiCfg.HandleListPlanograms)
	mux.HandleFunc("GET /api/planograms/{id}/pockets", apiCfg.HandleListPocketsForPlanogram)
	mux.HandleFunc("GET /api/planograms/{id}/locations", apiCfg.HandleListLocationsByPlanogram)
	mux.HandleFunc("GET /api/planograms/{locationID}/planograms", apiCfg.HandleGetPlanogramsByLocation)
	mux.HandleFunc("PUT /api/planograms/{id}/reassign", apiCfg.HandleReassignPlanogramToLocation)

	// Planogram Pocket routes
	mux.HandleFunc("GET /api/planogram-pockets/{id}", apiCfg.HandleGetPlanogramPocket)
	mux.HandleFunc("PUT /api/planogram-pockets/{id}", apiCfg.HandleUpdatePlanogramPocket)
	mux.HandleFunc("DELETE /api/planogram-pockets/{id}", apiCfg.HandleDeletePlanogramPocket)
	mux.HandleFunc("POST /api/planogram-pockets", apiCfg.HandleCreatePlanogramPocket)

	// Invoice routes
	mux.HandleFunc("POST /api/invoices", apiCfg.HandleCreateInvoice)
	mux.HandleFunc("GET /api/invoices/{id}", apiCfg.HandleGetInvoice)
	mux.HandleFunc("PUT /api/invoices/{id}", apiCfg.HandleUpdateInvoice)
	mux.HandleFunc("DELETE /api/invoices/{id}", apiCfg.HandleDeleteInvoice)

	// Run the server on the specified port
	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, mux)
}
