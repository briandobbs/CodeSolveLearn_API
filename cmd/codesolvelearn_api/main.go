package main

import (
	"CodeSolveLearn_API/controller"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: resolveAddr(), Handler: controller.Handler()}

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
	err := server.Shutdown(ctx)
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
