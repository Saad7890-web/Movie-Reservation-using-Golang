package server

import (
	"movie-reservation-system/internal/domain"
	"movie-reservation-system/internal/handler"
	"net/http"
)

func SetupRoutes(authHandler *handler.AuthHandler, adminHandler *handler.AdminHandler) http.Handler{
	mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	
	mux.Handle("/auth/signup", http.HandlerFunc(authHandler.SignUp))
	mux.Handle("/auth/login", http.HandlerFunc(authHandler.Login))


	mux.Handle("/admin/dashboard", ApplyMiddleware(
		http.HandlerFunc(adminHandler.Dashboard),
		JWTMiddleware,
		RequireRole(domain.RoleAdmin),
	))
	return mux

}