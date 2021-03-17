// Main function to run
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// load the environment file
	err := godotenv.Load("assets/.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	// Get port number
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Set a new client
	myClient := &http.Client{Timeout: 10 * time.Second}
	client := NewClient(myClient)
	fs := http.FileServer(http.Dir("assets"))
	// Handle requests
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/search", searchHandler(client))
	mux.HandleFunc("/", indexHandler)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Error in ListenAndServe")
		return
	}
}
