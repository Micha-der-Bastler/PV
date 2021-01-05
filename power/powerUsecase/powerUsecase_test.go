package powerUsecase_test

import (
	"github.com/michaderbastler/pv/domain"
	"github.com/michaderbastler/pv/domain/mocks"
	"github.com/michaderbastler/pv/power/powerUsecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetPower(t *testing.T) {
	/*** Arrange ***/
	powRepoMock := new(mocks.PowerRepositoryRest)
	powRepoMock.On("GetPower", "http://shelly1pm-BA0F5F").Return(domain.Power(100), nil)
	powUc := powerUsecase.NewPowerUsecase(powRepoMock)

	/*** Act ***/
	resPow, resErr := powUc.GetPower("http://shelly1pm-BA0F5F")

	/*** Assert ***/
	assert.Equal(t, domain.Power(100), resPow)
	assert.Equal(t, nil, resErr)
	// Verify that powUc.GetPower called the mocked powRepoMock.GetPower method as expected
	powRepoMock.AssertExpectations(t)
}
