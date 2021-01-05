package powerDeliveryRest_test

import (
	"github.com/labstack/echo"
	"github.com/michaderbastler/pv/domain"
	"github.com/michaderbastler/pv/domain/mocks"
	"github.com/michaderbastler/pv/power/powerDelivery/powerDeliveryRest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetPower_Routes(t *testing.T) {
	/*** Arrange ***/
	// Variables
	path := "/power"

	// Create the test object "power delivery" with a mocked "power usecase"
	e := echo.New()
	powUcMock := new(mocks.PowerUsecase)
	powUcMock.On("GetPower", mock.Anything).Return(domain.Power(100), nil)
	powerDeliveryRest.NewPowerDeliveryRest(e, powUcMock)

	// Setup the test table
	tt := []struct {
		name   string
		method string
		path   string
		status int
	}{
		{
			name:   "GET", // GET requests a resource
			method: http.MethodGet,
			path:   path,
			status: http.StatusOK,
		},
		{
			name:   "POST",
			method: http.MethodPost, // POST creates a resource with a unique ID to be determined by the server
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "PUT",
			method: http.MethodPut, // PUT replaces or creates a resource
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "DELETE",
			method: http.MethodDelete, // DELETE deletes the resource addressed by the URI
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "PATCH",
			method: http.MethodPatch, // PATCH partially updates the resource addressed by the URI
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "HEAD", // HEAD only supplies the HTTP header of a resource
			method: http.MethodHead,
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "OPTIONS", // OPTIONS provides the available communication options of a resource
			method: http.MethodOptions,
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "TRACE", // TRACE is used to track requests that run across multiple nodes
			method: http.MethodTrace,
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "CONNECT", // CONNECT is for connections in which proxy servers act dynamically as a tunnel
			method: http.MethodConnect,
			path:   path,
			status: http.StatusMethodNotAllowed,
		},
		{
			name:   "NOT_FOUND",
			method: http.MethodGet,
			path:   "/poker",
			status: http.StatusNotFound,
		},
	}

	// Iterate over the test table
	for _, tCase := range tt {
		t.Run(tCase.name, func(t *testing.T) {
			req := httptest.NewRequest(tCase.method, tCase.path, nil)
			responseRecorder := httptest.NewRecorder()

			/*** Act ***/
			e.ServeHTTP(responseRecorder, req)

			/*** Assert ***/
			assert.Equal(t, tCase.status, responseRecorder.Code)
			if tCase.name == "GET" {
				assert.Equal(t, "{\"power\":100}\n", responseRecorder.Body.String())
				// TODO: Is this really necessary? Testcase Get makes it true for all test cases
				// Verify that the powDelRest.GetPower handler called the mocked powUcMock.GetPower method as expected
				powUcMock.AssertExpectations(t)
			}
		})
	}
}
