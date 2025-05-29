package api

import (
	"encoding/json"
	"net/http"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/service"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.RegisterUserInput
	// Decode the JSON request body into the input struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := service.RegisterUser(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully", "user_id": user.ID.String()})
}

func (cfg *ApiConfig) HandleAuthenticateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.AuthenticateUserInput
	// Decode the JSON request body into the input struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := service.AuthenticateUser(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to authenticate user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User authenticated successfully", "user_id": user.ID.String()})
}

func (cfg *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Bad Request: User ID is required", http.StatusBadRequest)
		return
	}

	// Validate the UUID format
	uid, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Bad Request: Invalid User ID format", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByID(r.Context(), cfg.DbQueries, uid)
	if err != nil {
		http.Error(w, "Failed to get user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (cfg *ApiConfig) HandleUpdateUserName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.UpdateUserNameInput
	// Decode the JSON request body into the input struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := service.UpdateUserName(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully", "user_id": user.ID.String()})
}

func (cfg *ApiConfig) HandleUpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input service.UpdateUserPasswordInput
	// Decode the JSON request body into the input struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := service.UpdateUserPassword(r.Context(), cfg.DbQueries, input)
	if err != nil {
		http.Error(w, "Failed to update password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}

func (cfg *ApiConfig) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := service.ListUsers(r.Context(), cfg.DbQueries)
	if err != nil {
		http.Error(w, "Failed to list users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (cfg *ApiConfig) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := idFromURLAsUUID(r)
	if err != nil {
		http.Error(w, "Invalid User ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteUser(r.Context(), cfg.DbQueries, userID)
	if err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
