package main

import (
	"github.com/Micha-der-Bastler/pv/power/powerDelivery/powerDeliveryRest"
	"github.com/Micha-der-Bastler/pv/power/powerRepository/powerRepositoryRest"
	"github.com/Micha-der-Bastler/pv/power/powerUsecase"
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
