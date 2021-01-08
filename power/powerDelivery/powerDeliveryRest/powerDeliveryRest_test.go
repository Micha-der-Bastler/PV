package powerDeliveryRest_test

import (
	"errors"
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
	defaultPath := "/power"
	defaultMethod := http.MethodGet

	e := echo.New()

	// Setup the test table
	tt := []struct {
		name               string
		method             string
		path               string
		usecaseReturnValue domain.Power
		usecaseReturnError error
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "GET", // GET requests a resource
			method:             http.MethodGet,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "POST", // POST creates a resource with a unique ID to be determined by the server
			method:             http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "PUT", // PUT replaces or creates a resource
			method:             http.MethodPut,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "DELETE", // DELETE deletes the resource addressed by the URI
			method:             http.MethodDelete,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "PATCH", // PATCH partially updates the resource addressed by the URI
			method:             http.MethodPatch,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "HEAD", // HEAD only supplies the HTTP header of a resource
			method:             http.MethodHead,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "OPTIONS", // OPTIONS provides the available communication options of a resource
			method:             http.MethodOptions,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "TRACE", // TRACE is used to track requests that run across multiple nodes
			method:             http.MethodTrace,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "CONNECT", // CONNECT is for connections in which proxy servers act dynamically as a tunnel
			method:             http.MethodConnect,
			expectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			name:               "NOT_FOUND",
			method:             http.MethodGet,
			path:               "/poker",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "usecaseReturns100W",
			usecaseReturnValue: 100,
			expectedStatusCode: http.StatusOK,
			expectedBody:       "{\"power\":100}\n",
		},
		{
			name:               "usecaseReturns504",
			usecaseReturnError: errors.New("504 Gateway Timeout: error description"),
			expectedStatusCode: http.StatusGatewayTimeout,
			expectedBody:       "{\"message\":\"504 Gateway Timeout: error description\"}\n",
		},
		{
			name:               "usecaseReturns504",
			usecaseReturnError: errors.New("502 Bad Gateway: error description"),
			expectedStatusCode: http.StatusBadGateway,
			expectedBody:       "{\"message\":\"502 Bad Gateway: error description\"}\n",
		},
	}

	// Iterate over the test table
	for _, tCase := range tt {
		t.Run(tCase.name, func(t *testing.T) {
			// Fill empty test case fields with default values
			if tCase.path == "" {
				tCase.path = defaultPath
			}
			if tCase.method == "" {
				tCase.method = defaultMethod
			}

			// Create the test server
			req := httptest.NewRequest(tCase.method, tCase.path, nil)
			responseRecorder := httptest.NewRecorder()

			// Create the test object "power delivery" with a mocked "power usecase"
			powUcMock := new(mocks.PowerUsecase)
			powUcMock.On("GetPower", mock.Anything).Return(tCase.usecaseReturnValue, tCase.usecaseReturnError)
			powerDeliveryRest.NewPowerDeliveryRest(e, powUcMock)

			/*** Act ***/
			e.ServeHTTP(responseRecorder, req)

			/*** Assert ***/
			assert.Equal(t, tCase.expectedStatusCode, responseRecorder.Code)
			if tCase.expectedBody != "" {
				assert.Equal(t, tCase.expectedBody, responseRecorder.Body.String())
			}
		})
	}
}
