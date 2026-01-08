package server

import (
	"movie-reservation-system/internal/handler"
	"net/http"
)

func SetupRoutes(authHandler *handler.AuthHandler) http.Handler{
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})


	mux.HandleFunc("/auth/signup", authHandler.SignUp)
	mux.HandleFunc("/auth/login", authHandler.Login)

	return mux

}