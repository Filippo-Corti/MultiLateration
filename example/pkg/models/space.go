package models

import (
	"errors"
	ml "positioning/pkg/multilateration"
)

type Space struct {
	Width                   int
	Height                  int
	Stations                []*ml.StationData
	TargetPosition          ml.Position
	EstimatedTargetPosition ml.Position
}

func NewSpace(width, height int) *Space {
	return &Space{
		Width:    width,
		Height:   height,
		Stations: []*ml.StationData{},
	}
}

func (s *Space) inBounds(x, y float64) bool {
	return (x >= 0 && x <= float64(s.Width)) || (y >= 0 && y <= float64(s.Height))
}

func (s *Space) AddStation(x, y float64) error {
	if !s.inBounds(x, y) {
		return errors.New("station out of bounds")
	}

	newStation := ml.NewStationData(x, y, 0)
	s.Stations = append(s.Stations, newStation)
	return nil
}

func (s *Space) SetTarget(x, y float64) error {
	if !s.inBounds(x, y) {
		return errors.New("target out of bounds")
	}

	s.TargetPosition = ml.NewPosition(x, y)
	return nil
}

func (s *Space) SetEstimatedTarget(x, y float64) error {
	if !s.inBounds(x, y) {
		return errors.New("estimated target out of bounds")
	}

	s.EstimatedTargetPosition = ml.NewPosition(x, y)
	return nil
}
