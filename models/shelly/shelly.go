package shelly

import "github.com/michaderbastler/pv/domain"

type Meter struct {
	Power domain.Power
}

type Status struct {
	Meters []Meter
}
