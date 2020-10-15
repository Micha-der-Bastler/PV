package shelly

import "github.com/Micha-der-Bastler/pv/domain"

type Meter struct {
	Power domain.Power
}

type Status struct {
	Meters []Meter
}
