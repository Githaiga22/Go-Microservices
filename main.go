// main.go
package main

import (
	"log"
	"net/http"
	"micro/services/booking"
	"micro/services/customer"
	"micro/services/room"
)

func main() {
	// Initialize routes for each service
	booking.RegisterRoutes()
	customer.RegisterRoutes()  // Ensure this is implemented like booking
	room.RegisterRoutes()      // Ensure this is implemented like booking

	// Add a root handler for the base path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Hotel Booking API"))
	})

	log.Println("Starting hotel booking microservices...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
