package multilateration

import (
	"errors"

	"positioning/pkg/models"
)

const EPSILON = 10e-3

type Position struct {
	X float64
	Y float64
}

func NewPosition(x, y float64) Position {
	return Position{
		X: x,
		Y: y,
	}
}

// Check https://www.101computing.net/cell-phone-trilateration-algorithm/ for formulas
func ExactMultilateration(data []*models.Station) (Position, error) {
	if len(data) < 3 {
		return Position{}, errors.New("unsufficient data to find target")
	}

	A := 2*data[1].X - 2*data[0].X
	B := 2*data[1].Y - 2*data[0].Y
	C := data[0].DistToTarget*data[0].DistToTarget - data[1].DistToTarget*data[1].DistToTarget - data[0].X*data[0].X + data[1].X*data[1].X - data[0].Y*data[0].Y + data[1].Y*data[1].Y
	D := 2*data[2].X - 2*data[1].X
	E := 2*data[2].Y - 2*data[1].Y
	F := data[1].DistToTarget*data[1].DistToTarget - data[2].DistToTarget*data[2].DistToTarget - data[1].X*data[1].X + data[2].X*data[2].X - data[1].Y*data[1].Y + data[2].Y*data[2].Y

	x := (C*E - F*B) / (E*A - B*D)
	y := (C*D - A*F) / (B*D - A*E)

	targetPosition := NewPosition(x, y)

	return targetPosition, nil
}
