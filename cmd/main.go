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

	// Register routes
	mux.HandleFunc("POST /api/users/register", apiCfg.HandleRegisterUser)
	mux.HandleFunc("POST /api/users/authenticate", apiCfg.HandleAuthenticateUser)
	mux.HandleFunc("GET /api/users/{id}", apiCfg.HandleGetUser)
	mux.HandleFunc("PUT /api/users/{id}/name", apiCfg.HandleUpdateUserName)
	mux.HandleFunc("PUT /api/users/{id}/password", apiCfg.HandleUpdateUserPassword)
	mux.HandleFunc("GET /api/customers", apiCfg.HandleListCustomers)
	mux.HandleFunc("POST /api/customers", apiCfg.HandleCreateCustomer)
	mux.HandleFunc("GET /api/customers/{id}", apiCfg.HandleGetCustomer)
	mux.HandleFunc("PUT /api/customers/{id}", apiCfg.HandleUpdateCustomer)
	mux.HandleFunc("DELETE /api/customers/{id}", apiCfg.HandleDeleteCustomer)

	// Run the server on the specified port
	log.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, mux)
}
