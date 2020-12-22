package shelly

import "github.com/michaderbastler/pv/domain"

type Status struct {
	Relays          []Relay
	Meters          []Meter
	Temperature     float64
	Overtemperature bool
}

type Relay struct {
	IsOn bool `json:"ison"`
}

type Meter struct {
	Power domain.Power
}
