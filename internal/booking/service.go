package booking

import (
	"errors"
	"space-trouble/internal/spacex"
	"time"
)

type BookingRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	Birthday    time.Time `json:"birthday"`
	LaunchpadID string    `json:"launchpad_id"`
	Destination string    `json:"destination"`
	LaunchDate  time.Time `json:"launch_date"`
}

func ValidateBooking(req BookingRequest) error {
	//check if the destination is valid
	validDestinations := []string{"Mars", "Moon", "Pluto", "Asteroid Belt", "Europa", "Titan", "Ganymede"}
	isValidDestination := false
	for _, destination := range validDestinations {
		if req.Destination == destination {
			isValidDestination = true
			break
		}
	}
	if !isValidDestination {
		return errors.New("invalid destination")
	}

	if spacex.SpaceXHasLaunch(req.LaunchpadID, req.LaunchDate) {
		return errors.New("launchpad is not available on the selected date")
	}

	return nil
}

func CreateBooking(req BookingRequest) error {
	//save the booking in the database (repo logic)
	return SaveBooking(req)
}

func FetchBookings() ([]BookingRequest, error) {
	//Fetch all bookings from the database
	return GetAllBookings()
}
