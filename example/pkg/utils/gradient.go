package utils

import (
	"log/slog"
	"math"
	"math/rand"
)

func GradientDescent1V(gradientFunc func(float64) float64, stepSize float64, maxIters int, tolerance float64) (min float64) {
	min = rand.Float64() * 100 //Starting Point

	for i := 0; i < maxIters; i++ {
		diff := stepSize * gradientFunc(min)
		
		if math.Abs(diff) < tolerance {
			break
		}

		min -= diff
	}

	return
}


func GradientDescent2V(gradientFunc func(float64, float64) (float64, float64), stepSize float64, maxIters int, tolerance float64) (minX float64, minY float64) {
	minX, minY = rand.Float64() * 100, rand.Float64() * 100 //Starting Point

	for i := 0; i < maxIters; i++ {
		gradX, gradY := gradientFunc(minX, minY)
		diffX, diffY := stepSize * gradX, stepSize * gradY

		if math.Abs(diffX) < tolerance && math.Abs(diffY) < tolerance {
			break
		}

		slog.Info("Gradient Descent", "step", i, "x", minX, "y", minY)

		minX -= diffX
		minY -= diffY
	}
	
	return
}