package mocks

import "PV/domain"

type powerUsecase struct {
}

func NewPowerUsecase() domain.PowerUsecase {
	return &powerUsecase{}
}

func (powUc *powerUsecase) GetPower(baseUrl string) (domain.Power, error) {
	return domain.Power(100), nil
}
