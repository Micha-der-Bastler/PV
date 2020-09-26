package powerRepositoryRest

import (
	"PV/domain"
	"PV/models/shelly"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type powerRepositoryRest struct {
}

// NewPowerRepositoryRest creates an object that fulfilles the domain.PowerRepositorySensor interface
func NewPowerRepositoryRest() domain.PowerRepositorySensor {
	return &powerRepositoryRest{}
}

// GetPower requests the endpoint "/status" of the Shelly device with the given base URL and
// returns tho power of meter[0] of its response.
func (r *powerRepositoryRest) GetPower(baseURL string) (domain.Power, error) {
	// Object to store the JSON value in
	var status shelly.Status

	// Send HTTP GET request
	resp, err := http.Get(baseURL + "/status")
	if err != nil {
		logrus.Error("Error getting response from Shelly. ", err)
		return 0, err
	}
	defer resp.Body.Close()

	// Parse JSON response
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		logrus.Error("Error reading response body from Shelly. ", err)
		return 0, err
	}
	return status.Meters[0].Power, nil
}
