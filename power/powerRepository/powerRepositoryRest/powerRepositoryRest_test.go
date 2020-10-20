package powerRepositoryRest_test

import (
	"github.com/michaderbastler/pv/domain"
	"github.com/michaderbastler/pv/power/powerRepository/powerRepositoryRest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetPower_PosValues(t *testing.T) {
	// Arrange 1/2
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(resWri http.ResponseWriter, req *http.Request) {
		// Assert 1/2
		// Test request parameters
		assert.Equal(t, "/status", req.URL.Path)
		assert.Equal(t, "GET", req.Method)
		// Arrange 2/2
		// Send response
		resWri.Write([]byte(`{"wifi_sta":{"connected":true,"ssid":"SpaetzleMitSauceToGo (5 GHz)","ip":"192.168.1.87",` +
			`"rssi":-34},"cloud":{"enabled":false,"connected":false},"mqtt":{"connected":false},"time":"14:36",` +
			`"unixtime":1600439798,"serial":201,"has_update":false,"mac":"A4CF12BA0F5F","cfg_changed_cnt":0,` +
			`"actions_stats":{"skipped":0},"relays":[{"ison":true,"has_timer":false,"timer_started":0,"timer_duration":0,` +
			`"timer_remaining":0,"overpower":false,"source":"http"}],"meters":[{"power":4.55,"overpower":0.00,` +
			`"is_valid":true,"timestamp":1600439798,"counters":[4.764, 4.757, 4.746],"total":102}],"inputs":[{"input":0,` +
			`"event":"","event_cnt":0}],"ext_sensors":{},"ext_temperature":{},"ext_humidity":{},"temperature":43.41,` +
			`"overtemperature":false,"tmp":{"tC":43.41,"tF":110.13, "is_valid":true},"update":{"status":"idle",` +
			`"has_update":false,"new_version":"20200827-070450/v1.8.3@4a8bc427",` +
			`"old_version":"20200827-070450/v1.8.3@4a8bc427"},"ram_total":50712,"ram_free":38448,"fs_size":233681,` +
			`"fs_free":149847,"uptime":11006}`))
	}))
	// Close the server when test finishes
	defer server.Close()
	// Create test object
	pRR := powerRepositoryRest.NewPowerRepositoryRest()

	// Act
	res, err := pRR.GetPower(server.URL)

	// Assert 2/2
	assert.Equal(t, domain.Power(4.55), res)
	assert.Equal(t, nil, err)
}

func Test_GetPower_NegValues(t *testing.T) {
	// Arrange 1/2
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(resWri http.ResponseWriter, req *http.Request) {
		// Assert 1/2
		// Test request parameters
		assert.Equal(t, "/status", req.URL.Path)
		assert.Equal(t, "GET", req.Method)
		// Arrange 2/2
		// Send response
		resWri.Write([]byte(`{"wifi_sta":{"connected":true,"ssid":"SpaetzleMitSauceToGo (5 GHz)","ip":"192.168.1.87",` +
			`"rssi":-34},"cloud":{"enabled":false,"connected":false},"mqtt":{"connected":false},"time":"14:36",` +
			`"unixtime":1600439798,"serial":201,"has_update":false,"mac":"A4CF12BA0F5F","cfg_changed_cnt":0,` +
			`"actions_stats":{"skipped":0},"relays":[{"ison":true,"has_timer":false,"timer_started":0,"timer_duration":0,` +
			`"timer_remaining":0,"overpower":false,"source":"http"}],"meters":[{"power":-4.55,"overpower":0.00,` +
			`"is_valid":true,"timestamp":1600439798,"counters":[4.764, 4.757, 4.746],"total":102}],"inputs":[{"input":0,` +
			`"event":"","event_cnt":0}],"ext_sensors":{},"ext_temperature":{},"ext_humidity":{},"temperature":43.41,` +
			`"overtemperature":false,"tmp":{"tC":43.41,"tF":110.13, "is_valid":true},"update":{"status":"idle",` +
			`"has_update":false,"new_version":"20200827-070450/v1.8.3@4a8bc427",` +
			`"old_version":"20200827-070450/v1.8.3@4a8bc427"},"ram_total":50712,"ram_free":38448,"fs_size":233681,` +
			`"fs_free":149847,"uptime":11006}`))
	}))
	// Close the server when test finishes
	defer server.Close()
	// Create test object
	powRepoRest := powerRepositoryRest.NewPowerRepositoryRest()

	// Act
	res, err := powRepoRest.GetPower(server.URL)

	// Assert 2/2
	assert.Equal(t, domain.Power(-4.55), res)
	assert.Equal(t, nil, err)
}
