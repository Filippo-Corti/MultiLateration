package controllers

import (
	"image"
	"log/slog"
	"positioning/pkg/models"
	ml "positioning/pkg/multilateration"
	"positioning/pkg/views"
)

type SpaceController struct {
	Space     *models.Space
	SpaceView *views.SpaceView
}

func NewSpaceController(width, height int) *SpaceController {
	space := models.NewSpace(width, height)
	spaceView := views.NewSpaceView(space)
	return &SpaceController{
		Space:     space,
		SpaceView: spaceView,
	}
}

func (sc *SpaceController) RenderView(showStationsArea bool) *image.RGBA {
	return sc.SpaceView.Render(showStationsArea)
}

func (sc *SpaceController) AddStation(x, y float64) error {
	station := models.NewStation(x, y, 0)

	if err := sc.Space.AddStation(station); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (sc *SpaceController) SetTarget(x, y float64) error {
	target := models.NewTarget(x, y)

	if err := sc.Space.SetTarget(target); err != nil {
		slog.Error(err.Error())
		return err
	}

	for _, station := range sc.Space.Stations {
		station.DetectTarget(target)
	}

	return nil
}

func (sc *SpaceController) OperateExactMultilateration() {
	slog.Info("Starting Exact Distance Multilateration")

	targetPosition, err := ml.ExactMultilateration(sc.Space.Stations)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	sc.SpaceView.SetEstimatedTarget(targetPosition)
	slog.Info("Multilateration finished", "targetPosition", targetPosition)
}
