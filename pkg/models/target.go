package models

type Target struct {
	X float64
	Y float64
}

func NewTarget(x, y float64) *Target {
	return &Target{
		X: x,
		Y: y,
	}
}
