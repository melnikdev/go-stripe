package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/melnikdev/go-stripe/internal/config"
	"github.com/melnikdev/go-stripe/internal/database"
	"github.com/melnikdev/go-stripe/internal/server"
)

func main() {
	log.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := config.NewConfig()

	db, err := database.NewDB(config)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatal("Error migrating database")
	}

	log.Println("Database connected")

	server := server.NewServer(config, db)

	done := make(chan bool, 1)

	go gracefulShutdown(server, done)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	<-done
	log.Println("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}
