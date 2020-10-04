package powerUsecase_test

import (
	"PV/domain"
	"PV/domain/mocks"
	"PV/power/powerUsecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetPower(t *testing.T) {
	// Arrange
	powRepoMock := mocks.NewPowerRepositoryRest()
	powUc := powerUsecase.NewPowerUsecase(powRepoMock)

	// Act
	resPow, resErr := powUc.GetPower("")

	// Assert
	assert.Equal(t, domain.Power(100), resPow)
	assert.Equal(t, nil, resErr)
}
