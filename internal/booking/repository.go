package booking

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	dataSourceName := "host=localhost port=5432 user=postgres password=12qw!@QW dbname=spacetrip_db sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	//ttest the database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
}

func SaveBooking(req BookingRequest) error {
	_, err := db.Exec("INSERT INTO bookings (first_name, last_name, gender, birthday, launchpad_id, destination, launch_date) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		req.FirstName, req.LastName, req.Gender, req.Birthday, req.LaunchpadID, req.Destination, req.LaunchDate)
	return err
}

func GetAllBookings() ([]BookingRequest, error) {
	rows, err := db.Query("SELECT first_name, last_name, gender, birthday, launchpad_id, destination, launch_date FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []BookingRequest
	for rows.Next() {
		var b BookingRequest
		err := rows.Scan(&b.FirstName, &b.LastName, &b.Gender, &b.Birthday, &b.LaunchpadID, &b.Destination, &b.LaunchDate)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}

	return bookings, nil
}
