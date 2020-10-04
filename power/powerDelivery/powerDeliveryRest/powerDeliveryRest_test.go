package powerDeliveryRest_test

import (
	"PV/domain/mocks"
	"PV/power/powerDelivery/powerDeliveryRest"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetPower(t *testing.T) {
	// Arrange
	e := echo.New()
	powUcMock := mocks.NewPowerUsecase()
	powerDeliveryRest.NewPowerDeliveryRest(e, powUcMock)

	req := httptest.NewRequest("GET", "/power", nil)
	responseRecorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(responseRecorder, req)

	// Assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "{\"power\":100}\n", responseRecorder.Body.String())
}
