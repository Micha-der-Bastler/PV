package domain

type Power float64

type PowerTo struct {
	Pow Power `json:"power"`
}

type PowerUsecase interface {
	GetPower(baseUrl string) (Power, error)
}

type PowerRepositoryRest interface {
	GetPower(baseUrl string) (Power, error)
}
