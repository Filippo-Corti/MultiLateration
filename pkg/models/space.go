package models

import "errors"

type Space struct {
	Width    int
	Height   int
	Stations []*Station
	Target   *Target
}

func NewSpace(width, height int) *Space {
	return &Space{
		Width:    width,
		Height:   height,
		Stations: []*Station{},
	}
}

func (s *Space) inBounds(x, y float64) bool {
	return (x >= 0 && x <= float64(s.Width)) || (y >= 0 && y <= float64(s.Height))
}

func (s *Space) AddStation(station *Station) error {
	if !s.inBounds(station.X, station.Y) {
		return errors.New("station out of bounds")
	}

	s.Stations = append(s.Stations, station)
	return nil
}

func (s *Space) SetTarget(target *Target) error {
	if !s.inBounds(target.X, target.Y) {
		return errors.New("target out of bounds")
	}

	s.Target = target
	return nil
}
