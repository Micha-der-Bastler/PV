package main

import (
	"github.com/labstack/echo"
	"github.com/michaderbastler/pv/power/powerDelivery/powerDeliveryRest"
	"github.com/michaderbastler/pv/power/powerRepository/powerRepositoryRest"
	"github.com/michaderbastler/pv/power/powerUsecase"
)

func main() {
	e := echo.New()

	powRepoRest := powerRepositoryRest.NewPowerRepositoryRest()
	powUc := powerUsecase.NewPowerUsecase(powRepoRest)
	powerDeliveryRest.NewPowerDeliveryRest(e, powUc)

	e.Start(":8080")
}
