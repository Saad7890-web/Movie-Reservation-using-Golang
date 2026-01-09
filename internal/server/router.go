package server

import (
	"movie-reservation-system/internal/handler"
	"net/http"
)

func SetupRoutes(authHandler *handler.AuthHandler) http.Handler{
	mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	
	mux.Handle("/auth/signup", http.HandlerFunc(authHandler.SignUp))
	mux.Handle("/auth/login", http.HandlerFunc(authHandler.Login))

	return mux

}