package main

import (
	"github.com/briandobbs/CodeSolveLearn_API/internal/controller"
	"net/http"
	"os"
)

func main() {
	server := &http.Server{Addr: resolveAddr(), Handler: controller.Handler()}
}

func resolveAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":9080"
}
