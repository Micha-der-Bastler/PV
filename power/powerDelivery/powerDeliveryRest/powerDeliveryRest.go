package powerDeliveryRest

import (
	"PV/domain"
	"github.com/labstack/echo"
	"net/http"
)

// powerDeliveryRest represents the  REST delivery layer of the power domain.
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
		return con.JSON(http.StatusInternalServerError, "")
	}
	powTo := &domain.PowerTo{Pow: pow}
	return con.JSON(http.StatusOK, powTo)
}
