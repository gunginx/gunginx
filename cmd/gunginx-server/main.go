package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gunginx/gunginx/internal/engine"
	"github.com/joho/godotenv"

	"time"
)

type Env struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func loadEnv() (*Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file")
	}

	port := os.Getenv("GUNGINX_PORT")
	if port == "" {
		return nil, fmt.Errorf("GUNGINX_PORT environment variable not set")
	}

	readTimeout, errRead := strconv.Atoi(os.Getenv("LOAD_BALANCER_READ_TIMEOUT"))
	if errRead != nil {
		return nil, fmt.Errorf("LOAD_BALANCER_READ_TIMEOUT environment variable not set/invalid")
	}
	writeTimeout, errWrite := strconv.Atoi(os.Getenv("LOAD_BALANCER_WRITE_TIMEOUT"))
	if errWrite != nil {
		return nil, fmt.Errorf("LOAD_BALANCER_WRITE_TIMEOUT environment variable not set/invalid")
	}

	idleTimeout, errIdle := strconv.Atoi(os.Getenv("LOAD_BALANCER_IDLE_TIMEOUT"))
	if errIdle != nil {
		return nil, fmt.Errorf("LOAD_BALANCER_IDLE_TIMEOUT environment variable not set/invalid")
	}

	env := Env{
		Port:         port,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	return &env, nil
}

func main() {
	env, err := loadEnv()
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	pool := &engine.ServerPool{}
	pool.AddBackend("http://localhost:8081")
	pool.AddBackend("http://localhost:8082")
	pool.AddBackend("http://localhost:8083")

	// ServerPool implements the http.Handler interface(ServeHTTP method), so we can pass it directly to the server
	server := http.Server{
		Addr:         ":" + env.Port,
		Handler:      pool,
		ReadTimeout:  time.Duration(env.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(env.IdleTimeout) * time.Second,
	}

	fmt.Printf("Gunginx Load Balancer started at http://localhost:%s\n", env.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Gunginx crashed: %v", err)
	}
}
