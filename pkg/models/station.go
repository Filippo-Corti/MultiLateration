package models

import "positioning/pkg/utils"

type Station struct {
	X            float64
	Y            float64
	DistToTarget float64
}

func NewStation(x, y, d float64) *Station {
	return &Station{
		X:            x,
		Y:            y,
		DistToTarget: d,
	}
}

func (s *Station) DetectTarget(t *Target) {
	// Currently Exact Distance is considered
	// TODO Change to approximation

	s.DistToTarget = utils.Distance(t.X, t.Y, s.X, s.Y)
}
