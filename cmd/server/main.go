package main

import (
	"log"
	"net/http"
	"space-trouble/internal/booking"

	"github.com/gorilla/mux"
)

func main() {
	booking.InitDB()

	r := mux.NewRouter()

	//endpoints
	r.HandleFunc("/book", booking.BookTicket).Methods("POST")
	r.HandleFunc("/bookings", booking.GetBookings).Methods("GET")

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
