package powerRepositoryRest

import (
	"encoding/json"
	"fmt"
	"github.com/michaderbastler/pv/domain"
	"github.com/michaderbastler/pv/models/shelly"
	"github.com/sirupsen/logrus"
	"net/http"
)

// powerRepositoryRest represents the REST repository layer of the power domain.
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
		// Add additional information to the error
		// 504 Gateway Timeout --> Server acts as gateway and doesn't get the response from another server in time
		err = fmt.Errorf("%v %v: error getting response from shelly: %w",
			http.StatusGatewayTimeout, http.StatusText(http.StatusGatewayTimeout), err)
		logrus.Error(err)
		return 0, err
	}
	defer resp.Body.Close()

	// Parse JSON response
	err = json.NewDecoder(resp.Body).Decode(&status) // Decoder buffers the byte stream as []byte internally
	// before unmarshalling it into a Go value (= parsing)
	if err != nil {
		// Add additional information to the error
		// 502 Bad Gateway --> Server acts as gateway and received a invalid response from another server
		err = fmt.Errorf("%v %v: error parsing response body from shelly: %w",
			http.StatusBadGateway, http.StatusText(http.StatusBadGateway), err)
		logrus.Error(err)
		return 0, err
	}
	return status.Meters[0].Power, nil
}
