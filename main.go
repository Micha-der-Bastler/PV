package main

import (
	"PV/power/powerDelivery/powerDeliveryRest"
	"PV/power/powerRepository/powerRepositoryRest"
	"PV/power/powerUsecase"
	"github.com/labstack/echo"
)

func main() {
	// No change to be able to commit
	e := echo.New()

	powRepoRest := powerRepositoryRest.NewPowerRepositoryRest()
	powUc := powerUsecase.NewPowerUsecase(powRepoRest)
	powerDeliveryRest.NewPowerDeliveryRest(e, powUc)

	e.Start(":8080")
}
