package powerRepositoryRest

import (
	"encoding/json"
	"github.com/Micha-der-Bastler/pv/domain"
	"github.com/Micha-der-Bastler/pv/models/shelly"
	"github.com/sirupsen/logrus"
	"net/http"
)

// powerRepositoryRest represents the  REST repository layer of the power domain.
type powerRepositoryRest struct {
}

// NewPowerRepositoryRest returns a pointer to an object that fulfills the domain.PowerRepositoryRest interface.
func NewPowerRepositoryRest() domain.PowerRepositoryRest {
	return &powerRepositoryRest{}
}

// GetPower requests the endpoint "/status" of the Shelly device with the given base URL and
// returns the power of its meter[0].
func (r *powerRepositoryRest) GetPower(baseUrl string) (domain.Power, error) {
	// Object to store the JSON value in
	var status shelly.Status

	// Send HTTP GET request
	resp, err := http.Get(baseUrl + "/status")
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
