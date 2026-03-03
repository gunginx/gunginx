package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//to fetch port number from user
	if len(os.Args) < 2 {
		log.Fatal("Please provide a port number as an argument")
	}
	port := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[SERVER %s] Received request for: %s\n", port, r.URL.Path)

		fmt.Fprintf(w, "Hello from the server running on port %s\n", port)
	})

	http.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(5 * time.Second)
		fmt.Fprintf(w, "Server slept for 5 seconds")
		fmt.Printf("Server slept for 5 seconds")
	})

	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
