package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fepozopo/oatmeal-studios-backend/internal/auth"
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

	// Set the expires time for the JWT token
	tokenTime := 7200 * time.Second // 2 hours
	token, err := auth.MakeJWT(user.ID, cfg.TokenSecret, tokenTime)
	if err != nil {
		http.Error(w, "Failed to generate access token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)

	// Generate a refresh token and store it in the database
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		http.Error(w, "Failed to generate refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	refreshTokenInput := service.CreateRefreshTokenInput{
		Token:  refreshToken,
		UserID: user.ID,
	}
	err = service.CreateRefreshToken(r.Context(), cfg.DbQueries, refreshTokenInput)
	if err != nil {
		http.Error(w, "Failed to store refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userResponse := UserResponse{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Token:        token,
		RefreshToken: refreshToken,
	}

	// Return the user response with the token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func (cfg *ApiConfig) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, err := idFromURLAsUUID(r)
	if userID == uuid.Nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	user, err := service.GetUserByID(r.Context(), cfg.DbQueries, userID)
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

// HandleRefreshToken processes a request to refresh a user's access token using
// their refresh token. It expects the refresh token to be provided in the
// Authorization header of the request. If the token is missing or invalid,
// it returns an appropriate error response. If the token is valid, it
// generates a new access token and returns it in the response body. If there
// is an error generating the new token, it responds with a 500 status and an
// error message.
func (cfg *ApiConfig) HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	user, err := service.GetUserFromRefreshToken(r.Context(), cfg.DbQueries, refreshToken)
	if err != nil {
		http.Error(w, "Failed to get user from refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokenTime := 7200 * time.Second // 2 hours
	token, err := auth.MakeJWT(user.ID, cfg.TokenSecret, tokenTime)
	if err != nil {
		http.Error(w, "Failed to generate access token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := NewJWT{
		Token:     token,
		ExpiresIn: int64(tokenTime / time.Second),
	}
	json.NewEncoder(w).Encode(response)
}

// HandleRevokeRefresh processes a request to revoke a user's refresh token.
// It extracts the refresh token from the Authorization header, validates it,
// and marks the associated refresh token as revoked in the database. If the
// refresh token is missing or invalid, or if there is an error storing the
// revocation, it responds with an appropriate error status and error message.
func (cfg *ApiConfig) HandleRevokeRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	err = service.RevokeRefreshToken(r.Context(), cfg.DbQueries, refreshToken)
	if err != nil {
		http.Error(w, "Failed to revoke refresh token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
