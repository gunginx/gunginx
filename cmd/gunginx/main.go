package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	fmt.Println("Server running at http://localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}

}
