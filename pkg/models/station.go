package models

type Station struct {
	X float64
	Y float64
}

func NewStation(x, y float64) *Station {
	return &Station{
		X: x,
		Y: y,
	}
}
