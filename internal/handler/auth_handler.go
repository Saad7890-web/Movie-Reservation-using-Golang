package handler

import (
	"encoding/json"
	"movie-reservation-system/internal/security"
	"movie-reservation-system/internal/service"
	"net/http"
)


type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func(h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request){
	var req struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	err := h.authService.Signup(req.Name, req.Email, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func(h *AuthHandler)Login(w http.ResponseWriter, r *http.Request){
	var req struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := security.GenerateToken(user.ID, user.Role.Name)

	json.NewEncoder(w).Encode(map[string]string{
		"token":token,
	})
}