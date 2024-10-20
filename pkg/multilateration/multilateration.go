package multilateration

import (
	"errors"
	"math"

	"positioning/pkg/models"
	"positioning/pkg/utils"
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

func LeastSquaredMultilateration(data []*models.Station) Position {
	
	gradient := func(x, y float64) (resX, resY float64) { 

		for _, station := range data {
			resX += 2 * (x - station.X) - 2 * (x - station.X) * station.DistToTarget * 1 / math.Hypot(x - station.X, y - station.Y)
			resY += 2 * (y - station.Y) - 2 * (y - station.Y) * station.DistToTarget * 1 / math.Hypot(x - station.X, y - station.Y)
		}

		return
	}

	x, y := utils.GradientDescent2V(gradient, 0.1, 100, 1)
	return NewPosition(x, y)
}
