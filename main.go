package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is application designed by Snehal for Demo Purpose")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Application is healthy")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	port := ":8080"

	log.Println("========================================")
	log.Println("Go Backend Application Started")
	log.Println("Designed by Snehal for Demo Purpose")
	log.Println("Server running on port 8080")
	log.Println("========================================")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}


