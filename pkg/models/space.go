package models

import "errors"

type Space struct {
	Width    int
	Height   int
	Stations []*Station
}

func NewSpace(width, height int) *Space {
	return &Space{
		Width:    width,
		Height:   height,
		Stations: []*Station{},
	}
}

func (s *Space) AddStation(station *Station) error {
	if station.X < 0 || station.X > float64(s.Width) || station.Y < 0 || station.Y > float64(s.Height) {
		return errors.New("Station out of bounds")
	}

	s.Stations = append(s.Stations, station)
	return nil
}
