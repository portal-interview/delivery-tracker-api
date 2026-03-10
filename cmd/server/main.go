package main

import (
	"log"
	"net/http"
	"os"

	"github.com/portal-interview/delivery-tracker-api/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /deliveries", handler.ListDeliveries)
	mux.HandleFunc("GET /deliveries/{id}", handler.GetDelivery)
	mux.HandleFunc("PUT /deliveries/{id}/location", handler.UpdateLocation)

	log.Printf("Delivery Tracker API starting on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
