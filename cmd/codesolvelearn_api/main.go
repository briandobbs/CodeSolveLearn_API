package main

import (
	"CodeSolveLearn_API/db"
	"CodeSolveLearn_API/internal/controller"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize the database
	database, err := db.InitDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), 3306)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := &http.Server{Addr: resolveAddr(), Handler: controller.Handler(database)}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("server forced to shutdown: %s\n", err.Error())
	}

	log.Println("server exiting")
}

func resolveAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":9080"
}
