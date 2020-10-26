package powerUsecase_test

import (
	"github.com/michaderbastler/pv/domain"
	"github.com/michaderbastler/pv/domain/mocks"
	"github.com/michaderbastler/pv/power/powerUsecase"
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
	assert.Equal(t, domain.Power(101), resPow)
	assert.Equal(t, nil, resErr)
}
