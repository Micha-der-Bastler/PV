package powerUsecase

import "github.com/michaderbastler/pv/domain"

// powerUsecase represents the usecase layer of the power domain.
type powerUsecase struct {
	powRepoRest domain.PowerRepositoryRest
}

// NewPowerUsecase returns a pointer to an object that fulfills the domain.PowerUsecase interface.
func NewPowerUsecase(powRepoRest domain.PowerRepositoryRest) domain.PowerUsecase {
	return &powerUsecase{powRepoRest: powRepoRest}
}

// GetPower calls the method "GetPower" of the repository layer with the given "baseUrl" and returns tho power
// together with its error.
func (powUc *powerUsecase) GetPower(baseUrl string) (domain.Power, error) {
	return powUc.powRepoRest.GetPower(baseUrl)
}
