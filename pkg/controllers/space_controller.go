package controllers

import (
	"image"
	"log/slog"
	"positioning/pkg/models"
	"positioning/pkg/views"
)

type SpaceController struct {
	Space *models.Space
}

func NewSpaceController(width, height int) *SpaceController {
	space := models.NewSpace(width, height)
	return &SpaceController{
		Space: space,
	}
}

func (sc *SpaceController) AddStation(x, y float64) error {
	station := models.NewStation(x, y)

	if err := sc.Space.AddStation(station); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (sc *SpaceController) RenderView() *image.RGBA {
	spaceView := views.NewSpaceView(sc.Space)

	return spaceView.Render()
}