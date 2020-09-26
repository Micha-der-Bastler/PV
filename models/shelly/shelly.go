package shelly

import "PV/domain"

type Meter struct {
	Power domain.Power
}

type Status struct {
	Meters []Meter
}
