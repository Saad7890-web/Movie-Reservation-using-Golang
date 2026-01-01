package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/saad7890/movie-reservation/internal/config"
	"github.com/saad7890/movie-reservation/internal/db"
	httpapp "github.com/saad7890/movie-reservation/internal/http"
)

func main() {
	// Load env file
	if err := config.LoadEnv(".env"); err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}

	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// Init DB
	database, err := db.NewPostgres(cfg.DB)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer database.Close()

	log.Println("Database connected")

	// HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      httpapp.NewRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		log.Printf("%s running on http://localhost:%s", cfg.AppName, cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown failed: %v", err)
	}

	log.Println("Server stopped cleanly")
}
