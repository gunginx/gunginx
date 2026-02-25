package main

import (
	"fmt"
	"github.com/gunginx/gunginx/internal/engine"
	"log"
	"net/http"
)

func main() {
	pool := &engine.ServerPool{}

	pool.AddBackend("http://localhost:8081")
	pool.AddBackend("http://localhost:8082")
	pool.AddBackend("http://localhost:8083")

	// ServerPool implements the http.Handler interface(ServeHTTP method), so we can pass it directly to the server
	server := http.Server{
		Addr:    ":8080",
		Handler: pool,
	}

	fmt.Println("Gunginx Load Balancer started at http://localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Gunginx crashed: %v", err)
	}
}
