package controllers

import (
	"image"
	"log/slog"
	"math"
	"math/rand/v2"
	"positioning/pkg/models"
	ml "positioning/pkg/multilateration"
	"positioning/pkg/utils"
	"positioning/pkg/views"
)

type TargetDistanceCalculationMode int

const (
	DistanceExact TargetDistanceCalculationMode = iota
	DistanceWithNormalError
)

func getStdDevByCalculationMode(mode TargetDistanceCalculationMode) float64 {
	switch mode {
	case DistanceExact:
		return 0
	case DistanceWithNormalError:
		return 50
	default:
		return -1
	}
}

type SpaceController struct {
	Space     *models.Space
	SpaceView *views.SpaceView

	ExecutionMode TargetDistanceCalculationMode
}

func NewSpaceController(width, height int, distCalcMode TargetDistanceCalculationMode) *SpaceController {
	space := models.NewSpace(width, height)
	spaceView := views.NewSpaceView(space)
	return &SpaceController{
		Space:     space,
		SpaceView: spaceView,
		ExecutionMode: distCalcMode,
	}
}

func (sc *SpaceController) RenderView(showStationsArea bool) *image.RGBA {
	return sc.SpaceView.Render(showStationsArea)
}

func (sc *SpaceController) AddStation(x, y float64) error {
	if err := sc.Space.AddStation(x, y); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (sc *SpaceController) SetTarget(x, y float64) error {
	if err := sc.Space.SetTarget(x, y); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (sc *SpaceController) ScanForTarget() {
	stdDev := getStdDevByCalculationMode(sc.ExecutionMode)

	for _, station := range sc.Space.Stations {
		station.DistToTarget = detectDistanceWithError(station.Position, sc.Space.TargetPosition, stdDev)
	}
}

// stdDev = 0 <==> No Error
func detectDistanceWithError(a, b ml.Position, stdDev float64) float64 {
	exactDistance := utils.Distance(a.X, a.Y, b.X, b.Y)

	// Normal Distribution around exactDistance
	return exactDistance - math.Abs(rand.NormFloat64())*stdDev
}

func (sc *SpaceController) MultilaterateToFindTarget() {
	switch sc.ExecutionMode {
	case DistanceExact:
		sc.operateExactMultilateration()
	case DistanceWithNormalError:
		sc.operateMultiLaterationWithError()
	default:
		slog.Error("Should not happen")
	}
}

func (sc *SpaceController) operateExactMultilateration() {
	slog.Info("Starting Exact Distance Multilateration")

	targetPosition, err := ml.ExactMultilateration(sc.Space.Stations)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	sc.Space.EstimatedTargetPosition = targetPosition
	slog.Info("Multilateration finished", "targetPosition", targetPosition)
}

func (sc *SpaceController) operateMultiLaterationWithError() {
	slog.Info("Starting Multilateration with Error in Distance Value")

	targetPosition := ml.LeastSquaredMultilateration(sc.Space.Stations)

	sc.Space.EstimatedTargetPosition = targetPosition
	slog.Info("Multilateration finished", "targetPosition", targetPosition)

}
