package domain

type Power float64

type PowerRepositorySensor interface {
	GetPower(baseURL string) (Power, error)
}
