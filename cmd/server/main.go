package main

import (
	"log"
	"net/http"

	"movie-reservation-system/internal/config"
	"movie-reservation-system/internal/database"
	"movie-reservation-system/internal/server"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	srv := server.New(cfg.AppPort, mux)
	srv.Start()
}
