package powerDeliveryRest

import (
	"github.com/labstack/echo"
	"github.com/michaderbastler/pv/domain"
	"net/http"
	"strconv"
	"strings"
)

// ResponseError is used to return the error message as JSON object
type ResponseError struct {
	Message string `json:"message"`
}

// powerDeliveryRest represents the REST delivery layer of the power domain.
type powerDeliveryRest struct {
	powUc domain.PowerUsecase
}

// NewPowerDeliveryRest initializes the endpoints of the power domain by registering their routes.
func NewPowerDeliveryRest(e *echo.Echo, powUc domain.PowerUsecase) {
	powDelRest := &powerDeliveryRest{powUc: powUc}

	// Routes
	e.GET("/power", powDelRest.GetPower)
}

// Handler "GetPower" returns the current power as JSON.
func (powDelRest *powerDeliveryRest) GetPower(con echo.Context) error {
	pow, err := powDelRest.powUc.GetPower("http://shelly1pm-BA0F5F")
	if err != nil {
		return con.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	powTo := &domain.PowerTo{Pow: pow}
	return con.JSON(http.StatusOK, powTo)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	} else if strings.HasPrefix(err.Error(), strconv.Itoa(http.StatusGatewayTimeout)) {
		return http.StatusGatewayTimeout
	} else if strings.HasPrefix(err.Error(), strconv.Itoa(http.StatusBadGateway)) {
		return http.StatusBadGateway
	} else {
		return http.StatusInternalServerError
	}
}
