package models

import (
	"math"
	"math/rand"
	"positioning/pkg/utils"
)

type Station struct {
	X            float64
	Y            float64
	DistToTarget float64
}

const APPROXIMATION_STDDEV = 50;

func NewStation(x, y, d float64) *Station {
	return &Station{
		X:            x,
		Y:            y,
		DistToTarget: d,
	}
}

func (s *Station) DetectTarget(t *Target) {
	s.DistToTarget = utils.Distance(t.X, t.Y, s.X, s.Y)
}

func (s *Station) DetectTargetWithApproxError(t *Target) {
	exactDistance := utils.Distance(t.X, t.Y, s.X, s.Y)

	// Normal Distribution around exactDistance
	s.DistToTarget = - math.Abs(rand.NormFloat64()) * APPROXIMATION_STDDEV + exactDistance
}
