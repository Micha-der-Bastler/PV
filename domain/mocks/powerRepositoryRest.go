package mocks

import (
	"github.com/michaderbastler/pv/domain"
)

type powerRepositoryRest struct {
}

func NewPowerRepositoryRest() domain.PowerRepositoryRest {
	return &powerRepositoryRest{}
}

func (r *powerRepositoryRest) GetPower(baseUrl string) (domain.Power, error) {
	return domain.Power(100), nil
}
