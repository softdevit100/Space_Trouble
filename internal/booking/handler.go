package booking

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BookTicket(w http.ResponseWriter, r *http.Request) {
	var req BookingRequest
	// body, _ := io.ReadAll(r.Body)
	// fmt.Println("Request Body:", string(body))
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Decode error:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := ValidateBooking(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//we call the booking service to save the booking in the database
	if err := CreateBooking(req); err != nil {
		http.Error(w, "Unable to create booking", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Booking created successfully"})
	w.WriteHeader(http.StatusCreated)

}

func GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := FetchBookings()
	if err != nil {
		http.Error(w, "Unable to fetch bookings", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
