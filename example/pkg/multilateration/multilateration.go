package multilateration

import (
	"errors"
	"math"

	"positioning/pkg/utils"
)

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

// Multilateration to find the Target, using Station Data and the Least Squares Estimation
func LeastSquaredMultilateration(data []*StationData) Position {

	gradient := func(x, y float64) (resX, resY float64) {

		for _, station := range data {
			// Formulas are hand calculated by me
			resX += 2*(x-station.Position.X) - 2*(x-station.Position.X)*station.DistToTarget*1/math.Hypot(x-station.Position.X, y-station.Position.Y)
			resY += 2*(y-station.Position.Y) - 2*(y-station.Position.Y)*station.DistToTarget*1/math.Hypot(x-station.Position.X, y-station.Position.Y)
		}

		return
	}

	x, y := utils.GradientDescent2V(gradient, 0.1, 100, 0.01)
	return NewPosition(x, y)
}

// Multilateration to find the Target, using Station Data and supposing measured distances are exact
func ExactMultilateration(data []*StationData) (Position, error) {
	if len(data) < 3 {
		return Position{}, errors.New("unsufficient data to find target")
	}

	A := 2*data[1].Position.X - 2*data[0].Position.X
	B := 2*data[1].Position.Y - 2*data[0].Position.Y
	C := data[0].DistToTarget*data[0].DistToTarget - data[1].DistToTarget*data[1].DistToTarget -
		data[0].Position.X*data[0].Position.X + data[1].Position.X*data[1].Position.X -
		data[0].Position.Y*data[0].Position.Y + data[1].Position.Y*data[1].Position.Y
	D := 2*data[2].Position.X - 2*data[1].Position.X
	E := 2*data[2].Position.Y - 2*data[1].Position.Y
	F := data[1].DistToTarget*data[1].DistToTarget - data[2].DistToTarget*data[2].DistToTarget -
		data[1].Position.X*data[1].Position.X + data[2].Position.X*data[2].Position.X -
		data[1].Position.Y*data[1].Position.Y + data[2].Position.Y*data[2].Position.Y

	x := (C*E - F*B) / (E*A - B*D)
	y := (C*D - A*F) / (B*D - A*E)

	targetPosition := NewPosition(x, y)

	return targetPosition, nil
}
