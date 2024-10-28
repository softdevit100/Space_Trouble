package spacex

import (
	"encoding/json"
	"net/http"
	"time"
)

type SpaceXLaunch struct {
	LaunchpadID string    `json:"launchpad"`
	LaunchDate  time.Time `json:"date_utc"`
}

func SpaceXHasLaunch(launchpadID string, launchDate time.Time) bool {
	url := "https://api.spacexdata.com/v4/launches/upcoming"
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var launches []SpaceXLaunch
	if err := json.NewDecoder(resp.Body).Decode(&launches); err != nil {
		return false
	}

	for _, launch := range launches {
		if launch.LaunchpadID == launchpadID && launch.LaunchDate.Equal(launchDate) {
			return true
		}
	}
	return false
}
