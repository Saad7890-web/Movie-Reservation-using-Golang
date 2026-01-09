package main

import (
	"log"

	"movie-reservation-system/internal/config"
	"movie-reservation-system/internal/database"
	"movie-reservation-system/internal/handler"
	"movie-reservation-system/internal/repository/postgres"
	"movie-reservation-system/internal/server"
	"movie-reservation-system/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	adminHandler := handler.NewAdminHandler()
	
	mux := server.SetupRoutes(authHandler, adminHandler)
	
	srv := server.New(cfg.AppPort, mux)
	srv.Start()
}
